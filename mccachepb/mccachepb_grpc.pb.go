// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: mccachepb.proto

package mccachepb

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
	McCache_Get_FullMethodName = "/geecachepb.McCache/Get"
)

// McCacheClient is the client API for McCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type McCacheClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type mcCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewMcCacheClient(cc grpc.ClientConnInterface) McCacheClient {
	return &mcCacheClient{cc}
}

func (c *mcCacheClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, McCache_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// McCacheServer is the server API for McCache service.
// All implementations must embed UnimplementedMcCacheServer
// for forward compatibility
type McCacheServer interface {
	Get(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMcCacheServer()
}

// UnimplementedMcCacheServer must be embedded to have forward compatible implementations.
type UnimplementedMcCacheServer struct {
}

func (UnimplementedMcCacheServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMcCacheServer) mustEmbedUnimplementedMcCacheServer() {}

// UnsafeMcCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to McCacheServer will
// result in compilation errors.
type UnsafeMcCacheServer interface {
	mustEmbedUnimplementedMcCacheServer()
}

func RegisterMcCacheServer(s grpc.ServiceRegistrar, srv McCacheServer) {
	s.RegisterService(&McCache_ServiceDesc, srv)
}

func _McCache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McCacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: McCache_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McCacheServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// McCache_ServiceDesc is the grpc.ServiceDesc for McCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var McCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geecachepb.McCache",
	HandlerType: (*McCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _McCache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mccachepb.proto",
}
