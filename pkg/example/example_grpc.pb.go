// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: example.proto

package example

import (
	context "context"
	input "github.com/soustify/data-gateway-buffer-go/pkg/input"
	output "github.com/soustify/data-gateway-buffer-go/pkg/output"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExampleService_Paginate_FullMethodName = "/example.ExampleService/Paginate"
	ExampleService_Count_FullMethodName    = "/example.ExampleService/Count"
	ExampleService_Create_FullMethodName   = "/example.ExampleService/Create"
	ExampleService_Update_FullMethodName   = "/example.ExampleService/Update"
	ExampleService_Inactive_FullMethodName = "/example.ExampleService/Inactive"
	ExampleService_Active_FullMethodName   = "/example.ExampleService/Active"
	ExampleService_FindOne_FullMethodName  = "/example.ExampleService/FindOne"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	Paginate(ctx context.Context, in *input.PaginationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExampleResponse], error)
	Count(ctx context.Context, in *input.FilteredRequest, opts ...grpc.CallOption) (*output.CountResponse, error)
	Create(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse], error)
	Update(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse], error)
	Inactive(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse], error)
	Active(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse], error)
	FindOne(ctx context.Context, in *input.UUIDRequest, opts ...grpc.CallOption) (*ExampleResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Paginate(ctx context.Context, in *input.PaginationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExampleResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_Paginate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[input.PaginationRequest, ExampleResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_PaginateClient = grpc.ServerStreamingClient[ExampleResponse]

func (c *exampleServiceClient) Count(ctx context.Context, in *input.FilteredRequest, opts ...grpc.CallOption) (*output.CountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(output.CountResponse)
	err := c.cc.Invoke(ctx, ExampleService_Count_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) Create(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_Create_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExampleRequest, output.PersistenceDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_CreateClient = grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse]

func (c *exampleServiceClient) Update(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[2], ExampleService_Update_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExampleRequest, output.PersistenceDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_UpdateClient = grpc.ClientStreamingClient[ExampleRequest, output.PersistenceDataResponse]

func (c *exampleServiceClient) Inactive(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[3], ExampleService_Inactive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[input.UUIDRequest, output.StatusDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_InactiveClient = grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse]

func (c *exampleServiceClient) Active(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[4], ExampleService_Active_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[input.UUIDRequest, output.StatusDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ActiveClient = grpc.ClientStreamingClient[input.UUIDRequest, output.StatusDataResponse]

func (c *exampleServiceClient) FindOne(ctx context.Context, in *input.UUIDRequest, opts ...grpc.CallOption) (*ExampleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleResponse)
	err := c.cc.Invoke(ctx, ExampleService_FindOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility.
type ExampleServiceServer interface {
	Paginate(*input.PaginationRequest, grpc.ServerStreamingServer[ExampleResponse]) error
	Count(context.Context, *input.FilteredRequest) (*output.CountResponse, error)
	Create(grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]) error
	Update(grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]) error
	Inactive(grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]) error
	Active(grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]) error
	FindOne(context.Context, *input.UUIDRequest) (*ExampleResponse, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExampleServiceServer struct{}

func (UnimplementedExampleServiceServer) Paginate(*input.PaginationRequest, grpc.ServerStreamingServer[ExampleResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Paginate not implemented")
}
func (UnimplementedExampleServiceServer) Count(context.Context, *input.FilteredRequest) (*output.CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedExampleServiceServer) Create(grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedExampleServiceServer) Update(grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedExampleServiceServer) Inactive(grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Inactive not implemented")
}
func (UnimplementedExampleServiceServer) Active(grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Active not implemented")
}
func (UnimplementedExampleServiceServer) FindOne(context.Context, *input.UUIDRequest) (*ExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}
func (UnimplementedExampleServiceServer) testEmbeddedByValue()                        {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	// If the following call pancis, it indicates UnimplementedExampleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_Paginate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(input.PaginationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).Paginate(m, &grpc.GenericServerStream[input.PaginationRequest, ExampleResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_PaginateServer = grpc.ServerStreamingServer[ExampleResponse]

func _ExampleService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(input.FilteredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_Count_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Count(ctx, req.(*input.FilteredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Create(&grpc.GenericServerStream[ExampleRequest, output.PersistenceDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_CreateServer = grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]

func _ExampleService_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Update(&grpc.GenericServerStream[ExampleRequest, output.PersistenceDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_UpdateServer = grpc.ClientStreamingServer[ExampleRequest, output.PersistenceDataResponse]

func _ExampleService_Inactive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Inactive(&grpc.GenericServerStream[input.UUIDRequest, output.StatusDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_InactiveServer = grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]

func _ExampleService_Active_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Active(&grpc.GenericServerStream[input.UUIDRequest, output.StatusDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ActiveServer = grpc.ClientStreamingServer[input.UUIDRequest, output.StatusDataResponse]

func _ExampleService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(input.UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_FindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).FindOne(ctx, req.(*input.UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _ExampleService_Count_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _ExampleService_FindOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Paginate",
			Handler:       _ExampleService_Paginate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Create",
			Handler:       _ExampleService_Create_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Update",
			Handler:       _ExampleService_Update_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Inactive",
			Handler:       _ExampleService_Inactive_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Active",
			Handler:       _ExampleService_Active_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}
