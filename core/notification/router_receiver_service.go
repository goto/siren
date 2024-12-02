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

	selectors, selectorConfig, err := n.ReceiverSelectors.parseAndValidate()
	if err != nil {
		return nil, nil, err
	}

	rcvs, err := s.deps.ReceiverService.List(ctx, receiver.Filter{
		MultipleLabels: selectors,
		Expanded:       true,
	})
	if err != nil {
		return nil, nil, err
	}

	if len(rcvs) == 0 {
		return nil, nil, errors.ErrNotFound
	}

	if selectorConfig != nil && len(rcvs) > 1 {
		return nil, nil, errors.ErrInvalid.WithMsgf("config override could only be used to 1 receiver, but got %d receiver", len(rcvs))
	} else if selectorConfig != nil && len(rcvs) == 1 {
		// config override flow
		var rcvView = &subscription.ReceiverView{}
		rcvView.FromReceiver(rcvs[0])
		rcvView.Configurations = s.mergeReceiverConfig(rcvView.Configurations, selectorConfig)
		metaMessages = append(metaMessages, n.MetaMessage(*rcvView))

		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			ReceiverID:     rcvs[0].ID,
			AlertIDs:       n.AlertIDs,
		})
		return metaMessages, notificationLogs, nil
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

func (s *RouterReceiverService) mergeReceiverConfig(receiverConfig, selectorConfig map[string]any) map[string]any {
	// override the existing config with the one from API if there is config clash
	result := map[string]any{}
	for k, v := range receiverConfig {
		result[k] = v
	}
	for k, v := range selectorConfig {
		result[k] = v
	}
	return result
}
