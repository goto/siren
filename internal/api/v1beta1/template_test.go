package v1beta1_test

import (
	"context"
	"testing"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/api/mocks"
	"github.com/goto/siren/internal/api/v1beta1"
	"github.com/goto/siren/pkg/errors"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGRPCServer_ListTemplates(t *testing.T) {
	t.Run("should return list of all templates", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)

		dummyReq := &sirenv1beta1.ListTemplatesRequest{}
		dummyResult := []template.Template{
			{
				ID:   1,
				Name: "foo",
				Body: "bar",
				Tags: []string{"foo", "bar"},
				Variables: []template.Variable{
					{
						Name:        "foo",
						Type:        "bar",
						Default:     "",
						Description: "",
					},
				},
			},
		}

		mockedTemplateService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), template.Filter{}).
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListTemplates(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetTemplates()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetName())
		assert.Equal(t, "bar", res.GetTemplates()[0].GetBody())
		assert.Equal(t, 1, len(res.GetTemplates()[0].GetVariables()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetVariables()[0].GetName())
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return list of all templates matching particular tag", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)

		dummyReq := &sirenv1beta1.ListTemplatesRequest{
			Tag: "foo",
		}

		dummyResult := []template.Template{
			{
				ID:   1,
				Name: "foo",
				Body: "bar",
				Tags: []string{"foo", "bar"},
				Variables: []template.Variable{
					{
						Name:        "foo",
						Type:        "bar",
						Default:     "",
						Description: "",
					},
				},
			},
		}

		mockedTemplateService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), template.Filter{Tag: "foo"}).
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListTemplates(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetTemplates()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetName())
		assert.Equal(t, "bar", res.GetTemplates()[0].GetBody())
		assert.Equal(t, 1, len(res.GetTemplates()[0].GetVariables()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetVariables()[0].GetName())
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Internal if getting templates failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.ListTemplatesRequest{
			Tag: "foo",
		}
		mockedTemplateService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), template.Filter{Tag: "foo"}).
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.ListTemplates(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		mockedTemplateService.AssertExpectations(t)
	})
}

func TestGRPCServer_GetTemplate(t *testing.T) {
	t.Run("should return template by name", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		dummyResult := &template.Template{
			ID:   1,
			Name: "foo",
			Body: "bar",
			Tags: []string{"foo", "bar"},
			Variables: []template.Variable{
				{
					Name:        "foo",
					Type:        "bar",
					Default:     "",
					Description: "",
				},
			},
		}

		mockedTemplateService.EXPECT().GetByName(mock.AnythingOfType("context.todoCtx"), "foo").
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.GetTemplate(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetTemplate().GetId())
		assert.Equal(t, "foo", res.GetTemplate().GetName())
		assert.Equal(t, "bar", res.GetTemplate().GetBody())
		assert.Equal(t, "foo", res.GetTemplate().GetVariables()[0].GetName())
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Not Found if template does not exist", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().GetByName(mock.AnythingOfType("context.todoCtx"), "foo").
			Return(nil, errors.ErrNotFound).Once()
		res, err := dummyGRPCServer.GetTemplate(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = NotFound desc = requested entity not found")
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Internal if getting template by name failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().GetByName(mock.AnythingOfType("context.todoCtx"), "foo").
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.GetTemplate(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		mockedTemplateService.AssertExpectations(t)
	})
}

func TestGRPCServer_UpsertTemplate(t *testing.T) {
	dummyReq := &sirenv1beta1.UpsertTemplateRequest{
		Id:   1,
		Name: "foo",
		Body: "bar",
		Tags: []string{"foo", "bar"},
		Variables: []*sirenv1beta1.TemplateVariables{
			{
				Name:        "foo",
				Type:        "bar",
				Default:     "",
				Description: "",
			},
		},
	}
	tmpl := &template.Template{
		ID:   1,
		Name: "foo",
		Body: "bar",
		Tags: []string{"foo", "bar"},
		Variables: []template.Variable{
			{
				Name:        "foo",
				Type:        "bar",
				Default:     "",
				Description: "",
			},
		},
	}

	t.Run("should return template by name", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)

		mockedTemplateService.EXPECT().Upsert(mock.AnythingOfType("context.todoCtx"), tmpl).Run(func(_a0 context.Context, _a1 *template.Template) {
			_a1.ID = uint64(1)
		}).Return(nil).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetId())
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error AlreadyExists if upsert template return err conflict", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		mockedTemplateService.EXPECT().Upsert(mock.AnythingOfType("context.todoCtx"), tmpl).Return(errors.ErrConflict).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = AlreadyExists desc = an entity with conflicting identifier exists")
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Internal if upsert template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		mockedTemplateService.EXPECT().Upsert(mock.AnythingOfType("context.todoCtx"), tmpl).Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		mockedTemplateService.AssertExpectations(t)
	})
}

func TestGRPCServer_DeleteTemplate(t *testing.T) {
	t.Run("should delete template", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.DeleteTemplateRequest{
			Name: "foo",
		}

		mockedTemplateService.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), "foo").
			Return(nil).Once()
		res, err := dummyGRPCServer.DeleteTemplate(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, &sirenv1beta1.DeleteTemplateResponse{}, res)
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Internal if deleting template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		dummyReq := &sirenv1beta1.DeleteTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), "foo").
			Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.DeleteTemplate(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		mockedTemplateService.AssertExpectations(t)
	})
}

func TestGRPCServer_RenderTemplate(t *testing.T) {
	dummyReq := &sirenv1beta1.RenderTemplateRequest{
		Name: "foo",
		Variables: map[string]string{
			"foo": "bar",
		},
	}

	t.Run("should render template", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)

		mockedTemplateService.EXPECT().Render(mock.AnythingOfType("context.todoCtx"), "foo", dummyReq.GetVariables()).
			Return("random", nil).Once()
		res, err := dummyGRPCServer.RenderTemplate(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "random", res.GetBody())
		mockedTemplateService.AssertExpectations(t)
	})

	t.Run("should return error Internal if rendering template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{TemplateService: mockedTemplateService})
		require.NoError(t, err)
		mockedTemplateService.EXPECT().Render(mock.AnythingOfType("context.todoCtx"), "foo", dummyReq.GetVariables()).
			Return("", errors.New("random error")).Once()
		res, err := dummyGRPCServer.RenderTemplate(context.TODO(), dummyReq)
		assert.Empty(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
		mockedTemplateService.AssertExpectations(t)
	})
}
