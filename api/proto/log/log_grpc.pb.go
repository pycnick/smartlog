// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// SmartLogV1Client is the client API for SmartLogV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartLogV1Client interface {
	SendLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*LogStatus, error)
}

type smartLogV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSmartLogV1Client(cc grpc.ClientConnInterface) SmartLogV1Client {
	return &smartLogV1Client{cc}
}

func (c *smartLogV1Client) SendLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*LogStatus, error) {
	out := new(LogStatus)
	err := c.cc.Invoke(ctx, "/SmartLogV1/SendLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartLogV1Server is the server API for SmartLogV1 service.
// All implementations must embed UnimplementedSmartLogV1Server
// for forward compatibility
type SmartLogV1Server interface {
	SendLog(context.Context, *Log) (*LogStatus, error)
	mustEmbedUnimplementedSmartLogV1Server()
}

// UnimplementedSmartLogV1Server must be embedded to have forward compatible implementations.
type UnimplementedSmartLogV1Server struct {
}

func (UnimplementedSmartLogV1Server) SendLog(context.Context, *Log) (*LogStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLog not implemented")
}
func (UnimplementedSmartLogV1Server) mustEmbedUnimplementedSmartLogV1Server() {}

// UnsafeSmartLogV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartLogV1Server will
// result in compilation errors.
type UnsafeSmartLogV1Server interface {
	mustEmbedUnimplementedSmartLogV1Server()
}

func RegisterSmartLogV1Server(s grpc.ServiceRegistrar, srv SmartLogV1Server) {
	s.RegisterService(&SmartLogV1_ServiceDesc, srv)
}

func _SmartLogV1_SendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartLogV1Server).SendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartLogV1/SendLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartLogV1Server).SendLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartLogV1_ServiceDesc is the grpc.ServiceDesc for SmartLogV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartLogV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SmartLogV1",
	HandlerType: (*SmartLogV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLog",
			Handler:    _SmartLogV1_SendLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}