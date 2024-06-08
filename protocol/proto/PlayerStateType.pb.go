// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerStateType.proto

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

type PlayerStateType int32

const (
	PlayerStateType_PLAYING_STATE_NONE       PlayerStateType = 0
	PlayerStateType_PLAYING_ROGUE_COSMOS     PlayerStateType = 1
	PlayerStateType_PLAYING_ROGUE_CHESS      PlayerStateType = 2
	PlayerStateType_PLAYING_ROGUE_CHESS_NOUS PlayerStateType = 3
	PlayerStateType_PLAYING_CHALLENGE_MEMORY PlayerStateType = 4
	PlayerStateType_PLAYING_CHALLENGE_STORY  PlayerStateType = 5
)

// Enum value maps for PlayerStateType.
var (
	PlayerStateType_name = map[int32]string{
		0: "PLAYING_STATE_NONE",
		1: "PLAYING_ROGUE_COSMOS",
		2: "PLAYING_ROGUE_CHESS",
		3: "PLAYING_ROGUE_CHESS_NOUS",
		4: "PLAYING_CHALLENGE_MEMORY",
		5: "PLAYING_CHALLENGE_STORY",
	}
	PlayerStateType_value = map[string]int32{
		"PLAYING_STATE_NONE":       0,
		"PLAYING_ROGUE_COSMOS":     1,
		"PLAYING_ROGUE_CHESS":      2,
		"PLAYING_ROGUE_CHESS_NOUS": 3,
		"PLAYING_CHALLENGE_MEMORY": 4,
		"PLAYING_CHALLENGE_STORY":  5,
	}
)

func (x PlayerStateType) Enum() *PlayerStateType {
	p := new(PlayerStateType)
	*p = x
	return p
}

func (x PlayerStateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerStateType) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerStateType_proto_enumTypes[0].Descriptor()
}

func (PlayerStateType) Type() protoreflect.EnumType {
	return &file_PlayerStateType_proto_enumTypes[0]
}

func (x PlayerStateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerStateType.Descriptor instead.
func (PlayerStateType) EnumDescriptor() ([]byte, []int) {
	return file_PlayerStateType_proto_rawDescGZIP(), []int{0}
}

var File_PlayerStateType_proto protoreflect.FileDescriptor

var file_PlayerStateType_proto_rawDesc = []byte{
	0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb5, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x53, 0x4d, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43,
	0x48, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e,
	0x47, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x4e, 0x4f,
	0x55, 0x53, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x5f,
	0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59,
	0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48,
	0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x05, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PlayerStateType_proto_rawDescOnce sync.Once
	file_PlayerStateType_proto_rawDescData = file_PlayerStateType_proto_rawDesc
)

func file_PlayerStateType_proto_rawDescGZIP() []byte {
	file_PlayerStateType_proto_rawDescOnce.Do(func() {
		file_PlayerStateType_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerStateType_proto_rawDescData)
	})
	return file_PlayerStateType_proto_rawDescData
}

var file_PlayerStateType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerStateType_proto_goTypes = []interface{}{
	(PlayerStateType)(0), // 0: PlayerStateType
}
var file_PlayerStateType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerStateType_proto_init() }
func file_PlayerStateType_proto_init() {
	if File_PlayerStateType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerStateType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerStateType_proto_goTypes,
		DependencyIndexes: file_PlayerStateType_proto_depIdxs,
		EnumInfos:         file_PlayerStateType_proto_enumTypes,
	}.Build()
	File_PlayerStateType_proto = out.File
	file_PlayerStateType_proto_rawDesc = nil
	file_PlayerStateType_proto_goTypes = nil
	file_PlayerStateType_proto_depIdxs = nil
}