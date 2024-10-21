// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueCellUpdateReason.proto

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

type ChessRogueCellUpdateReason int32

const (
	ChessRogueCellUpdateReason_CHESS_ROGUE_CELL_UPDATE_REASON_NONE     ChessRogueCellUpdateReason = 0
	ChessRogueCellUpdateReason_CHESS_ROGUE_CELL_UPDATE_REASON_MODIFIER ChessRogueCellUpdateReason = 1
)

// Enum value maps for ChessRogueCellUpdateReason.
var (
	ChessRogueCellUpdateReason_name = map[int32]string{
		0: "CHESS_ROGUE_CELL_UPDATE_REASON_NONE",
		1: "CHESS_ROGUE_CELL_UPDATE_REASON_MODIFIER",
	}
	ChessRogueCellUpdateReason_value = map[string]int32{
		"CHESS_ROGUE_CELL_UPDATE_REASON_NONE":     0,
		"CHESS_ROGUE_CELL_UPDATE_REASON_MODIFIER": 1,
	}
)

func (x ChessRogueCellUpdateReason) Enum() *ChessRogueCellUpdateReason {
	p := new(ChessRogueCellUpdateReason)
	*p = x
	return p
}

func (x ChessRogueCellUpdateReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChessRogueCellUpdateReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ChessRogueCellUpdateReason_proto_enumTypes[0].Descriptor()
}

func (ChessRogueCellUpdateReason) Type() protoreflect.EnumType {
	return &file_ChessRogueCellUpdateReason_proto_enumTypes[0]
}

func (x ChessRogueCellUpdateReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChessRogueCellUpdateReason.Descriptor instead.
func (ChessRogueCellUpdateReason) EnumDescriptor() ([]byte, []int) {
	return file_ChessRogueCellUpdateReason_proto_rawDescGZIP(), []int{0}
}

var File_ChessRogueCellUpdateReason_proto protoreflect.FileDescriptor

var file_ChessRogueCellUpdateReason_proto_rawDesc = []byte{
	0x0a, 0x20, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x72, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x65, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x23, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f,
	0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x48, 0x45,
	0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x49,
	0x46, 0x49, 0x45, 0x52, 0x10, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueCellUpdateReason_proto_rawDescOnce sync.Once
	file_ChessRogueCellUpdateReason_proto_rawDescData = file_ChessRogueCellUpdateReason_proto_rawDesc
)

func file_ChessRogueCellUpdateReason_proto_rawDescGZIP() []byte {
	file_ChessRogueCellUpdateReason_proto_rawDescOnce.Do(func() {
		file_ChessRogueCellUpdateReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueCellUpdateReason_proto_rawDescData)
	})
	return file_ChessRogueCellUpdateReason_proto_rawDescData
}

var file_ChessRogueCellUpdateReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChessRogueCellUpdateReason_proto_goTypes = []interface{}{
	(ChessRogueCellUpdateReason)(0), // 0: ChessRogueCellUpdateReason
}
var file_ChessRogueCellUpdateReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueCellUpdateReason_proto_init() }
func file_ChessRogueCellUpdateReason_proto_init() {
	if File_ChessRogueCellUpdateReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChessRogueCellUpdateReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueCellUpdateReason_proto_goTypes,
		DependencyIndexes: file_ChessRogueCellUpdateReason_proto_depIdxs,
		EnumInfos:         file_ChessRogueCellUpdateReason_proto_enumTypes,
	}.Build()
	File_ChessRogueCellUpdateReason_proto = out.File
	file_ChessRogueCellUpdateReason_proto_rawDesc = nil
	file_ChessRogueCellUpdateReason_proto_goTypes = nil
	file_ChessRogueCellUpdateReason_proto_depIdxs = nil
}
