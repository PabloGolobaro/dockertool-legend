// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: stats.proto

package protoStats

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

// ContainerStatsServiceClient is the client API for ContainerStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerStatsServiceClient interface {
	GetStatsStream(ctx context.Context, in *GetStatsMessage, opts ...grpc.CallOption) (ContainerStatsService_GetStatsStreamClient, error)
}

type containerStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerStatsServiceClient(cc grpc.ClientConnInterface) ContainerStatsServiceClient {
	return &containerStatsServiceClient{cc}
}

func (c *containerStatsServiceClient) GetStatsStream(ctx context.Context, in *GetStatsMessage, opts ...grpc.CallOption) (ContainerStatsService_GetStatsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContainerStatsService_ServiceDesc.Streams[0], "/protoStats.ContainerStatsService/getStatsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerStatsServiceGetStatsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContainerStatsService_GetStatsStreamClient interface {
	Recv() (*Stats, error)
	grpc.ClientStream
}

type containerStatsServiceGetStatsStreamClient struct {
	grpc.ClientStream
}

func (x *containerStatsServiceGetStatsStreamClient) Recv() (*Stats, error) {
	m := new(Stats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContainerStatsServiceServer is the server API for ContainerStatsService service.
// All implementations must embed UnimplementedContainerStatsServiceServer
// for forward compatibility
type ContainerStatsServiceServer interface {
	GetStatsStream(*GetStatsMessage, ContainerStatsService_GetStatsStreamServer) error
	mustEmbedUnimplementedContainerStatsServiceServer()
}

// UnimplementedContainerStatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContainerStatsServiceServer struct {
}

func (UnimplementedContainerStatsServiceServer) GetStatsStream(*GetStatsMessage, ContainerStatsService_GetStatsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatsStream not implemented")
}
func (UnimplementedContainerStatsServiceServer) mustEmbedUnimplementedContainerStatsServiceServer() {}

// UnsafeContainerStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerStatsServiceServer will
// result in compilation errors.
type UnsafeContainerStatsServiceServer interface {
	mustEmbedUnimplementedContainerStatsServiceServer()
}

func RegisterContainerStatsServiceServer(s grpc.ServiceRegistrar, srv ContainerStatsServiceServer) {
	s.RegisterService(&ContainerStatsService_ServiceDesc, srv)
}

func _ContainerStatsService_GetStatsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStatsMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerStatsServiceServer).GetStatsStream(m, &containerStatsServiceGetStatsStreamServer{stream})
}

type ContainerStatsService_GetStatsStreamServer interface {
	Send(*Stats) error
	grpc.ServerStream
}

type containerStatsServiceGetStatsStreamServer struct {
	grpc.ServerStream
}

func (x *containerStatsServiceGetStatsStreamServer) Send(m *Stats) error {
	return x.ServerStream.SendMsg(m)
}

// ContainerStatsService_ServiceDesc is the grpc.ServiceDesc for ContainerStatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerStatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoStats.ContainerStatsService",
	HandlerType: (*ContainerStatsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getStatsStream",
			Handler:       _ContainerStatsService_GetStatsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stats.proto",
}
