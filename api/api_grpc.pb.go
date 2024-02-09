// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api.proto

package api

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

const (
	ConnectionService_Connection_FullMethodName = "/api.ConnectionService/Connection"
)

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	Connection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) Connection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, ConnectionService_Connection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	Connection(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) Connection(context.Context, *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connection not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_Connection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Connection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectionService_Connection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Connection(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connection",
			Handler:    _ConnectionService_Connection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	USISendService_Send_FullMethodName = "/api.USISendService/Send"
)

// USISendServiceClient is the client API for USISendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type USISendServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type uSISendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUSISendServiceClient(cc grpc.ClientConnInterface) USISendServiceClient {
	return &uSISendServiceClient{cc}
}

func (c *uSISendServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, USISendService_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// USISendServiceServer is the server API for USISendService service.
// All implementations must embed UnimplementedUSISendServiceServer
// for forward compatibility
type USISendServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	mustEmbedUnimplementedUSISendServiceServer()
}

// UnimplementedUSISendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUSISendServiceServer struct {
}

func (UnimplementedUSISendServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedUSISendServiceServer) mustEmbedUnimplementedUSISendServiceServer() {}

// UnsafeUSISendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to USISendServiceServer will
// result in compilation errors.
type UnsafeUSISendServiceServer interface {
	mustEmbedUnimplementedUSISendServiceServer()
}

func RegisterUSISendServiceServer(s grpc.ServiceRegistrar, srv USISendServiceServer) {
	s.RegisterService(&USISendService_ServiceDesc, srv)
}

func _USISendService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USISendServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: USISendService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USISendServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// USISendService_ServiceDesc is the grpc.ServiceDesc for USISendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var USISendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.USISendService",
	HandlerType: (*USISendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _USISendService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	USIReceiveService_Receive_FullMethodName = "/api.USIReceiveService/Receive"
)

// USIReceiveServiceClient is the client API for USIReceiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type USIReceiveServiceClient interface {
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (USIReceiveService_ReceiveClient, error)
}

type uSIReceiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUSIReceiveServiceClient(cc grpc.ClientConnInterface) USIReceiveServiceClient {
	return &uSIReceiveServiceClient{cc}
}

func (c *uSIReceiveServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (USIReceiveService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &USIReceiveService_ServiceDesc.Streams[0], USIReceiveService_Receive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uSIReceiveServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type USIReceiveService_ReceiveClient interface {
	Recv() (*ReceiveResponse, error)
	grpc.ClientStream
}

type uSIReceiveServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *uSIReceiveServiceReceiveClient) Recv() (*ReceiveResponse, error) {
	m := new(ReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// USIReceiveServiceServer is the server API for USIReceiveService service.
// All implementations must embed UnimplementedUSIReceiveServiceServer
// for forward compatibility
type USIReceiveServiceServer interface {
	Receive(*ReceiveRequest, USIReceiveService_ReceiveServer) error
	mustEmbedUnimplementedUSIReceiveServiceServer()
}

// UnimplementedUSIReceiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUSIReceiveServiceServer struct {
}

func (UnimplementedUSIReceiveServiceServer) Receive(*ReceiveRequest, USIReceiveService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedUSIReceiveServiceServer) mustEmbedUnimplementedUSIReceiveServiceServer() {}

// UnsafeUSIReceiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to USIReceiveServiceServer will
// result in compilation errors.
type UnsafeUSIReceiveServiceServer interface {
	mustEmbedUnimplementedUSIReceiveServiceServer()
}

func RegisterUSIReceiveServiceServer(s grpc.ServiceRegistrar, srv USIReceiveServiceServer) {
	s.RegisterService(&USIReceiveService_ServiceDesc, srv)
}

func _USIReceiveService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(USIReceiveServiceServer).Receive(m, &uSIReceiveServiceReceiveServer{stream})
}

type USIReceiveService_ReceiveServer interface {
	Send(*ReceiveResponse) error
	grpc.ServerStream
}

type uSIReceiveServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *uSIReceiveServiceReceiveServer) Send(m *ReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

// USIReceiveService_ServiceDesc is the grpc.ServiceDesc for USIReceiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var USIReceiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.USIReceiveService",
	HandlerType: (*USIReceiveServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _USIReceiveService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}