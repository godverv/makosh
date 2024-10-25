// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: grpc/makosh_api.proto

package example_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MakoshAPI_Version_FullMethodName = "/makosh_api.makoshAPI/Version"
)

// MakoshAPIClient is the client API for MakoshAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MakoshAPIClient interface {
	Version(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type makoshAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMakoshAPIClient(cc grpc.ClientConnInterface) MakoshAPIClient {
	return &makoshAPIClient{cc}
}

func (c *makoshAPIClient) Version(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, MakoshAPI_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MakoshAPIServer is the server API for MakoshAPI service.
// All implementations must embed UnimplementedMakoshAPIServer
// for forward compatibility.
type MakoshAPIServer interface {
	Version(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedMakoshAPIServer()
}

// UnimplementedMakoshAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMakoshAPIServer struct{}

func (UnimplementedMakoshAPIServer) Version(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedMakoshAPIServer) mustEmbedUnimplementedMakoshAPIServer() {}
func (UnimplementedMakoshAPIServer) testEmbeddedByValue()                   {}

// UnsafeMakoshAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MakoshAPIServer will
// result in compilation errors.
type UnsafeMakoshAPIServer interface {
	mustEmbedUnimplementedMakoshAPIServer()
}

func RegisterMakoshAPIServer(s grpc.ServiceRegistrar, srv MakoshAPIServer) {
	// If the following call pancis, it indicates UnimplementedMakoshAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MakoshAPI_ServiceDesc, srv)
}

func _MakoshAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MakoshAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MakoshAPI_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MakoshAPIServer).Version(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MakoshAPI_ServiceDesc is the grpc.ServiceDesc for MakoshAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MakoshAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "makosh_api.makoshAPI",
	HandlerType: (*MakoshAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _MakoshAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/makosh_api.proto",
}
