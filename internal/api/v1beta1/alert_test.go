package v1beta1_test

import (
	"context"
	"testing"
	"time"

	"github.com/odpf/salt/log"
	"github.com/odpf/siren/core/alert"
	"github.com/odpf/siren/internal/api"
	"github.com/odpf/siren/internal/api/mocks"
	"github.com/odpf/siren/internal/api/v1beta1"
	"github.com/odpf/siren/pkg/errors"
	sirenv1beta1 "github.com/odpf/siren/proto/odpf/siren/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGRPCServer_ListAlerts(t *testing.T) {
	t.Run("should return alert history objects", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		timenow := time.Now()
		dummyAlerts := []alert.Alert{{
			ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "bar", MetricValue: "30", Rule: "bar",
			TriggeredAt: timenow,
		}}
		mockedAlertService.EXPECT().List(mock.AnythingOfType("*context.emptyCtx"), alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    100,
			EndTime:      200,
		}).Return(dummyAlerts, nil).Once()
		dummyGRPCServer := v1beta1.NewGRPCServer(nil, nil, &api.Deps{AlertService: mockedAlertService})

		dummyReq := &sirenv1beta1.ListAlertsRequest{
			ResourceName: "foo",
			ProviderId:   1,
			StartTime:    100,
			EndTime:      200,
		}
		res, err := dummyGRPCServer.ListAlerts(context.Background(), dummyReq)
		assert.Equal(t, 1, len(res.GetAlerts()))
		assert.Equal(t, uint64(1), res.GetAlerts()[0].GetId())
		assert.Equal(t, "foo", res.GetAlerts()[0].GetResourceName())
		assert.Equal(t, "bar", res.GetAlerts()[0].GetMetricName())
		assert.Equal(t, "30", res.GetAlerts()[0].GetMetricValue())
		assert.Equal(t, "CRITICAL", res.GetAlerts()[0].GetSeverity())
		assert.Equal(t, "bar", res.GetAlerts()[0].GetRule())
		assert.Nil(t, err)
		mockedAlertService.AssertExpectations(t)
	})

	t.Run("should return error Internal if getting alert history failed", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		dummyGRPCServer := v1beta1.NewGRPCServer(nil, log.NewNoop(), &api.Deps{AlertService: mockedAlertService})

		mockedAlertService.EXPECT().List(mock.AnythingOfType("*context.emptyCtx"), alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    100,
			EndTime:      200,
		}).Return(nil, errors.New("random error")).Once()

		dummyReq := &sirenv1beta1.ListAlertsRequest{
			ResourceName: "foo",
			ProviderId:   1,
			StartTime:    100,
			EndTime:      200,
		}
		res, err := dummyGRPCServer.ListAlerts(context.Background(), dummyReq)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		assert.Nil(t, res)
		mockedAlertService.AssertExpectations(t)
	})
}

func TestGRPCServer_CreateAlertHistory(t *testing.T) {
	timenow := timestamppb.New(time.Now())
	payload := []*alert.Alert{
		{
			ProviderID:   1,
			ResourceName: "foo",
			MetricName:   "bar",
			MetricValue:  "30",
			Severity:     "CRITICAL",
			Rule:         "random",
			TriggeredAt:  timenow.AsTime(),
		},
	}
	dummyReq := &sirenv1beta1.CreateCortexAlertsRequest{
		ProviderId: 1,
		Alerts: []*sirenv1beta1.CortexAlert{
			{
				Status: "foo",
				Labels: map[string]string{
					"severity": "CRITICAL",
				},
				Annotations: map[string]string{
					"resource":    "foo",
					"template":    "random",
					"metricName":  "bar",
					"metricValue": "30",
				},
				StartsAt: timenow,
			},
		},
	}

	t.Run("should create alerts objects", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		mockNotificationService := new(mocks.NotificationService)

		dummyAlerts := []alert.Alert{{
			ID:           1,
			ProviderID:   1,
			ResourceName: "foo",
			MetricName:   "bar",
			MetricValue:  "30",
			Severity:     "CRITICAL",
			Rule:         "random",
			TriggeredAt:  timenow.AsTime(),
		}}
		mockedAlertService.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), payload).
			Return(dummyAlerts, nil).Once()
		mockNotificationService.EXPECT().DispatchToSubscribers(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(nil)

		dummyGRPCServer := v1beta1.NewGRPCServer(nil, log.NewNoop(), &api.Deps{AlertService: mockedAlertService, NotificationService: mockNotificationService})

		res, err := dummyGRPCServer.CreateCortexAlerts(context.Background(), dummyReq)
		assert.Equal(t, 1, len(res.GetAlerts()))
		assert.Equal(t, uint64(1), res.GetAlerts()[0].GetId())
		assert.Equal(t, "foo", res.GetAlerts()[0].GetResourceName())
		assert.Equal(t, "random", res.GetAlerts()[0].GetRule())
		assert.Equal(t, "bar", res.GetAlerts()[0].GetMetricName())
		assert.Equal(t, "30", res.GetAlerts()[0].GetMetricValue())
		assert.Equal(t, "CRITICAL", res.GetAlerts()[0].GetSeverity())
		assert.Nil(t, err)
		mockedAlertService.AssertExpectations(t)
	})

	t.Run("should create alerts for resolved alerts", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		mockNotificationService := new(mocks.NotificationService)

		dummyReq := &sirenv1beta1.CreateCortexAlertsRequest{
			ProviderId: 1,
			Alerts: []*sirenv1beta1.CortexAlert{
				{
					Status: "resolved",
					Labels: map[string]string{
						"severity": "CRITICAL",
					},
					Annotations: map[string]string{
						"resource":    "foo",
						"template":    "random",
						"metricName":  "bar",
						"metricValue": "30",
					},
					StartsAt: timenow,
				},
			},
		}
		payload := []*alert.Alert{
			{
				ProviderID:   1,
				ResourceName: "foo",
				MetricName:   "bar",
				MetricValue:  "30",
				Severity:     "resolved",
				Rule:         "random",
				TriggeredAt:  timenow.AsTime(),
			},
		}
		dummyAlerts := []alert.Alert{{
			ID:           1,
			ProviderID:   1,
			ResourceName: "foo",
			MetricName:   "bar",
			MetricValue:  "30",
			Severity:     "resolved",
			Rule:         "random",
			TriggeredAt:  timenow.AsTime(),
		}}
		mockedAlertService.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), payload).
			Return(dummyAlerts, nil).Once()
		mockNotificationService.EXPECT().DispatchToSubscribers(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(nil)

		dummyGRPCServer := v1beta1.NewGRPCServer(nil, log.NewNoop(), &api.Deps{AlertService: mockedAlertService, NotificationService: mockNotificationService})

		res, err := dummyGRPCServer.CreateCortexAlerts(context.Background(), dummyReq)
		assert.Equal(t, 1, len(res.GetAlerts()))
		assert.Equal(t, uint64(1), res.GetAlerts()[0].GetId())
		assert.Equal(t, "foo", res.GetAlerts()[0].GetResourceName())
		assert.Equal(t, "random", res.GetAlerts()[0].GetRule())
		assert.Equal(t, "bar", res.GetAlerts()[0].GetMetricName())
		assert.Equal(t, "30", res.GetAlerts()[0].GetMetricValue())
		assert.Equal(t, "resolved", res.GetAlerts()[0].GetSeverity())
		assert.Nil(t, err)
		mockedAlertService.AssertExpectations(t)
	})

	t.Run("should return error Internal if getting alert history failed", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		dummyGRPCServer := v1beta1.NewGRPCServer(nil, log.NewNoop(), &api.Deps{AlertService: mockedAlertService})

		mockedAlertService.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), payload).
			Return(nil, errors.New("random error")).Once()

		res, err := dummyGRPCServer.CreateCortexAlerts(context.Background(), dummyReq)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		assert.Nil(t, res)
		mockedAlertService.AssertExpectations(t)
	})

	t.Run("should insert valid alerts and should not return error if parameters are missing", func(t *testing.T) {
		mockedAlertService := &mocks.AlertService{}
		mockNotificationService := new(mocks.NotificationService)

		dummyReq := &sirenv1beta1.CreateCortexAlertsRequest{
			ProviderId: 1,
			Alerts: []*sirenv1beta1.CortexAlert{
				{
					Status: "foo",
					Labels: map[string]string{
						"severity": "CRITICAL",
					},
					Annotations: map[string]string{
						"resource":    "foo",
						"metricName":  "bar",
						"metricValue": "30",
					},
					StartsAt: timenow,
				},
				{
					Status: "foo",
					Labels: map[string]string{
						"severity": "CRITICAL",
					},
					Annotations: map[string]string{
						"resource":    "foo",
						"template":    "random",
						"metricName":  "bar",
						"metricValue": "30",
					},
					StartsAt: timenow,
				},
			},
		}
		dummyAlerts := []alert.Alert{{
			ProviderID:   1,
			ResourceName: "foo",
			MetricName:   "bar",
			MetricValue:  "30",
			Rule:         "random",
			Severity:     "CRITICAL",
			TriggeredAt:  time.Now(),
		}}

		dummyGRPCServer := v1beta1.NewGRPCServer(nil, log.NewNoop(), &api.Deps{AlertService: mockedAlertService, NotificationService: mockNotificationService})

		mockedAlertService.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), payload).
			Return(dummyAlerts, nil).Once()
		mockNotificationService.EXPECT().DispatchToSubscribers(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(nil)

		res, err := dummyGRPCServer.CreateCortexAlerts(context.Background(), dummyReq)
		assert.Equal(t, 1, len(res.GetAlerts()))
		assert.Equal(t, uint64(1), res.GetAlerts()[0].GetProviderId())
		assert.Equal(t, "foo", res.GetAlerts()[0].GetResourceName())
		assert.Equal(t, "random", res.GetAlerts()[0].GetRule())
		assert.Equal(t, "bar", res.GetAlerts()[0].GetMetricName())
		assert.Equal(t, "30", res.GetAlerts()[0].GetMetricValue())
		assert.Equal(t, "CRITICAL", res.GetAlerts()[0].GetSeverity())
		assert.Nil(t, err)
		mockedAlertService.AssertExpectations(t)
	})
}