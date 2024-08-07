package v1beta1

import (
	"context"
	"fmt"

	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/internal/api"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GRPCServer) ListNamespaces(ctx context.Context, _ *sirenv1beta1.ListNamespacesRequest) (*sirenv1beta1.ListNamespacesResponse, error) {
	namespaces, err := s.namespaceService.List(ctx)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	items := []*sirenv1beta1.Namespace{}
	for _, namespace := range namespaces {
		credentials, err := structpb.NewStruct(namespace.Credentials)
		if err != nil {
			return nil, api.GenerateRPCErr(s.logger, fmt.Errorf("failed to fetch namespace credentials: %w", err))
		}

		item := &sirenv1beta1.Namespace{
			Id:          namespace.ID,
			Urn:         namespace.URN,
			Name:        namespace.Name,
			Credentials: credentials,
			Labels:      namespace.Labels,
			Provider:    namespace.Provider.ID,
			CreatedAt:   timestamppb.New(namespace.CreatedAt),
			UpdatedAt:   timestamppb.New(namespace.UpdatedAt),
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListNamespacesResponse{
		Namespaces: items,
	}, nil
}

func (s *GRPCServer) CreateNamespace(ctx context.Context, req *sirenv1beta1.CreateNamespaceRequest) (*sirenv1beta1.CreateNamespaceResponse, error) {
	ns := &namespace.Namespace{
		Provider: provider.Provider{
			ID: req.GetProvider(),
		},
		URN:         req.GetUrn(),
		Name:        req.GetName(),
		Credentials: req.GetCredentials().AsMap(),
		Labels:      req.GetLabels(),
	}
	err := s.namespaceService.Create(ctx, ns)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.CreateNamespaceResponse{
		Id: ns.ID,
	}, nil
}

func (s *GRPCServer) GetNamespace(ctx context.Context, req *sirenv1beta1.GetNamespaceRequest) (*sirenv1beta1.GetNamespaceResponse, error) {
	namespace, err := s.namespaceService.Get(ctx, req.GetId())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	credentials, err := structpb.NewStruct(namespace.Credentials)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, fmt.Errorf("failed to fetch namespace credentials: %w", err))
	}

	return &sirenv1beta1.GetNamespaceResponse{
		Namespace: &sirenv1beta1.Namespace{
			Id:          namespace.ID,
			Urn:         namespace.URN,
			Name:        namespace.Name,
			Credentials: credentials,
			Labels:      namespace.Labels,
			Provider:    namespace.Provider.ID,
			CreatedAt:   timestamppb.New(namespace.CreatedAt),
			UpdatedAt:   timestamppb.New(namespace.UpdatedAt),
		},
	}, nil
}

func (s *GRPCServer) UpdateNamespace(ctx context.Context, req *sirenv1beta1.UpdateNamespaceRequest) (*sirenv1beta1.UpdateNamespaceResponse, error) {
	ns := &namespace.Namespace{
		ID: req.GetId(),
		Provider: provider.Provider{
			ID: req.GetProvider(),
		},
		Name:        req.GetName(),
		Credentials: req.GetCredentials().AsMap(),
		Labels:      req.GetLabels(),
	}
	err := s.namespaceService.Update(ctx, ns)
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.UpdateNamespaceResponse{
		Id: ns.ID,
	}, nil
}

func (s *GRPCServer) DeleteNamespace(ctx context.Context, req *sirenv1beta1.DeleteNamespaceRequest) (*sirenv1beta1.DeleteNamespaceResponse, error) {
	err := s.namespaceService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, api.GenerateRPCErr(s.logger, err)
	}

	return &sirenv1beta1.DeleteNamespaceResponse{}, nil
}
