// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accesssvc.proto

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

type IssueRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueRequest) Reset()         { *m = IssueRequest{} }
func (m *IssueRequest) String() string { return proto.CompactTextString(m) }
func (*IssueRequest) ProtoMessage()    {}
func (*IssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{0}
}

func (m *IssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueRequest.Unmarshal(m, b)
}
func (m *IssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueRequest.Marshal(b, m, deterministic)
}
func (m *IssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueRequest.Merge(m, src)
}
func (m *IssueRequest) XXX_Size() int {
	return xxx_messageInfo_IssueRequest.Size(m)
}
func (m *IssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueRequest proto.InternalMessageInfo

func (m *IssueRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type IssueResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode            string   `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorDescription     string   `protobuf:"bytes,3,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	AccessToken          string   `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType            string   `protobuf:"bytes,5,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,6,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueResponse) Reset()         { *m = IssueResponse{} }
func (m *IssueResponse) String() string { return proto.CompactTextString(m) }
func (*IssueResponse) ProtoMessage()    {}
func (*IssueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{1}
}

func (m *IssueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueResponse.Unmarshal(m, b)
}
func (m *IssueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueResponse.Marshal(b, m, deterministic)
}
func (m *IssueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueResponse.Merge(m, src)
}
func (m *IssueResponse) XXX_Size() int {
	return xxx_messageInfo_IssueResponse.Size(m)
}
func (m *IssueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueResponse proto.InternalMessageInfo

func (m *IssueResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *IssueResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *IssueResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *IssueResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *IssueResponse) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *IssueResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

type PeekRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeekRequest) Reset()         { *m = PeekRequest{} }
func (m *PeekRequest) String() string { return proto.CompactTextString(m) }
func (*PeekRequest) ProtoMessage()    {}
func (*PeekRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{2}
}

func (m *PeekRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeekRequest.Unmarshal(m, b)
}
func (m *PeekRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeekRequest.Marshal(b, m, deterministic)
}
func (m *PeekRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeekRequest.Merge(m, src)
}
func (m *PeekRequest) XXX_Size() int {
	return xxx_messageInfo_PeekRequest.Size(m)
}
func (m *PeekRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeekRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeekRequest proto.InternalMessageInfo

func (m *PeekRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type PeekResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode            string   `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorDescription     string   `protobuf:"bytes,3,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	Session              *Session `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeekResponse) Reset()         { *m = PeekResponse{} }
func (m *PeekResponse) String() string { return proto.CompactTextString(m) }
func (*PeekResponse) ProtoMessage()    {}
func (*PeekResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{3}
}

func (m *PeekResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeekResponse.Unmarshal(m, b)
}
func (m *PeekResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeekResponse.Marshal(b, m, deterministic)
}
func (m *PeekResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeekResponse.Merge(m, src)
}
func (m *PeekResponse) XXX_Size() int {
	return xxx_messageInfo_PeekResponse.Size(m)
}
func (m *PeekResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeekResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeekResponse proto.InternalMessageInfo

func (m *PeekResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PeekResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *PeekResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *PeekResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type RevokeRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeRequest) Reset()         { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()    {}
func (*RevokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{4}
}

func (m *RevokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeRequest.Unmarshal(m, b)
}
func (m *RevokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeRequest.Marshal(b, m, deterministic)
}
func (m *RevokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeRequest.Merge(m, src)
}
func (m *RevokeRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeRequest.Size(m)
}
func (m *RevokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeRequest proto.InternalMessageInfo

func (m *RevokeRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RevokeResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode            string   `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorDescription     string   `protobuf:"bytes,3,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeResponse) Reset()         { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()    {}
func (*RevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{5}
}

func (m *RevokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResponse.Unmarshal(m, b)
}
func (m *RevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResponse.Marshal(b, m, deterministic)
}
func (m *RevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResponse.Merge(m, src)
}
func (m *RevokeResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeResponse.Size(m)
}
func (m *RevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResponse proto.InternalMessageInfo

func (m *RevokeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RevokeResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *RevokeResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type Session struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RedirectUri          string   `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	GrantedScopes        []string `protobuf:"bytes,5,rep,name=granted_scopes,json=grantedScopes,proto3" json:"granted_scopes,omitempty"`
	AccessClaimsJson     string   `protobuf:"bytes,6,opt,name=access_claims_json,json=accessClaimsJson,proto3" json:"access_claims_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_3077aed999763cf9, []int{6}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Session) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Session) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *Session) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Session) GetGrantedScopes() []string {
	if m != nil {
		return m.GrantedScopes
	}
	return nil
}

func (m *Session) GetAccessClaimsJson() string {
	if m != nil {
		return m.AccessClaimsJson
	}
	return ""
}

func init() {
	proto.RegisterType((*IssueRequest)(nil), "pb.IssueRequest")
	proto.RegisterType((*IssueResponse)(nil), "pb.IssueResponse")
	proto.RegisterType((*PeekRequest)(nil), "pb.PeekRequest")
	proto.RegisterType((*PeekResponse)(nil), "pb.PeekResponse")
	proto.RegisterType((*RevokeRequest)(nil), "pb.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "pb.RevokeResponse")
	proto.RegisterType((*Session)(nil), "pb.Session")
}

func init() { proto.RegisterFile("accesssvc.proto", fileDescriptor_3077aed999763cf9) }

var fileDescriptor_3077aed999763cf9 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xd9, 0xe6, 0xab, 0x1e, 0x27, 0x6d, 0x3a, 0x27, 0x2b, 0x08, 0x29, 0x58, 0xaa, 0x14,
	0xa9, 0x28, 0x82, 0x20, 0x1e, 0x00, 0x95, 0x4b, 0x38, 0x21, 0xa7, 0x9c, 0xad, 0x64, 0x3d, 0x42,
	0xdb, 0x14, 0xef, 0xb2, 0xe3, 0x58, 0xf4, 0x59, 0xb8, 0xf2, 0x4c, 0xf0, 0x3a, 0xc8, 0xbb, 0x6b,
	0x62, 0x40, 0x42, 0xea, 0xa5, 0x47, 0xff, 0xfe, 0x3b, 0xfb, 0x9f, 0xaf, 0x35, 0x9c, 0x6f, 0xa5,
	0x24, 0x66, 0xae, 0xe5, 0xd2, 0x58, 0x5d, 0x69, 0x3c, 0x31, 0xbb, 0xf4, 0x0d, 0x8c, 0xd7, 0xcc,
	0x07, 0xca, 0xe8, 0xcb, 0x81, 0xb8, 0xc2, 0x4b, 0x18, 0x31, 0x31, 0x2b, 0x5d, 0x26, 0x62, 0x2e,
	0x16, 0xf1, 0x2a, 0x5e, 0x9a, 0xdd, 0x72, 0xe3, 0x51, 0xd6, 0x6a, 0xe9, 0x0f, 0x01, 0x93, 0x10,
	0xc7, 0x46, 0x97, 0x4c, 0x98, 0xc0, 0x88, 0x0f, 0xce, 0xc0, 0x05, 0x9e, 0x66, 0xed, 0x27, 0x3e,
	0x03, 0x20, 0x6b, 0xb5, 0xcd, 0xa5, 0x2e, 0x28, 0x39, 0x99, 0x8b, 0x45, 0x94, 0x45, 0x8e, 0x5c,
	0xeb, 0x82, 0xf0, 0x0a, 0x2e, 0xbc, 0x5c, 0x10, 0x4b, 0xab, 0x4c, 0xd5, 0x78, 0xf7, 0xdc, 0xa9,
	0xa9, 0x13, 0xde, 0x1d, 0x39, 0x3e, 0x87, 0xb1, 0xaf, 0x22, 0xaf, 0xf4, 0x9e, 0xca, 0xa4, 0xef,
	0xce, 0xc5, 0x9e, 0xdd, 0x34, 0xa8, 0xb1, 0x73, 0x5a, 0x5e, 0xdd, 0x1b, 0x4a, 0x06, 0xde, 0xce,
	0x91, 0x9b, 0x7b, 0x43, 0x2e, 0x9b, 0xaf, 0x46, 0x59, 0xe2, 0x5c, 0x95, 0xc9, 0x70, 0x2e, 0x16,
	0xbd, 0x2c, 0x0a, 0x64, 0x5d, 0xa6, 0x2f, 0x21, 0xfe, 0x40, 0xb4, 0x6f, 0xdb, 0xf1, 0xb7, 0x9f,
	0xf8, 0xc7, 0x2f, 0xfd, 0x26, 0x60, 0xec, 0x43, 0x1e, 0xb5, 0x13, 0x9d, 0x41, 0xf5, 0xff, 0x33,
	0xa8, 0x15, 0x4c, 0x32, 0xaa, 0xf5, 0x9e, 0x1e, 0x50, 0x51, 0x0d, 0x67, 0x6d, 0xcc, 0x63, 0x96,
	0x94, 0xfe, 0x14, 0x30, 0x0a, 0x05, 0x34, 0xf7, 0x5a, 0x9f, 0x71, 0xae, 0x8a, 0x90, 0x64, 0x14,
	0xc8, 0xba, 0xc0, 0xa7, 0x10, 0xc9, 0x3b, 0x45, 0xa5, 0x53, 0xbd, 0xeb, 0xa9, 0x07, 0xeb, 0xa2,
	0x29, 0xd1, 0x52, 0xa1, 0x2c, 0xc9, 0x2a, 0x3f, 0x58, 0x15, 0xfc, 0xe2, 0x96, 0x7d, 0xb4, 0xca,
	0x17, 0xb4, 0xbb, 0x25, 0x59, 0x85, 0x15, 0x6a, 0x3f, 0xf1, 0x12, 0xce, 0x3e, 0xd9, 0x6d, 0x59,
	0x51, 0x91, 0xb3, 0xd4, 0x86, 0x38, 0x19, 0xcc, 0x7b, 0x8b, 0x28, 0x9b, 0x04, 0xba, 0x71, 0x10,
	0x5f, 0x00, 0x86, 0x36, 0xca, 0xbb, 0xad, 0xfa, 0xcc, 0xf9, 0x2d, 0x6b, 0xbf, 0x4e, 0x51, 0x36,
	0xf5, 0xca, 0xb5, 0x13, 0xde, 0xb3, 0x2e, 0x57, 0xdf, 0x05, 0xe0, 0xdb, 0x63, 0x87, 0x37, 0x64,
	0x6b, 0x25, 0x09, 0x97, 0x30, 0x70, 0x8f, 0x08, 0xa7, 0xcd, 0xec, 0xba, 0xef, 0x70, 0x76, 0xd1,
	0x21, 0x7e, 0x08, 0xe9, 0x13, 0xbc, 0x82, 0x7e, 0xb3, 0x69, 0x78, 0xde, 0x88, 0x9d, 0x35, 0x9d,
	0x4d, 0x8f, 0xe0, 0xf7, 0xe1, 0x57, 0x30, 0xf4, 0x53, 0x44, 0x77, 0xd7, 0x1f, 0x5b, 0x30, 0xc3,
	0x2e, 0x6a, 0x43, 0x76, 0x43, 0xf7, 0x5f, 0x78, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x44,
	0x6a, 0xe8, 0x2a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessTokenServiceClient is the client API for AccessTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessTokenServiceClient interface {
	Issue(ctx context.Context, in *IssueRequest, opts ...grpc.CallOption) (*IssueResponse, error)
	Peek(ctx context.Context, in *PeekRequest, opts ...grpc.CallOption) (*PeekResponse, error)
	Revoke(ctx context.Context, in *RevokeRequest, opts ...grpc.CallOption) (*RevokeResponse, error)
}

type accessTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessTokenServiceClient(cc *grpc.ClientConn) AccessTokenServiceClient {
	return &accessTokenServiceClient{cc}
}

func (c *accessTokenServiceClient) Issue(ctx context.Context, in *IssueRequest, opts ...grpc.CallOption) (*IssueResponse, error) {
	out := new(IssueResponse)
	err := c.cc.Invoke(ctx, "/pb.AccessTokenService/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessTokenServiceClient) Peek(ctx context.Context, in *PeekRequest, opts ...grpc.CallOption) (*PeekResponse, error) {
	out := new(PeekResponse)
	err := c.cc.Invoke(ctx, "/pb.AccessTokenService/Peek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessTokenServiceClient) Revoke(ctx context.Context, in *RevokeRequest, opts ...grpc.CallOption) (*RevokeResponse, error) {
	out := new(RevokeResponse)
	err := c.cc.Invoke(ctx, "/pb.AccessTokenService/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessTokenServiceServer is the server API for AccessTokenService service.
type AccessTokenServiceServer interface {
	Issue(context.Context, *IssueRequest) (*IssueResponse, error)
	Peek(context.Context, *PeekRequest) (*PeekResponse, error)
	Revoke(context.Context, *RevokeRequest) (*RevokeResponse, error)
}

// UnimplementedAccessTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessTokenServiceServer struct {
}

func (*UnimplementedAccessTokenServiceServer) Issue(ctx context.Context, req *IssueRequest) (*IssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (*UnimplementedAccessTokenServiceServer) Peek(ctx context.Context, req *PeekRequest) (*PeekResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peek not implemented")
}
func (*UnimplementedAccessTokenServiceServer) Revoke(ctx context.Context, req *RevokeRequest) (*RevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}

func RegisterAccessTokenServiceServer(s *grpc.Server, srv AccessTokenServiceServer) {
	s.RegisterService(&_AccessTokenService_serviceDesc, srv)
}

func _AccessTokenService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccessTokenService/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).Issue(ctx, req.(*IssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessTokenService_Peek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeekRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).Peek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccessTokenService/Peek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).Peek(ctx, req.(*PeekRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessTokenService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccessTokenService/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).Revoke(ctx, req.(*RevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AccessTokenService",
	HandlerType: (*AccessTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _AccessTokenService_Issue_Handler,
		},
		{
			MethodName: "Peek",
			Handler:    _AccessTokenService_Peek_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _AccessTokenService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accesssvc.proto",
}
