package lark

import (
	"context"

	"github.com/goto/siren/pkg/secret"
)

type Encryptor interface {
	Encrypt(str secret.MaskableString) (secret.MaskableString, error)
	Decrypt(str secret.MaskableString) (secret.MaskableString, error)
}

type LarkCaller interface {
	GetWorkspaceChannels(ctx context.Context, clientID, clientSecret secret.MaskableString) ([]Channel, error)
	Notify(ctx context.Context, conf NotificationConfig, message Message) error
}
