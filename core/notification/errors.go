package notification

import (
	"errors"
)

var (
	ErrNoMessage                   = errors.New("no message sent, probably because not matching any subscription or receiver")
	ErrRouteSubscriberNoMatchFound = errors.New("not matching any subscription")
)
