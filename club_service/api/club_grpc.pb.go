// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ClubServiceClient is the client API for ClubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClubServiceClient interface {
	GetById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Club, error)
}

type clubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClubServiceClient(cc grpc.ClientConnInterface) ClubServiceClient {
	return &clubServiceClient{cc}
}

func (c *clubServiceClient) GetById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Club, error) {
	out := new(Club)
	err := c.cc.Invoke(ctx, "/api.ClubService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServiceServer is the server API for ClubService service.
// All implementations should embed UnimplementedClubServiceServer
// for forward compatibility
type ClubServiceServer interface {
	GetById(context.Context, *Request) (*Club, error)
}

// UnimplementedClubServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClubServiceServer struct {
}

func (UnimplementedClubServiceServer) GetById(context.Context, *Request) (*Club, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

// UnsafeClubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClubServiceServer will
// result in compilation errors.
type UnsafeClubServiceServer interface {
	mustEmbedUnimplementedClubServiceServer()
}

func RegisterClubServiceServer(s grpc.ServiceRegistrar, srv ClubServiceServer) {
	s.RegisterService(&ClubService_ServiceDesc, srv)
}

func _ClubService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ClubService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).GetById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ClubService_ServiceDesc is the grpc.ServiceDesc for ClubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ClubService",
	HandlerType: (*ClubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _ClubService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "club.proto",
}
