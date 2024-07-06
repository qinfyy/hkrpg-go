// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournInfo.proto

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

type RogueTournInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueTournSaveList       []*RogueTournSaveList       `protobuf:"bytes,11,rep,name=rogue_tourn_save_list,json=rogueTournSaveList,proto3" json:"rogue_tourn_save_list,omitempty"`
	RogueTournAreaInfo       []*RogueTournAreaInfo       `protobuf:"bytes,10,rep,name=rogue_tourn_area_info,json=rogueTournAreaInfo,proto3" json:"rogue_tourn_area_info,omitempty"`
	InspirationCircuit       *InspirationCircuitInfo     `protobuf:"bytes,2,opt,name=inspiration_circuit,json=inspirationCircuit,proto3" json:"inspiration_circuit,omitempty"`
	RogueTournSeasonInfo     *RogueTournSeasonInfo       `protobuf:"bytes,4,opt,name=rogue_tourn_season_info,json=rogueTournSeasonInfo,proto3" json:"rogue_tourn_season_info,omitempty"`
	ExtraScoreInfo           *ExtraScoreInfo             `protobuf:"bytes,14,opt,name=extra_score_info,json=extraScoreInfo,proto3" json:"extra_score_info,omitempty"`
	RogueTournExpInfo        *RogueTournExpInfo          `protobuf:"bytes,3,opt,name=rogue_tourn_exp_info,json=rogueTournExpInfo,proto3" json:"rogue_tourn_exp_info,omitempty"`
	RogueTournHandbook       *RogueTournHandbookInfo     `protobuf:"bytes,8,opt,name=rogue_tourn_handbook,json=rogueTournHandbook,proto3" json:"rogue_tourn_handbook,omitempty"`
	RogueTournDifficultyInfo []*RogueTournDifficultyInfo `protobuf:"bytes,9,rep,name=rogue_tourn_difficulty_info,json=rogueTournDifficultyInfo,proto3" json:"rogue_tourn_difficulty_info,omitempty"`
}

func (x *RogueTournInfo) Reset() {
	*x = RogueTournInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournInfo) ProtoMessage() {}

func (x *RogueTournInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournInfo.ProtoReflect.Descriptor instead.
func (*RogueTournInfo) Descriptor() ([]byte, []int) {
	return file_RogueTournInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournInfo) GetRogueTournSaveList() []*RogueTournSaveList {
	if x != nil {
		return x.RogueTournSaveList
	}
	return nil
}

func (x *RogueTournInfo) GetRogueTournAreaInfo() []*RogueTournAreaInfo {
	if x != nil {
		return x.RogueTournAreaInfo
	}
	return nil
}

func (x *RogueTournInfo) GetInspirationCircuit() *InspirationCircuitInfo {
	if x != nil {
		return x.InspirationCircuit
	}
	return nil
}

func (x *RogueTournInfo) GetRogueTournSeasonInfo() *RogueTournSeasonInfo {
	if x != nil {
		return x.RogueTournSeasonInfo
	}
	return nil
}

func (x *RogueTournInfo) GetExtraScoreInfo() *ExtraScoreInfo {
	if x != nil {
		return x.ExtraScoreInfo
	}
	return nil
}

func (x *RogueTournInfo) GetRogueTournExpInfo() *RogueTournExpInfo {
	if x != nil {
		return x.RogueTournExpInfo
	}
	return nil
}

func (x *RogueTournInfo) GetRogueTournHandbook() *RogueTournHandbookInfo {
	if x != nil {
		return x.RogueTournHandbook
	}
	return nil
}

func (x *RogueTournInfo) GetRogueTournDifficultyInfo() []*RogueTournDifficultyInfo {
	if x != nil {
		return x.RogueTournDifficultyInfo
	}
	return nil
}

var File_RogueTournInfo_proto protoreflect.FileDescriptor

var file_RogueTournInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x61, 0x76, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x45, 0x78, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x49, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdd, 0x04, 0x0a, 0x0e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x15, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f,
	0x75, 0x72, 0x6e, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x15,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x12, 0x4c,
	0x0a, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x10,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x45, 0x78, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x45, 0x78, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x14,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x58, 0x0a, 0x1b, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x18, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RogueTournInfo_proto_rawDescOnce sync.Once
	file_RogueTournInfo_proto_rawDescData = file_RogueTournInfo_proto_rawDesc
)

func file_RogueTournInfo_proto_rawDescGZIP() []byte {
	file_RogueTournInfo_proto_rawDescOnce.Do(func() {
		file_RogueTournInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournInfo_proto_rawDescData)
	})
	return file_RogueTournInfo_proto_rawDescData
}

var file_RogueTournInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournInfo_proto_goTypes = []interface{}{
	(*RogueTournInfo)(nil),           // 0: RogueTournInfo
	(*RogueTournSaveList)(nil),       // 1: RogueTournSaveList
	(*RogueTournAreaInfo)(nil),       // 2: RogueTournAreaInfo
	(*InspirationCircuitInfo)(nil),   // 3: InspirationCircuitInfo
	(*RogueTournSeasonInfo)(nil),     // 4: RogueTournSeasonInfo
	(*ExtraScoreInfo)(nil),           // 5: ExtraScoreInfo
	(*RogueTournExpInfo)(nil),        // 6: RogueTournExpInfo
	(*RogueTournHandbookInfo)(nil),   // 7: RogueTournHandbookInfo
	(*RogueTournDifficultyInfo)(nil), // 8: RogueTournDifficultyInfo
}
var file_RogueTournInfo_proto_depIdxs = []int32{
	1, // 0: RogueTournInfo.rogue_tourn_save_list:type_name -> RogueTournSaveList
	2, // 1: RogueTournInfo.rogue_tourn_area_info:type_name -> RogueTournAreaInfo
	3, // 2: RogueTournInfo.inspiration_circuit:type_name -> InspirationCircuitInfo
	4, // 3: RogueTournInfo.rogue_tourn_season_info:type_name -> RogueTournSeasonInfo
	5, // 4: RogueTournInfo.extra_score_info:type_name -> ExtraScoreInfo
	6, // 5: RogueTournInfo.rogue_tourn_exp_info:type_name -> RogueTournExpInfo
	7, // 6: RogueTournInfo.rogue_tourn_handbook:type_name -> RogueTournHandbookInfo
	8, // 7: RogueTournInfo.rogue_tourn_difficulty_info:type_name -> RogueTournDifficultyInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RogueTournInfo_proto_init() }
func file_RogueTournInfo_proto_init() {
	if File_RogueTournInfo_proto != nil {
		return
	}
	file_RogueTournAreaInfo_proto_init()
	file_RogueTournSaveList_proto_init()
	file_RogueTournHandbookInfo_proto_init()
	file_ExtraScoreInfo_proto_init()
	file_RogueTournExpInfo_proto_init()
	file_InspirationCircuitInfo_proto_init()
	file_RogueTournDifficultyInfo_proto_init()
	file_RogueTournSeasonInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournInfo); i {
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
			RawDescriptor: file_RogueTournInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournInfo_proto_goTypes,
		DependencyIndexes: file_RogueTournInfo_proto_depIdxs,
		MessageInfos:      file_RogueTournInfo_proto_msgTypes,
	}.Build()
	File_RogueTournInfo_proto = out.File
	file_RogueTournInfo_proto_rawDesc = nil
	file_RogueTournInfo_proto_goTypes = nil
	file_RogueTournInfo_proto_depIdxs = nil
}