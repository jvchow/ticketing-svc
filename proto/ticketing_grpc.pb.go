// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: proto/ticketing.proto

package train

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

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error)
	GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Receipt, error)
	ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/train.TicketService/PurchaseTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/train.TicketService/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error) {
	out := new(SeatResponse)
	err := c.cc.Invoke(ctx, "/train.TicketService/ViewSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/train.TicketService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/train.TicketService/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*Receipt, error)
	GetReceipt(context.Context, *UserRequest) (*Receipt, error)
	ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error)
	RemoveUser(context.Context, *UserRequest) (*StatusResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*StatusResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetReceipt(context.Context, *UserRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTicketServiceServer) ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeats not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *UserRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train.TicketService/PurchaseTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train.TicketService/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ViewSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ViewSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train.TicketService/ViewSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ViewSeats(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train.TicketService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train.TicketService/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "train.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TicketService_GetReceipt_Handler,
		},
		{
			MethodName: "ViewSeats",
			Handler:    _TicketService_ViewSeats_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticketing.proto",
}
