package notification

import (
	"context"
	"fmt"
	"strconv"

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

	var rcvs []receiver.Receiver
	var userConfigs map[uint64]map[string]interface{}

	// Check if any selector contains a config
	hasConfig := false
	for _, selector := range n.ReceiverSelectors {
		if _, ok := selector["config"]; ok {
			hasConfig = true
			break
		}
	}

	if hasConfig {
		// Handle case when config is provided
		rcvs, userConfigs, err = s.handleConfigCase(ctx, n.ReceiverSelectors)
	} else {
		// Handle case when only receiver IDs are provided
		rcvs, err = s.handleNonConfigCase(ctx, n.ReceiverSelectors)
	}

	if err != nil {
		return nil, nil, err
	}

	if len(rcvs) == 0 {
		return nil, nil, errors.ErrNotFound
	}

	// Check if the number of receivers exceeds the max messages receiver flow
	if len(rcvs) > s.deps.Cfg.MaxMessagesReceiverFlow {
		return nil, nil, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages receiver flow threshold %d. this will spam and broadcast to %d channel. found %d receiver selectors passed, you might want to check your receiver selectors configuration", len(rcvs), s.deps.Cfg.MaxMessagesReceiverFlow, len(rcvs), len(n.ReceiverSelectors))
	}

	for _, rcv := range rcvs {
		rcvView := &subscription.ReceiverView{}
		rcvView.FromReceiver(rcv)

		if config, ok := userConfigs[rcv.ID]; ok {
			// Merge user-provided config with receiver config
			for k, v := range config {
				rcvView.Configurations[k] = v
			}
		}

		metaMessage := n.MetaMessage(*rcvView)
		metaMessage.NotificationIDs = []string{n.ID}

		metaMessages = append(metaMessages, metaMessage)
		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	return metaMessages, notificationLogs, nil
}

func (s *RouterReceiverService) handleConfigCase(ctx context.Context, selectors []map[string]any) ([]receiver.Receiver, map[uint64]map[string]interface{}, error) {
	var receiverIDs []uint64
	userConfigs := make(map[uint64]map[string]interface{})

	for _, selector := range selectors {
		if idStr, ok := selector["id"].(string); ok {
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				return nil, nil, errors.ErrInvalid.WithMsgf("invalid receiver id: %s", idStr)
			}
			receiverIDs = append(receiverIDs, id)
		}
		if config, ok := selector["config"].(map[string]interface{}); ok {
			for _, id := range receiverIDs {
				if userConfigs[id] == nil {
					userConfigs[id] = make(map[string]interface{})
				}
				for k, v := range config {
					userConfigs[id][k] = v
				}
			}
		}
	}

	rcvs := make([]receiver.Receiver, 0, len(receiverIDs))
	for _, id := range receiverIDs {
		rcv, err := s.deps.ReceiverService.Get(ctx, id)
		if err != nil {
			return nil, nil, err
		}
		rcvs = append(rcvs, *rcv)
	}

	return rcvs, userConfigs, nil
}

func (s *RouterReceiverService) handleNonConfigCase(ctx context.Context, selectors []map[string]any) ([]receiver.Receiver, error) {
	convertedSelectors := make([]map[string]string, len(selectors))
	for i, selector := range selectors {
		convertedSelectors[i] = make(map[string]string)
		for k, v := range selector {
			convertedSelectors[i][k] = fmt.Sprint(v)
		}
	}
	return s.deps.ReceiverService.List(ctx, receiver.Filter{
		MultipleLabels: convertedSelectors,
		Expanded:       true,
	})
}
