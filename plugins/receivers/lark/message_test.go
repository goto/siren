package lark

import (
	"testing"
)

func TestMessage_BuildLarkMessage(t *testing.T) {
	tests := []struct {
		name    string
		message Message
		wantErr bool
	}{
		{
			name: "should build all message options if all fields in message present",
			message: Message{
				Channel:   "channel", // won't be included
				Color:     "blue",
				Text:      "text",
				Username:  "username",
				IconEmoji: ":emoji:",
				IconURL:   "icon_url",
				LinkNames: true, // won't be included
				Elements: []CardElement{
					{
						Title:   "title",
						Pretext: "pretext",
						Text:    "text",
						Actions: []Action{
							{Tag: "tag", URL: "url", Type: "type", Text: Text{Tag: "tag", Content: "content"}},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.message.BuildLarkMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.BuildLarkMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
