package v1beta1_test

import (
	"context"
	"testing"
	"time"

	"github.com/goto/salt/log"
	"github.com/goto/siren/pkg/errors"

	"github.com/goto/siren/core/rule"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/api/mocks"
	"github.com/goto/siren/internal/api/v1beta1"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGRPCServer_ListRules(t *testing.T) {
	dummyPayload := &sirenv1beta1.ListRulesRequest{
		Name:              "foo",
		Namespace:         "test",
		GroupName:         "foo",
		Template:          "foo",
		ProviderNamespace: 1,
	}

	t.Run("should return stored rules", func(t *testing.T) {
		ctx := context.TODO()
		mockedRuleService := &mocks.RuleService{}
		dummyResult := []rule.Rule{
			{
				Name:      "foo",
				Enabled:   true,
				GroupName: "foo",
				Namespace: "test",
				Template:  "foo",
				Variables: []rule.RuleVariable{
					{
						Name:        "foo",
						Type:        "int",
						Value:       "bar",
						Description: "",
					},
				},
				ProviderNamespace: 1,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
		}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{RuleService: mockedRuleService})
		require.NoError(t, err)
		mockedRuleService.EXPECT().List(ctx, rule.Filter{
			Name:         dummyPayload.Name,
			Namespace:    dummyPayload.Namespace,
			GroupName:    dummyPayload.GroupName,
			TemplateName: dummyPayload.Template,
			NamespaceID:  dummyPayload.ProviderNamespace,
		}).
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListRules(ctx, dummyPayload)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetRules()))
		assert.Equal(t, "foo", res.GetRules()[0].GetName())
		assert.Equal(t, "test", res.GetRules()[0].GetNamespace())
		assert.Equal(t, true, res.GetRules()[0].GetEnabled())
		assert.Equal(t, 1, len(res.GetRules()[0].GetVariables()))
		mockedRuleService.AssertExpectations(t)
	})

	t.Run("should return error Internal if getting rules failed", func(t *testing.T) {
		ctx := context.TODO()
		mockedRuleService := &mocks.RuleService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{RuleService: mockedRuleService})
		require.NoError(t, err)
		mockedRuleService.EXPECT().List(ctx, rule.Filter{
			Name:         dummyPayload.Name,
			Namespace:    dummyPayload.Namespace,
			GroupName:    dummyPayload.GroupName,
			TemplateName: dummyPayload.Template,
			NamespaceID:  dummyPayload.ProviderNamespace,
		}).Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.ListRules(ctx, dummyPayload)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_UpdateRules(t *testing.T) {
	testID := uint64(88)
	dummyPayload := &rule.Rule{
		Enabled:   true,
		GroupName: "foo",
		Namespace: "test",
		Template:  "foo",
		Variables: []rule.RuleVariable{
			{
				Name:        "foo",
				Type:        "int",
				Value:       "bar",
				Description: "",
			},
		},
		ProviderNamespace: 1,
	}

	dummyReq := &sirenv1beta1.UpdateRuleRequest{
		Enabled:   true,
		GroupName: "foo",
		Namespace: "test",
		Template:  "foo",
		Variables: []*sirenv1beta1.Variables{
			{
				Name:        "foo",
				Type:        "int",
				Value:       "bar",
				Description: "",
			},
		},
		ProviderNamespace: 1,
	}

	t.Run("should update rule", func(t *testing.T) {
		ctx := context.TODO()
		mockedRuleService := &mocks.RuleService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{RuleService: mockedRuleService})
		require.NoError(t, err)
		dummyResult := &rule.Rule{}
		*dummyResult = *dummyPayload
		dummyResult.Enabled = false
		dummyResult.Name = "foo"
		dummyResult.ID = testID

		mockedRuleService.EXPECT().Upsert(ctx, dummyPayload).Run(func(ctx context.Context, r *rule.Rule) {
			r.ID = testID
		}).Return(nil).Once()
		res, err := dummyGRPCServer.UpdateRule(ctx, dummyReq)
		assert.Nil(t, err)

		assert.Equal(t, testID, res.GetRule().GetId())
		mockedRuleService.AssertExpectations(t)
	})

	t.Run("should return error AlreadyExist if update rules return err conflict", func(t *testing.T) {
		ctx := context.TODO()
		mockedRuleService := &mocks.RuleService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{RuleService: mockedRuleService})
		require.NoError(t, err)
		mockedRuleService.EXPECT().Upsert(ctx, dummyPayload).Return(errors.ErrConflict).Once()

		res, err := dummyGRPCServer.UpdateRule(ctx, dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = AlreadyExists desc = an entity with conflicting identifier exists")
	})

	t.Run("should return error Internal if getting rules failed", func(t *testing.T) {
		ctx := context.TODO()
		mockedRuleService := &mocks.RuleService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{RuleService: mockedRuleService})
		require.NoError(t, err)
		mockedRuleService.EXPECT().Upsert(ctx, dummyPayload).Return(errors.New("random error")).Once()

		res, err := dummyGRPCServer.UpdateRule(ctx, dummyReq)

		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}
