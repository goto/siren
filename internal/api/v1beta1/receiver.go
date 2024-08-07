package v1beta1

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/secret"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
)

func (s *GRPCServer) ListReceivers(ctx context.Context, req *sirenv1beta1.ListReceiversRequest) (*sirenv1beta1.ListReceiversResponse, error) {
	var filterMultipleLabels []map[string]string

	if len(req.GetLabels()) > 0 {
		filterMultipleLabels = []map[string]string{req.GetLabels()}
	}

	receivers, err := s.receiverService.List(ctx, receiver.Filter{
		MultipleLabels: filterMultipleLabels,
	})
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1beta1.Receiver{}
	for _, rcv := range receivers {
		configurations, err := structpb.NewStruct(sanitizeConfigMap(rcv.Configurations))
		if err != nil {
			return nil, api.GenerateRPCErr(s.logger, err)
		}

		item := &sirenv1beta1.Receiver{
			Id:             rcv.ID,
			Name:           rcv.Name,
			Type:           rcv.Type,
			Configurations: configurations,
			Labels:         rcv.Labels,
			ParentId:       rcv.ParentID,
			CreatedAt:      timestamppb.New(rcv.CreatedAt),
			UpdatedAt:      timestamppb.New(rcv.UpdatedAt),
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListReceiversResponse{
		Receivers: items,
	}, nil
}

func (s *GRPCServer) CreateReceiver(ctx context.Context, req *sirenv1beta1.CreateReceiverRequest) (*sirenv1beta1.CreateReceiverResponse, error) {
	if !receiver.IsTypeSupported(req.GetType()) {
		return nil, api.GenerateRPCErr(s.logger, errors.ErrInvalid.WithMsgf("unsupported type %s", req.GetType()))
	}

	rcv := &receiver.Receiver{
		Name:           req.GetName(),
		Type:           req.GetType(),
		Labels:         req.GetLabels(),
		Configurations: req.GetConfigurations().AsMap(),
		ParentID:       req.GetParentId(),
	}

	err := s.receiverService.Create(ctx, rcv)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.CreateReceiverResponse{
		Id: rcv.ID,
	}, nil
}

func (s *GRPCServer) GetReceiver(ctx context.Context, req *sirenv1beta1.GetReceiverRequest) (*sirenv1beta1.GetReceiverResponse, error) {
	rcv, err := s.receiverService.Get(ctx, req.GetId(), receiver.GetWithData())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	data, err := structpb.NewStruct(rcv.Data)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	configuration, err := structpb.NewStruct(sanitizeConfigMap(rcv.Configurations))
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.GetReceiverResponse{
		Receiver: &sirenv1beta1.Receiver{
			Id:             rcv.ID,
			Name:           rcv.Name,
			Type:           rcv.Type,
			Labels:         rcv.Labels,
			Configurations: configuration,
			Data:           data,
			ParentId:       rcv.ParentID,
			CreatedAt:      timestamppb.New(rcv.CreatedAt),
			UpdatedAt:      timestamppb.New(rcv.UpdatedAt),
		},
	}, nil
}

func (s *GRPCServer) UpdateReceiver(ctx context.Context, req *sirenv1beta1.UpdateReceiverRequest) (*sirenv1beta1.UpdateReceiverResponse, error) {
	rcv := &receiver.Receiver{
		ID:             req.GetId(),
		Name:           req.GetName(),
		Labels:         req.GetLabels(),
		Configurations: req.GetConfigurations().AsMap(),
		ParentID:       req.GetParentId(),
	}

	if err := s.receiverService.Update(ctx, rcv); err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.UpdateReceiverResponse{
		Id: rcv.ID,
	}, nil
}

func (s *GRPCServer) DeleteReceiver(ctx context.Context, req *sirenv1beta1.DeleteReceiverRequest) (*sirenv1beta1.DeleteReceiverResponse, error) {
	err := s.receiverService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.DeleteReceiverResponse{}, nil
}

// sanitizeConfigMap does all sanitization to present receiver configurations to the user
func sanitizeConfigMap(receiverConfigMap map[string]any) map[string]any {
	var newConfigMap = make(map[string]any)
	for k, v := range receiverConfigMap {
		// sanitize maskable string. convert `secret.MaskableString` to string to be compatible with structpb
		if reflect.TypeOf(v) == reflect.TypeOf(secret.MaskableString("")) {
			newConfigMap[k] = fmt.Sprintf("%v", v)
		} else {
			newConfigMap[k] = v
		}
	}
	return newConfigMap
}
