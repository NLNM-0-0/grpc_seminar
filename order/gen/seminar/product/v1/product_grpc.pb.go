// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: seminar/product/v1/product.proto

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
	ProductService_PostProduct_FullMethodName      = "/seminar.product.v1.ProductService/PostProduct"
	ProductService_GetProduct_FullMethodName       = "/seminar.product.v1.ProductService/GetProduct"
	ProductService_GetProductByIds_FullMethodName  = "/seminar.product.v1.ProductService/GetProductByIds"
	ProductService_RateProductByIds_FullMethodName = "/seminar.product.v1.ProductService/RateProductByIds"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetProductByIds(ctx context.Context, in *GetProductsByIdsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetProductsByIdsResponse], error)
	RateProductByIds(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateProductByIdsRequest, RateProductByIdsResponse], error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostProductResponse)
	err := c.cc.Invoke(ctx, ProductService_PostProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductByIds(ctx context.Context, in *GetProductsByIdsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetProductsByIdsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], ProductService_GetProductByIds_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetProductsByIdsRequest, GetProductsByIdsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_GetProductByIdsClient = grpc.ServerStreamingClient[GetProductsByIdsResponse]

func (c *productServiceClient) RateProductByIds(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateProductByIdsRequest, RateProductByIdsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], ProductService_RateProductByIds_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RateProductByIdsRequest, RateProductByIdsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_RateProductByIdsClient = grpc.BidiStreamingClient[RateProductByIdsRequest, RateProductByIdsResponse]

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility.
type ProductServiceServer interface {
	PostProduct(context.Context, *PostProductRequest) (*PostProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetProductByIds(*GetProductsByIdsRequest, grpc.ServerStreamingServer[GetProductsByIdsResponse]) error
	RateProductByIds(grpc.BidiStreamingServer[RateProductByIdsRequest, RateProductByIdsResponse]) error
}

// UnimplementedProductServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) PostProduct(context.Context, *PostProductRequest) (*PostProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProductByIds(*GetProductsByIdsRequest, grpc.ServerStreamingServer[GetProductsByIdsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetProductByIds not implemented")
}
func (UnimplementedProductServiceServer) RateProductByIds(grpc.BidiStreamingServer[RateProductByIdsRequest, RateProductByIdsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method RateProductByIds not implemented")
}
func (UnimplementedProductServiceServer) testEmbeddedByValue() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_PostProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).PostProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_PostProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).PostProduct(ctx, req.(*PostProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductByIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProductsByIdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).GetProductByIds(m, &grpc.GenericServerStream[GetProductsByIdsRequest, GetProductsByIdsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_GetProductByIdsServer = grpc.ServerStreamingServer[GetProductsByIdsResponse]

func _ProductService_RateProductByIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).RateProductByIds(&grpc.GenericServerStream[RateProductByIdsRequest, RateProductByIdsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_RateProductByIdsServer = grpc.BidiStreamingServer[RateProductByIdsRequest, RateProductByIdsResponse]

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seminar.product.v1.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostProduct",
			Handler:    _ProductService_PostProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProductByIds",
			Handler:       _ProductService_GetProductByIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RateProductByIds",
			Handler:       _ProductService_RateProductByIds_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "seminar/product/v1/product.proto",
}
