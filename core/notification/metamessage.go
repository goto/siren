package notification

import "github.com/goto/siren/core/subscription"

type MetaMessage struct {
	ReceiverID      uint64
	SubscriptionIDs []uint64
	ReceiverType    string
	Notification    Notification
	ReceiverConfigs map[string]any
	MergedLabels    map[string][]string
}

func buildMetaMessage(n Notification, rcvView subscription.ReceiverView) MetaMessage {
	return MetaMessage{
		ReceiverID:      rcvView.ID,
		SubscriptionIDs: []uint64{rcvView.SubscriptionID},
		ReceiverType:    rcvView.Type,
		Notification:    n,
		ReceiverConfigs: rcvView.Configurations,
	}
}
