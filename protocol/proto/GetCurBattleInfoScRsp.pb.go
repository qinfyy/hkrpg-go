// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetCurBattleInfoScRsp.proto

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

type GetCurBattleInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastEndStatus BattleEndStatus  `protobuf:"varint,1,opt,name=last_end_status,json=lastEndStatus,proto3,enum=BattleEndStatus" json:"last_end_status,omitempty"`
	LastEventId   uint32           `protobuf:"varint,4,opt,name=last_event_id,json=lastEventId,proto3" json:"last_event_id,omitempty"`
	Retcode       uint32           `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ABBCICNNNJH   *GFOBLINGFBL     `protobuf:"bytes,10,opt,name=ABBCICNNNJH,proto3" json:"ABBCICNNNJH,omitempty"`
	BattleInfo    *SceneBattleInfo `protobuf:"bytes,5,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
}

func (x *GetCurBattleInfoScRsp) Reset() {
	*x = GetCurBattleInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCurBattleInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurBattleInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurBattleInfoScRsp) ProtoMessage() {}

func (x *GetCurBattleInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCurBattleInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurBattleInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetCurBattleInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetCurBattleInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurBattleInfoScRsp) GetLastEndStatus() BattleEndStatus {
	if x != nil {
		return x.LastEndStatus
	}
	return BattleEndStatus_BATTLE_END_NONE
}

func (x *GetCurBattleInfoScRsp) GetLastEventId() uint32 {
	if x != nil {
		return x.LastEventId
	}
	return 0
}

func (x *GetCurBattleInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCurBattleInfoScRsp) GetABBCICNNNJH() *GFOBLINGFBL {
	if x != nil {
		return x.ABBCICNNNJH
	}
	return nil
}

func (x *GetCurBattleInfoScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

var File_GetCurBattleInfoScRsp_proto protoreflect.FileDescriptor

var file_GetCurBattleInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47,
	0x46, 0x4f, 0x42, 0x4c, 0x49, 0x4e, 0x47, 0x46, 0x42, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x2e, 0x0a, 0x0b, 0x41, 0x42, 0x42, 0x43, 0x49, 0x43, 0x4e, 0x4e, 0x4e, 0x4a, 0x48, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x46, 0x4f, 0x42, 0x4c, 0x49, 0x4e, 0x47,
	0x46, 0x42, 0x4c, 0x52, 0x0b, 0x41, 0x42, 0x42, 0x43, 0x49, 0x43, 0x4e, 0x4e, 0x4e, 0x4a, 0x48,
	0x12, 0x31, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCurBattleInfoScRsp_proto_rawDescOnce sync.Once
	file_GetCurBattleInfoScRsp_proto_rawDescData = file_GetCurBattleInfoScRsp_proto_rawDesc
)

func file_GetCurBattleInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetCurBattleInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetCurBattleInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCurBattleInfoScRsp_proto_rawDescData)
	})
	return file_GetCurBattleInfoScRsp_proto_rawDescData
}

var file_GetCurBattleInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCurBattleInfoScRsp_proto_goTypes = []interface{}{
	(*GetCurBattleInfoScRsp)(nil), // 0: GetCurBattleInfoScRsp
	(BattleEndStatus)(0),          // 1: BattleEndStatus
	(*GFOBLINGFBL)(nil),           // 2: GFOBLINGFBL
	(*SceneBattleInfo)(nil),       // 3: SceneBattleInfo
}
var file_GetCurBattleInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetCurBattleInfoScRsp.last_end_status:type_name -> BattleEndStatus
	2, // 1: GetCurBattleInfoScRsp.ABBCICNNNJH:type_name -> GFOBLINGFBL
	3, // 2: GetCurBattleInfoScRsp.battle_info:type_name -> SceneBattleInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetCurBattleInfoScRsp_proto_init() }
func file_GetCurBattleInfoScRsp_proto_init() {
	if File_GetCurBattleInfoScRsp_proto != nil {
		return
	}
	file_GFOBLINGFBL_proto_init()
	file_BattleEndStatus_proto_init()
	file_SceneBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCurBattleInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurBattleInfoScRsp); i {
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
			RawDescriptor: file_GetCurBattleInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCurBattleInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetCurBattleInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetCurBattleInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetCurBattleInfoScRsp_proto = out.File
	file_GetCurBattleInfoScRsp_proto_rawDesc = nil
	file_GetCurBattleInfoScRsp_proto_goTypes = nil
	file_GetCurBattleInfoScRsp_proto_depIdxs = nil
}
