// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: cmd.server.proto

package proto

import (
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

type ServiceConnectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3,enum=proto.ServerType" json:"server_type,omitempty"`
	AppId      string     `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Addr       string     `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Port       string     `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ServiceConnectionReq) Reset() {
	*x = ServiceConnectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConnectionReq) ProtoMessage() {}

func (x *ServiceConnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConnectionReq.ProtoReflect.Descriptor instead.
func (*ServiceConnectionReq) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceConnectionReq) GetServerType() ServerType {
	if x != nil {
		return x.ServerType
	}
	return ServerType_SERVICE_NONE
}

func (x *ServiceConnectionReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ServiceConnectionReq) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ServiceConnectionReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type ServiceConnectionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3,enum=proto.ServerType" json:"server_type,omitempty"`
	AppId      string     `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *ServiceConnectionRsp) Reset() {
	*x = ServiceConnectionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConnectionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConnectionRsp) ProtoMessage() {}

func (x *ServiceConnectionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConnectionRsp.ProtoReflect.Descriptor instead.
func (*ServiceConnectionRsp) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceConnectionRsp) GetServerType() ServerType {
	if x != nil {
		return x.ServerType
	}
	return ServerType_SERVICE_NONE
}

func (x *ServiceConnectionRsp) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type PlayerLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid uint32 `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
}

func (x *PlayerLoginReq) Reset() {
	*x = PlayerLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginReq) ProtoMessage() {}

func (x *PlayerLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginReq.ProtoReflect.Descriptor instead.
func (*PlayerLoginReq) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerLoginReq) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

type PlayerLoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid uint32 `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
}

func (x *PlayerLoginRsp) Reset() {
	*x = PlayerLoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginRsp) ProtoMessage() {}

func (x *PlayerLoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginRsp.ProtoReflect.Descriptor instead.
func (*PlayerLoginRsp) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerLoginRsp) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

type PlayerToGameByGateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType int64  `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	PlayerBin   []byte `protobuf:"bytes,2,opt,name=player_bin,json=playerBin,proto3" json:"player_bin,omitempty"`
}

func (x *PlayerToGameByGateReq) Reset() {
	*x = PlayerToGameByGateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerToGameByGateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerToGameByGateReq) ProtoMessage() {}

func (x *PlayerToGameByGateReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerToGameByGateReq.ProtoReflect.Descriptor instead.
func (*PlayerToGameByGateReq) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerToGameByGateReq) GetMessageType() int64 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *PlayerToGameByGateReq) GetPlayerBin() []byte {
	if x != nil {
		return x.PlayerBin
	}
	return nil
}

type PlayerToGameByGateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType int64  `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	PlayerBin   []byte `protobuf:"bytes,2,opt,name=player_bin,json=playerBin,proto3" json:"player_bin,omitempty"`
}

func (x *PlayerToGameByGateRsp) Reset() {
	*x = PlayerToGameByGateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerToGameByGateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerToGameByGateRsp) ProtoMessage() {}

func (x *PlayerToGameByGateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerToGameByGateRsp.ProtoReflect.Descriptor instead.
func (*PlayerToGameByGateRsp) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerToGameByGateRsp) GetMessageType() int64 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *PlayerToGameByGateRsp) GetPlayerBin() []byte {
	if x != nil {
		return x.PlayerBin
	}
	return nil
}

type GetServerOuterAddrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3,enum=proto.ServerType" json:"server_type,omitempty"`
	AppId      string     `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	PlayerNum  uint64     `protobuf:"varint,3,opt,name=player_num,json=playerNum,proto3" json:"player_num,omitempty"`
}

func (x *GetServerOuterAddrReq) Reset() {
	*x = GetServerOuterAddrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerOuterAddrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerOuterAddrReq) ProtoMessage() {}

func (x *GetServerOuterAddrReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerOuterAddrReq.ProtoReflect.Descriptor instead.
func (*GetServerOuterAddrReq) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{6}
}

func (x *GetServerOuterAddrReq) GetServerType() ServerType {
	if x != nil {
		return x.ServerType
	}
	return ServerType_SERVICE_NONE
}

func (x *GetServerOuterAddrReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *GetServerOuterAddrReq) GetPlayerNum() uint64 {
	if x != nil {
		return x.PlayerNum
	}
	return 0
}

type GetServerOuterAddrRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType ServerType `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3,enum=proto.ServerType" json:"server_type,omitempty"`
	Addr       string     `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Port       string     `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GetServerOuterAddrRsp) Reset() {
	*x = GetServerOuterAddrRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerOuterAddrRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerOuterAddrRsp) ProtoMessage() {}

func (x *GetServerOuterAddrRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerOuterAddrRsp.ProtoReflect.Descriptor instead.
func (*GetServerOuterAddrRsp) Descriptor() ([]byte, []int) {
	return file_cmd_server_proto_rawDescGZIP(), []int{7}
}

func (x *GetServerOuterAddrRsp) GetServerType() ServerType {
	if x != nil {
		return x.ServerType
	}
	return ServerType_SERVICE_NONE
}

func (x *GetServerOuterAddrRsp) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GetServerOuterAddrRsp) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

var File_cmd_server_proto protoreflect.FileDescriptor

var file_cmd_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a,
	0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x61, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x12, 0x32, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x22, 0x59, 0x0a,
	0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x47,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x22, 0x59, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x54, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x47, 0x61, 0x74, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x42, 0x69, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x52, 0x73, 0x70,
	0x12, 0x32, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_server_proto_rawDescOnce sync.Once
	file_cmd_server_proto_rawDescData = file_cmd_server_proto_rawDesc
)

func file_cmd_server_proto_rawDescGZIP() []byte {
	file_cmd_server_proto_rawDescOnce.Do(func() {
		file_cmd_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_server_proto_rawDescData)
	})
	return file_cmd_server_proto_rawDescData
}

var file_cmd_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_server_proto_goTypes = []interface{}{
	(*ServiceConnectionReq)(nil),  // 0: proto.ServiceConnectionReq
	(*ServiceConnectionRsp)(nil),  // 1: proto.ServiceConnectionRsp
	(*PlayerLoginReq)(nil),        // 2: proto.PlayerLoginReq
	(*PlayerLoginRsp)(nil),        // 3: proto.PlayerLoginRsp
	(*PlayerToGameByGateReq)(nil), // 4: proto.PlayerToGameByGateReq
	(*PlayerToGameByGateRsp)(nil), // 5: proto.PlayerToGameByGateRsp
	(*GetServerOuterAddrReq)(nil), // 6: proto.GetServerOuterAddrReq
	(*GetServerOuterAddrRsp)(nil), // 7: proto.GetServerOuterAddrRsp
	(ServerType)(0),               // 8: proto.ServerType
}
var file_cmd_server_proto_depIdxs = []int32{
	8, // 0: proto.ServiceConnectionReq.server_type:type_name -> proto.ServerType
	8, // 1: proto.ServiceConnectionRsp.server_type:type_name -> proto.ServerType
	8, // 2: proto.GetServerOuterAddrReq.server_type:type_name -> proto.ServerType
	8, // 3: proto.GetServerOuterAddrRsp.server_type:type_name -> proto.ServerType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cmd_server_proto_init() }
func file_cmd_server_proto_init() {
	if File_cmd_server_proto != nil {
		return
	}
	file_enum_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConnectionReq); i {
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
		file_cmd_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConnectionRsp); i {
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
		file_cmd_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginReq); i {
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
		file_cmd_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginRsp); i {
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
		file_cmd_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerToGameByGateReq); i {
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
		file_cmd_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerToGameByGateRsp); i {
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
		file_cmd_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerOuterAddrReq); i {
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
		file_cmd_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerOuterAddrRsp); i {
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
			RawDescriptor: file_cmd_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_server_proto_goTypes,
		DependencyIndexes: file_cmd_server_proto_depIdxs,
		MessageInfos:      file_cmd_server_proto_msgTypes,
	}.Build()
	File_cmd_server_proto = out.File
	file_cmd_server_proto_rawDesc = nil
	file_cmd_server_proto_goTypes = nil
	file_cmd_server_proto_depIdxs = nil
}
