package lark_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goto/salt/log"
	"github.com/goto/siren/pkg/secret"
	"github.com/goto/siren/plugins/receivers/lark"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetWorkspaceChannels(t *testing.T) {
	var clientId = secret.MaskableString("test-id")
	var clientSecret = secret.MaskableString("test-secret")

	t.Run("return error when failed to fetch joined channel list", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)

		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		channels, err := c.GetWorkspaceChannels(context.Background(), clientId, clientSecret)

		assert.EqualError(t, err, "failed to fetch joined channel list: unexpected end of JSON input")
		assert.Empty(t, channels)

		testServer.Close()
	})
}

func TestClient_NotifyChannel(t *testing.T) {
	var clientId = secret.MaskableString("test-id")
	var clientSecret = secret.MaskableString("test-secret")

	t.Run("return error when message receiver type is wrong", func(t *testing.T) {
		c := lark.NewClient(lark.AppConfig{}, log.NewNoop())
		err := c.Notify(context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
			},
			lark.Message{})

		assert.EqualError(t, err, "unknown receiver type \"\"")
	})

	t.Run("return error when failed to fetch joined channel list", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		err := c.Notify(
			context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
				SubscriptionConfig: lark.SubscriptionConfig{
					ChannelType: lark.TypeChannelChannel,
				},
			},
			lark.Message{})

		assert.EqualError(t, err, "unexpected end of JSON input")

		testServer.Close()
	})

	t.Run("return error when app is not part of the channel", func(t *testing.T) {
		var chatId = "123"
		var channelName = "test"
		var pageToken = ""
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "//open-apis/im/v1/chats" {
				data := larkim.ListChatRespData{Items: []*larkim.ListChat{{ChatId: &chatId, Name: &channelName}}, PageToken: &pageToken}
				respStruct := larkim.ListChatResp{
					Data: &data,
				}

				respByte, _ := json.Marshal(respStruct)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respByte)
				return
			}
			if r.URL.Path == "//open-apis/auth/v3/tenant_access_token/internal" {
				respByte := []byte(`{
					"code":                0,
					"expire":              7158,
					"msg":                 "ok",
					"tenant_access_token": "t-g1111c9DMTQF65O4AQAWCBBE3TRRCLAVU4B5MY27"
				}`)

				w.Write(respByte)
				return
			} else {
				w.Write([]byte(`{"ok":true}`))
			}
		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		err := c.Notify(
			context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
				SubscriptionConfig: lark.SubscriptionConfig{
					ChannelType: lark.TypeChannelChannel,
				},
			},
			lark.Message{})

		assert.EqualError(t, err, "app is not part of the channel \"\"")

		testServer.Close()
	})

	t.Run("return nil error when notify is succeed through channel", func(t *testing.T) {
		var chatId = "123"
		var channelName = "test"
		var pageToken = ""
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "//open-apis/im/v1/chats" {
				data := larkim.ListChatRespData{Items: []*larkim.ListChat{{ChatId: &chatId, Name: &channelName}}, PageToken: &pageToken}
				respStruct := larkim.ListChatResp{
					Data: &data,
				}

				respByte, _ := json.Marshal(respStruct)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respByte)
				return
			}
			if r.URL.Path == "//open-apis/auth/v3/tenant_access_token/internal" {
				respByte := []byte(`{
					"code":                0,
					"expire":              7158,
					"msg":                 "ok",
					"tenant_access_token": "t-g1111c9DMTQF65O4AQAWCBBE3TRRCLAVU4B5MY27"
				}`)

				w.Write(respByte)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"ok":true}`))
			}
		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		err := c.Notify(
			context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
				SubscriptionConfig: lark.SubscriptionConfig{
					ChannelType: lark.TypeChannelChannel,
				},
			},
			lark.Message{
				Channel: "test",
			})

		assert.NoError(t, err)

		testServer.Close()
	})
}

func TestClient_NotifyUser(t *testing.T) {
	var clientId = secret.MaskableString("test-id")
	var clientSecret = secret.MaskableString("test-secret")

	t.Run("return error when failed to get user for an email", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "//open-apis/contact/v3/users/batch_get_id" {
				data := larkcontact.BatchGetIdUserRespData{UserList: []*larkcontact.UserContactInfo{{Email: nil}}}
				respStruct := larkcontact.BatchGetIdUserResp{
					Data: &data,
				}

				respByte, _ := json.Marshal(respStruct)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respByte)
				return
			}
			if r.URL.Path == "//open-apis/auth/v3/tenant_access_token/internal" {
				respByte := []byte(`{
					"code":                0,
					"expire":              7158,
					"msg":                 "ok",
					"tenant_access_token": "t-g1111c9DMTQF65O4AQAWCBBE3TRRCLAVU4B5MY27"
				}`)

				w.Write(respByte)
				return
			} else {
				w.Write([]byte(`{"ok":true}`))
			}
		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		err := c.Notify(
			context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
				SubscriptionConfig: lark.SubscriptionConfig{
					ChannelType: lark.TypeChannelUser,
				},
			},
			lark.Message{})

		assert.EqualError(t, err, "failed to get id for \"\"")

		testServer.Close()
	})

	t.Run("return nil error when notify is succeed through user", func(t *testing.T) {
		var userid = "123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "//open-apis/contact/v3/users/batch_get_id" {
				data := larkcontact.BatchGetIdUserRespData{UserList: []*larkcontact.UserContactInfo{{UserId: &userid}}}
				respStruct := larkcontact.BatchGetIdUserResp{
					Data: &data,
				}

				respByte, _ := json.Marshal(respStruct)
				w.Header().Set("Content-Type", "application/json")
				w.Write(respByte)
				return
			}
			if r.URL.Path == "//open-apis/auth/v3/tenant_access_token/internal" {
				respByte := []byte(`{
					"code":                0,
					"expire":              7158,
					"msg":                 "ok",
					"tenant_access_token": "t-g1111c9DMTQF65O4AQAWCBBE3TRRCLAVU4B5MY27"
				}`)

				w.Write(respByte)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"ok":true}`))
			}
		}))

		c := lark.NewClient(lark.AppConfig{APIHost: testServer.URL}, log.NewNoop())
		err := c.Notify(
			context.Background(),
			lark.NotificationConfig{
				ReceiverConfig: lark.ReceiverConfig{
					ClientID:     clientId,
					ClientSecret: clientSecret,
				},
				SubscriptionConfig: lark.SubscriptionConfig{
					ChannelType: lark.TypeChannelUser,
					ChannelName: "email@email.com",
				},
			},
			lark.Message{})

		assert.NoError(t, err)

		testServer.Close()
	})
}
