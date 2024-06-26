package subscription

import (
	"context"

	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/pkg/errors"
)

type LogService interface {
	ListSubscriptionIDsBySilenceID(ctx context.Context, silenceID string) ([]int64, error)
}

type NamespaceService interface {
	List(context.Context) ([]namespace.Namespace, error)
	Create(context.Context, *namespace.Namespace) error
	Get(context.Context, uint64) (*namespace.Namespace, error)
	Update(context.Context, *namespace.Namespace) error
	Delete(context.Context, uint64) error
}

type ReceiverService interface {
	List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error)
	PostHookDBTransformConfigs(ctx context.Context, receiverType string, configs map[string]any) (map[string]any, error)
}

// Service handles business logic
type Service struct {
	repository                  Repository
	logService                  LogService
	namespaceService            NamespaceService
	receiverService             ReceiverService
	subscriptionReceiverService SubscriptionReceiverService
}

// NewService returns service struct
func NewService(
	repository Repository,
	logService LogService,
	namespaceService NamespaceService,
	receiverService ReceiverService,
	subscriptionReceiverService SubscriptionReceiverService,
) *Service {
	svc := &Service{
		repository:                  repository,
		logService:                  logService,
		namespaceService:            namespaceService,
		receiverService:             receiverService,
		subscriptionReceiverService: subscriptionReceiverService,
	}

	return svc
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Subscription, error) {

	if flt.SilenceID != "" {
		subscriptionIDs, err := s.logService.ListSubscriptionIDsBySilenceID(ctx, flt.SilenceID)
		if err != nil {
			return nil, err
		}
		flt.IDs = subscriptionIDs
	}

	subscriptions, err := s.repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (s *Service) Create(ctx context.Context, sub *Subscription) error {
	if err := s.repository.Create(ctx, sub); err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.Is(err, ErrRelation) {
			return errors.ErrNotFound.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Get(ctx context.Context, id uint64) (*Subscription, error) {
	subscription, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.As(err, new(NotFoundError)) {
			return nil, errors.ErrNotFound.WithMsgf(err.Error())
		}
		return nil, err
	}

	return subscription, nil
}

func (s *Service) Update(ctx context.Context, sub *Subscription) error {
	if err := s.repository.Update(ctx, sub); err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.Is(err, ErrRelation) {
			return errors.ErrNotFound.WithMsgf(err.Error())
		}
		if errors.As(err, new(NotFoundError)) {
			return errors.ErrNotFound.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id uint64) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *Service) MatchByLabels(ctx context.Context, namespaceID uint64, notificationLabels map[string]string) ([]Subscription, error) {
	// fetch all subscriptions by matching labels.
	subscriptionsByLabels, err := s.repository.List(ctx, Filter{
		NamespaceID:       namespaceID,
		NotificationMatch: notificationLabels,
	})
	if err != nil {
		return nil, err
	}

	if len(subscriptionsByLabels) == 0 {
		return nil, nil
	}

	receiversMap, err := CreateReceiversMap(ctx, s.receiverService, subscriptionsByLabels)
	if err != nil {
		return nil, err
	}

	subscriptionsByLabels, err = AssignReceivers(receiversMap, subscriptionsByLabels)
	if err != nil {
		return nil, err
	}

	return subscriptionsByLabels, nil
}

func CreateReceiversMap(ctx context.Context, receiverService ReceiverService, subscriptions []Subscription) (map[uint64]*receiver.Receiver, error) {
	receiversMap := map[uint64]*receiver.Receiver{}
	for _, subs := range subscriptions {
		for _, rcv := range subs.Receivers {
			if rcv.ID != 0 {
				receiversMap[rcv.ID] = nil
			}
		}
	}

	// empty receivers map
	if len(receiversMap) == 0 {
		return nil, errors.New("no receivers found in subscription")
	}

	listOfReceiverIDs := []uint64{}
	for k := range receiversMap {
		listOfReceiverIDs = append(listOfReceiverIDs, k)
	}

	filteredReceivers, err := receiverService.List(ctx, receiver.Filter{
		ReceiverIDs: listOfReceiverIDs,
		Expanded:    true,
	})
	if err != nil {
		return nil, err
	}

	for i, rcv := range filteredReceivers {
		receiversMap[rcv.ID] = &filteredReceivers[i]
	}

	nilReceivers := []uint64{}
	for id, rcv := range receiversMap {
		if rcv == nil {
			nilReceivers = append(nilReceivers, id)
			continue
		}
	}

	if len(nilReceivers) > 0 {
		return nil, errors.ErrInvalid.WithMsgf("receiver id %v don't exist", nilReceivers)
	}

	return receiversMap, nil
}

func AssignReceivers(receiversMap map[uint64]*receiver.Receiver, subscriptions []Subscription) ([]Subscription, error) {
	for is := range subscriptions {
		for ir, subsRcv := range subscriptions[is].Receivers {
			if mappedRcv := receiversMap[subsRcv.ID]; mappedRcv == nil {
				return nil, errors.ErrInvalid.WithMsgf("receiver id %d not found", subsRcv.ID)
			}
			mergedConfigMap := MergeConfigsMap(subsRcv.Configuration, receiversMap[subsRcv.ID].Configurations)
			subscriptions[is].Receivers[ir].ID = receiversMap[subsRcv.ID].ID
			subscriptions[is].Receivers[ir].Type = receiversMap[subsRcv.ID].Type
			subscriptions[is].Receivers[ir].Configuration = mergedConfigMap
		}
	}

	return subscriptions, nil
}

func MergeConfigsMap(subscriptionConfigMap map[string]any, receiverConfigsMap map[string]any) map[string]any {
	var newConfigMap = make(map[string]any)
	for k, v := range subscriptionConfigMap {
		newConfigMap[k] = v
	}
	for k, v := range receiverConfigsMap {
		newConfigMap[k] = v
	}
	return newConfigMap
}
