// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournCurGameInfo.proto

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

type RogueTournCurGameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleInfo            *ChessRogueMiracleInfo  `protobuf:"bytes,13,opt,name=miracle_info,json=miracleInfo,proto3" json:"miracle_info,omitempty"`
	UnlockValue            *KeywordUnlockValue     `protobuf:"bytes,2,opt,name=unlock_value,json=unlockValue,proto3" json:"unlock_value,omitempty"`
	ItemValue              *RogueGameItemValue     `protobuf:"bytes,3,opt,name=item_value,json=itemValue,proto3" json:"item_value,omitempty"`
	TournFormulaInfo       *RogueTournFormulaInfo  `protobuf:"bytes,6,opt,name=tourn_formula_info,json=tournFormulaInfo,proto3" json:"tourn_formula_info,omitempty"`
	Level                  *RogueTournLevelInfo    `protobuf:"bytes,11,opt,name=level,proto3" json:"level,omitempty"`
	ONAIHEPGDKM            *ALONKPBGLGI            `protobuf:"bytes,5,opt,name=ONAIHEPGDKM,proto3" json:"ONAIHEPGDKM,omitempty"`
	RogueTournGameAreaInfo *RogueTournGameAreaInfo `protobuf:"bytes,10,opt,name=rogue_tourn_game_area_info,json=rogueTournGameAreaInfo,proto3" json:"rogue_tourn_game_area_info,omitempty"`
	MCIBAOIDNIO            *GHKBGNFCEDK            `protobuf:"bytes,14,opt,name=MCIBAOIDNIO,proto3" json:"MCIBAOIDNIO,omitempty"`
	Lineup                 *RogueTournVirtualItem  `protobuf:"bytes,15,opt,name=lineup,proto3" json:"lineup,omitempty"`
	Buff                   *ChessRogueBuffInfo     `protobuf:"bytes,4,opt,name=buff,proto3" json:"buff,omitempty"`
}

func (x *RogueTournCurGameInfo) Reset() {
	*x = RogueTournCurGameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournCurGameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournCurGameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournCurGameInfo) ProtoMessage() {}

func (x *RogueTournCurGameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournCurGameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournCurGameInfo.ProtoReflect.Descriptor instead.
func (*RogueTournCurGameInfo) Descriptor() ([]byte, []int) {
	return file_RogueTournCurGameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournCurGameInfo) GetMiracleInfo() *ChessRogueMiracleInfo {
	if x != nil {
		return x.MiracleInfo
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetUnlockValue() *KeywordUnlockValue {
	if x != nil {
		return x.UnlockValue
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetItemValue() *RogueGameItemValue {
	if x != nil {
		return x.ItemValue
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetTournFormulaInfo() *RogueTournFormulaInfo {
	if x != nil {
		return x.TournFormulaInfo
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetLevel() *RogueTournLevelInfo {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetONAIHEPGDKM() *ALONKPBGLGI {
	if x != nil {
		return x.ONAIHEPGDKM
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetRogueTournGameAreaInfo() *RogueTournGameAreaInfo {
	if x != nil {
		return x.RogueTournGameAreaInfo
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetMCIBAOIDNIO() *GHKBGNFCEDK {
	if x != nil {
		return x.MCIBAOIDNIO
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetLineup() *RogueTournVirtualItem {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *RogueTournCurGameInfo) GetBuff() *ChessRogueBuffInfo {
	if x != nil {
		return x.Buff
	}
	return nil
}

var File_RogueTournCurGameInfo_proto protoreflect.FileDescriptor

var file_RogueTournCurGameInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4c, 0x4f, 0x4e, 0x4b, 0x50, 0x42, 0x47, 0x4c,
	0x47, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x48, 0x4b, 0x42, 0x47, 0x4e,
	0x46, 0x43, 0x45, 0x44, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x04, 0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x43, 0x75, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0c, 0x6d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x44, 0x0a, 0x12, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75,
	0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4e, 0x41, 0x49, 0x48, 0x45, 0x50, 0x47,
	0x44, 0x4b, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4c, 0x4f, 0x4e,
	0x4b, 0x50, 0x42, 0x47, 0x4c, 0x47, 0x49, 0x52, 0x0b, 0x4f, 0x4e, 0x41, 0x49, 0x48, 0x45, 0x50,
	0x47, 0x44, 0x4b, 0x4d, 0x12, 0x53, 0x0a, 0x1a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f,
	0x75, 0x72, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x16, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x43, 0x49,
	0x42, 0x41, 0x4f, 0x49, 0x44, 0x4e, 0x49, 0x4f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x47, 0x48, 0x4b, 0x42, 0x47, 0x4e, 0x46, 0x43, 0x45, 0x44, 0x4b, 0x52, 0x0b, 0x4d, 0x43,
	0x49, 0x42, 0x41, 0x4f, 0x49, 0x44, 0x4e, 0x49, 0x4f, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x75, 0x66,
	0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x75,
	0x66, 0x66, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournCurGameInfo_proto_rawDescOnce sync.Once
	file_RogueTournCurGameInfo_proto_rawDescData = file_RogueTournCurGameInfo_proto_rawDesc
)

func file_RogueTournCurGameInfo_proto_rawDescGZIP() []byte {
	file_RogueTournCurGameInfo_proto_rawDescOnce.Do(func() {
		file_RogueTournCurGameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournCurGameInfo_proto_rawDescData)
	})
	return file_RogueTournCurGameInfo_proto_rawDescData
}

var file_RogueTournCurGameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournCurGameInfo_proto_goTypes = []interface{}{
	(*RogueTournCurGameInfo)(nil),  // 0: RogueTournCurGameInfo
	(*ChessRogueMiracleInfo)(nil),  // 1: ChessRogueMiracleInfo
	(*KeywordUnlockValue)(nil),     // 2: KeywordUnlockValue
	(*RogueGameItemValue)(nil),     // 3: RogueGameItemValue
	(*RogueTournFormulaInfo)(nil),  // 4: RogueTournFormulaInfo
	(*RogueTournLevelInfo)(nil),    // 5: RogueTournLevelInfo
	(*ALONKPBGLGI)(nil),            // 6: ALONKPBGLGI
	(*RogueTournGameAreaInfo)(nil), // 7: RogueTournGameAreaInfo
	(*GHKBGNFCEDK)(nil),            // 8: GHKBGNFCEDK
	(*RogueTournVirtualItem)(nil),  // 9: RogueTournVirtualItem
	(*ChessRogueBuffInfo)(nil),     // 10: ChessRogueBuffInfo
}
var file_RogueTournCurGameInfo_proto_depIdxs = []int32{
	1,  // 0: RogueTournCurGameInfo.miracle_info:type_name -> ChessRogueMiracleInfo
	2,  // 1: RogueTournCurGameInfo.unlock_value:type_name -> KeywordUnlockValue
	3,  // 2: RogueTournCurGameInfo.item_value:type_name -> RogueGameItemValue
	4,  // 3: RogueTournCurGameInfo.tourn_formula_info:type_name -> RogueTournFormulaInfo
	5,  // 4: RogueTournCurGameInfo.level:type_name -> RogueTournLevelInfo
	6,  // 5: RogueTournCurGameInfo.ONAIHEPGDKM:type_name -> ALONKPBGLGI
	7,  // 6: RogueTournCurGameInfo.rogue_tourn_game_area_info:type_name -> RogueTournGameAreaInfo
	8,  // 7: RogueTournCurGameInfo.MCIBAOIDNIO:type_name -> GHKBGNFCEDK
	9,  // 8: RogueTournCurGameInfo.lineup:type_name -> RogueTournVirtualItem
	10, // 9: RogueTournCurGameInfo.buff:type_name -> ChessRogueBuffInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_RogueTournCurGameInfo_proto_init() }
func file_RogueTournCurGameInfo_proto_init() {
	if File_RogueTournCurGameInfo_proto != nil {
		return
	}
	file_ChessRogueMiracleInfo_proto_init()
	file_RogueTournVirtualItem_proto_init()
	file_RogueTournGameAreaInfo_proto_init()
	file_ALONKPBGLGI_proto_init()
	file_GHKBGNFCEDK_proto_init()
	file_RogueTournLevelInfo_proto_init()
	file_RogueGameItemValue_proto_init()
	file_RogueTournFormulaInfo_proto_init()
	file_KeywordUnlockValue_proto_init()
	file_ChessRogueBuffInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournCurGameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournCurGameInfo); i {
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
			RawDescriptor: file_RogueTournCurGameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournCurGameInfo_proto_goTypes,
		DependencyIndexes: file_RogueTournCurGameInfo_proto_depIdxs,
		MessageInfos:      file_RogueTournCurGameInfo_proto_msgTypes,
	}.Build()
	File_RogueTournCurGameInfo_proto = out.File
	file_RogueTournCurGameInfo_proto_rawDesc = nil
	file_RogueTournCurGameInfo_proto_goTypes = nil
	file_RogueTournCurGameInfo_proto_depIdxs = nil
}
