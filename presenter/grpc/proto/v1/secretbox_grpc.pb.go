// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: secretbox.proto

package v1

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
	SecretBoxService_GetSecret_FullMethodName = "/presenter.grpc.proto.v1.SecretBoxService/GetSecret"
)

// SecretBoxServiceClient is the client API for SecretBoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretBoxServiceClient interface {
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
}

type secretBoxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretBoxServiceClient(cc grpc.ClientConnInterface) SecretBoxServiceClient {
	return &secretBoxServiceClient{cc}
}

func (c *secretBoxServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, SecretBoxService_GetSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretBoxServiceServer is the server API for SecretBoxService service.
// All implementations must embed UnimplementedSecretBoxServiceServer
// for forward compatibility.
type SecretBoxServiceServer interface {
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	mustEmbedUnimplementedSecretBoxServiceServer()
}

// UnimplementedSecretBoxServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSecretBoxServiceServer struct{}

func (UnimplementedSecretBoxServiceServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretBoxServiceServer) mustEmbedUnimplementedSecretBoxServiceServer() {}
func (UnimplementedSecretBoxServiceServer) testEmbeddedByValue()                          {}

// UnsafeSecretBoxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretBoxServiceServer will
// result in compilation errors.
type UnsafeSecretBoxServiceServer interface {
	mustEmbedUnimplementedSecretBoxServiceServer()
}

func RegisterSecretBoxServiceServer(s grpc.ServiceRegistrar, srv SecretBoxServiceServer) {
	// If the following call pancis, it indicates UnimplementedSecretBoxServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SecretBoxService_ServiceDesc, srv)
}

func _SecretBoxService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretBoxServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretBoxService_GetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretBoxServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretBoxService_ServiceDesc is the grpc.ServiceDesc for SecretBoxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretBoxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "presenter.grpc.proto.v1.SecretBoxService",
	HandlerType: (*SecretBoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretBoxService_GetSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secretbox.proto",
}
