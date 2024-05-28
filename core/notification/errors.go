package notification

import (
	"errors"
	"fmt"
)

var ErrNoMessage = errors.New("no message found")

type DispatchError struct {
	Status string
	Err    error
}

func (de *DispatchError) Error() string {
	return fmt.Sprintf("status %s: err %v", de.Status, de.Err)
}
