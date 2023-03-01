// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResultType int32

const (
	ResultType_OK                    ResultType = 0
	ResultType_REQUEST_INVALID       ResultType = 1000
	ResultType_SESSION_INVALID       ResultType = 1001
	ResultType_SESSION_EXPIRE        ResultType = 1002
	ResultType_SIZE_LIMITED          ResultType = 1003
	ResultType_DB_ERROR              ResultType = 1004
	ResultType_NO_CHANGED            ResultType = 1005
	ResultType_NETWORK_ERROR         ResultType = 1006
	ResultType_DATA_ERROR            ResultType = 1007
	ResultType_ENCRYPT_ERROR         ResultType = 1008 /// lỗi khi chưa thiết lập kết nối authen
	ResultType_INTERNAL_SERVER_ERROR ResultType = 1009
	ResultType_STRANGE_REQUEST       ResultType = 1010
)

// Enum value maps for ResultType.
var (
	ResultType_name = map[int32]string{
		0:    "OK",
		1000: "REQUEST_INVALID",
		1001: "SESSION_INVALID",
		1002: "SESSION_EXPIRE",
		1003: "SIZE_LIMITED",
		1004: "DB_ERROR",
		1005: "NO_CHANGED",
		1006: "NETWORK_ERROR",
		1007: "DATA_ERROR",
		1008: "ENCRYPT_ERROR",
		1009: "INTERNAL_SERVER_ERROR",
		1010: "STRANGE_REQUEST",
	}
	ResultType_value = map[string]int32{
		"OK":                    0,
		"REQUEST_INVALID":       1000,
		"SESSION_INVALID":       1001,
		"SESSION_EXPIRE":        1002,
		"SIZE_LIMITED":          1003,
		"DB_ERROR":              1004,
		"NO_CHANGED":            1005,
		"NETWORK_ERROR":         1006,
		"DATA_ERROR":            1007,
		"ENCRYPT_ERROR":         1008,
		"INTERNAL_SERVER_ERROR": 1009,
		"STRANGE_REQUEST":       1010,
	}
)

func (x ResultType) Enum() *ResultType {
	p := new(ResultType)
	*p = x
	return p
}

func (x ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (ResultType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultType.Descriptor instead.
func (ResultType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type EncodeType int32

const (
	EncodeType_NONE EncodeType = 0  // không mã hóa playload
	EncodeType_XOR  EncodeType = 32 // mã hóa với xor key ( sv sử dụng xor key của client và client sử dụng xor key của server để decode payload )
	EncodeType_RSA  EncodeType = 64 // mã hóa với rsa public key của nhau để encode data
	EncodeType_AES  EncodeType = 96 // with key = 32 byte (AES-256-CBC)
)

// Enum value maps for EncodeType.
var (
	EncodeType_name = map[int32]string{
		0:  "NONE",
		32: "XOR",
		64: "RSA",
		96: "AES",
	}
	EncodeType_value = map[string]int32{
		"NONE": 0,
		"XOR":  32,
		"RSA":  64,
		"AES":  96,
	}
)

func (x EncodeType) Enum() *EncodeType {
	p := new(EncodeType)
	*p = x
	return p
}

func (x EncodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (EncodeType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x EncodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodeType.Descriptor instead.
func (EncodeType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type Platform int32

const (
	Platform_IOS     Platform = 0
	Platform_ANDROID Platform = 1
	Platform_WEB     Platform = 2
	Platform_WINDOW  Platform = 3
	Platform_MAC     Platform = 4
	Platform_LINUX   Platform = 5
	Platform_OTHER   Platform = 99
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0:  "IOS",
		1:  "ANDROID",
		2:  "WEB",
		3:  "WINDOW",
		4:  "MAC",
		5:  "LINUX",
		99: "OTHER",
	}
	Platform_value = map[string]int32{
		"IOS":     0,
		"ANDROID": 1,
		"WEB":     2,
		"WINDOW":  3,
		"MAC":     4,
		"LINUX":   5,
		"OTHER":   99,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[2].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[2]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type Group int32

const (
	Group_UNUSE      Group = 0
	Group_CONNECTION Group = 1 // ==> connect.api.proto
	Group_AUTHEN     Group = 2 // ==> authen.api.proto
	Group_ECHO       Group = 3 // ==> echo.api.proto
	Group_MARKET     Group = 4 // ==> market.api.proto
	Group_PNS        Group = 5 // apns.api.proto // send push notification ios,android,window,web
)

// Enum value maps for Group.
var (
	Group_name = map[int32]string{
		0: "UNUSE",
		1: "CONNECTION",
		2: "AUTHEN",
		3: "ECHO",
		4: "MARKET",
		5: "PNS",
	}
	Group_value = map[string]int32{
		"UNUSE":      0,
		"CONNECTION": 1,
		"AUTHEN":     2,
		"ECHO":       3,
		"MARKET":     4,
		"PNS":        5,
	}
)

func (x Group) Enum() *Group {
	p := new(Group)
	*p = x
	return p
}

func (x Group) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Group) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[3].Descriptor()
}

func (Group) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[3]
}

func (x Group) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Group.Descriptor instead.
func (Group) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

type Role int32

const (
	Role_ADMIN    Role = 0
	Role_PROVIDER Role = 1
	Role_USER     Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ADMIN",
		1: "PROVIDER",
		2: "USER",
	}
	Role_value = map[string]int32{
		"ADMIN":    0,
		"PROVIDER": 1,
		"USER":     2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[4].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[4]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionId  string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Expire     uint64 `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
	LastActive uint64 `protobuf:"varint,4,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
	Role       uint32 `protobuf:"varint,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Session) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Session) GetExpire() uint64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *Session) GetLastActive() uint64 {
	if x != nil {
		return x.LastActive
	}
	return 0
}

func (x *Session) GetRole() uint32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        uint32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Group       Group      `protobuf:"varint,2,opt,name=group,proto3,enum=Group" json:"group,omitempty"`
	Session     *Session   `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`                         // client không set khi sử dụng http tcp ws (http sử dụng V-Authorization ở Header mỗi request), đối với grpc và udp cần gửi thêm session_id để sv kiểm tra tính xác thực user
	Protocol    uint32     `protobuf:"varint,4,opt,name=protocol,proto3" json:"protocol,omitempty"`                      // client không set
	PayloadType uint32     `protobuf:"varint,5,opt,name=payloadType,proto3" json:"payloadType,omitempty"`                // client không set
	BinRequest  []byte     `protobuf:"bytes,6,opt,name=bin_request,json=binRequest,proto3" json:"bin_request,omitempty"` // client không set
	Request     *anypb.Any `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Request) GetGroup() Group {
	if x != nil {
		return x.Group
	}
	return Group_UNUSE
}

func (x *Request) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *Request) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *Request) GetPayloadType() uint32 {
	if x != nil {
		return x.PayloadType
	}
	return 0
}

func (x *Request) GetBinRequest() []byte {
	if x != nil {
		return x.BinRequest
	}
	return nil
}

func (x *Request) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   uint32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Reply    *anypb.Any `protobuf:"bytes,3,opt,name=reply,proto3" json:"reply,omitempty"`
	BinReply []byte     `protobuf:"bytes,4,opt,name=bin_reply,json=binReply,proto3" json:"bin_reply,omitempty"` // trường này luôn luôn null khi được trả về client
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Reply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Reply) GetReply() *anypb.Any {
	if x != nil {
		return x.Reply
	}
	return nil
}

func (x *Reply) GetBinReply() []byte {
	if x != nil {
		return x.BinReply
	}
	return nil
}

type Receive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       uint32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ServerTime uint64     `protobuf:"varint,2,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	Receive    *anypb.Any `protobuf:"bytes,3,opt,name=receive,proto3" json:"receive,omitempty"`
}

func (x *Receive) Reset() {
	*x = Receive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receive) ProtoMessage() {}

func (x *Receive) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receive.ProtoReflect.Descriptor instead.
func (*Receive) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Receive) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Receive) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *Receive) GetReceive() *anypb.Any {
	if x != nil {
		return x.Receive
	}
	return nil
}

type DiscoveryService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DiscoveryService) Reset() {
	*x = DiscoveryService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryService) ProtoMessage() {}

func (x *DiscoveryService) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryService.ProtoReflect.Descriptor instead.
func (*DiscoveryService) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *DiscoveryService) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DiscoveryService_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscoveryService_Request) Reset() {
	*x = DiscoveryService_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryService_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryService_Request) ProtoMessage() {}

func (x *DiscoveryService_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryService_Request.ProtoReflect.Descriptor instead.
func (*DiscoveryService_Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4, 0}
}

type DiscoveryService_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []uint32 `protobuf:"varint,1,rep,packed,name=types,proto3" json:"types,omitempty"`
}

func (x *DiscoveryService_Reply) Reset() {
	*x = DiscoveryService_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryService_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryService_Reply) ProtoMessage() {}

func (x *DiscoveryService_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryService_Reply.ProtoReflect.Descriptor instead.
func (*DiscoveryService_Reply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4, 1}
}

func (x *DiscoveryService_Reply) GetTypes() []uint32 {
	if x != nil {
		return x.Types
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x6e, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2a, 0xf3, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x0f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xe8, 0x07, 0x12,
	0x14, 0x0a, 0x0f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0xe9, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x10, 0xea, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x49,
	0x5a, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xeb, 0x07, 0x12, 0x0d, 0x0a,
	0x08, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xec, 0x07, 0x12, 0x0f, 0x0a, 0x0a,
	0x4e, 0x4f, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xed, 0x07, 0x12, 0x12, 0x0a,
	0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xee,
	0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0xef, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xf0, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0xf1, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xf2, 0x07, 0x2a, 0x31, 0x0a, 0x0a, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x58, 0x4f, 0x52, 0x10, 0x20, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41,
	0x10, 0x40, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x45, 0x53, 0x10, 0x60, 0x2a, 0x54, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x57, 0x45, 0x42, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x49, 0x4e, 0x55, 0x58, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10,
	0x63, 0x2a, 0x4d, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e,
	0x55, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x43, 0x48, 0x4f, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4e, 0x53, 0x10, 0x05,
	0x2a, 0x29, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(ResultType)(0),                  // 0: ResultType
	(EncodeType)(0),                  // 1: EncodeType
	(Platform)(0),                    // 2: Platform
	(Group)(0),                       // 3: Group
	(Role)(0),                        // 4: Role
	(*Session)(nil),                  // 5: Session
	(*Request)(nil),                  // 6: Request
	(*Reply)(nil),                    // 7: Reply
	(*Receive)(nil),                  // 8: Receive
	(*DiscoveryService)(nil),         // 9: DiscoveryService
	(*DiscoveryService_Request)(nil), // 10: DiscoveryService.Request
	(*DiscoveryService_Reply)(nil),   // 11: DiscoveryService.Reply
	(*anypb.Any)(nil),                // 12: google.protobuf.Any
}
var file_api_proto_depIdxs = []int32{
	3,  // 0: Request.group:type_name -> Group
	5,  // 1: Request.session:type_name -> Session
	12, // 2: Request.request:type_name -> google.protobuf.Any
	12, // 3: Reply.reply:type_name -> google.protobuf.Any
	12, // 4: Receive.receive:type_name -> google.protobuf.Any
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receive); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryService_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryService_Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
