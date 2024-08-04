// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BEBEALGJPLI.proto

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

type BEBEALGJPLI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChessRogueBuffInfo    *ChessRogueBuff    `protobuf:"bytes,7,opt,name=chess_rogue_buff_info,json=chessRogueBuffInfo,proto3" json:"chess_rogue_buff_info,omitempty"`
	Status                RogueStatus        `protobuf:"varint,11,opt,name=status,proto3,enum=RogueStatus" json:"status,omitempty"`
	BaseAvatarIdList      []uint32           `protobuf:"varint,3,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	TrialAvatarIdList     []uint32           `protobuf:"varint,4,rep,packed,name=trial_avatar_id_list,json=trialAvatarIdList,proto3" json:"trial_avatar_id_list,omitempty"`
	GGHMFLALMKN           uint32             `protobuf:"varint,15,opt,name=GGHMFLALMKN,proto3" json:"GGHMFLALMKN,omitempty"`
	ChessRogueMiracleInfo *ChessRogueMiracle `protobuf:"bytes,12,opt,name=chess_rogue_miracle_info,json=chessRogueMiracleInfo,proto3" json:"chess_rogue_miracle_info,omitempty"`
	AHDIECBFBIF           uint32             `protobuf:"varint,8,opt,name=AHDIECBFBIF,proto3" json:"AHDIECBFBIF,omitempty"`
	GCAMGEFODNA           uint32             `protobuf:"varint,13,opt,name=GCAMGEFODNA,proto3" json:"GCAMGEFODNA,omitempty"`
	MapId                 uint32             `protobuf:"varint,9,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (x *BEBEALGJPLI) Reset() {
	*x = BEBEALGJPLI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BEBEALGJPLI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BEBEALGJPLI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BEBEALGJPLI) ProtoMessage() {}

func (x *BEBEALGJPLI) ProtoReflect() protoreflect.Message {
	mi := &file_BEBEALGJPLI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BEBEALGJPLI.ProtoReflect.Descriptor instead.
func (*BEBEALGJPLI) Descriptor() ([]byte, []int) {
	return file_BEBEALGJPLI_proto_rawDescGZIP(), []int{0}
}

func (x *BEBEALGJPLI) GetChessRogueBuffInfo() *ChessRogueBuff {
	if x != nil {
		return x.ChessRogueBuffInfo
	}
	return nil
}

func (x *BEBEALGJPLI) GetStatus() RogueStatus {
	if x != nil {
		return x.Status
	}
	return RogueStatus_ROGUE_STATUS_NONE
}

func (x *BEBEALGJPLI) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *BEBEALGJPLI) GetTrialAvatarIdList() []uint32 {
	if x != nil {
		return x.TrialAvatarIdList
	}
	return nil
}

func (x *BEBEALGJPLI) GetGGHMFLALMKN() uint32 {
	if x != nil {
		return x.GGHMFLALMKN
	}
	return 0
}

func (x *BEBEALGJPLI) GetChessRogueMiracleInfo() *ChessRogueMiracle {
	if x != nil {
		return x.ChessRogueMiracleInfo
	}
	return nil
}

func (x *BEBEALGJPLI) GetAHDIECBFBIF() uint32 {
	if x != nil {
		return x.AHDIECBFBIF
	}
	return 0
}

func (x *BEBEALGJPLI) GetGCAMGEFODNA() uint32 {
	if x != nil {
		return x.GCAMGEFODNA
	}
	return 0
}

func (x *BEBEALGJPLI) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

var File_BEBEALGJPLI_proto protoreflect.FileDescriptor

var file_BEBEALGJPLI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x45, 0x42, 0x45, 0x41, 0x4c, 0x47, 0x4a, 0x50, 0x4c, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x0b, 0x42, 0x45, 0x42, 0x45, 0x41, 0x4c,
	0x47, 0x4a, 0x50, 0x4c, 0x49, 0x12, 0x42, 0x0a, 0x15, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x12, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2d, 0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x61,
	0x73, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x14, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x47, 0x47, 0x48, 0x4d, 0x46, 0x4c, 0x41, 0x4c, 0x4d, 0x4b, 0x4e, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x47, 0x48, 0x4d, 0x46, 0x4c, 0x41, 0x4c, 0x4d, 0x4b,
	0x4e, 0x12, 0x4b, 0x0a, 0x18, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x15, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x48, 0x44, 0x49, 0x45, 0x43, 0x42, 0x46, 0x42, 0x49, 0x46, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x48, 0x44, 0x49, 0x45, 0x43, 0x42, 0x46, 0x42, 0x49, 0x46,
	0x12, 0x20, 0x0a, 0x0b, 0x47, 0x43, 0x41, 0x4d, 0x47, 0x45, 0x46, 0x4f, 0x44, 0x4e, 0x41, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x43, 0x41, 0x4d, 0x47, 0x45, 0x46, 0x4f, 0x44,
	0x4e, 0x41, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BEBEALGJPLI_proto_rawDescOnce sync.Once
	file_BEBEALGJPLI_proto_rawDescData = file_BEBEALGJPLI_proto_rawDesc
)

func file_BEBEALGJPLI_proto_rawDescGZIP() []byte {
	file_BEBEALGJPLI_proto_rawDescOnce.Do(func() {
		file_BEBEALGJPLI_proto_rawDescData = protoimpl.X.CompressGZIP(file_BEBEALGJPLI_proto_rawDescData)
	})
	return file_BEBEALGJPLI_proto_rawDescData
}

var file_BEBEALGJPLI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BEBEALGJPLI_proto_goTypes = []interface{}{
	(*BEBEALGJPLI)(nil),       // 0: BEBEALGJPLI
	(*ChessRogueBuff)(nil),    // 1: ChessRogueBuff
	(RogueStatus)(0),          // 2: RogueStatus
	(*ChessRogueMiracle)(nil), // 3: ChessRogueMiracle
}
var file_BEBEALGJPLI_proto_depIdxs = []int32{
	1, // 0: BEBEALGJPLI.chess_rogue_buff_info:type_name -> ChessRogueBuff
	2, // 1: BEBEALGJPLI.status:type_name -> RogueStatus
	3, // 2: BEBEALGJPLI.chess_rogue_miracle_info:type_name -> ChessRogueMiracle
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_BEBEALGJPLI_proto_init() }
func file_BEBEALGJPLI_proto_init() {
	if File_BEBEALGJPLI_proto != nil {
		return
	}
	file_ChessRogueBuff_proto_init()
	file_ChessRogueMiracle_proto_init()
	file_RogueStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BEBEALGJPLI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BEBEALGJPLI); i {
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
			RawDescriptor: file_BEBEALGJPLI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BEBEALGJPLI_proto_goTypes,
		DependencyIndexes: file_BEBEALGJPLI_proto_depIdxs,
		MessageInfos:      file_BEBEALGJPLI_proto_msgTypes,
	}.Build()
	File_BEBEALGJPLI_proto = out.File
	file_BEBEALGJPLI_proto_rawDesc = nil
	file_BEBEALGJPLI_proto_goTypes = nil
	file_BEBEALGJPLI_proto_depIdxs = nil
}
