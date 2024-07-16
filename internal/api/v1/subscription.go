package v1

import (
	"context"

	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/internal/api"
	sirenv1 "github.com/goto/siren/proto/gotocompany/siren/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GRPCServer) ListSubscriptions(ctx context.Context, req *sirenv1.ListSubscriptionsRequest) (*sirenv1.ListSubscriptionsResponse, error) {
	// NOTE: only string value that is queryable with this approach
	var metadataQuery = map[string]any{}
	if len(req.GetMetadata()) == 0 {
		metadataQuery = nil
	} else {
		for k, v := range req.GetMetadata() {
			metadataQuery[k] = v
		}
	}

	var (
		subscriptions []subscription.Subscription
		err           error
	)

	subscriptions, err = s.subscriptionService.ListV2(ctx, subscription.Filter{
		NamespaceID:                req.GetNamespaceId(),
		SilenceID:                  req.GetSilenceId(),
		Metadata:                   metadataQuery,
		Match:                      req.GetMatch(),
		NotificationMatch:          req.GetNotificationMatch(),
		ReceiverID:                 req.GetReceiverId(),
		SubscriptionReceiverLabels: req.GetSubscriptionReceiverLabels(),
	})
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1.Subscription{}

	for _, sub := range subscriptions {

		subscriptionReceiverRelationsPB := make([]*sirenv1.SubscriptionReceiverRelation, 0)
		for _, item := range sub.ReceiversRelation {
			subscriptionReceiverRelationsPB = append(subscriptionReceiverRelationsPB, &sirenv1.SubscriptionReceiverRelation{
				ReceiverId: item.ReceiverID,
				Labels:     item.Labels,
				CreatedAt:  timestamppb.New(item.CreatedAt),
				UpdatedAt:  timestamppb.New(item.UpdatedAt),
			})
		}

		metadata, err := structpb.NewStruct(sub.Metadata)
		if err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}

		item := &sirenv1.Subscription{
			Id:                sub.ID,
			Urn:               sub.URN,
			Namespace:         sub.Namespace,
			Match:             sub.Match,
			ReceiversRelation: subscriptionReceiverRelationsPB,
			Metadata:          metadata,
			CreatedBy:         sub.CreatedBy,
			UpdatedBy:         sub.UpdatedBy,
			CreatedAt:         timestamppb.New(sub.CreatedAt),
			UpdatedAt:         timestamppb.New(sub.UpdatedAt),
		}
		items = append(items, item)
	}
	return &sirenv1.ListSubscriptionsResponse{
		Subscriptions: items,
	}, nil
}
