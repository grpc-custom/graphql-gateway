// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/photo_share/proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	GithubLogin          string   `protobuf:"bytes,1,opt,name=github_login,json=githubLogin,proto3" json:"github_login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetGithubLogin() string {
	if m != nil {
		return m.GithubLogin
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type TotalUsersResponse struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalUsersResponse) Reset()         { *m = TotalUsersResponse{} }
func (m *TotalUsersResponse) String() string { return proto.CompactTextString(m) }
func (*TotalUsersResponse) ProtoMessage()    {}
func (*TotalUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{2}
}

func (m *TotalUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalUsersResponse.Unmarshal(m, b)
}
func (m *TotalUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalUsersResponse.Marshal(b, m, deterministic)
}
func (m *TotalUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalUsersResponse.Merge(m, src)
}
func (m *TotalUsersResponse) XXX_Size() int {
	return xxx_messageInfo_TotalUsersResponse.Size(m)
}
func (m *TotalUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalUsersResponse proto.InternalMessageInfo

func (m *TotalUsersResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type AllUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllUsersResponse) Reset()         { *m = AllUsersResponse{} }
func (m *AllUsersResponse) String() string { return proto.CompactTextString(m) }
func (*AllUsersResponse) ProtoMessage()    {}
func (*AllUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{3}
}

func (m *AllUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllUsersResponse.Unmarshal(m, b)
}
func (m *AllUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllUsersResponse.Marshal(b, m, deterministic)
}
func (m *AllUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllUsersResponse.Merge(m, src)
}
func (m *AllUsersResponse) XXX_Size() int {
	return xxx_messageInfo_AllUsersResponse.Size(m)
}
func (m *AllUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllUsersResponse proto.InternalMessageInfo

func (m *AllUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type LoginRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{4}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type GithubAuthRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubAuthRequest) Reset()         { *m = GithubAuthRequest{} }
func (m *GithubAuthRequest) String() string { return proto.CompactTextString(m) }
func (*GithubAuthRequest) ProtoMessage()    {}
func (*GithubAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{5}
}

func (m *GithubAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubAuthRequest.Unmarshal(m, b)
}
func (m *GithubAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubAuthRequest.Marshal(b, m, deterministic)
}
func (m *GithubAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubAuthRequest.Merge(m, src)
}
func (m *GithubAuthRequest) XXX_Size() int {
	return xxx_messageInfo_GithubAuthRequest.Size(m)
}
func (m *GithubAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GithubAuthRequest proto.InternalMessageInfo

func (m *GithubAuthRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GithubAuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubAuthResponse) Reset()         { *m = GithubAuthResponse{} }
func (m *GithubAuthResponse) String() string { return proto.CompactTextString(m) }
func (*GithubAuthResponse) ProtoMessage()    {}
func (*GithubAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb90d003b8ad82f2, []int{6}
}

func (m *GithubAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubAuthResponse.Unmarshal(m, b)
}
func (m *GithubAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubAuthResponse.Marshal(b, m, deterministic)
}
func (m *GithubAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubAuthResponse.Merge(m, src)
}
func (m *GithubAuthResponse) XXX_Size() int {
	return xxx_messageInfo_GithubAuthResponse.Size(m)
}
func (m *GithubAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GithubAuthResponse proto.InternalMessageInfo

func (m *GithubAuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GithubAuthResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*TotalUsersResponse)(nil), "user.TotalUsersResponse")
	proto.RegisterType((*AllUsersResponse)(nil), "user.AllUsersResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*GithubAuthRequest)(nil), "user.GithubAuthRequest")
	proto.RegisterType((*GithubAuthResponse)(nil), "user.GithubAuthResponse")
}

func init() {
	proto.RegisterFile("example/photo_share/proto/user/user.proto", fileDescriptor_bb90d003b8ad82f2)
}

var fileDescriptor_bb90d003b8ad82f2 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x8a, 0x6c, 0xdc, 0x89, 0x29, 0xe9, 0x50, 0x5c, 0xe1, 0x42, 0x71, 0x97, 0x42, 0x3f,
	0x4c, 0x24, 0x48, 0x73, 0xca, 0x2d, 0x87, 0x52, 0x08, 0x2d, 0xb4, 0x6e, 0x73, 0x0e, 0x6b, 0x77,
	0x2a, 0x85, 0x4a, 0x5e, 0x65, 0x77, 0x95, 0x36, 0x7f, 0xc7, 0xff, 0xc1, 0x3f, 0xad, 0xf7, 0xb0,
	0xb3, 0x92, 0x2c, 0xec, 0xe4, 0x62, 0xe6, 0xcd, 0xbc, 0x79, 0xeb, 0x37, 0x4f, 0xf0, 0x9e, 0xfe,
	0xc9, 0xa2, 0xcc, 0x29, 0x29, 0x33, 0x65, 0xd5, 0x95, 0xc9, 0xa4, 0xa6, 0xa4, 0xd4, 0xca, 0xaa,
	0xa4, 0x32, 0xa4, 0xf9, 0x27, 0x66, 0x8c, 0xa1, 0xab, 0x27, 0xa7, 0xe9, 0xb5, 0xcd, 0xaa, 0x45,
	0xbc, 0x54, 0x45, 0x92, 0xea, 0x72, 0x79, 0xbc, 0xac, 0x8c, 0xe5, 0x5a, 0x96, 0xd9, 0x4d, 0x7e,
	0x9c, 0x4a, 0x4b, 0x7f, 0xe5, 0x5d, 0x83, 0xfd, 0xee, 0xe4, 0x65, 0xaa, 0x54, 0x9a, 0xd7, 0xca,
	0x8b, 0xea, 0x77, 0x42, 0x45, 0x69, 0xef, 0xfc, 0x50, 0x5c, 0x42, 0x78, 0x69, 0x48, 0xe3, 0x6b,
	0x18, 0x79, 0xf1, 0xab, 0x5c, 0xa5, 0xd7, 0xab, 0xa8, 0x37, 0xed, 0xbd, 0x7b, 0x32, 0x3f, 0xf4,
	0xbd, 0x2f, 0xae, 0x85, 0x08, 0xe1, 0x4a, 0x16, 0x14, 0x05, 0x3c, 0xe2, 0x1a, 0xc7, 0x30, 0x90,
	0xb7, 0xd2, 0x4a, 0x1d, 0x1d, 0x70, 0xb7, 0x46, 0x22, 0x86, 0x91, 0x93, 0x9d, 0x93, 0x29, 0xd5,
	0xca, 0x10, 0xbe, 0x02, 0x76, 0xc0, 0xb2, 0x87, 0x27, 0x10, 0xb3, 0x35, 0x66, 0x70, 0x5f, 0x7c,
	0x00, 0xfc, 0xa9, 0xac, 0xcc, 0x5d, 0xcb, 0xb4, 0x5b, 0xcf, 0xa1, 0x6f, 0x5d, 0x97, 0xd7, 0xfa,
	0x73, 0x0f, 0xc4, 0x29, 0x1c, 0x9d, 0xe7, 0x3b, 0xcc, 0x29, 0xf4, 0x9d, 0x8e, 0x89, 0x7a, 0xd3,
	0x83, 0x9d, 0x07, 0xfc, 0x40, 0xbc, 0x81, 0x11, 0xdb, 0x98, 0xd3, 0x4d, 0x45, 0xc6, 0x3a, 0xed,
	0xae, 0x53, 0x0f, 0xc4, 0x5b, 0x78, 0xf6, 0x99, 0x2d, 0x9f, 0x57, 0x36, 0x6b, 0xa8, 0x08, 0xe1,
	0x52, 0xfd, 0xa2, 0x9a, 0xc9, 0xb5, 0xb8, 0x00, 0xec, 0x12, 0xbb, 0x7f, 0xf8, 0x0f, 0xb5, 0xa2,
	0x0c, 0x5a, 0xf3, 0xc1, 0xc3, 0xe6, 0x4f, 0xfe, 0x07, 0x00, 0x0e, 0xfe, 0x20, 0x7d, 0x4b, 0x1a,
	0xcf, 0x20, 0xf8, 0x4a, 0x38, 0x8e, 0x7d, 0x6c, 0x71, 0x13, 0x5b, 0xfc, 0xc9, 0xc5, 0x36, 0xc1,
	0xce, 0x7a, 0xfd, 0xac, 0x18, 0xae, 0x37, 0xb3, 0x10, 0x82, 0x82, 0xf0, 0x1b, 0xc0, 0xf6, 0x8e,
	0x8f, 0x6a, 0x44, 0x5e, 0x63, 0xff, 0xe2, 0xe2, 0x68, 0xbd, 0x99, 0x8d, 0x00, 0xec, 0x56, 0xe3,
	0x02, 0x86, 0xcd, 0xb5, 0x1f, 0xd5, 0x1b, 0x7b, 0xbd, 0xdd, 0x54, 0xc4, 0xd3, 0xf5, 0x66, 0x06,
	0x30, 0x94, 0xcd, 0xfe, 0x59, 0xfd, 0xb1, 0xd5, 0x1e, 0xba, 0x79, 0x3c, 0xe8, 0x0b, 0xd6, 0x9b,
	0xd9, 0xc0, 0x1f, 0x0f, 0xbf, 0x03, 0x6c, 0x0f, 0x8e, 0x2f, 0x3c, 0x7b, 0x2f, 0xab, 0xc6, 0xda,
	0x7e, 0x36, 0xde, 0x1a, 0x42, 0xda, 0x4e, 0x16, 0x03, 0xb6, 0xf1, 0xf1, 0x3e, 0x00, 0x00, 0xff,
	0xff, 0x8f, 0x45, 0x73, 0xe9, 0x88, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServerClient interface {
	Me(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserResponse, error)
	TotalUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TotalUsersResponse, error)
	AllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponse, error)
	User(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GithubAuth(ctx context.Context, in *GithubAuthRequest, opts ...grpc.CallOption) (*GithubAuthResponse, error)
}

type userServerClient struct {
	cc *grpc.ClientConn
}

func NewUserServerClient(cc *grpc.ClientConn) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) Me(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserServer/Me", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) TotalUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TotalUsersResponse, error) {
	out := new(TotalUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserServer/TotalUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) AllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponse, error) {
	out := new(AllUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserServer/AllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) User(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserServer/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) GithubAuth(ctx context.Context, in *GithubAuthRequest, opts ...grpc.CallOption) (*GithubAuthResponse, error) {
	out := new(GithubAuthResponse)
	err := c.cc.Invoke(ctx, "/user.UserServer/GithubAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
type UserServerServer interface {
	Me(context.Context, *empty.Empty) (*UserResponse, error)
	TotalUsers(context.Context, *empty.Empty) (*TotalUsersResponse, error)
	AllUsers(context.Context, *empty.Empty) (*AllUsersResponse, error)
	User(context.Context, *LoginRequest) (*UserResponse, error)
	GithubAuth(context.Context, *GithubAuthRequest) (*GithubAuthResponse, error)
}

func RegisterUserServerServer(s *grpc.Server, srv UserServerServer) {
	s.RegisterService(&_UserServer_serviceDesc, srv)
}

func _UserServer_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServer/Me",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).Me(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_TotalUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).TotalUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServer/TotalUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).TotalUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_AllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).AllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServer/AllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).AllUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServer/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).User(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_GithubAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GithubAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GithubAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServer/GithubAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GithubAuth(ctx, req.(*GithubAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Me",
			Handler:    _UserServer_Me_Handler,
		},
		{
			MethodName: "TotalUsers",
			Handler:    _UserServer_TotalUsers_Handler,
		},
		{
			MethodName: "AllUsers",
			Handler:    _UserServer_AllUsers_Handler,
		},
		{
			MethodName: "User",
			Handler:    _UserServer_User_Handler,
		},
		{
			MethodName: "GithubAuth",
			Handler:    _UserServer_GithubAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/photo_share/proto/user/user.proto",
}