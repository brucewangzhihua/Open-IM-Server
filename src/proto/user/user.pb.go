// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

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

type CommonResp struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResp) Reset()         { *m = CommonResp{} }
func (m *CommonResp) String() string { return proto.CompactTextString(m) }
func (*CommonResp) ProtoMessage()    {}
func (*CommonResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *CommonResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResp.Unmarshal(m, b)
}
func (m *CommonResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResp.Marshal(b, m, deterministic)
}
func (m *CommonResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResp.Merge(m, src)
}
func (m *CommonResp) XXX_Size() int {
	return xxx_messageInfo_CommonResp.Size(m)
}
func (m *CommonResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResp.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResp proto.InternalMessageInfo

func (m *CommonResp) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *CommonResp) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type GetUserInfoReq struct {
	UserIDList           []string `protobuf:"bytes,1,rep,name=userIDList,proto3" json:"userIDList,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	OperationID          string   `protobuf:"bytes,3,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(m, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

func (m *GetUserInfoReq) GetUserIDList() []string {
	if m != nil {
		return m.UserIDList
	}
	return nil
}

func (m *GetUserInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetUserInfoReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type GetUserInfoResp struct {
	ErrorCode            int32       `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMsg             string      `protobuf:"bytes,2,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	Data                 []*UserInfo `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUserInfoResp) Reset()         { *m = GetUserInfoResp{} }
func (m *GetUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoResp) ProtoMessage()    {}
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *GetUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoResp.Unmarshal(m, b)
}
func (m *GetUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoResp.Marshal(b, m, deterministic)
}
func (m *GetUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoResp.Merge(m, src)
}
func (m *GetUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoResp.Size(m)
}
func (m *GetUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoResp proto.InternalMessageInfo

func (m *GetUserInfoResp) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetUserInfoResp) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *GetUserInfoResp) GetData() []*UserInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Gender               int32    `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Birth                string   `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Ex                   string   `protobuf:"bytes,8,opt,name=ex,proto3" json:"ex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserInfo) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfo) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

type LogoutReq struct {
	OperationID          string   `protobuf:"bytes,1,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *LogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReq.Unmarshal(m, b)
}
func (m *LogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReq.Marshal(b, m, deterministic)
}
func (m *LogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReq.Merge(m, src)
}
func (m *LogoutReq) XXX_Size() int {
	return xxx_messageInfo_LogoutReq.Size(m)
}
func (m *LogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReq proto.InternalMessageInfo

func (m *LogoutReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *LogoutReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateUserInfoReq struct {
	Icon                 string   `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               int32    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile               string   `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Birth                string   `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Ex                   string   `protobuf:"bytes,7,opt,name=ex,proto3" json:"ex,omitempty"`
	Token                string   `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	OperationID          string   `protobuf:"bytes,9,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfoReq) Reset()         { *m = UpdateUserInfoReq{} }
func (m *UpdateUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoReq) ProtoMessage()    {}
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{5}
}

func (m *UpdateUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoReq.Unmarshal(m, b)
}
func (m *UpdateUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoReq.Merge(m, src)
}
func (m *UpdateUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoReq.Size(m)
}
func (m *UpdateUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoReq proto.InternalMessageInfo

func (m *UpdateUserInfoReq) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UpdateUserInfoReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserInfoReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UpdateUserInfoReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UpdateUserInfoReq) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

func (m *UpdateUserInfoReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateUserInfoReq) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

func (m *UpdateUserInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateUserInfoReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonResp)(nil), "user.CommonResp")
	proto.RegisterType((*GetUserInfoReq)(nil), "user.GetUserInfoReq")
	proto.RegisterType((*GetUserInfoResp)(nil), "user.GetUserInfoResp")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*LogoutReq)(nil), "user.LogoutReq")
	proto.RegisterType((*UpdateUserInfoReq)(nil), "user.UpdateUserInfoReq")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0xa5, 0x27, 0x1f, 0x9b, 0x54, 0x20, 0xbb, 0x16, 0xab, 0x36, 0x83, 0x48, 0xc8, 0x29, 0xa7,
	0x15, 0xd6, 0x9b, 0x7b, 0x73, 0x06, 0x65, 0x60, 0x45, 0x08, 0xec, 0xc5, 0x5b, 0xc6, 0x94, 0x99,
	0xe0, 0x24, 0x1d, 0x3b, 0x3d, 0x30, 0x37, 0xff, 0x92, 0xff, 0xc8, 0xbf, 0x22, 0xdd, 0xc9, 0x4c,
	0x32, 0x33, 0xd1, 0x8b, 0x97, 0x50, 0xef, 0x75, 0xa8, 0xaa, 0xf7, 0xaa, 0x0a, 0xae, 0x77, 0x2d,
	0xc9, 0x37, 0xfa, 0x73, 0xd7, 0x48, 0xa1, 0x04, 0xda, 0x3a, 0x8e, 0x3f, 0x00, 0x2c, 0x44, 0x55,
	0x89, 0x3a, 0xa5, 0xb6, 0xc1, 0x57, 0xe0, 0x93, 0x94, 0x42, 0x2e, 0x44, 0x4e, 0x9c, 0x45, 0x2c,
	0x71, 0xd2, 0x81, 0xc0, 0x39, 0x78, 0x06, 0x7c, 0x6a, 0x0b, 0x3e, 0x8b, 0x58, 0xe2, 0xa7, 0x47,
	0x1c, 0x6f, 0x20, 0xfc, 0x48, 0xea, 0xa9, 0x25, 0xb9, 0xaa, 0xbf, 0x89, 0x94, 0x7e, 0xe0, 0x6b,
	0x00, 0x5d, 0x61, 0xb5, 0x7c, 0x2c, 0x5b, 0xc5, 0x59, 0x64, 0x25, 0x7e, 0x3a, 0x62, 0xf0, 0x16,
	0x1c, 0x25, 0xbe, 0x53, 0xdd, 0xa7, 0xea, 0x00, 0x46, 0x10, 0x7c, 0x6e, 0x48, 0x66, 0xaa, 0x14,
	0xf5, 0x6a, 0xc9, 0x2d, 0xf3, 0x36, 0xa6, 0x62, 0x01, 0xd7, 0x27, 0x95, 0xfe, 0xa7, 0x6d, 0x8c,
	0xc1, 0x5e, 0x66, 0x2a, 0xe3, 0x56, 0x64, 0x25, 0xc1, 0x7d, 0x78, 0x67, 0xfc, 0x39, 0xe6, 0x36,
	0x6f, 0xf1, 0x2f, 0x06, 0xde, 0x81, 0xc2, 0x1b, 0xb0, 0x76, 0x65, 0x6e, 0x8a, 0xf8, 0xa9, 0x0e,
	0x11, 0xc1, 0xae, 0xb3, 0x8a, 0xfa, 0xd4, 0x26, 0xd6, 0x5c, 0xf9, 0x55, 0xd4, 0x7d, 0xfb, 0x26,
	0xc6, 0x17, 0xe0, 0x16, 0x54, 0xe7, 0x24, 0xb9, 0x6d, 0x3a, 0xec, 0x91, 0xe6, 0x2b, 0xb1, 0x2e,
	0xb7, 0xc4, 0x1d, 0xf3, 0x77, 0x8f, 0xb4, 0x3f, 0xeb, 0x52, 0xaa, 0x0d, 0x77, 0x3b, 0x7f, 0x0c,
	0xd0, 0x2c, 0x55, 0x59, 0xb9, 0xe5, 0x57, 0x1d, 0x6b, 0x00, 0x86, 0x30, 0xa3, 0x3d, 0xf7, 0x0c,
	0x35, 0xa3, 0x7d, 0xbc, 0x00, 0xff, 0x51, 0x14, 0x62, 0xa7, 0xf4, 0x20, 0xce, 0x2c, 0x65, 0x17,
	0x96, 0x4e, 0x8f, 0x22, 0xfe, 0xcd, 0xe0, 0xd9, 0x53, 0x93, 0x67, 0x8a, 0xc6, 0x63, 0x3d, 0x48,
	0x63, 0x23, 0x69, 0x53, 0x16, 0x0c, 0x72, 0xad, 0xbf, 0xc8, 0xb5, 0xa7, 0xe5, 0x3a, 0x93, 0x72,
	0xdd, 0x4b, 0xb9, 0x57, 0x07, 0xb9, 0x43, 0xff, 0xde, 0x3f, 0x56, 0xc9, 0xbf, 0xd0, 0x7d, 0xff,
	0x13, 0xcc, 0x11, 0xe0, 0x3b, 0x08, 0x8a, 0x61, 0xa5, 0xf0, 0xb6, 0x5b, 0x83, 0xd3, 0x7d, 0x9e,
	0x3f, 0x9f, 0x60, 0xdb, 0x06, 0x1f, 0x20, 0x3c, 0x35, 0x09, 0x5f, 0xf6, 0x5b, 0x74, 0x6e, 0xdd,
	0xfc, 0xa6, 0x7b, 0x18, 0xee, 0xed, 0x7d, 0xf0, 0xc5, 0xd7, 0xd4, 0x83, 0xfe, 0xac, 0x5d, 0x73,
	0x97, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6d, 0x34, 0xa2, 0xaa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*CommonResp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/user.user/getUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/user.user/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*CommonResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (*UnimplementedUserServer) UpdateUserInfo(ctx context.Context, req *UpdateUserInfoReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _User_UpdateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
