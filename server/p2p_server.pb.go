// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: p2p_server.proto

package server

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MsgType int32

const (
	MsgType_HEARTBEAT_REQ  MsgType = 0
	MsgType_HEARTBEAT_RESP MsgType = 10
	MsgType_LOGIN_REQ      MsgType = 1
	MsgType_LOGIN_RESP     MsgType = 11
	MsgType_LOGOUT_REQ     MsgType = 2
	MsgType_LOGOUT_RESP    MsgType = 12
	MsgType_PUNCH_REQ      MsgType = 3
	MsgType_PUNCH_RESP     MsgType = 13
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0:  "HEARTBEAT_REQ",
		10: "HEARTBEAT_RESP",
		1:  "LOGIN_REQ",
		11: "LOGIN_RESP",
		2:  "LOGOUT_REQ",
		12: "LOGOUT_RESP",
		3:  "PUNCH_REQ",
		13: "PUNCH_RESP",
	}
	MsgType_value = map[string]int32{
		"HEARTBEAT_REQ":  0,
		"HEARTBEAT_RESP": 10,
		"LOGIN_REQ":      1,
		"LOGIN_RESP":     11,
		"LOGOUT_REQ":     2,
		"LOGOUT_RESP":    12,
		"PUNCH_REQ":      3,
		"PUNCH_RESP":     13,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_server_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_p2p_server_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{0}
}

type RetCode int32

const (
	RetCode_OK           RetCode = 0
	RetCode_FAIL         RetCode = 1
	RetCode_UserNotFound RetCode = 2
)

// Enum value maps for RetCode.
var (
	RetCode_name = map[int32]string{
		0: "OK",
		1: "FAIL",
		2: "UserNotFound",
	}
	RetCode_value = map[string]int32{
		"OK":           0,
		"FAIL":         1,
		"UserNotFound": 2,
	}
)

func (x RetCode) Enum() *RetCode {
	p := new(RetCode)
	*p = x
	return p
}

func (x RetCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetCode) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_server_proto_enumTypes[1].Descriptor()
}

func (RetCode) Type() protoreflect.EnumType {
	return &file_p2p_server_proto_enumTypes[1]
}

func (x RetCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetCode.Descriptor instead.
func (RetCode) EnumDescriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{1}
}

type MsgHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MsgType `protobuf:"varint,1,opt,name=type,proto3,enum=MsgType" json:"type,omitempty"`
}

func (x *MsgHeader) Reset() {
	*x = MsgHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHeader) ProtoMessage() {}

func (x *MsgHeader) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHeader.ProtoReflect.Descriptor instead.
func (*MsgHeader) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{0}
}

func (x *MsgHeader) GetType() MsgType {
	if x != nil {
		return x.Type
	}
	return MsgType_HEARTBEAT_REQ
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *MsgHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Body   []byte     `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetHeader() *MsgHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Msg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type HeartbeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *HeartbeatReq) Reset() {
	*x = HeartbeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReq) ProtoMessage() {}

func (x *HeartbeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReq.ProtoReflect.Descriptor instead.
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatReq) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *HeartbeatReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HeartbeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result RetCode `protobuf:"varint,1,opt,name=Result,proto3,enum=RetCode" json:"Result,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ID     uint64  `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *HeartbeatResp) Reset() {
	*x = HeartbeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResp) ProtoMessage() {}

func (x *HeartbeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResp.ProtoReflect.Descriptor instead.
func (*HeartbeatResp) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResp) GetResult() RetCode {
	if x != nil {
		return x.Result
	}
	return RetCode_OK
}

func (x *HeartbeatResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *HeartbeatResp) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

// login
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result RetCode `protobuf:"varint,1,opt,name=Result,proto3,enum=RetCode" json:"Result,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ID     uint64  `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetResult() RetCode {
	if x != nil {
		return x.Result
	}
	return RetCode_OK
}

func (x *LoginResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginResp) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

// logout
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutReq) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type LogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result RetCode `protobuf:"varint,1,opt,name=Result,proto3,enum=RetCode" json:"Result,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *LogoutResp) Reset() {
	*x = LogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResp) ProtoMessage() {}

func (x *LogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResp.ProtoReflect.Descriptor instead.
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{7}
}

func (x *LogoutResp) GetResult() RetCode {
	if x != nil {
		return x.Result
	}
	return RetCode_OK
}

func (x *LogoutResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// punch
type PunchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TargetID int64  `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
}

func (x *PunchReq) Reset() {
	*x = PunchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunchReq) ProtoMessage() {}

func (x *PunchReq) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunchReq.ProtoReflect.Descriptor instead.
func (*PunchReq) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{8}
}

func (x *PunchReq) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PunchReq) GetTargetID() int64 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

type PunchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result RetCode `protobuf:"varint,1,opt,name=Result,proto3,enum=RetCode" json:"Result,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	IP     string  `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Port   uint32  `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *PunchResp) Reset() {
	*x = PunchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunchResp) ProtoMessage() {}

func (x *PunchResp) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunchResp.ProtoReflect.Descriptor instead.
func (*PunchResp) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{9}
}

func (x *PunchResp) GetResult() RetCode {
	if x != nil {
		return x.Result
	}
	return RetCode_OK
}

func (x *PunchResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PunchResp) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *PunchResp) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// 客户端收到打洞请求
type GetPunchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP   string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *GetPunchReq) Reset() {
	*x = GetPunchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPunchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPunchReq) ProtoMessage() {}

func (x *GetPunchReq) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPunchReq.ProtoReflect.Descriptor instead.
func (*GetPunchReq) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{10}
}

func (x *GetPunchReq) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *GetPunchReq) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_p2p_server_proto protoreflect.FileDescriptor

var file_p2p_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x32, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x30, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x53,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x08, 0x2e, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x1e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x1b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x40, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x08, 0x2e, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0x36, 0x0a, 0x08, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x09, 0x50,
	0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x2a, 0x8f, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x51,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x4e, 0x43, 0x48, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x10, 0x0d, 0x2a, 0x2d, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_server_proto_rawDescOnce sync.Once
	file_p2p_server_proto_rawDescData = file_p2p_server_proto_rawDesc
)

func file_p2p_server_proto_rawDescGZIP() []byte {
	file_p2p_server_proto_rawDescOnce.Do(func() {
		file_p2p_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_server_proto_rawDescData)
	})
	return file_p2p_server_proto_rawDescData
}

var file_p2p_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_p2p_server_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_p2p_server_proto_goTypes = []interface{}{
	(MsgType)(0),          // 0: MsgType
	(RetCode)(0),          // 1: RetCode
	(*MsgHeader)(nil),     // 2: MsgHeader
	(*Msg)(nil),           // 3: Msg
	(*HeartbeatReq)(nil),  // 4: HeartbeatReq
	(*HeartbeatResp)(nil), // 5: HeartbeatResp
	(*LoginReq)(nil),      // 6: LoginReq
	(*LoginResp)(nil),     // 7: LoginResp
	(*LogoutReq)(nil),     // 8: LogoutReq
	(*LogoutResp)(nil),    // 9: LogoutResp
	(*PunchReq)(nil),      // 10: PunchReq
	(*PunchResp)(nil),     // 11: PunchResp
	(*GetPunchReq)(nil),   // 12: GetPunchReq
}
var file_p2p_server_proto_depIdxs = []int32{
	0, // 0: MsgHeader.type:type_name -> MsgType
	2, // 1: Msg.Header:type_name -> MsgHeader
	1, // 2: HeartbeatResp.Result:type_name -> RetCode
	1, // 3: LoginResp.Result:type_name -> RetCode
	1, // 4: LogoutResp.Result:type_name -> RetCode
	1, // 5: PunchResp.Result:type_name -> RetCode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_p2p_server_proto_init() }
func file_p2p_server_proto_init() {
	if File_p2p_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHeader); i {
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
		file_p2p_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_p2p_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatReq); i {
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
		file_p2p_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResp); i {
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
		file_p2p_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_p2p_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_p2p_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_p2p_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResp); i {
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
		file_p2p_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunchReq); i {
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
		file_p2p_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunchResp); i {
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
		file_p2p_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPunchReq); i {
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
			RawDescriptor: file_p2p_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_server_proto_goTypes,
		DependencyIndexes: file_p2p_server_proto_depIdxs,
		EnumInfos:         file_p2p_server_proto_enumTypes,
		MessageInfos:      file_p2p_server_proto_msgTypes,
	}.Build()
	File_p2p_server_proto = out.File
	file_p2p_server_proto_rawDesc = nil
	file_p2p_server_proto_goTypes = nil
	file_p2p_server_proto_depIdxs = nil
}
