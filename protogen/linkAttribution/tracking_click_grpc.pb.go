// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: LinkAttributionProto/tracking_click.proto

package linkAttribution

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

// TrackingClickServiceClient is the client API for TrackingClickService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackingClickServiceClient interface {
	TrackingClickCreate(ctx context.Context, in *TrackingClickCreateRequest, opts ...grpc.CallOption) (*TrackingClickUpsertResponse, error)
}

type trackingClickServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingClickServiceClient(cc grpc.ClientConnInterface) TrackingClickServiceClient {
	return &trackingClickServiceClient{cc}
}

func (c *trackingClickServiceClient) TrackingClickCreate(ctx context.Context, in *TrackingClickCreateRequest, opts ...grpc.CallOption) (*TrackingClickUpsertResponse, error) {
	out := new(TrackingClickUpsertResponse)
	err := c.cc.Invoke(ctx, "/chat.TrackingClickService/TrackingClickCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingClickServiceServer is the server API for TrackingClickService service.
// All implementations must embed UnimplementedTrackingClickServiceServer
// for forward compatibility
type TrackingClickServiceServer interface {
	TrackingClickCreate(context.Context, *TrackingClickCreateRequest) (*TrackingClickUpsertResponse, error)
	mustEmbedUnimplementedTrackingClickServiceServer()
}

// UnimplementedTrackingClickServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrackingClickServiceServer struct {
}

func (UnimplementedTrackingClickServiceServer) TrackingClickCreate(context.Context, *TrackingClickCreateRequest) (*TrackingClickUpsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackingClickCreate not implemented")
}
func (UnimplementedTrackingClickServiceServer) mustEmbedUnimplementedTrackingClickServiceServer() {}

// UnsafeTrackingClickServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackingClickServiceServer will
// result in compilation errors.
type UnsafeTrackingClickServiceServer interface {
	mustEmbedUnimplementedTrackingClickServiceServer()
}

func RegisterTrackingClickServiceServer(s grpc.ServiceRegistrar, srv TrackingClickServiceServer) {
	s.RegisterService(&TrackingClickService_ServiceDesc, srv)
}

func _TrackingClickService_TrackingClickCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackingClickCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingClickServiceServer).TrackingClickCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.TrackingClickService/TrackingClickCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingClickServiceServer).TrackingClickCreate(ctx, req.(*TrackingClickCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackingClickService_ServiceDesc is the grpc.ServiceDesc for TrackingClickService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackingClickService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.TrackingClickService",
	HandlerType: (*TrackingClickServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrackingClickCreate",
			Handler:    _TrackingClickService_TrackingClickCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LinkAttributionProto/tracking_click.proto",
}
