// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: proto/mensajes.proto

package proto

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

// IntercambiosClient is the client API for Intercambios service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntercambiosClient interface {
	Notificar(ctx context.Context, in *ContiReq, opts ...grpc.CallOption) (*Confirmacion, error)
	Guardar(ctx context.Context, in *OMSReq, opts ...grpc.CallOption) (*Confirmacion, error)
	Nombres(ctx context.Context, in *ONUReq, opts ...grpc.CallOption) (*ONUResp, error)
	Buscar(ctx context.Context, in *OMSONUReq, opts ...grpc.CallOption) (*DTNResp, error)
}

type intercambiosClient struct {
	cc grpc.ClientConnInterface
}

func NewIntercambiosClient(cc grpc.ClientConnInterface) IntercambiosClient {
	return &intercambiosClient{cc}
}

func (c *intercambiosClient) Notificar(ctx context.Context, in *ContiReq, opts ...grpc.CallOption) (*Confirmacion, error) {
	out := new(Confirmacion)
	err := c.cc.Invoke(ctx, "/grpc.Intercambios/Notificar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intercambiosClient) Guardar(ctx context.Context, in *OMSReq, opts ...grpc.CallOption) (*Confirmacion, error) {
	out := new(Confirmacion)
	err := c.cc.Invoke(ctx, "/grpc.Intercambios/Guardar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intercambiosClient) Nombres(ctx context.Context, in *ONUReq, opts ...grpc.CallOption) (*ONUResp, error) {
	out := new(ONUResp)
	err := c.cc.Invoke(ctx, "/grpc.Intercambios/Nombres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intercambiosClient) Buscar(ctx context.Context, in *OMSONUReq, opts ...grpc.CallOption) (*DTNResp, error) {
	out := new(DTNResp)
	err := c.cc.Invoke(ctx, "/grpc.Intercambios/Buscar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntercambiosServer is the server API for Intercambios service.
// All implementations must embed UnimplementedIntercambiosServer
// for forward compatibility
type IntercambiosServer interface {
	Notificar(context.Context, *ContiReq) (*Confirmacion, error)
	Guardar(context.Context, *OMSReq) (*Confirmacion, error)
	Nombres(context.Context, *ONUReq) (*ONUResp, error)
	Buscar(context.Context, *OMSONUReq) (*DTNResp, error)
	mustEmbedUnimplementedIntercambiosServer()
}

// UnimplementedIntercambiosServer must be embedded to have forward compatible implementations.
type UnimplementedIntercambiosServer struct {
}

func (UnimplementedIntercambiosServer) Notificar(context.Context, *ContiReq) (*Confirmacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notificar not implemented")
}
func (UnimplementedIntercambiosServer) Guardar(context.Context, *OMSReq) (*Confirmacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Guardar not implemented")
}
func (UnimplementedIntercambiosServer) Nombres(context.Context, *ONUReq) (*ONUResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nombres not implemented")
}
func (UnimplementedIntercambiosServer) Buscar(context.Context, *OMSONUReq) (*DTNResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buscar not implemented")
}
func (UnimplementedIntercambiosServer) mustEmbedUnimplementedIntercambiosServer() {}

// UnsafeIntercambiosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntercambiosServer will
// result in compilation errors.
type UnsafeIntercambiosServer interface {
	mustEmbedUnimplementedIntercambiosServer()
}

func RegisterIntercambiosServer(s grpc.ServiceRegistrar, srv IntercambiosServer) {
	s.RegisterService(&Intercambios_ServiceDesc, srv)
}

func _Intercambios_Notificar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambiosServer).Notificar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Intercambios/Notificar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambiosServer).Notificar(ctx, req.(*ContiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Intercambios_Guardar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OMSReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambiosServer).Guardar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Intercambios/Guardar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambiosServer).Guardar(ctx, req.(*OMSReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Intercambios_Nombres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ONUReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambiosServer).Nombres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Intercambios/Nombres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambiosServer).Nombres(ctx, req.(*ONUReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Intercambios_Buscar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OMSONUReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntercambiosServer).Buscar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Intercambios/Buscar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntercambiosServer).Buscar(ctx, req.(*OMSONUReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Intercambios_ServiceDesc is the grpc.ServiceDesc for Intercambios service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Intercambios_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Intercambios",
	HandlerType: (*IntercambiosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notificar",
			Handler:    _Intercambios_Notificar_Handler,
		},
		{
			MethodName: "Guardar",
			Handler:    _Intercambios_Guardar_Handler,
		},
		{
			MethodName: "Nombres",
			Handler:    _Intercambios_Nombres_Handler,
		},
		{
			MethodName: "Buscar",
			Handler:    _Intercambios_Buscar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mensajes.proto",
}
