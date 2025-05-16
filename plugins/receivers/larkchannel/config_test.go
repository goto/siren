package larkchannel_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/goto/siren/pkg/secret"
	"github.com/goto/siren/plugins/receivers/lark"
	"github.com/goto/siren/plugins/receivers/larkchannel"
)

func TestReceiverConfig(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		testCases := []struct {
			name    string
			c       larkchannel.ReceiverConfig
			wantErr bool
		}{
			{
				name:    "return error if one of required field is missing",
				wantErr: true,
			},
			{
				name: "return nil if all required fields are present",
				c: larkchannel.ReceiverConfig{
					ChannelName: "a-channel",
				},
				wantErr: false,
			},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if err := tc.c.Validate(); (err != nil) != tc.wantErr {
					t.Errorf("ReceiverConfig.Validate() error = %v, wantErr %v", err, tc.wantErr)
				}
			})
		}
	})
}

func TestNotificationConfig(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		testCases := []struct {
			name    string
			c       larkchannel.NotificationConfig
			wantErr bool
		}{
			{
				name:    "return error if one of required field is missing",
				wantErr: true,
			},
			{
				name: "return nil if all required fields are present",
				c: larkchannel.NotificationConfig{
					ReceiverConfig: larkchannel.ReceiverConfig{
						LarkReceiverConfig: lark.ReceiverConfig{
							ClientID:     "foo",
							ClientSecret: "foo",
						},
						ChannelName: "a-channel",
					},
				},
				wantErr: false,
			},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if err := tc.c.Validate(); (err != nil) != tc.wantErr {
					t.Errorf("NotificationConfig.Validate() error = %v, wantErr %v", err, tc.wantErr)
				}
			})
		}
	})

	t.Run("AsMap", func(t *testing.T) {
		nc := larkchannel.NotificationConfig{
			ReceiverConfig: larkchannel.ReceiverConfig{
				LarkReceiverConfig: lark.ReceiverConfig{
					ClientID:      secret.MaskableString("foo"),
					ClientSecret:  "foo",
					ValidDuration: time.Duration(0),
				},
				ChannelName: "channel",
			},
		}

		if diff := cmp.Diff(map[string]any{
			"channel_name":  "channel",
			"channel_type":  "",
			"client_id":     secret.MaskableString("foo"),
			"client_secret": secret.MaskableString("foo"),
		}, nc.AsMap()); diff != "" {
			t.Errorf("result not match\n%v", diff)
		}
	})
}
