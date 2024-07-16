package notification

import (
	"context"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/pkg/errors"
)

type RouterReceiverService struct {
	deps Deps
}

func NewRouterReceiverService(
	deps Deps,
) *RouterReceiverService {
	return &RouterReceiverService{
		deps: deps,
	}
}

func (s *RouterReceiverService) PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
	if len(n.ReceiverSelectors) > s.deps.Cfg.MaxNumReceiverSelectors {
		return nil, nil, errors.ErrInvalid.WithMsgf("number of receiver selectors should be less than or equal threshold %d", s.deps.Cfg.MaxNumReceiverSelectors)
	}

	rcvs, err := s.deps.ReceiverService.List(ctx, receiver.Filter{
		MultipleLabels: n.ReceiverSelectors,
		Expanded:       true,
	})
	if err != nil {
		return nil, nil, err
	}

	if len(rcvs) == 0 {
		return nil, nil, errors.ErrNotFound
	}

	for _, rcv := range rcvs {
		var rcvView = &subscription.ReceiverView{}
		rcvView.FromReceiver(rcv)
		metaMessages = append(metaMessages, n.MetaMessage(*rcvView))

		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	var metaMessagesNum = len(metaMessages)
	if metaMessagesNum > s.deps.Cfg.MaxMessagesReceiverFlow {
		return nil, nil, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages receiver flow threshold %d. this will spam and broadcast to %d channel. found %d receiver selectors passed, you might want to check your receiver selectors configuration", metaMessagesNum, s.deps.Cfg.MaxMessagesReceiverFlow, metaMessagesNum, len(n.ReceiverSelectors))
	}

	return metaMessages, notificationLogs, nil
}
