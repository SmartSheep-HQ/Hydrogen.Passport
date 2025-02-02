// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: third_client.proto

package proto

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
	ThirdClientService_GetThirdClient_FullMethodName = "/proto.ThirdClientService/GetThirdClient"
)

// ThirdClientServiceClient is the client API for ThirdClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdClientServiceClient interface {
	GetThirdClient(ctx context.Context, in *GetThirdClientRequest, opts ...grpc.CallOption) (*GetThirdClientResponse, error)
}

type thirdClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdClientServiceClient(cc grpc.ClientConnInterface) ThirdClientServiceClient {
	return &thirdClientServiceClient{cc}
}

func (c *thirdClientServiceClient) GetThirdClient(ctx context.Context, in *GetThirdClientRequest, opts ...grpc.CallOption) (*GetThirdClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetThirdClientResponse)
	err := c.cc.Invoke(ctx, ThirdClientService_GetThirdClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdClientServiceServer is the server API for ThirdClientService service.
// All implementations must embed UnimplementedThirdClientServiceServer
// for forward compatibility.
type ThirdClientServiceServer interface {
	GetThirdClient(context.Context, *GetThirdClientRequest) (*GetThirdClientResponse, error)
	mustEmbedUnimplementedThirdClientServiceServer()
}

// UnimplementedThirdClientServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedThirdClientServiceServer struct{}

func (UnimplementedThirdClientServiceServer) GetThirdClient(context.Context, *GetThirdClientRequest) (*GetThirdClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThirdClient not implemented")
}
func (UnimplementedThirdClientServiceServer) mustEmbedUnimplementedThirdClientServiceServer() {}
func (UnimplementedThirdClientServiceServer) testEmbeddedByValue()                            {}

// UnsafeThirdClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdClientServiceServer will
// result in compilation errors.
type UnsafeThirdClientServiceServer interface {
	mustEmbedUnimplementedThirdClientServiceServer()
}

func RegisterThirdClientServiceServer(s grpc.ServiceRegistrar, srv ThirdClientServiceServer) {
	// If the following call pancis, it indicates UnimplementedThirdClientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ThirdClientService_ServiceDesc, srv)
}

func _ThirdClientService_GetThirdClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThirdClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdClientServiceServer).GetThirdClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThirdClientService_GetThirdClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdClientServiceServer).GetThirdClient(ctx, req.(*GetThirdClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdClientService_ServiceDesc is the grpc.ServiceDesc for ThirdClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ThirdClientService",
	HandlerType: (*ThirdClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThirdClient",
			Handler:    _ThirdClientService_GetThirdClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "third_client.proto",
}
