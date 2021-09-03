// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package app_data_monitoring_bp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AlertServiceClient is the client API for AlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertServiceClient interface {
	CreateAlertNotification(ctx context.Context, in *SlackAlertRequest, opts ...grpc.CallOption) (*SlackAlertResponse, error)
}

type alertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertServiceClient(cc grpc.ClientConnInterface) AlertServiceClient {
	return &alertServiceClient{cc}
}

func (c *alertServiceClient) CreateAlertNotification(ctx context.Context, in *SlackAlertRequest, opts ...grpc.CallOption) (*SlackAlertResponse, error) {
	out := new(SlackAlertResponse)
	err := c.cc.Invoke(ctx, "/app_data_monitoring.v1.AlertService/CreateAlertNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServiceServer is the server API for AlertService service.
// All implementations should embed UnimplementedAlertServiceServer
// for forward compatibility
type AlertServiceServer interface {
	CreateAlertNotification(context.Context, *SlackAlertRequest) (*SlackAlertResponse, error)
}

// UnimplementedAlertServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAlertServiceServer struct {
}

func (UnimplementedAlertServiceServer) CreateAlertNotification(context.Context, *SlackAlertRequest) (*SlackAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertNotification not implemented")
}

// UnsafeAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServiceServer will
// result in compilation errors.
type UnsafeAlertServiceServer interface {
	mustEmbedUnimplementedAlertServiceServer()
}

func RegisterAlertServiceServer(s grpc.ServiceRegistrar, srv AlertServiceServer) {
	s.RegisterService(&_AlertService_serviceDesc, srv)
}

func _AlertService_CreateAlertNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlackAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).CreateAlertNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app_data_monitoring.v1.AlertService/CreateAlertNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).CreateAlertNotification(ctx, req.(*SlackAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app_data_monitoring.v1.AlertService",
	HandlerType: (*AlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlertNotification",
			Handler:    _AlertService_CreateAlertNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/app_data_monitoring_bp/alert.proto",
}
