package v1_test

import (
	"context"
	"testing"
	"time"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/api/mocks"
	v1 "github.com/goto/siren/internal/api/v1"
	"github.com/goto/siren/pkg/errors"
	sirenv1 "github.com/goto/siren/proto/gotocompany/siren/v1"
	"github.com/stretchr/testify/assert"
)

var (
	configuration = map[string]any{
		"foo": "bar",
	}

	match = map[string]string{
		"foo": "bar",
	}

	subMetadata = map[string]any{
		"foo1_metadata": "bar1_metadata",
		"foo2_metadata": float64(1),
		"foo3_metadata": true,
	}

	creator = "user@gotocompany.com"
)

func TestGRPCServer_ListSubscriptions(t *testing.T) {
	t.Run("should return list of all subscriptions", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}

		dummyGRPCServer := v1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{SubscriptionService: mockedSubscriptionService})

		dummyResult := []subscription.Subscription{
			{
				ID:        1,
				URN:       "foo",
				Namespace: 1,
				Receivers: []subscription.Receiver{{ID: 1, Configuration: configuration}},
				Match:     match,
				Metadata:  subMetadata,
				CreatedBy: creator,
				UpdatedBy: creator,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		mockedSubscriptionService.EXPECT().ListV2(context.TODO(), subscription.Filter{}).Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListSubscriptions(context.TODO(), &sirenv1.ListSubscriptionsRequest{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetSubscriptions()))
		assert.Equal(t, uint64(1), res.GetSubscriptions()[0].GetId())
		assert.Equal(t, "bar", res.GetSubscriptions()[0].GetMatch()["foo"])
		assert.Equal(t, true, res.GetSubscriptions()[0].GetMetadata().AsMap()["foo3_metadata"])
		assert.Equal(t, creator, res.GetSubscriptions()[0].GetUpdatedBy())
	})

	t.Run("should return error Internal if getting subscriptions fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}

		dummyGRPCServer := v1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{SubscriptionService: mockedSubscriptionService})

		mockedSubscriptionService.EXPECT().ListV2(context.TODO(), subscription.Filter{}).Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.ListSubscriptions(context.TODO(), &sirenv1.ListSubscriptionsRequest{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}
