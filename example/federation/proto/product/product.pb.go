// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/federation/proto/product/product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-custom/graphql-gateway/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Product struct {
	Upc                  string   `protobuf:"bytes,1,opt,name=upc,proto3" json:"upc,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5f96e24a42c505, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetUpc() string {
	if m != nil {
		return m.Upc
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type TopProductsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopProductsRequest) Reset()         { *m = TopProductsRequest{} }
func (m *TopProductsRequest) String() string { return proto.CompactTextString(m) }
func (*TopProductsRequest) ProtoMessage()    {}
func (*TopProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5f96e24a42c505, []int{1}
}

func (m *TopProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopProductsRequest.Unmarshal(m, b)
}
func (m *TopProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopProductsRequest.Marshal(b, m, deterministic)
}
func (m *TopProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopProductsRequest.Merge(m, src)
}
func (m *TopProductsRequest) XXX_Size() int {
	return xxx_messageInfo_TopProductsRequest.Size(m)
}
func (m *TopProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopProductsRequest proto.InternalMessageInfo

type TopProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TopProductsResponse) Reset()         { *m = TopProductsResponse{} }
func (m *TopProductsResponse) String() string { return proto.CompactTextString(m) }
func (*TopProductsResponse) ProtoMessage()    {}
func (*TopProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5f96e24a42c505, []int{2}
}

func (m *TopProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopProductsResponse.Unmarshal(m, b)
}
func (m *TopProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopProductsResponse.Marshal(b, m, deterministic)
}
func (m *TopProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopProductsResponse.Merge(m, src)
}
func (m *TopProductsResponse) XXX_Size() int {
	return xxx_messageInfo_TopProductsResponse.Size(m)
}
func (m *TopProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopProductsResponse proto.InternalMessageInfo

func (m *TopProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type GetProductRequest struct {
	Upc                  string   `protobuf:"bytes,1,opt,name=upc,proto3" json:"upc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5f96e24a42c505, []int{3}
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

func (m *GetProductRequest) GetUpc() string {
	if m != nil {
		return m.Upc
	}
	return ""
}

type GetProductResponse struct {
	Products             *Product `protobuf:"bytes,1,opt,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductResponse) Reset()         { *m = GetProductResponse{} }
func (m *GetProductResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductResponse) ProtoMessage()    {}
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5f96e24a42c505, []int{4}
}

func (m *GetProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductResponse.Unmarshal(m, b)
}
func (m *GetProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductResponse.Marshal(b, m, deterministic)
}
func (m *GetProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductResponse.Merge(m, src)
}
func (m *GetProductResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductResponse.Size(m)
}
func (m *GetProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductResponse proto.InternalMessageInfo

func (m *GetProductResponse) GetProducts() *Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "product.Product")
	proto.RegisterType((*TopProductsRequest)(nil), "product.TopProductsRequest")
	proto.RegisterType((*TopProductsResponse)(nil), "product.TopProductsResponse")
	proto.RegisterType((*GetProductRequest)(nil), "product.GetProductRequest")
	proto.RegisterType((*GetProductResponse)(nil), "product.GetProductResponse")
}

func init() {
	proto.RegisterFile("example/federation/proto/product/product.proto", fileDescriptor_dc5f96e24a42c505)
}

var fileDescriptor_dc5f96e24a42c505 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x41, 0x40, 0x07, 0x25, 0x30, 0x62, 0xd2, 0x14, 0x0f, 0x64, 0x4f, 0x24, 0x0d, 0xc5,
	0x20, 0x89, 0x89, 0x47, 0x2e, 0x26, 0x78, 0x31, 0xd5, 0x8b, 0xd1, 0xc4, 0x94, 0x65, 0x84, 0x26,
	0x40, 0x97, 0x76, 0x0b, 0xfa, 0x3b, 0xfd, 0x87, 0x7e, 0x85, 0x3f, 0x65, 0x68, 0xb7, 0x50, 0x03,
	0x07, 0x2f, 0xed, 0xcc, 0xbc, 0x7d, 0x6f, 0xdf, 0xdb, 0x0c, 0x58, 0xf4, 0xe5, 0x2c, 0xc4, 0x9c,
	0x7a, 0x9f, 0x34, 0x21, 0xdf, 0x91, 0xae, 0xb7, 0xec, 0x09, 0xdf, 0x93, 0xde, 0xf6, 0x3b, 0x09,
	0xb9, 0xcc, 0xfe, 0x56, 0x32, 0xc5, 0x8a, 0x6a, 0x8d, 0xc1, 0xd4, 0x95, 0xb3, 0x70, 0x6c, 0x71,
	0x6f, 0xd1, 0x9b, 0xfa, 0x82, 0x77, 0x79, 0x18, 0xc8, 0xa4, 0x76, 0xc4, 0x6c, 0x35, 0xef, 0x4e,
	0x1d, 0x49, 0x1b, 0xe7, 0x3b, 0xeb, 0x53, 0x3a, 0x7b, 0x85, 0xca, 0x53, 0x2a, 0x80, 0x75, 0x28,
	0x86, 0x82, 0xeb, 0x5a, 0x5b, 0xeb, 0x9c, 0xd9, 0xdb, 0x12, 0x11, 0x4e, 0x96, 0xce, 0x82, 0xf4,
	0x42, 0x32, 0x4a, 0x6a, 0x6c, 0x42, 0x49, 0xf8, 0x2e, 0x27, 0xbd, 0xd8, 0xd6, 0x3a, 0x25, 0x3b,
	0x6d, 0xee, 0x31, 0x8a, 0xcd, 0xda, 0x5e, 0x6a, 0xcb, 0x66, 0x4d, 0xc0, 0x17, 0x4f, 0xa8, 0x51,
	0x60, 0xd3, 0x2a, 0xa4, 0x40, 0xb2, 0x47, 0xb8, 0xfc, 0x33, 0x0d, 0x84, 0xb7, 0x0c, 0x08, 0x07,
	0x70, 0xaa, 0x82, 0x04, 0xba, 0xd6, 0x2e, 0x76, 0xaa, 0xfd, 0xba, 0x95, 0x05, 0x55, 0x87, 0x87,
	0xe5, 0x28, 0x36, 0x0b, 0x37, 0x9a, 0xbd, 0x3b, 0xc9, 0xee, 0xa0, 0xf1, 0x40, 0x52, 0xe1, 0xea,
	0x06, 0x64, 0xb9, 0x1c, 0xc3, 0x7a, 0x14, 0x9b, 0xe7, 0x23, 0x50, 0x8c, 0x0f, 0x77, 0x92, 0x24,
	0x63, 0x23, 0xc0, 0x3c, 0xf1, 0xa8, 0x09, 0xed, 0x7f, 0x26, 0xfa, 0x3f, 0x1a, 0xd4, 0x14, 0xfa,
	0x4c, 0xfe, 0xda, 0xe5, 0x84, 0x6f, 0x50, 0xcd, 0x85, 0xc4, 0xd6, 0x4e, 0xe5, 0xf0, 0x41, 0x8c,
	0xeb, 0xe3, 0x60, 0x6a, 0x89, 0x35, 0xa2, 0xd8, 0xbc, 0x80, 0xaa, 0xcc, 0xa9, 0xbd, 0x03, 0xec,
	0xbd, 0xa3, 0xb1, 0xa3, 0x1f, 0xbc, 0x84, 0xd1, 0x3a, 0x8a, 0x29, 0xe5, 0xab, 0x28, 0x36, 0x1b,
	0xac, 0x6c, 0xd3, 0xda, 0xa5, 0xcd, 0x30, 0x5b, 0xa3, 0x71, 0x39, 0xd9, 0x8b, 0xdb, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x36, 0xe0, 0xe8, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	TopProducts(ctx context.Context, in *TopProductsRequest, opts ...grpc.CallOption) (*TopProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) TopProducts(ctx context.Context, in *TopProductsRequest, opts ...grpc.CallOption) (*TopProductsResponse, error) {
	out := new(TopProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/TopProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	TopProducts(context.Context, *TopProductsRequest) (*TopProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_TopProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).TopProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/TopProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).TopProducts(ctx, req.(*TopProductsRequest))
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
		FullMethod: "/product.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TopProducts",
			Handler:    _ProductService_TopProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/federation/proto/product/product.proto",
}