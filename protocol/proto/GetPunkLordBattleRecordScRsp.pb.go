// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPunkLordBattleRecordScRsp.proto

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

type GetPunkLordBattleRecordScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode          uint32                  `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BattleRecordList []*PunkLordBattleRecord `protobuf:"bytes,13,rep,name=battle_record_list,json=battleRecordList,proto3" json:"battle_record_list,omitempty"`
	BattleReplayList []*PunkLordBattleReplay `protobuf:"bytes,15,rep,name=battle_replay_list,json=battleReplayList,proto3" json:"battle_replay_list,omitempty"`
	MonsterKey       *PunkLordMonsterKey     `protobuf:"bytes,7,opt,name=monster_key,json=monsterKey,proto3" json:"monster_key,omitempty"`
}

func (x *GetPunkLordBattleRecordScRsp) Reset() {
	*x = GetPunkLordBattleRecordScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPunkLordBattleRecordScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPunkLordBattleRecordScRsp) ProtoMessage() {}

func (x *GetPunkLordBattleRecordScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPunkLordBattleRecordScRsp.ProtoReflect.Descriptor instead.
func (*GetPunkLordBattleRecordScRsp) Descriptor() ([]byte, []int) {
	return file_GetPunkLordBattleRecordScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPunkLordBattleRecordScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPunkLordBattleRecordScRsp) GetBattleRecordList() []*PunkLordBattleRecord {
	if x != nil {
		return x.BattleRecordList
	}
	return nil
}

func (x *GetPunkLordBattleRecordScRsp) GetBattleReplayList() []*PunkLordBattleReplay {
	if x != nil {
		return x.BattleReplayList
	}
	return nil
}

func (x *GetPunkLordBattleRecordScRsp) GetMonsterKey() *PunkLordMonsterKey {
	if x != nil {
		return x.MonsterKey
	}
	return nil
}

var File_GetPunkLordBattleRecordScRsp_proto protoreflect.FileDescriptor

var file_GetPunkLordBattleRecordScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x50, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x43, 0x0a, 0x12, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x12, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetPunkLordBattleRecordScRsp_proto_rawDescOnce sync.Once
	file_GetPunkLordBattleRecordScRsp_proto_rawDescData = file_GetPunkLordBattleRecordScRsp_proto_rawDesc
)

func file_GetPunkLordBattleRecordScRsp_proto_rawDescGZIP() []byte {
	file_GetPunkLordBattleRecordScRsp_proto_rawDescOnce.Do(func() {
		file_GetPunkLordBattleRecordScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPunkLordBattleRecordScRsp_proto_rawDescData)
	})
	return file_GetPunkLordBattleRecordScRsp_proto_rawDescData
}

var file_GetPunkLordBattleRecordScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPunkLordBattleRecordScRsp_proto_goTypes = []interface{}{
	(*GetPunkLordBattleRecordScRsp)(nil), // 0: GetPunkLordBattleRecordScRsp
	(*PunkLordBattleRecord)(nil),         // 1: PunkLordBattleRecord
	(*PunkLordBattleReplay)(nil),         // 2: PunkLordBattleReplay
	(*PunkLordMonsterKey)(nil),           // 3: PunkLordMonsterKey
}
var file_GetPunkLordBattleRecordScRsp_proto_depIdxs = []int32{
	1, // 0: GetPunkLordBattleRecordScRsp.battle_record_list:type_name -> PunkLordBattleRecord
	2, // 1: GetPunkLordBattleRecordScRsp.battle_replay_list:type_name -> PunkLordBattleReplay
	3, // 2: GetPunkLordBattleRecordScRsp.monster_key:type_name -> PunkLordMonsterKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetPunkLordBattleRecordScRsp_proto_init() }
func file_GetPunkLordBattleRecordScRsp_proto_init() {
	if File_GetPunkLordBattleRecordScRsp_proto != nil {
		return
	}
	file_PunkLordBattleReplay_proto_init()
	file_PunkLordBattleRecord_proto_init()
	file_PunkLordMonsterKey_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPunkLordBattleRecordScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPunkLordBattleRecordScRsp); i {
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
			RawDescriptor: file_GetPunkLordBattleRecordScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPunkLordBattleRecordScRsp_proto_goTypes,
		DependencyIndexes: file_GetPunkLordBattleRecordScRsp_proto_depIdxs,
		MessageInfos:      file_GetPunkLordBattleRecordScRsp_proto_msgTypes,
	}.Build()
	File_GetPunkLordBattleRecordScRsp_proto = out.File
	file_GetPunkLordBattleRecordScRsp_proto_rawDesc = nil
	file_GetPunkLordBattleRecordScRsp_proto_goTypes = nil
	file_GetPunkLordBattleRecordScRsp_proto_depIdxs = nil
}
