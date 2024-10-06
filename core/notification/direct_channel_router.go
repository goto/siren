package notification

import (
	"context"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/pkg/errors"
)

const ReceiverTypeDirect = "direct"

type ReceiverView struct {
	ID      int
	Type    string
	Channel string
}

type DirectChannelRouter struct {
	deps Deps
}

func NewDirectChannelRouter(deps Deps) *DirectChannelRouter {
	return &DirectChannelRouter{deps: deps}
}

func (r *DirectChannelRouter) PrepareMetaMessages(ctx context.Context, n Notification) ([]MetaMessage, []log.Notification, error) {
    if len(n.ReceiverSelectors) > r.deps.Cfg.MaxNumReceiverSelectors {
        return nil, nil, errors.ErrInvalid.WithMsgf("number of receiver selectors should be less than or equal threshold %d", r.deps.Cfg.MaxNumReceiverSelectors)
    }

    var metaMessages []MetaMessage
    var notificationLogs []log.Notification

    for _, selector := range n.ReceiverSelectors {
        var channel string
        var receiverType string

        if slackChannel, ok := selector["channel_name"].(string); ok {
            channel = slackChannel
            receiverType = "slack_channel"
        } else if directChannel, ok := selector["channel"].(string); ok {
            channel = directChannel
            receiverType = ReceiverTypeDirect
        } else {
            return nil, nil, errors.ErrInvalid.WithMsgf("missing or invalid channel in receiver selector")
        }

        rcvView := ReceiverView{
            ID:      0, // Using 0 as a placeholder for direct channels
            Type:    receiverType,
            Channel: channel,
        }

        metaMessages = append(metaMessages, n.MetaMessage(subscription.ReceiverView{
            ID:   uint64(rcvView.ID),
            Type: rcvView.Type,
        }))

        notificationLogs = append(notificationLogs, log.Notification{
            NamespaceID:    n.NamespaceID,
            NotificationID: n.ID,
            ReceiverID:     0,
            AlertIDs:       n.AlertIDs,
        })
    }

    if len(metaMessages) == 0 {
        return nil, nil, errors.ErrNotFound
    }

    if len(metaMessages) > r.deps.Cfg.MaxMessagesReceiverFlow {
        return nil, nil, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages direct channel flow threshold %d. this will spam and broadcast to %d channels. found %d receiver selectors passed, you might want to check your receiver selectors configuration", len(metaMessages), r.deps.Cfg.MaxMessagesReceiverFlow, len(metaMessages), len(n.ReceiverSelectors))
    }

    return metaMessages, notificationLogs, nil
}
