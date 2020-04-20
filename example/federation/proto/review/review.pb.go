// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/federation/proto/review/review.proto

package review

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

type Review struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	AuthorId             string   `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	ProductId            string   `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d6327e502ce189, []int{0}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Review) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Review) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Review) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type ListUserReviewsRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	First                int32    `protobuf:"varint,2,opt,name=first,proto3" json:"first,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserReviewsRequest) Reset()         { *m = ListUserReviewsRequest{} }
func (m *ListUserReviewsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserReviewsRequest) ProtoMessage()    {}
func (*ListUserReviewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d6327e502ce189, []int{1}
}

func (m *ListUserReviewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserReviewsRequest.Unmarshal(m, b)
}
func (m *ListUserReviewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserReviewsRequest.Marshal(b, m, deterministic)
}
func (m *ListUserReviewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserReviewsRequest.Merge(m, src)
}
func (m *ListUserReviewsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserReviewsRequest.Size(m)
}
func (m *ListUserReviewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserReviewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserReviewsRequest proto.InternalMessageInfo

func (m *ListUserReviewsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListUserReviewsRequest) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

type ListUserReviewsResponse struct {
	Reviews              []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListUserReviewsResponse) Reset()         { *m = ListUserReviewsResponse{} }
func (m *ListUserReviewsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserReviewsResponse) ProtoMessage()    {}
func (*ListUserReviewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d6327e502ce189, []int{2}
}

func (m *ListUserReviewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserReviewsResponse.Unmarshal(m, b)
}
func (m *ListUserReviewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserReviewsResponse.Marshal(b, m, deterministic)
}
func (m *ListUserReviewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserReviewsResponse.Merge(m, src)
}
func (m *ListUserReviewsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserReviewsResponse.Size(m)
}
func (m *ListUserReviewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserReviewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserReviewsResponse proto.InternalMessageInfo

func (m *ListUserReviewsResponse) GetReviews() []*Review {
	if m != nil {
		return m.Reviews
	}
	return nil
}

type ListProductReviewsRequest struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProductReviewsRequest) Reset()         { *m = ListProductReviewsRequest{} }
func (m *ListProductReviewsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProductReviewsRequest) ProtoMessage()    {}
func (*ListProductReviewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d6327e502ce189, []int{3}
}

func (m *ListProductReviewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductReviewsRequest.Unmarshal(m, b)
}
func (m *ListProductReviewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductReviewsRequest.Marshal(b, m, deterministic)
}
func (m *ListProductReviewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductReviewsRequest.Merge(m, src)
}
func (m *ListProductReviewsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProductReviewsRequest.Size(m)
}
func (m *ListProductReviewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductReviewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductReviewsRequest proto.InternalMessageInfo

func (m *ListProductReviewsRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type ListProductReviewsResponse struct {
	Reviews              []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListProductReviewsResponse) Reset()         { *m = ListProductReviewsResponse{} }
func (m *ListProductReviewsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProductReviewsResponse) ProtoMessage()    {}
func (*ListProductReviewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d6327e502ce189, []int{4}
}

func (m *ListProductReviewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductReviewsResponse.Unmarshal(m, b)
}
func (m *ListProductReviewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductReviewsResponse.Marshal(b, m, deterministic)
}
func (m *ListProductReviewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductReviewsResponse.Merge(m, src)
}
func (m *ListProductReviewsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProductReviewsResponse.Size(m)
}
func (m *ListProductReviewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductReviewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductReviewsResponse proto.InternalMessageInfo

func (m *ListProductReviewsResponse) GetReviews() []*Review {
	if m != nil {
		return m.Reviews
	}
	return nil
}

func init() {
	proto.RegisterType((*Review)(nil), "review.Review")
	proto.RegisterType((*ListUserReviewsRequest)(nil), "review.ListUserReviewsRequest")
	proto.RegisterType((*ListUserReviewsResponse)(nil), "review.ListUserReviewsResponse")
	proto.RegisterType((*ListProductReviewsRequest)(nil), "review.ListProductReviewsRequest")
	proto.RegisterType((*ListProductReviewsResponse)(nil), "review.ListProductReviewsResponse")
}

func init() {
	proto.RegisterFile("example/federation/proto/review/review.proto", fileDescriptor_c9d6327e502ce189)
}

var fileDescriptor_c9d6327e502ce189 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xd2, 0x24, 0x5d, 0x3e, 0x60, 0x43, 0x46, 0x1a, 0x21, 0x42, 0xd0, 0x19, 0x21, 0x55,
	0x0a, 0x6b, 0xa6, 0xc1, 0x15, 0x57, 0xa8, 0x12, 0x17, 0x1d, 0x08, 0xa1, 0x20, 0x2e, 0xb8, 0x42,
	0x69, 0xec, 0xa5, 0x96, 0xb6, 0x39, 0xb5, 0x9d, 0x96, 0xbe, 0x02, 0x8f, 0x91, 0x77, 0xc8, 0x63,
	0xf1, 0x0e, 0x28, 0x76, 0x52, 0xf5, 0x0f, 0x2e, 0x76, 0x15, 0xfb, 0x3b, 0x27, 0xe7, 0x9c, 0xef,
	0xc8, 0xf0, 0x86, 0xfe, 0x4a, 0x6f, 0x8b, 0x1b, 0x1a, 0x5f, 0x53, 0x42, 0x45, 0xaa, 0x18, 0xbf,
	0x8b, 0x0b, 0xc1, 0x15, 0x8f, 0x05, 0x5d, 0x30, 0xba, 0x6c, 0x3f, 0x23, 0x3d, 0x43, 0x9e, 0xb9,
	0x85, 0xef, 0x72, 0xa6, 0x66, 0xe5, 0x74, 0x94, 0xf1, 0xdb, 0x38, 0x17, 0x45, 0x76, 0x9e, 0x95,
	0x52, 0xe9, 0x73, 0x5a, 0xcc, 0xe6, 0x37, 0xe7, 0x79, 0xaa, 0xe8, 0x32, 0x5d, 0x75, 0x77, 0xf3,
	0x37, 0xfe, 0x6d, 0x81, 0x97, 0x68, 0x01, 0x74, 0x0c, 0x36, 0x23, 0x81, 0x35, 0xb0, 0x86, 0x7e,
	0x62, 0x33, 0x82, 0x10, 0x38, 0x53, 0x4e, 0x56, 0x81, 0xad, 0x27, 0xfa, 0x8c, 0x5e, 0x81, 0x9f,
	0x96, 0x6a, 0xc6, 0xc5, 0x4f, 0x46, 0x82, 0x5e, 0x03, 0x8c, 0xbd, 0xaa, 0x8e, 0xec, 0x0f, 0x56,
	0x72, 0x64, 0x80, 0x09, 0x41, 0xaf, 0x01, 0x0a, 0xc1, 0x49, 0x99, 0xa9, 0x86, 0xe5, 0x6c, 0xb1,
	0xfc, 0x16, 0x99, 0x90, 0xf7, 0x8f, 0xab, 0x3a, 0x7a, 0xb8, 0x76, 0xb7, 0x19, 0xc1, 0x3f, 0xe0,
	0xf4, 0x33, 0x93, 0xea, 0xbb, 0xa4, 0xc2, 0x4c, 0x65, 0x42, 0xe7, 0x25, 0x95, 0x0a, 0x9d, 0x41,
	0xbf, 0x94, 0x54, 0xbb, 0xea, 0x80, 0xe3, 0xa3, 0xaa, 0x8e, 0x9c, 0x2b, 0x9b, 0x91, 0xc4, 0x6b,
	0x80, 0x09, 0x41, 0xcf, 0xc1, 0xbd, 0x66, 0x42, 0x2a, 0x9d, 0xd7, 0x35, 0x86, 0x03, 0x37, 0x31,
	0x43, 0xfc, 0x09, 0x9e, 0xee, 0x49, 0xcb, 0x82, 0xdf, 0x49, 0x8a, 0x2e, 0xa0, 0x6f, 0x2a, 0x94,
	0x81, 0x35, 0xe8, 0x0d, 0x1f, 0x5c, 0x1e, 0x8f, 0xda, 0x82, 0x0d, 0xd3, 0x48, 0x5d, 0x58, 0x49,
	0x47, 0xc3, 0x1f, 0xe1, 0x59, 0x23, 0xf6, 0xd5, 0xac, 0xb2, 0x13, 0x75, 0xb8, 0xb5, 0xbd, 0x49,
	0xeb, 0x57, 0x75, 0xe4, 0x5e, 0xf5, 0xca, 0x22, 0xdb, 0x28, 0x00, 0x7f, 0x81, 0xf0, 0x90, 0xcc,
	0x7d, 0x63, 0x5d, 0xfe, 0xb1, 0xe0, 0x91, 0xc1, 0xbe, 0x51, 0xb1, 0x60, 0x19, 0x45, 0x39, 0x9c,
	0xec, 0x6c, 0x8d, 0x5e, 0x74, 0x2a, 0x87, 0x9b, 0x0e, 0x5f, 0xfe, 0x13, 0x37, 0xb9, 0xf0, 0x93,
	0xaa, 0x8e, 0x4e, 0xb0, 0xd3, 0x40, 0xe3, 0xce, 0x1a, 0x49, 0x40, 0xfb, 0xab, 0xa0, 0xb3, 0x4d,
	0xad, 0x83, 0x6d, 0x85, 0xf8, 0x7f, 0x94, 0xd6, 0xf1, 0xb4, 0xaa, 0x23, 0x84, 0xfb, 0x2d, 0xba,
	0x36, 0x9d, 0x7a, 0xfa, 0x09, 0xbf, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x26, 0x4c, 0x32, 0x99,
	0x30, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	ListUserReviews(ctx context.Context, in *ListUserReviewsRequest, opts ...grpc.CallOption) (*ListUserReviewsResponse, error)
	ListProductReviews(ctx context.Context, in *ListProductReviewsRequest, opts ...grpc.CallOption) (*ListProductReviewsResponse, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) ListUserReviews(ctx context.Context, in *ListUserReviewsRequest, opts ...grpc.CallOption) (*ListUserReviewsResponse, error) {
	out := new(ListUserReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.ReviewService/ListUserReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ListProductReviews(ctx context.Context, in *ListProductReviewsRequest, opts ...grpc.CallOption) (*ListProductReviewsResponse, error) {
	out := new(ListProductReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.ReviewService/ListProductReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	ListUserReviews(context.Context, *ListUserReviewsRequest) (*ListUserReviewsResponse, error)
	ListProductReviews(context.Context, *ListProductReviewsRequest) (*ListProductReviewsResponse, error)
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_ListUserReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListUserReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/ListUserReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListUserReviews(ctx, req.(*ListUserReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ListProductReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListProductReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/ListProductReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListProductReviews(ctx, req.(*ListProductReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUserReviews",
			Handler:    _ReviewService_ListUserReviews_Handler,
		},
		{
			MethodName: "ListProductReviews",
			Handler:    _ReviewService_ListProductReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/federation/proto/review/review.proto",
}