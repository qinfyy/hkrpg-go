// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordMonsterBasicInfo.proto

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

type PunkLordMonsterBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint32            `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	MonsterId   uint32            `protobuf:"varint,2,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	ConfigId    uint32            `protobuf:"varint,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	WorldLevel  uint32            `protobuf:"varint,4,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	CreateTime  int64             `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LeftHp      uint32            `protobuf:"varint,6,opt,name=left_hp,json=leftHp,proto3" json:"left_hp,omitempty"`
	AttackerNum uint32            `protobuf:"varint,7,opt,name=attacker_num,json=attackerNum,proto3" json:"attacker_num,omitempty"`
	ShareType   PunkLordShareType `protobuf:"varint,8,opt,name=share_type,json=shareType,proto3,enum=PunkLordShareType" json:"share_type,omitempty"`
	KNACMDEJFAF bool              `protobuf:"varint,9,opt,name=KNACMDEJFAF,proto3" json:"KNACMDEJFAF,omitempty"`
}

func (x *PunkLordMonsterBasicInfo) Reset() {
	*x = PunkLordMonsterBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PunkLordMonsterBasicInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunkLordMonsterBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunkLordMonsterBasicInfo) ProtoMessage() {}

func (x *PunkLordMonsterBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PunkLordMonsterBasicInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunkLordMonsterBasicInfo.ProtoReflect.Descriptor instead.
func (*PunkLordMonsterBasicInfo) Descriptor() ([]byte, []int) {
	return file_PunkLordMonsterBasicInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PunkLordMonsterBasicInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetLeftHp() uint32 {
	if x != nil {
		return x.LeftHp
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetAttackerNum() uint32 {
	if x != nil {
		return x.AttackerNum
	}
	return 0
}

func (x *PunkLordMonsterBasicInfo) GetShareType() PunkLordShareType {
	if x != nil {
		return x.ShareType
	}
	return PunkLordShareType_PUNK_LORD_SHARE_TYPE_NONE
}

func (x *PunkLordMonsterBasicInfo) GetKNACMDEJFAF() bool {
	if x != nil {
		return x.KNACMDEJFAF
	}
	return false
}

var File_PunkLordMonsterBasicInfo_proto protoreflect.FileDescriptor

var file_PunkLordMonsterBasicInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x18, 0x50, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x68,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x66, 0x74, 0x48, 0x70, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72,
	0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4e, 0x41, 0x43, 0x4d, 0x44, 0x45,
	0x4a, 0x46, 0x41, 0x46, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x4e, 0x41, 0x43,
	0x4d, 0x44, 0x45, 0x4a, 0x46, 0x41, 0x46, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PunkLordMonsterBasicInfo_proto_rawDescOnce sync.Once
	file_PunkLordMonsterBasicInfo_proto_rawDescData = file_PunkLordMonsterBasicInfo_proto_rawDesc
)

func file_PunkLordMonsterBasicInfo_proto_rawDescGZIP() []byte {
	file_PunkLordMonsterBasicInfo_proto_rawDescOnce.Do(func() {
		file_PunkLordMonsterBasicInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordMonsterBasicInfo_proto_rawDescData)
	})
	return file_PunkLordMonsterBasicInfo_proto_rawDescData
}

var file_PunkLordMonsterBasicInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PunkLordMonsterBasicInfo_proto_goTypes = []interface{}{
	(*PunkLordMonsterBasicInfo)(nil), // 0: PunkLordMonsterBasicInfo
	(PunkLordShareType)(0),           // 1: PunkLordShareType
}
var file_PunkLordMonsterBasicInfo_proto_depIdxs = []int32{
	1, // 0: PunkLordMonsterBasicInfo.share_type:type_name -> PunkLordShareType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PunkLordMonsterBasicInfo_proto_init() }
func file_PunkLordMonsterBasicInfo_proto_init() {
	if File_PunkLordMonsterBasicInfo_proto != nil {
		return
	}
	file_PunkLordShareType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PunkLordMonsterBasicInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunkLordMonsterBasicInfo); i {
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
			RawDescriptor: file_PunkLordMonsterBasicInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordMonsterBasicInfo_proto_goTypes,
		DependencyIndexes: file_PunkLordMonsterBasicInfo_proto_depIdxs,
		MessageInfos:      file_PunkLordMonsterBasicInfo_proto_msgTypes,
	}.Build()
	File_PunkLordMonsterBasicInfo_proto = out.File
	file_PunkLordMonsterBasicInfo_proto_rawDesc = nil
	file_PunkLordMonsterBasicInfo_proto_goTypes = nil
	file_PunkLordMonsterBasicInfo_proto_depIdxs = nil
}