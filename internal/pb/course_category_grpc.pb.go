// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/course_category.proto

package pb

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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	CreateCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamClient, error)
	CreateCategoryStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamBidirectionalClient, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CreateCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[0], "/pb.CategoryService/CreateCategoryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateCategoryStreamClient{stream}
	return x, nil
}

type CategoryService_CreateCategoryStreamClient interface {
	Send(*CreateCategoryRequest) error
	CloseAndRecv() (*CategoriesResponse, error)
	grpc.ClientStream
}

type categoryServiceCreateCategoryStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateCategoryStreamClient) Send(m *CreateCategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamClient) CloseAndRecv() (*CategoriesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CategoriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) CreateCategoryStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[1], "/pb.CategoryService/CreateCategoryStreamBidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateCategoryStreamBidirectionalClient{stream}
	return x, nil
}

type CategoryService_CreateCategoryStreamBidirectionalClient interface {
	Send(*CreateCategoryRequest) error
	Recv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceCreateCategoryStreamBidirectionalClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateCategoryStreamBidirectionalClient) Send(m *CreateCategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamBidirectionalClient) Recv() (*CategoryResponse, error) {
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	GetCategory(context.Context, *GetCategoryRequest) (*CategoryResponse, error)
	GetCategories(context.Context, *EmptyRequest) (*CategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)
	CreateCategoryStream(CategoryService_CreateCategoryStreamServer) error
	CreateCategoryStreamBidirectional(CategoryService_CreateCategoryStreamBidirectionalServer) error
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategories(context.Context, *EmptyRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CreateCategoryStream(CategoryService_CreateCategoryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCategoryStream not implemented")
}
func (UnimplementedCategoryServiceServer) CreateCategoryStreamBidirectional(CategoryService_CreateCategoryStreamBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCategoryStreamBidirectional not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategories(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CreateCategoryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).CreateCategoryStream(&categoryServiceCreateCategoryStreamServer{stream})
}

type CategoryService_CreateCategoryStreamServer interface {
	SendAndClose(*CategoriesResponse) error
	Recv() (*CreateCategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceCreateCategoryStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateCategoryStreamServer) SendAndClose(m *CategoriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamServer) Recv() (*CreateCategoryRequest, error) {
	m := new(CreateCategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CategoryService_CreateCategoryStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).CreateCategoryStreamBidirectional(&categoryServiceCreateCategoryStreamBidirectionalServer{stream})
}

type CategoryService_CreateCategoryStreamBidirectionalServer interface {
	Send(*CategoryResponse) error
	Recv() (*CreateCategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceCreateCategoryStreamBidirectionalServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateCategoryStreamBidirectionalServer) Send(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamBidirectionalServer) Recv() (*CreateCategoryRequest, error) {
	m := new(CreateCategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategory",
			Handler:    _CategoryService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _CategoryService_GetCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCategoryStream",
			Handler:       _CategoryService_CreateCategoryStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateCategoryStreamBidirectional",
			Handler:       _CategoryService_CreateCategoryStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/course_category.proto",
}

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
	GetCourses(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CoursesResponse, error)
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
	CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error)
	CreateCourseStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirectionalClient, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourses(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[0], "/pb.CourseService/CreateCourseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamClient{stream}
	return x, nil
}

type CourseService_CreateCourseStreamClient interface {
	Send(*CreateCourseRequest) error
	CloseAndRecv() (*CoursesResponse, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamClient) CloseAndRecv() (*CoursesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CoursesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseServiceClient) CreateCourseStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[1], "/pb.CourseService/CreateCourseStreamBidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamBidirectionalClient{stream}
	return x, nil
}

type CourseService_CreateCourseStreamBidirectionalClient interface {
	Send(*CreateCourseRequest) error
	Recv() (*CourseResponse, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamBidirectionalClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamBidirectionalClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirectionalClient) Recv() (*CourseResponse, error) {
	m := new(CourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	GetCourse(context.Context, *GetCourseRequest) (*CourseResponse, error)
	GetCourses(context.Context, *EmptyRequest) (*CoursesResponse, error)
	CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error)
	CreateCourseStream(CourseService_CreateCourseStreamServer) error
	CreateCourseStreamBidirectional(CourseService_CreateCourseStreamBidirectionalServer) error
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) GetCourse(context.Context, *GetCourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServiceServer) GetCourses(context.Context, *EmptyRequest) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStream(CourseService_CreateCourseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStream not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStreamBidirectional(CourseService_CreateCourseStreamBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStreamBidirectional not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourse(ctx, req.(*GetCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourses(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_CreateCourseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStream(&courseServiceCreateCourseStreamServer{stream})
}

type CourseService_CreateCourseStreamServer interface {
	SendAndClose(*CoursesResponse) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamServer) SendAndClose(m *CoursesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CourseService_CreateCourseStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStreamBidirectional(&courseServiceCreateCourseStreamBidirectionalServer{stream})
}

type CourseService_CreateCourseStreamBidirectionalServer interface {
	Send(*CourseResponse) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamBidirectionalServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamBidirectionalServer) Send(m *CourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirectionalServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourse",
			Handler:    _CourseService_GetCourse_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _CourseService_GetCourses_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCourseStream",
			Handler:       _CourseService_CreateCourseStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateCourseStreamBidirectional",
			Handler:       _CourseService_CreateCourseStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/course_category.proto",
}
