package provider_test

import (
	"context"
	"testing"
	"time"

	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/core/provider/mocks"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var supportedTypes = []string{"cortex"}

func TestList(t *testing.T) {
	ctx := context.TODO()
	credentials := make(pgc.StringAnyMap)
	credentials["foo"] = "cortex"
	labels := make(pgc.StringStringMap)
	labels["foo"] = "cortex"

	t.Run("should call repository List method and return result in domain's type", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		dummyProviders := []provider.Provider{
			{
				ID:          10,
				Host:        "foo",
				Type:        "cortex",
				Name:        "foo",
				Credentials: credentials,
				Labels:      labels,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}
		repositoryMock.EXPECT().List(mock.AnythingOfType("context.todoCtx"), provider.Filter{}).Return(dummyProviders, nil).Once()
		result, err := dummyService.List(ctx, provider.Filter{})
		assert.Nil(t, err)
		assert.Equal(t, len(dummyProviders), len(result))
		assert.Equal(t, dummyProviders[0].Name, result[0].Name)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository List method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().List(mock.AnythingOfType("context.todoCtx"), provider.Filter{}).Return(nil, errors.New("random error")).Once()
		result, err := dummyService.List(ctx, provider.Filter{})
		assert.Nil(t, result)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	credentials := make(pgc.StringAnyMap)
	credentials["foo"] = "cortex"
	labels := make(pgc.StringStringMap)
	labels["foo"] = "cortex"
	timenow := time.Now()
	dummyProviderID := uint64(10)
	dummyProvider := &provider.Provider{
		ID:          dummyProviderID,
		Host:        "foo",
		Type:        "cortex",
		Name:        "foo",
		Credentials: credentials,
		Labels:      labels,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}

	t.Run("should call repository Create method and return result in domain's type", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(nil).Once()
		err := dummyService.Create(ctx, dummyProvider)
		assert.Nil(t, err)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Create method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(errors.New("random error")).Once()
		err := dummyService.Create(ctx, dummyProvider)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Create method and return conflict error if duplicated", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(provider.ErrDuplicate).Once()
		err := dummyService.Create(ctx, dummyProvider)
		assert.EqualError(t, err, "urn already exist")
		repositoryMock.AssertExpectations(t)
	})
}

func TestGetProvider(t *testing.T) {
	ctx := context.TODO()
	dummyProviderID := uint64(10)
	credentials := make(pgc.StringAnyMap)
	credentials["foo"] = "cortex"
	labels := make(pgc.StringStringMap)
	labels["foo"] = "cortex"
	timenow := time.Now()
	dummyProvider := &provider.Provider{
		ID:          dummyProviderID,
		Host:        "foo",
		Type:        "cortex",
		Name:        "foo",
		Credentials: credentials,
		Labels:      labels,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}

	t.Run("should call repository Get method and return result in domain's type", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), dummyProviderID).Return(dummyProvider, nil).Once()
		result, err := dummyService.Get(ctx, dummyProviderID)
		assert.Nil(t, err)
		assert.Equal(t, dummyProvider, result)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Get method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), dummyProviderID).Return(nil, errors.New("random error")).Once()
		result, err := dummyService.Get(ctx, dummyProviderID)
		assert.Empty(t, result)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Get method and return error if repository return not found error", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), dummyProviderID).Return(nil, provider.NotFoundError{}).Once()
		result, err := dummyService.Get(ctx, dummyProviderID)
		assert.Empty(t, result)
		assert.EqualError(t, err, "provider not found")
		repositoryMock.AssertExpectations(t)
	})
}

func TestUpdateProvider(t *testing.T) {
	ctx := context.TODO()
	dummyProviderID := uint64(10)
	timenow := time.Now()
	credentials := make(pgc.StringAnyMap)
	credentials["foo"] = "cortex"
	labels := make(pgc.StringStringMap)
	labels["foo"] = "cortex"
	dummyProvider := &provider.Provider{
		ID:          dummyProviderID,
		Host:        "foo",
		Type:        "cortex",
		Name:        "foo",
		Credentials: credentials,
		Labels:      labels,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}

	t.Run("should call repository Update method and return result in domain's type", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(nil).Once()
		err := dummyService.Update(ctx, dummyProvider)
		assert.Nil(t, err)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Update method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(errors.New("random error")).Once()
		err := dummyService.Update(ctx, dummyProvider)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Update method and return error not found if repository return not found error", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(provider.NotFoundError{}).Once()
		err := dummyService.Update(ctx, dummyProvider)
		assert.EqualError(t, err, "provider not found")
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Update method and return conflict error if repository return duplicate error", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), dummyProvider).Return(provider.ErrDuplicate).Once()
		err := dummyService.Update(ctx, dummyProvider)
		assert.EqualError(t, err, "urn already exist")
		repositoryMock.AssertExpectations(t)
	})
}

func TestDeleteProvider(t *testing.T) {
	ctx := context.TODO()
	providerID := uint64(10)

	t.Run("should call repository Delete method and return nil if no error", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), providerID).Return(nil).Once()
		err := dummyService.Delete(ctx, providerID)
		assert.Nil(t, err)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Delete method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := provider.NewService(repositoryMock, supportedTypes)
		repositoryMock.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), providerID).Return(errors.New("random error")).Once()
		err := dummyService.Delete(ctx, providerID)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})
}
