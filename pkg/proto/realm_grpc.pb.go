// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: realm.proto

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
	RealmService_ListAvailableRealm_FullMethodName   = "/proto.RealmService/ListAvailableRealm"
	RealmService_ListOwnedRealm_FullMethodName       = "/proto.RealmService/ListOwnedRealm"
	RealmService_ListRealm_FullMethodName            = "/proto.RealmService/ListRealm"
	RealmService_GetRealm_FullMethodName             = "/proto.RealmService/GetRealm"
	RealmService_ListRealmMember_FullMethodName      = "/proto.RealmService/ListRealmMember"
	RealmService_GetRealmMember_FullMethodName       = "/proto.RealmService/GetRealmMember"
	RealmService_CheckRealmMemberPerm_FullMethodName = "/proto.RealmService/CheckRealmMemberPerm"
)

// RealmServiceClient is the client API for RealmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealmServiceClient interface {
	ListAvailableRealm(ctx context.Context, in *LookupUserRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error)
	ListOwnedRealm(ctx context.Context, in *LookupUserRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error)
	ListRealm(ctx context.Context, in *ListRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error)
	GetRealm(ctx context.Context, in *LookupRealmRequest, opts ...grpc.CallOption) (*RealmInfo, error)
	ListRealmMember(ctx context.Context, in *RealmMemberLookupRequest, opts ...grpc.CallOption) (*ListRealmMemberResponse, error)
	GetRealmMember(ctx context.Context, in *RealmMemberLookupRequest, opts ...grpc.CallOption) (*RealmMemberInfo, error)
	CheckRealmMemberPerm(ctx context.Context, in *CheckRealmPermRequest, opts ...grpc.CallOption) (*CheckRealmPermResponse, error)
}

type realmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRealmServiceClient(cc grpc.ClientConnInterface) RealmServiceClient {
	return &realmServiceClient{cc}
}

func (c *realmServiceClient) ListAvailableRealm(ctx context.Context, in *LookupUserRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRealmResponse)
	err := c.cc.Invoke(ctx, RealmService_ListAvailableRealm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) ListOwnedRealm(ctx context.Context, in *LookupUserRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRealmResponse)
	err := c.cc.Invoke(ctx, RealmService_ListOwnedRealm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) ListRealm(ctx context.Context, in *ListRealmRequest, opts ...grpc.CallOption) (*ListRealmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRealmResponse)
	err := c.cc.Invoke(ctx, RealmService_ListRealm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) GetRealm(ctx context.Context, in *LookupRealmRequest, opts ...grpc.CallOption) (*RealmInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealmInfo)
	err := c.cc.Invoke(ctx, RealmService_GetRealm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) ListRealmMember(ctx context.Context, in *RealmMemberLookupRequest, opts ...grpc.CallOption) (*ListRealmMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRealmMemberResponse)
	err := c.cc.Invoke(ctx, RealmService_ListRealmMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) GetRealmMember(ctx context.Context, in *RealmMemberLookupRequest, opts ...grpc.CallOption) (*RealmMemberInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealmMemberInfo)
	err := c.cc.Invoke(ctx, RealmService_GetRealmMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) CheckRealmMemberPerm(ctx context.Context, in *CheckRealmPermRequest, opts ...grpc.CallOption) (*CheckRealmPermResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckRealmPermResponse)
	err := c.cc.Invoke(ctx, RealmService_CheckRealmMemberPerm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealmServiceServer is the server API for RealmService service.
// All implementations must embed UnimplementedRealmServiceServer
// for forward compatibility.
type RealmServiceServer interface {
	ListAvailableRealm(context.Context, *LookupUserRealmRequest) (*ListRealmResponse, error)
	ListOwnedRealm(context.Context, *LookupUserRealmRequest) (*ListRealmResponse, error)
	ListRealm(context.Context, *ListRealmRequest) (*ListRealmResponse, error)
	GetRealm(context.Context, *LookupRealmRequest) (*RealmInfo, error)
	ListRealmMember(context.Context, *RealmMemberLookupRequest) (*ListRealmMemberResponse, error)
	GetRealmMember(context.Context, *RealmMemberLookupRequest) (*RealmMemberInfo, error)
	CheckRealmMemberPerm(context.Context, *CheckRealmPermRequest) (*CheckRealmPermResponse, error)
	mustEmbedUnimplementedRealmServiceServer()
}

// UnimplementedRealmServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRealmServiceServer struct{}

func (UnimplementedRealmServiceServer) ListAvailableRealm(context.Context, *LookupUserRealmRequest) (*ListRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableRealm not implemented")
}
func (UnimplementedRealmServiceServer) ListOwnedRealm(context.Context, *LookupUserRealmRequest) (*ListRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOwnedRealm not implemented")
}
func (UnimplementedRealmServiceServer) ListRealm(context.Context, *ListRealmRequest) (*ListRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRealm not implemented")
}
func (UnimplementedRealmServiceServer) GetRealm(context.Context, *LookupRealmRequest) (*RealmInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealm not implemented")
}
func (UnimplementedRealmServiceServer) ListRealmMember(context.Context, *RealmMemberLookupRequest) (*ListRealmMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRealmMember not implemented")
}
func (UnimplementedRealmServiceServer) GetRealmMember(context.Context, *RealmMemberLookupRequest) (*RealmMemberInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealmMember not implemented")
}
func (UnimplementedRealmServiceServer) CheckRealmMemberPerm(context.Context, *CheckRealmPermRequest) (*CheckRealmPermResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRealmMemberPerm not implemented")
}
func (UnimplementedRealmServiceServer) mustEmbedUnimplementedRealmServiceServer() {}
func (UnimplementedRealmServiceServer) testEmbeddedByValue()                      {}

// UnsafeRealmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealmServiceServer will
// result in compilation errors.
type UnsafeRealmServiceServer interface {
	mustEmbedUnimplementedRealmServiceServer()
}

func RegisterRealmServiceServer(s grpc.ServiceRegistrar, srv RealmServiceServer) {
	// If the following call pancis, it indicates UnimplementedRealmServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RealmService_ServiceDesc, srv)
}

func _RealmService_ListAvailableRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupUserRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).ListAvailableRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_ListAvailableRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).ListAvailableRealm(ctx, req.(*LookupUserRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_ListOwnedRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupUserRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).ListOwnedRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_ListOwnedRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).ListOwnedRealm(ctx, req.(*LookupUserRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_ListRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).ListRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_ListRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).ListRealm(ctx, req.(*ListRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_GetRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).GetRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_GetRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).GetRealm(ctx, req.(*LookupRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_ListRealmMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RealmMemberLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).ListRealmMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_ListRealmMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).ListRealmMember(ctx, req.(*RealmMemberLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_GetRealmMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RealmMemberLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).GetRealmMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_GetRealmMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).GetRealmMember(ctx, req.(*RealmMemberLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_CheckRealmMemberPerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRealmPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).CheckRealmMemberPerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_CheckRealmMemberPerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).CheckRealmMemberPerm(ctx, req.(*CheckRealmPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealmService_ServiceDesc is the grpc.ServiceDesc for RealmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RealmService",
	HandlerType: (*RealmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAvailableRealm",
			Handler:    _RealmService_ListAvailableRealm_Handler,
		},
		{
			MethodName: "ListOwnedRealm",
			Handler:    _RealmService_ListOwnedRealm_Handler,
		},
		{
			MethodName: "ListRealm",
			Handler:    _RealmService_ListRealm_Handler,
		},
		{
			MethodName: "GetRealm",
			Handler:    _RealmService_GetRealm_Handler,
		},
		{
			MethodName: "ListRealmMember",
			Handler:    _RealmService_ListRealmMember_Handler,
		},
		{
			MethodName: "GetRealmMember",
			Handler:    _RealmService_GetRealmMember_Handler,
		},
		{
			MethodName: "CheckRealmMemberPerm",
			Handler:    _RealmService_CheckRealmMemberPerm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realm.proto",
}
