package lark

import (
	"context"
	"fmt"

	"github.com/goto/salt/log"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/httpclient"
	"github.com/goto/siren/pkg/retry"
	"github.com/goto/siren/pkg/secret"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

const (
	defaultLarkAPIHost = "https://open.larksuite.com"
	larkPathOAuth      = "/open-apis/auth/v3/tenant_access_token/internal/"
)

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientOption func(*Client)

// ClientWithHTTPClient assigns custom http client when creating a lark client
func ClientWithHTTPClient(httpClient *httpclient.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// ClientWithRetrier wraps client call with retrier
func ClientWithRetrier(runner retry.Runner) ClientOption {
	return func(c *Client) {
		// note: for now retry only happen in send message context method
		c.retrier = runner
	}
}

type Client struct {
	cfg        AppConfig
	httpClient *httpclient.Client
	retrier    retry.Runner
	logger     log.Logger
}

// NewClient is a constructor to create lark client.
// this version uses lark v3 SDK.
func NewClient(cfg AppConfig, logger log.Logger, opts ...ClientOption) *Client {
	c := &Client{
		cfg:    cfg,
		logger: logger,
	}
	for _, opt := range opts {
		opt(c)
	}

	if cfg.APIHost == "" {
		c.cfg.APIHost = defaultLarkAPIHost
	}

	// sanitize
	c.cfg.APIHost = c.cfg.APIHost + "/"

	if c.httpClient == nil {
		c.httpClient = httpclient.New(cfg.HTTPClient)
	}

	return c
}

// GetWorkspaceChannels fetches list of joined channel of a client
func (c *Client) GetWorkspaceChannels(ctx context.Context, clientID, clientSecret secret.MaskableString) ([]Channel, error) {
	var client = lark.NewClient(clientID.UnmaskedString(), clientSecret.UnmaskedString(), lark.WithOpenBaseUrl(c.cfg.APIHost))

	joinedChannelList, err := c.getJoinedChannelsList(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch joined channel list: %w", err)
	}

	result := make([]Channel, 0)
	for _, c := range joinedChannelList {
		result = append(result, Channel{
			ID:   *c.ChatId,
			Name: *c.Name,
		})
	}
	return result, nil
}

func (c *Client) Notify(ctx context.Context, conf NotificationConfig, message Message) error {
	if c.retrier != nil {
		if err := c.retrier.Run(ctx, func(ctx context.Context) error {
			return c.notify(ctx, conf, message)
		}); err != nil {
			return err
		}
	}
	return c.notify(ctx, conf, message)
}

// Notify sends message to a specific lark channel
func (c *Client) notify(ctx context.Context, conf NotificationConfig, message Message) error {

	var client = lark.NewClient(conf.ClientID.UnmaskedString(), conf.ClientSecret.UnmaskedString(), lark.WithOpenBaseUrl(c.cfg.APIHost))

	var channelID string
	var receiverType string
	switch conf.ChannelType {
	case TypeChannelChannel:
		joinedChannelList, err := c.getJoinedChannelsList(ctx, client)
		if err != nil {
			if err := c.checkLarkErrorRetryable(err); errors.As(err, new(retry.RetryableError)) {
				return err
			}
			return fmt.Errorf("failed to fetch joined channel list: %w", err)
		}
		channelID = searchChannelId(joinedChannelList, message.Channel)
		if channelID == "" {
			return fmt.Errorf("app is not part of the channel %q", message.Channel)
		}
		receiverType = "chat_id"
	case TypeChannelUser:
		_, err := c.getUserByEmail(ctx, message.Channel, client)
		if err != nil {
			if err.Error() == "users_not_found" {
				return fmt.Errorf("failed to get id for %q", message.Channel)
			}
			return c.checkLarkErrorRetryable(err)
		}
		channelID = message.Channel
		receiverType = "email"
	default:
		return fmt.Errorf("unknown receiver type %q", conf.ChannelType)
	}

	msgOptions, err := message.BuildLarkMessage()
	if err != nil {
		return err
	}

	if err := c.sendMessageContext(ctx, client, receiverType, channelID, msgOptions); err != nil {
		if err := c.checkLarkErrorRetryable(err); errors.As(err, new(retry.RetryableError)) {
			return err
		}
		return fmt.Errorf("failed to send message to %q: %w", message.Channel, err)
	}

	return nil
}

func (c *Client) sendMessageContext(ctx context.Context, client *lark.Client, receiverType, channelID string, msgOpts string) error {
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(receiverType).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(channelID).
			MsgType(`interactive`).
			Content(msgOpts).
			Build()).
		Build()

	resp, err := client.Im.Message.Create(ctx, req)
	if err != nil {
		return c.checkLarkErrorRetryable(err)
	}

	if !resp.Success() {
		return fmt.Errorf("failed to send message: %s", larkcore.Prettify(resp))
	}

	c.logger.Debug("Lark client.Im.Message.Create response", "response", larkcore.Prettify(resp.Data))
	return nil
}

func (c *Client) checkLarkErrorRetryable(err error) error {
	return retry.RetryableError{Err: err}
}

func (c *Client) getJoinedChannelsList(ctx context.Context, client *lark.Client) (list []*larkim.ListChat, err error) {
	list = []*larkim.ListChat{}
	curr := ""

	for {
		req := larkim.NewListChatReqBuilder().Limit(1000).PageToken(curr).Build()
		resp, err := client.Im.Chat.List(ctx, req)
		if err != nil {
			return list, err
		}

		if !resp.Success() {
			return nil, fmt.Errorf("failed to fetch channel list: %s", larkcore.Prettify(resp))
		}

		list = append(list, resp.Data.Items...)

		if resp.Data.PageToken == nil || *resp.Data.PageToken == "" {
			break
		}
		curr = *resp.Data.PageToken
	}

	return list, nil
}

func searchChannelId(channels []*larkim.ListChat, channelName string) string {
	for _, c := range channels {
		if *c.Name == channelName {
			return *c.ChatId
		}
	}
	return ""
}

func (c *Client) getUserByEmail(ctx context.Context, email string, client *lark.Client) (string, error) {

	req := larkcontact.NewBatchGetIdUserReqBuilder().
		Body(larkcontact.NewBatchGetIdUserReqBodyBuilder().
			Emails([]string{email}).
			IncludeResigned(true).
			Build()).
		Build()
	resp, err := client.Contact.User.BatchGetId(ctx, req)
	if err != nil {
		return "", err
	}

	if !resp.Success() {
		return "", fmt.Errorf("failed to fetch user by email: %s", larkcore.Prettify(resp))
	}

	userinfo := resp.Data.UserList[len(resp.Data.UserList)-1]
	if userinfo.UserId == nil {
		return "", fmt.Errorf("users_not_found")
	}
	return *userinfo.UserId, nil
}
