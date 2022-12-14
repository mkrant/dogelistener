// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
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

// DogeServerClient is the client API for DogeServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DogeServerClient interface {
	// Sends a greeting
	Connect(ctx context.Context, opts ...grpc.CallOption) (DogeServer_ConnectClient, error)
}

type dogeServerClient struct {
	cc grpc.ClientConnInterface
}

func NewDogeServerClient(cc grpc.ClientConnInterface) DogeServerClient {
	return &dogeServerClient{cc}
}

func (c *dogeServerClient) Connect(ctx context.Context, opts ...grpc.CallOption) (DogeServer_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &DogeServer_ServiceDesc.Streams[0], "/DogeServer/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &dogeServerConnectClient{stream}
	return x, nil
}

type DogeServer_ConnectClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type dogeServerConnectClient struct {
	grpc.ClientStream
}

func (x *dogeServerConnectClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dogeServerConnectClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DogeServerServer is the server API for DogeServer service.
// All implementations must embed UnimplementedDogeServerServer
// for forward compatibility
type DogeServerServer interface {
	// Sends a greeting
	Connect(DogeServer_ConnectServer) error
	mustEmbedUnimplementedDogeServerServer()
}

// UnimplementedDogeServerServer must be embedded to have forward compatible implementations.
type UnimplementedDogeServerServer struct {
}

func (UnimplementedDogeServerServer) Connect(DogeServer_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedDogeServerServer) mustEmbedUnimplementedDogeServerServer() {}

// UnsafeDogeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DogeServerServer will
// result in compilation errors.
type UnsafeDogeServerServer interface {
	mustEmbedUnimplementedDogeServerServer()
}

func RegisterDogeServerServer(s grpc.ServiceRegistrar, srv DogeServerServer) {
	s.RegisterService(&DogeServer_ServiceDesc, srv)
}

func _DogeServer_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DogeServerServer).Connect(&dogeServerConnectServer{stream})
}

type DogeServer_ConnectServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type dogeServerConnectServer struct {
	grpc.ServerStream
}

func (x *dogeServerConnectServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dogeServerConnectServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DogeServer_ServiceDesc is the grpc.ServiceDesc for DogeServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DogeServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DogeServer",
	HandlerType: (*DogeServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _DogeServer_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
