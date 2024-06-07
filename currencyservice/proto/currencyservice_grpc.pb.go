// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/currencyservice.proto

package microshopping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CurrencyService_GetSupportedCurrencies_FullMethodName = "/microshopping.CurrencyService/GetSupportedCurrencies"
	CurrencyService_Convert_FullMethodName                = "/microshopping.CurrencyService/Convert"
)

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	GetSupportedCurrencies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) GetSupportedCurrencies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSupportedCurrenciesResponse)
	err := c.cc.Invoke(ctx, CurrencyService_GetSupportedCurrencies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Money)
	err := c.cc.Invoke(ctx, CurrencyService_Convert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations should embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	GetSupportedCurrencies(context.Context, *emptypb.Empty) (*GetSupportedCurrenciesResponse, error)
	Convert(context.Context, *CurrencyConversionRequest) (*Money, error)
}

// UnimplementedCurrencyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) GetSupportedCurrencies(context.Context, *emptypb.Empty) (*GetSupportedCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedCurrencies not implemented")
}
func (UnimplementedCurrencyServiceServer) Convert(context.Context, *CurrencyConversionRequest) (*Money, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_GetSupportedCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrencyService_GetSupportedCurrencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyConversionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrencyService_Convert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Convert(ctx, req.(*CurrencyConversionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microshopping.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportedCurrencies",
			Handler:    _CurrencyService_GetSupportedCurrencies_Handler,
		},
		{
			MethodName: "Convert",
			Handler:    _CurrencyService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/currencyservice.proto",
}
