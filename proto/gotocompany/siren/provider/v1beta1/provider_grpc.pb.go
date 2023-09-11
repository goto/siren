// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gotocompany/siren/provider/v1beta1/provider.proto

package sirenproviderv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProviderServiceClient is the client API for ProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServiceClient interface {
	SyncRuntimeConfig(ctx context.Context, in *SyncRuntimeConfigRequest, opts ...grpc.CallOption) (*SyncRuntimeConfigResponse, error)
	UpsertRule(ctx context.Context, in *UpsertRuleRequest, opts ...grpc.CallOption) (*UpsertRuleResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	TransformToAlerts(ctx context.Context, in *TransformToAlertsRequest, opts ...grpc.CallOption) (*TransformToAlertsResponse, error)
}

type providerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceClient(cc grpc.ClientConnInterface) ProviderServiceClient {
	return &providerServiceClient{cc}
}

func (c *providerServiceClient) SyncRuntimeConfig(ctx context.Context, in *SyncRuntimeConfigRequest, opts ...grpc.CallOption) (*SyncRuntimeConfigResponse, error) {
	out := new(SyncRuntimeConfigResponse)
	err := c.cc.Invoke(ctx, "/gotocompany.siren.provider.v1beta1.ProviderService/SyncRuntimeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) UpsertRule(ctx context.Context, in *UpsertRuleRequest, opts ...grpc.CallOption) (*UpsertRuleResponse, error) {
	out := new(UpsertRuleResponse)
	err := c.cc.Invoke(ctx, "/gotocompany.siren.provider.v1beta1.ProviderService/UpsertRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/gotocompany.siren.provider.v1beta1.ProviderService/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) TransformToAlerts(ctx context.Context, in *TransformToAlertsRequest, opts ...grpc.CallOption) (*TransformToAlertsResponse, error) {
	out := new(TransformToAlertsResponse)
	err := c.cc.Invoke(ctx, "/gotocompany.siren.provider.v1beta1.ProviderService/TransformToAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServiceServer is the server API for ProviderService service.
// All implementations must embed UnimplementedProviderServiceServer
// for forward compatibility
type ProviderServiceServer interface {
	SyncRuntimeConfig(context.Context, *SyncRuntimeConfigRequest) (*SyncRuntimeConfigResponse, error)
	UpsertRule(context.Context, *UpsertRuleRequest) (*UpsertRuleResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	TransformToAlerts(context.Context, *TransformToAlertsRequest) (*TransformToAlertsResponse, error)
	mustEmbedUnimplementedProviderServiceServer()
}

// UnimplementedProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServer struct {
}

func (UnimplementedProviderServiceServer) SyncRuntimeConfig(context.Context, *SyncRuntimeConfigRequest) (*SyncRuntimeConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRuntimeConfig not implemented")
}
func (UnimplementedProviderServiceServer) UpsertRule(context.Context, *UpsertRuleRequest) (*UpsertRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertRule not implemented")
}
func (UnimplementedProviderServiceServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedProviderServiceServer) TransformToAlerts(context.Context, *TransformToAlertsRequest) (*TransformToAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformToAlerts not implemented")
}
func (UnimplementedProviderServiceServer) mustEmbedUnimplementedProviderServiceServer() {}

// UnsafeProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServiceServer will
// result in compilation errors.
type UnsafeProviderServiceServer interface {
	mustEmbedUnimplementedProviderServiceServer()
}

func RegisterProviderServiceServer(s grpc.ServiceRegistrar, srv ProviderServiceServer) {
	s.RegisterService(&ProviderService_ServiceDesc, srv)
}

func _ProviderService_SyncRuntimeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRuntimeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).SyncRuntimeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotocompany.siren.provider.v1beta1.ProviderService/SyncRuntimeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).SyncRuntimeConfig(ctx, req.(*SyncRuntimeConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_UpsertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).UpsertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotocompany.siren.provider.v1beta1.ProviderService/UpsertRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).UpsertRule(ctx, req.(*UpsertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotocompany.siren.provider.v1beta1.ProviderService/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_TransformToAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformToAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).TransformToAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotocompany.siren.provider.v1beta1.ProviderService/TransformToAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).TransformToAlerts(ctx, req.(*TransformToAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderService_ServiceDesc is the grpc.ServiceDesc for ProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gotocompany.siren.provider.v1beta1.ProviderService",
	HandlerType: (*ProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncRuntimeConfig",
			Handler:    _ProviderService_SyncRuntimeConfig_Handler,
		},
		{
			MethodName: "UpsertRule",
			Handler:    _ProviderService_UpsertRule_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _ProviderService_SetConfig_Handler,
		},
		{
			MethodName: "TransformToAlerts",
			Handler:    _ProviderService_TransformToAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gotocompany/siren/provider/v1beta1/provider.proto",
}
