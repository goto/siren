package v1beta1_test

import (
	"context"
	"testing"
	"time"

	"github.com/goto/salt/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/internal/api/mocks"
	"github.com/goto/siren/internal/api/v1beta1"
	"github.com/goto/siren/pkg/errors"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestGRPCServer_ListReceiver(t *testing.T) {
	configurations := make(map[string]any)
	configurations["foo"] = "bar"
	labels := make(map[string]string)
	labels["foo"] = "bar"
	dummyResult := []receiver.Receiver{
		{
			ID:             1,
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	t.Run("should return list of all receiver", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)

		mockedReceiverService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), receiver.Filter{}).Return(dummyResult, nil).Once()

		res, err := dummyGRPCServer.ListReceivers(context.TODO(), &sirenv1beta1.ListReceiversRequest{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetReceivers()))
		assert.Equal(t, uint64(1), res.GetReceivers()[0].GetId())
		assert.Equal(t, "foo", res.GetReceivers()[0].GetName())
		assert.Equal(t, "bar", res.GetReceivers()[0].GetType())
		assert.Equal(t, "bar", res.GetReceivers()[0].GetConfigurations().AsMap()["foo"])
		assert.Equal(t, "bar", res.GetReceivers()[0].GetLabels()["foo"])
	})

	t.Run("should return error Internal if getting providers failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), receiver.Filter{}).
			Return(nil, errors.New("random error"))

		res, err := dummyGRPCServer.ListReceivers(context.TODO(), &sirenv1beta1.ListReceiversRequest{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})

	t.Run("should return error Internal if NewStruct conversion failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		configurations["foo"] = string([]byte{0xff})
		dummyResult := []receiver.Receiver{
			{
				ID:             1,
				Name:           "foo",
				Type:           "bar",
				Labels:         labels,
				Configurations: configurations,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		}

		mockedReceiverService.EXPECT().List(mock.AnythingOfType("context.todoCtx"), receiver.Filter{}).
			Return(dummyResult, nil)
		res, err := dummyGRPCServer.ListReceivers(context.TODO(), &sirenv1beta1.ListReceiversRequest{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_CreateReceiver(t *testing.T) {
	configurations := make(map[string]any)
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"
	labels := make(map[string]string)
	labels["foo"] = "bar"
	generatedID := uint64(77)

	configurationsData, _ := structpb.NewStruct(configurations)
	dummyReq := &sirenv1beta1.CreateReceiverRequest{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurationsData,
	}
	payload := &receiver.Receiver{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("Should create a receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), payload).Run(func(ctx context.Context, rcv *receiver.Receiver) {
			rcv.ID = generatedID
		}).Return(nil).Once()
		res, err := dummyGRPCServer.CreateReceiver(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, generatedID, res.GetId())
	})

	t.Run("should return error Invalid Argument if create receiver failed with err invalid", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)

		mockedReceiverService.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), payload).Return(errors.ErrInvalid).Once()

		res, err := dummyGRPCServer.CreateReceiver(context.TODO(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = request is not valid")
		assert.Nil(t, res)
	})

	t.Run("should return error Internal if creating receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), payload).Return(errors.New("random error")).Once()

		res, err := dummyGRPCServer.CreateReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})

}

func TestGRPCServer_GetReceiver(t *testing.T) {
	configurations := make(map[string]any)
	configurations["foo"] = "bar"
	labels := make(map[string]string)
	labels["foo"] = "bar"

	receiverId := uint64(1)
	dummyReq := &sirenv1beta1.GetReceiverRequest{
		Id: 1,
	}
	payload := &receiver.Receiver{
		Name:           "foo",
		Type:           "bar",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("should return a receiver", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), receiverId, mock.AnythingOfType("receiver.GetOption")).
			Return(payload, nil).Once()

		res, err := dummyGRPCServer.GetReceiver(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "foo", res.GetReceiver().GetName())
		assert.Equal(t, "bar", res.GetReceiver().GetType())
		assert.Equal(t, "bar", res.GetReceiver().GetLabels()["foo"])
		assert.Equal(t, "bar", res.GetReceiver().GetConfigurations().AsMap()["foo"])
	})

	t.Run("should return error Not Found if no receiver found", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), receiverId, mock.AnythingOfType("receiver.GetOption")).
			Return(nil, errors.ErrNotFound).Once()

		res, err := dummyGRPCServer.GetReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = NotFound desc = requested entity not found")
	})

	t.Run("should return error Internal if getting receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), receiverId, mock.AnythingOfType("receiver.GetOption")).
			Return(payload, errors.New("random error")).Once()

		res, err := dummyGRPCServer.GetReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})

	t.Run("should return error Internal if NewStruct conversion of configuration failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)

		configurations["foo"] = string([]byte{0xff})
		payload := &receiver.Receiver{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
		}

		mockedReceiverService.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), receiverId, mock.AnythingOfType("receiver.GetOption")).
			Return(payload, nil)
		res, err := dummyGRPCServer.GetReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})

	t.Run("should return error Internal if data NewStruct conversion of data failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		data := make(map[string]any)
		data["channels"] = string([]byte{0xff})
		payload := &receiver.Receiver{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
			Data:           data,
		}

		mockedReceiverService.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), receiverId, mock.AnythingOfType("receiver.GetOption")).
			Return(payload, nil)
		res, err := dummyGRPCServer.GetReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_UpdateReceiver(t *testing.T) {
	configurations := make(map[string]any)
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"

	labels := make(map[string]string)
	labels["foo"] = "bar"

	configurationsData, _ := structpb.NewStruct(configurations)
	dummyReq := &sirenv1beta1.UpdateReceiverRequest{
		Id:             uint64(22),
		Name:           "foo",
		Labels:         labels,
		Configurations: configurationsData,
	}
	payload := &receiver.Receiver{
		ID:             uint64(22),
		Name:           "foo",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("should update receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), payload).Run(func(ctx context.Context, rcv *receiver.Receiver) {
			rcv.ID = payload.ID
		}).Return(nil).Once()

		res, err := dummyGRPCServer.UpdateReceiver(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, payload.ID, res.GetId())
	})

	t.Run("should return error Invalid Argument if update receiver return invalid error", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)

		mockedReceiverService.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), payload).Return(errors.ErrInvalid).Once()

		res, err := dummyGRPCServer.UpdateReceiver(context.TODO(), dummyReq)
		assert.EqualError(t, err, "rpc error: code = InvalidArgument desc = request is not valid")
		assert.Nil(t, res)
	})

	t.Run("should return error Internal if updating receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}

		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), payload).Return(errors.New("random error"))

		res, err := dummyGRPCServer.UpdateReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_DeleteReceiver(t *testing.T) {
	providerId := uint64(10)
	dummyReq := &sirenv1beta1.DeleteReceiverRequest{
		Id: uint64(10),
	}

	t.Run("should delete receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), providerId).
			Return(nil).Once()

		res, err := dummyGRPCServer.DeleteReceiver(context.TODO(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "", res.String())
	})

	t.Run("should return error Internal if deleting receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer, err := v1beta1.NewGRPCServer(log.NewNoop(), api.HeadersConfig{}, &api.Deps{ReceiverService: mockedReceiverService})
		require.NoError(t, err)
		mockedReceiverService.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), providerId).
			Return(errors.New("random error")).Once()

		res, err := dummyGRPCServer.DeleteReceiver(context.TODO(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}
