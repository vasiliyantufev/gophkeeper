// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: gophkeeper.proto

package gophkeeper

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
	Gophkeeper_HandleAuthentication_FullMethodName       = "/api.Gophkeeper/HandleAuthentication"
	Gophkeeper_HandleRegistration_FullMethodName         = "/api.Gophkeeper/HandleRegistration"
	Gophkeeper_HandleCreateText_FullMethodName           = "/api.Gophkeeper/HandleCreateText"
	Gophkeeper_HandleGetListText_FullMethodName          = "/api.Gophkeeper/HandleGetListText"
	Gophkeeper_HandleGetNodeText_FullMethodName          = "/api.Gophkeeper/HandleGetNodeText"
	Gophkeeper_HandlePing_FullMethodName                 = "/api.Gophkeeper/HandlePing"
	Gophkeeper_HandleUserExist_FullMethodName            = "/api.Gophkeeper/HandleUserExist"
	Gophkeeper_HandleCreateCard_FullMethodName           = "/api.Gophkeeper/HandleCreateCard"
	Gophkeeper_HandleGetNodeCard_FullMethodName          = "/api.Gophkeeper/HandleGetNodeCard"
	Gophkeeper_HandleGetListCard_FullMethodName          = "/api.Gophkeeper/HandleGetListCard"
	Gophkeeper_HandleCreateLoginPassword_FullMethodName  = "/api.Gophkeeper/HandleCreateLoginPassword"
	Gophkeeper_HandleGetNodeLoginPassword_FullMethodName = "/api.Gophkeeper/HandleGetNodeLoginPassword"
	Gophkeeper_HandleGetListLoginPassword_FullMethodName = "/api.Gophkeeper/HandleGetListLoginPassword"
)

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	HandleAuthentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error)
	HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error)
	HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error)
	HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	HandleUserExist(ctx context.Context, in *UserExistRequest, opts ...grpc.CallOption) (*UserExistResponse, error)
	HandleCreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error)
	HandleGetNodeCard(ctx context.Context, in *GetNodeCardRequest, opts ...grpc.CallOption) (*GetNodeCardResponse, error)
	HandleGetListCard(ctx context.Context, in *GetListCardRequest, opts ...grpc.CallOption) (*GetListCardResponse, error)
	HandleCreateLoginPassword(ctx context.Context, in *CreateLoginPasswordRequest, opts ...grpc.CallOption) (*CreateLoginPasswordResponse, error)
	HandleGetNodeLoginPassword(ctx context.Context, in *GetNodeLoginPasswordRequest, opts ...grpc.CallOption) (*GetNodeLoginPasswordResponse, error)
	HandleGetListLoginPassword(ctx context.Context, in *GetListLoginPasswordRequest, opts ...grpc.CallOption) (*GetListLoginPasswordResponse, error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) HandleAuthentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error) {
	out := new(CreateTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleCreateText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error) {
	out := new(GetListTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetListText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error) {
	out := new(GetNodeTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetNodeText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandlePing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleUserExist(ctx context.Context, in *UserExistRequest, opts ...grpc.CallOption) (*UserExistResponse, error) {
	out := new(UserExistResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleUserExist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleCreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error) {
	out := new(CreateCardResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleCreateCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetNodeCard(ctx context.Context, in *GetNodeCardRequest, opts ...grpc.CallOption) (*GetNodeCardResponse, error) {
	out := new(GetNodeCardResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetNodeCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetListCard(ctx context.Context, in *GetListCardRequest, opts ...grpc.CallOption) (*GetListCardResponse, error) {
	out := new(GetListCardResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetListCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleCreateLoginPassword(ctx context.Context, in *CreateLoginPasswordRequest, opts ...grpc.CallOption) (*CreateLoginPasswordResponse, error) {
	out := new(CreateLoginPasswordResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleCreateLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetNodeLoginPassword(ctx context.Context, in *GetNodeLoginPasswordRequest, opts ...grpc.CallOption) (*GetNodeLoginPasswordResponse, error) {
	out := new(GetNodeLoginPasswordResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetNodeLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetListLoginPassword(ctx context.Context, in *GetListLoginPasswordRequest, opts ...grpc.CallOption) (*GetListLoginPasswordResponse, error) {
	out := new(GetListLoginPasswordResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetListLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility
type GophkeeperServer interface {
	HandleAuthentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error)
	HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error)
	HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error)
	HandlePing(context.Context, *PingRequest) (*PingResponse, error)
	HandleUserExist(context.Context, *UserExistRequest) (*UserExistResponse, error)
	HandleCreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error)
	HandleGetNodeCard(context.Context, *GetNodeCardRequest) (*GetNodeCardResponse, error)
	HandleGetListCard(context.Context, *GetListCardRequest) (*GetListCardResponse, error)
	HandleCreateLoginPassword(context.Context, *CreateLoginPasswordRequest) (*CreateLoginPasswordResponse, error)
	HandleGetNodeLoginPassword(context.Context, *GetNodeLoginPasswordRequest) (*GetNodeLoginPasswordResponse, error)
	HandleGetListLoginPassword(context.Context, *GetListLoginPasswordRequest) (*GetListLoginPasswordResponse, error)
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophkeeperServer struct {
}

func (UnimplementedGophkeeperServer) HandleAuthentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAuthentication not implemented")
}
func (UnimplementedGophkeeperServer) HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRegistration not implemented")
}
func (UnimplementedGophkeeperServer) HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetListText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetNodeText not implemented")
}
func (UnimplementedGophkeeperServer) HandlePing(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePing not implemented")
}
func (UnimplementedGophkeeperServer) HandleUserExist(context.Context, *UserExistRequest) (*UserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUserExist not implemented")
}
func (UnimplementedGophkeeperServer) HandleCreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateCard not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetNodeCard(context.Context, *GetNodeCardRequest) (*GetNodeCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetNodeCard not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetListCard(context.Context, *GetListCardRequest) (*GetListCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetListCard not implemented")
}
func (UnimplementedGophkeeperServer) HandleCreateLoginPassword(context.Context, *CreateLoginPasswordRequest) (*CreateLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateLoginPassword not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetNodeLoginPassword(context.Context, *GetNodeLoginPasswordRequest) (*GetNodeLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetNodeLoginPassword not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetListLoginPassword(context.Context, *GetListLoginPasswordRequest) (*GetListLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetListLoginPassword not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_HandleAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleAuthentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleRegistration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleCreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleCreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleCreateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleCreateText(ctx, req.(*CreateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetListText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetListText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetListText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetListText(ctx, req.(*GetListTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetNodeText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetNodeText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, req.(*GetNodeTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandlePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandlePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandlePing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandlePing(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleUserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleUserExist(ctx, req.(*UserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleCreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleCreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleCreateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleCreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetNodeCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetNodeCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetNodeCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetNodeCard(ctx, req.(*GetNodeCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetListCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetListCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetListCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetListCard(ctx, req.(*GetListCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleCreateLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleCreateLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleCreateLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleCreateLoginPassword(ctx, req.(*CreateLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetNodeLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetNodeLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetNodeLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetNodeLoginPassword(ctx, req.(*GetNodeLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetListLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetListLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetListLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetListLoginPassword(ctx, req.(*GetListLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleAuthentication",
			Handler:    _Gophkeeper_HandleAuthentication_Handler,
		},
		{
			MethodName: "HandleRegistration",
			Handler:    _Gophkeeper_HandleRegistration_Handler,
		},
		{
			MethodName: "HandleCreateText",
			Handler:    _Gophkeeper_HandleCreateText_Handler,
		},
		{
			MethodName: "HandleGetListText",
			Handler:    _Gophkeeper_HandleGetListText_Handler,
		},
		{
			MethodName: "HandleGetNodeText",
			Handler:    _Gophkeeper_HandleGetNodeText_Handler,
		},
		{
			MethodName: "HandlePing",
			Handler:    _Gophkeeper_HandlePing_Handler,
		},
		{
			MethodName: "HandleUserExist",
			Handler:    _Gophkeeper_HandleUserExist_Handler,
		},
		{
			MethodName: "HandleCreateCard",
			Handler:    _Gophkeeper_HandleCreateCard_Handler,
		},
		{
			MethodName: "HandleGetNodeCard",
			Handler:    _Gophkeeper_HandleGetNodeCard_Handler,
		},
		{
			MethodName: "HandleGetListCard",
			Handler:    _Gophkeeper_HandleGetListCard_Handler,
		},
		{
			MethodName: "HandleCreateLoginPassword",
			Handler:    _Gophkeeper_HandleCreateLoginPassword_Handler,
		},
		{
			MethodName: "HandleGetNodeLoginPassword",
			Handler:    _Gophkeeper_HandleGetNodeLoginPassword_Handler,
		},
		{
			MethodName: "HandleGetListLoginPassword",
			Handler:    _Gophkeeper_HandleGetListLoginPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper.proto",
}
