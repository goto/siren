package notification

import (
	"time"

	"github.com/odpf/siren/pkg/errors"
)

// Notification is a model of notification
type Notification struct {
	ID                  string                 `json:"id"`
	Data                map[string]interface{} `json:"data"`
	Labels              map[string]string      `json:"labels"`
	ValidDurationString string                 `json:"valid_duration"`
	Template            string                 `json:"template"`
	CreatedAt           time.Time
}

// ToMessage transforms Notification model to one or several Messages
func (n Notification) ToMessage(receiverType string, notificationConfigMap map[string]interface{}) (*Message, error) {
	var (
		expiryDuration time.Duration
		err            error
	)

	if n.ValidDurationString != "" {
		expiryDuration, err = time.ParseDuration(n.ValidDurationString)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf(err.Error())
		}
	}

	nm := &Message{}
	nm.Initialize(
		n,
		receiverType,
		notificationConfigMap,
		InitWithExpiryDuration(expiryDuration),
	)

	return nm, nil
}