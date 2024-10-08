package lark

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Color     string        `yaml:"templatecolor,omitempty" json:"templatecolor,omitempty"  mapstructure:"templatecolor"`
	Channel   string        `yaml:"channel,omitempty" json:"channel,omitempty"  mapstructure:"channel"`
	Text      string        `yaml:"text,omitempty" json:"text,omitempty"  mapstructure:"text"`
	Username  string        `yaml:"username,omitempty" json:"username,omitempty"  mapstructure:"username"`
	IconEmoji string        `yaml:"icon_emoji,omitempty" json:"icon_emoji,omitempty" mapstructure:"icon_emoji"`
	IconURL   string        `yaml:"icon_url,omitempty" json:"icon_url,omitempty"  mapstructure:"icon_url"`
	LinkNames bool          `yaml:"link_names,omitempty" json:"link_names,omitempty"  mapstructure:"link_names"`
	Elements  []CardElement `yaml:"elements,omitempty" json:"attachments,omitempty" mapstructure:"elements"`
}

func (m Message) BuildLarkMessage() (string, error) {
	message := MessageCard{
		Config:   Config{WideScreenMode: true},
		Elements: []Element{},
		Header:   Header{},
	}
	message.Header.Template = m.Color
	message.Header.Title = Text{Content: m.Text, Tag: "lark_md"}
	message.Elements = append(message.Elements, Element{Tag: "div", Text: Text{Content: m.IconEmoji, Tag: "lark_md"}})

	for _, a := range m.Elements {
		element, err := a.ToLark()
		if err != nil {
			return "", fmt.Errorf("failed to parse lark element: %w", err)
		}
		message.Elements = append(message.Elements, Element{Tag: "div", Text: Text{Content: element.Pretext, Tag: "lark_md"}})
		message.Elements = append(message.Elements, Element{Tag: "div", Text: Text{Content: element.Text, Tag: "lark_md"}})

		for _, action := range element.Actions {
			message.Elements = append(message.Elements, Element{Tag: "action", Actions: []Action{{Tag: action.Tag, URL: action.URL, Type: action.Type, Text: Text{Tag: action.Text.Tag, Content: action.Text.Content}}}}) // actions
		}
	}
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonData), nil
}

type CardElement struct {
	Title   string   `json:"title,omitempty"`
	Pretext string   `json:"pretext,omitempty"`
	Text    string   `json:"text,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

func (ma CardElement) ToLark() (*CardElement, error) {
	gaBlob, err := json.Marshal(ma)
	if err != nil {
		return nil, err
	}

	ga := &CardElement{}
	if err := json.Unmarshal(gaBlob, &ga); err != nil {
		return nil, err
	}

	return ga, nil
}

type MessageCard struct {
	Config   Config    `json:"config"`
	Elements []Element `json:"elements"`
	Header   Header    `json:"header"`
}

type Config struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type Element struct {
	Tag     string   `json:"tag"`
	Text    Text     `json:"text,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Text struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Action struct {
	Tag  string `json:"tag"`
	Text Text   `json:"text"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Header struct {
	Template string `json:"template"`
	Title    Text   `json:"title"`
}
