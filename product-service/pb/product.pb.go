// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetProductRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ProductResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductRequest) Reset()         { *m = CreateProductRequest{} }
func (m *CreateProductRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductRequest) ProtoMessage()    {}
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *CreateProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductRequest.Unmarshal(m, b)
}
func (m *CreateProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductRequest.Marshal(b, m, deterministic)
}
func (m *CreateProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductRequest.Merge(m, src)
}
func (m *CreateProductRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductRequest.Size(m)
}
func (m *CreateProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductRequest proto.InternalMessageInfo

func (m *CreateProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductRequest)(nil), "pb.GetProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "pb.ProductResponse")
	proto.RegisterType((*CreateProductRequest)(nil), "pb.CreateProductRequest")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe6,
	0x12, 0x74, 0x4f, 0x2d, 0x09, 0x80, 0x88, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x99,
	0x72, 0xf1, 0xc3, 0x55, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xa2, 0x2b, 0x11, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xb4, 0xb8,
	0x44, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0xd1, 0x8c, 0x87, 0xa9, 0x65, 0x44, 0xa8, 0x35, 0x3a,
	0xc2, 0xc8, 0xc5, 0x0e, 0x55, 0x26, 0x64, 0xc5, 0xc5, 0x85, 0x70, 0x93, 0x90, 0xa8, 0x5e, 0x41,
	0x92, 0x1e, 0x86, 0x1b, 0xa5, 0x84, 0x41, 0xc2, 0x68, 0xae, 0x52, 0x62, 0x10, 0xb2, 0xe6, 0xe2,
	0x46, 0xa8, 0x35, 0x22, 0x51, 0xb3, 0x03, 0x17, 0x2f, 0x8a, 0x83, 0x85, 0x24, 0x40, 0xea, 0xb0,
	0xf9, 0x01, 0x87, 0x09, 0x49, 0x6c, 0xe0, 0x90, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0x1c, 0x6d, 0x9a, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProduct2(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Product/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProduct2(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Product/GetProduct2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Product/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	GetProduct(context.Context, *GetProductRequest) (*ProductResponse, error)
	GetProduct2(context.Context, *GetProductRequest) (*ProductResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*ProductResponse, error)
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) GetProduct(ctx context.Context, req *GetProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedProductServer) GetProduct2(ctx context.Context, req *GetProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct2 not implemented")
}
func (*UnimplementedProductServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Product/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProduct2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Product/GetProduct2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct2(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Product/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "GetProduct2",
			Handler:    _Product_GetProduct2_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}