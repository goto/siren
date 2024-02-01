package alert_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/alert/mocks"
	"github.com/goto/siren/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_List(t *testing.T) {
	ctx := context.TODO()

	t.Run("should call repository List method with proper arguments and return result in domain's type", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := alert.NewService(repositoryMock, nil, nil)
		timenow := time.Now()
		dummyAlerts := []alert.Alert{
			{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "20",
				Rule: "bar", TriggeredAt: timenow},
			{ID: 2, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "0",
				Rule: "bar", TriggeredAt: timenow},
		}
		repositoryMock.EXPECT().List(mock.AnythingOfType("context.todoCtx"), alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    0,
			EndTime:      100,
		}).Return(dummyAlerts, nil)
		actualAlerts, err := dummyService.List(ctx, alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    0,
			EndTime:      100,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, actualAlerts)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository List method with proper arguments if endtime is zero", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := alert.NewService(repositoryMock, nil, nil)
		timenow := time.Now()
		dummyAlerts := []alert.Alert{
			{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "20",
				Rule: "bar", TriggeredAt: timenow},
			{ID: 2, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "0",
				Rule: "bar", TriggeredAt: timenow},
		}
		repositoryMock.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(dummyAlerts, nil)
		actualAlerts, err := dummyService.List(ctx, alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    0,
			EndTime:      0,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, actualAlerts)
		repositoryMock.AssertNotCalled(t, "Get", "foo", uint64(1), uint64(0))
	})

	t.Run("should call repository List method and handle errors", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := alert.NewService(repositoryMock, nil, nil)
		repositoryMock.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.Anything).
			Return(nil, errors.New("random error"))
		actualAlerts, err := dummyService.List(ctx, alert.Filter{
			ProviderID:   1,
			ResourceName: "foo",
			StartTime:    0,
			EndTime:      0,
		})
		assert.EqualError(t, err, "random error")
		assert.Nil(t, actualAlerts)
	})
}

func TestService_CreateAlerts(t *testing.T) {
	var (
		ctx               = context.TODO()
		timenow           = time.Now()
		testType          = "test"
		alertsToBeCreated = map[string]any{
			"alerts": []map[string]any{
				{
					"annotations": map[string]any{
						"metricName":  "bar",
						"metricValue": "30",
						"resource":    "foo",
						"template":    "random",
					},
					"labels": map[string]any{
						"severity": "foo",
					},
					"startsAt": timenow.String(),
					"status":   "foo",
				},
			},
		}
	)

	var testCases = []struct {
		name              string
		setup             func(*mocks.Repository, *mocks.AlertTransformer)
		alertsToBeCreated map[string]any
		expectedAlerts    []alert.Alert
		expectedFiringLen int
		wantErr           bool
	}{
		{
			name: "should return error if TransformToAlerts return error",
			setup: func(ar *mocks.Repository, at *mocks.AlertTransformer) {
				at.EXPECT().TransformToAlerts(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]interface {}")).Return(nil, 0, errors.New("some error"))
			},
			alertsToBeCreated: alertsToBeCreated,
			wantErr:           true,
		},
		{
			name: "should call repository Create method with proper arguments",
			setup: func(ar *mocks.Repository, at *mocks.AlertTransformer) {
				at.EXPECT().TransformToAlerts(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]interface {}")).Return([]alert.Alert{
					{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "lag", MetricValue: "20",
						Rule: "lagHigh", TriggeredAt: timenow},
				}, 1, nil)
				ar.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("alert.Alert")).Return(alert.Alert{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "lag", MetricValue: "20",
					Rule: "lagHigh", TriggeredAt: timenow}, nil)
			},
			alertsToBeCreated: alertsToBeCreated,
			expectedFiringLen: 1,
			expectedAlerts: []alert.Alert{
				{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "lag", MetricValue: "20",
					Rule: "lagHigh", TriggeredAt: timenow},
			},
		},
		{
			name: "should return error not found if repository return err relation",
			setup: func(ar *mocks.Repository, at *mocks.AlertTransformer) {
				at.EXPECT().TransformToAlerts(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]interface {}")).Return([]alert.Alert{
					{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "lag", MetricValue: "20",
						Rule: "lagHigh", TriggeredAt: timenow},
				}, 1, nil)
				ar.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(alert.Alert{}, alert.ErrRelation)
			},
			alertsToBeCreated: alertsToBeCreated,
			wantErr:           true,
		},
		{
			name: "should handle errors from repository",
			setup: func(ar *mocks.Repository, at *mocks.AlertTransformer) {
				at.EXPECT().TransformToAlerts(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]interface {}")).Return([]alert.Alert{
					{ID: 1, ProviderID: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "lag", MetricValue: "20",
						Rule: "lagHigh", TriggeredAt: timenow},
				}, 1, nil)
				ar.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.Anything).Return(alert.Alert{}, errors.New("random error"))
			},
			alertsToBeCreated: alertsToBeCreated,
			wantErr:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				repositoryMock       = &mocks.Repository{}
				alertTransformerMock = &mocks.AlertTransformer{}
			)

			if tc.setup != nil {
				tc.setup(repositoryMock, alertTransformerMock)
			}

			svc := alert.NewService(repositoryMock, nil, map[string]alert.AlertTransformer{
				testType: alertTransformerMock,
			})
			actualAlerts, firingLen, err := svc.CreateAlerts(ctx, testType, 1, 1, tc.alertsToBeCreated)
			if tc.wantErr {
				if err == nil {
					t.Error("error should not be nil")
				}
			} else {
				if diff := cmp.Diff(actualAlerts, tc.expectedAlerts); diff != "" {
					t.Errorf("result not equal, diff are %+v", diff)
				}
				assert.Equal(t, tc.expectedFiringLen, firingLen)
			}
		})
	}
}
