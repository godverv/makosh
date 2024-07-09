// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: grpc/makosh-be_api.proto

package makosh_be

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MakoshBeAPI_Version_FullMethodName         = "/makosh_be_api.makoshBeAPI/Version"
	MakoshBeAPI_ListEndpoints_FullMethodName   = "/makosh_be_api.makoshBeAPI/ListEndpoints"
	MakoshBeAPI_UpsertEndpoints_FullMethodName = "/makosh_be_api.makoshBeAPI/UpsertEndpoints"
)

// MakoshBeAPIClient is the client API for MakoshBeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MakoshBeAPIClient interface {
	Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error)
	ListEndpoints(ctx context.Context, in *ListEndpoints_Request, opts ...grpc.CallOption) (*ListEndpoints_Response, error)
	UpsertEndpoints(ctx context.Context, in *UpsertEndpoints_Request, opts ...grpc.CallOption) (*UpsertEndpoints_Response, error)
}

type makoshBeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMakoshBeAPIClient(cc grpc.ClientConnInterface) MakoshBeAPIClient {
	return &makoshBeAPIClient{cc}
}

func (c *makoshBeAPIClient) Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Version_Response)
	err := c.cc.Invoke(ctx, MakoshBeAPI_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *makoshBeAPIClient) ListEndpoints(ctx context.Context, in *ListEndpoints_Request, opts ...grpc.CallOption) (*ListEndpoints_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEndpoints_Response)
	err := c.cc.Invoke(ctx, MakoshBeAPI_ListEndpoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *makoshBeAPIClient) UpsertEndpoints(ctx context.Context, in *UpsertEndpoints_Request, opts ...grpc.CallOption) (*UpsertEndpoints_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertEndpoints_Response)
	err := c.cc.Invoke(ctx, MakoshBeAPI_UpsertEndpoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MakoshBeAPIServer is the server API for MakoshBeAPI service.
// All implementations must embed UnimplementedMakoshBeAPIServer
// for forward compatibility
type MakoshBeAPIServer interface {
	Version(context.Context, *Version_Request) (*Version_Response, error)
	ListEndpoints(context.Context, *ListEndpoints_Request) (*ListEndpoints_Response, error)
	UpsertEndpoints(context.Context, *UpsertEndpoints_Request) (*UpsertEndpoints_Response, error)
	mustEmbedUnimplementedMakoshBeAPIServer()
}

// UnimplementedMakoshBeAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMakoshBeAPIServer struct {
}

func (UnimplementedMakoshBeAPIServer) Version(context.Context, *Version_Request) (*Version_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedMakoshBeAPIServer) ListEndpoints(context.Context, *ListEndpoints_Request) (*ListEndpoints_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEndpoints not implemented")
}
func (UnimplementedMakoshBeAPIServer) UpsertEndpoints(context.Context, *UpsertEndpoints_Request) (*UpsertEndpoints_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertEndpoints not implemented")
}
func (UnimplementedMakoshBeAPIServer) mustEmbedUnimplementedMakoshBeAPIServer() {}

// UnsafeMakoshBeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MakoshBeAPIServer will
// result in compilation errors.
type UnsafeMakoshBeAPIServer interface {
	mustEmbedUnimplementedMakoshBeAPIServer()
}

func RegisterMakoshBeAPIServer(s grpc.ServiceRegistrar, srv MakoshBeAPIServer) {
	s.RegisterService(&MakoshBeAPI_ServiceDesc, srv)
}

func _MakoshBeAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Version_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MakoshBeAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MakoshBeAPI_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MakoshBeAPIServer).Version(ctx, req.(*Version_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MakoshBeAPI_ListEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndpoints_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MakoshBeAPIServer).ListEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MakoshBeAPI_ListEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MakoshBeAPIServer).ListEndpoints(ctx, req.(*ListEndpoints_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MakoshBeAPI_UpsertEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertEndpoints_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MakoshBeAPIServer).UpsertEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MakoshBeAPI_UpsertEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MakoshBeAPIServer).UpsertEndpoints(ctx, req.(*UpsertEndpoints_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MakoshBeAPI_ServiceDesc is the grpc.ServiceDesc for MakoshBeAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MakoshBeAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "makosh_be_api.makoshBeAPI",
	HandlerType: (*MakoshBeAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _MakoshBeAPI_Version_Handler,
		},
		{
			MethodName: "ListEndpoints",
			Handler:    _MakoshBeAPI_ListEndpoints_Handler,
		},
		{
			MethodName: "UpsertEndpoints",
			Handler:    _MakoshBeAPI_UpsertEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/makosh-be_api.proto",
}
