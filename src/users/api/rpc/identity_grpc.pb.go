// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity_grpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	presenter "github.com/timoth-y/chainmetric-contracts/src/users/api/presenter"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("identity_grpc.proto", fileDescriptor_98a4d650020a3907) }

var fileDescriptor_98a4d650020a3907 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x10, 0xc7, 0x49, 0x73, 0x48, 0x2c, 0x84, 0x15, 0x2c, 0x22, 0xa4, 0xb2, 0x12, 0xb2, 0x0b, 0x5a,
	0x88, 0xd8, 0x09, 0x57, 0xd8, 0x9e, 0x58, 0x68, 0x23, 0xc9, 0x3a, 0x6e, 0x06, 0xb2, 0x3b, 0x7b,
	0xb3, 0x13, 0xe1, 0xde, 0xd0, 0xd2, 0x37, 0x50, 0xf2, 0x24, 0x62, 0x62, 0xc4, 0xea, 0x48, 0x39,
	0x33, 0xff, 0x0f, 0x7e, 0x93, 0x1f, 0xe3, 0x0b, 0x04, 0x41, 0xd9, 0x3d, 0x3b, 0x8e, 0x56, 0x47,
	0x26, 0x21, 0x75, 0x66, 0xdb, 0x1a, 0x83, 0x07, 0x61, 0xb4, 0x7a, 0x16, 0x4c, 0x37, 0x6d, 0x29,
	0x08, 0xd7, 0x56, 0x52, 0x71, 0xd8, 0x27, 0xe0, 0x34, 0xad, 0x8b, 0x53, 0x47, 0xe4, 0x3a, 0x30,
	0xe3, 0xd4, 0xf4, 0xaf, 0x06, 0x7c, 0x9c, 0x3d, 0x17, 0x9f, 0x59, 0x7e, 0x74, 0xf7, 0x1b, 0x73,
	0x0f, 0xfc, 0x86, 0x16, 0xd4, 0x36, 0x3f, 0x60, 0x70, 0x98, 0x04, 0x58, 0x5d, 0xeb, 0x3d, 0x85,
	0x91, 0x21, 0x41, 0x10, 0x60, 0xbd, 0x19, 0xf5, 0x5c, 0x0b, 0x52, 0xd8, 0xc0, 0xb6, 0x87, 0x24,
	0xc5, 0xf9, 0x32, 0xeb, 0x43, 0x02, 0x56, 0x8f, 0xf9, 0x0a, 0x02, 0x53, 0xd7, 0xa9, 0xab, 0x65,
	0xae, 0xf5, 0xa8, 0xf6, 0x10, 0x64, 0xae, 0x3b, 0xd1, 0x13, 0xa7, 0x9e, 0x39, 0xf5, 0xfa, 0x87,
	0xf3, 0xf6, 0xe6, 0x7d, 0x28, 0xb3, 0x8f, 0xa1, 0xcc, 0xbe, 0x86, 0x32, 0x7b, 0xaa, 0x1c, 0x4a,
	0xdb, 0x37, 0xda, 0x92, 0x37, 0x82, 0x9e, 0xa4, 0xad, 0x76, 0xe6, 0x5f, 0x63, 0xf5, 0xf7, 0x44,
	0xc3, 0xd1, 0x36, 0xab, 0x31, 0xec, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x5e, 0x64, 0x17,
	0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityServiceClient interface {
	Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.User, error)
	Enroll(ctx context.Context, in *presenter.EnrollmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityServiceClient struct {
	cc *grpc.ClientConn
}

func NewIdentityServiceClient(cc *grpc.ClientConn) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) Register(ctx context.Context, in *presenter.RegistrationRequest, opts ...grpc.CallOption) (*presenter.User, error) {
	out := new(presenter.User)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.proto.contracts.IdentityService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) Enroll(ctx context.Context, in *presenter.EnrollmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chainmetric.identity.proto.contracts.IdentityService/enroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
type IdentityServiceServer interface {
	Register(context.Context, *presenter.RegistrationRequest) (*presenter.User, error)
	Enroll(context.Context, *presenter.EnrollmentRequest) (*emptypb.Empty, error)
}

// UnimplementedIdentityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (*UnimplementedIdentityServiceServer) Register(ctx context.Context, req *presenter.RegistrationRequest) (*presenter.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedIdentityServiceServer) Enroll(ctx context.Context, req *presenter.EnrollmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enroll not implemented")
}

func RegisterIdentityServiceServer(s *grpc.Server, srv IdentityServiceServer) {
	s.RegisterService(&_IdentityService_serviceDesc, srv)
}

func _IdentityService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.proto.contracts.IdentityService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Register(ctx, req.(*presenter.RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_Enroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(presenter.EnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Enroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmetric.identity.proto.contracts.IdentityService/Enroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Enroll(ctx, req.(*presenter.EnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainmetric.identity.proto.contracts.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _IdentityService_Register_Handler,
		},
		{
			MethodName: "enroll",
			Handler:    _IdentityService_Enroll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity_grpc.proto",
}
