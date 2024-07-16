package notification

import (
	"context"

	"github.com/goto/siren/core/log"
)

type RouterSubscriberService struct {
	deps Deps
}

func NewRouterSubscriberService(
	deps Deps,
) *RouterSubscriberService {
	return &RouterSubscriberService{
		deps: deps,
	}
}

func (s *RouterSubscriberService) PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
	receiversView, err := s.deps.SubscriptionService.MatchByLabelsV2(ctx, n.NamespaceID, n.Labels)
	if err != nil {
		return nil, nil, err
	}

	if len(receiversView) == 0 {
		return nil, nil, ErrRouteSubscriberNoMatchFound
	}

	for _, rcv := range receiversView {
		metaMessages = append(metaMessages, n.MetaMessage(rcv))

		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			SubscriptionID: rcv.SubscriptionID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	return metaMessages, notificationLogs, nil
}
