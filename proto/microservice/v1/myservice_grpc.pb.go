// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package myservice_proto

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

// SserviceServiceClient is the client API for SserviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SserviceServiceClient interface {
	ScheduleReminder(ctx context.Context, in *ScheduleReminderRequest, opts ...grpc.CallOption) (*ScheduleReminderResponse, error)
}

type sserviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSserviceServiceClient(cc grpc.ClientConnInterface) SserviceServiceClient {
	return &sserviceServiceClient{cc}
}

func (c *sserviceServiceClient) ScheduleReminder(ctx context.Context, in *ScheduleReminderRequest, opts ...grpc.CallOption) (*ScheduleReminderResponse, error) {
	out := new(ScheduleReminderResponse)
	err := c.cc.Invoke(ctx, "/microservice.sservice.v1.SserviceService/ScheduleReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SserviceServiceServer is the server API for SserviceService service.
// All implementations must embed UnimplementedSserviceServiceServer
// for forward compatibility
type SserviceServiceServer interface {
	ScheduleReminder(context.Context, *ScheduleReminderRequest) (*ScheduleReminderResponse, error)
	mustEmbedUnimplementedSserviceServiceServer()
}

// UnimplementedSserviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSserviceServiceServer struct {
}

func (UnimplementedSserviceServiceServer) ScheduleReminder(context.Context, *ScheduleReminderRequest) (*ScheduleReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleReminder not implemented")
}
func (UnimplementedSserviceServiceServer) mustEmbedUnimplementedSserviceServiceServer() {}

// UnsafeSserviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SserviceServiceServer will
// result in compilation errors.
type UnsafeSserviceServiceServer interface {
	mustEmbedUnimplementedSserviceServiceServer()
}

func RegisterSserviceServiceServer(s grpc.ServiceRegistrar, srv SserviceServiceServer) {
	s.RegisterService(&SserviceService_ServiceDesc, srv)
}

func _SserviceService_ScheduleReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SserviceServiceServer).ScheduleReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservice.sservice.v1.SserviceService/ScheduleReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SserviceServiceServer).ScheduleReminder(ctx, req.(*ScheduleReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SserviceService_ServiceDesc is the grpc.ServiceDesc for SserviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SserviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservice.sservice.v1.SserviceService",
	HandlerType: (*SserviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleReminder",
			Handler:    _SserviceService_ScheduleReminder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/microservice/v1/myservice.proto",
}
