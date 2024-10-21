// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NMABOGNBIPH.proto

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

type NMABOGNBIPH int32

const (
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_NONE           NMABOGNBIPH = 0
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_TURN_SETTLE    NMABOGNBIPH = 1
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_STATUS_UPGRADE NMABOGNBIPH = 2
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION         NMABOGNBIPH = 3
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION_HINT    NMABOGNBIPH = 4
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_STORY          NMABOGNBIPH = 5
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_EXAM_BONUS     NMABOGNBIPH = 6
	NMABOGNBIPH_SWORD_TRAIN_GAME_SOURCE_TYPE_DIALOGUE       NMABOGNBIPH = 7
)

// Enum value maps for NMABOGNBIPH.
var (
	NMABOGNBIPH_name = map[int32]string{
		0: "SWORD_TRAIN_GAME_SOURCE_TYPE_NONE",
		1: "SWORD_TRAIN_GAME_SOURCE_TYPE_TURN_SETTLE",
		2: "SWORD_TRAIN_GAME_SOURCE_TYPE_STATUS_UPGRADE",
		3: "SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION",
		4: "SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION_HINT",
		5: "SWORD_TRAIN_GAME_SOURCE_TYPE_STORY",
		6: "SWORD_TRAIN_GAME_SOURCE_TYPE_EXAM_BONUS",
		7: "SWORD_TRAIN_GAME_SOURCE_TYPE_DIALOGUE",
	}
	NMABOGNBIPH_value = map[string]int32{
		"SWORD_TRAIN_GAME_SOURCE_TYPE_NONE":           0,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_TURN_SETTLE":    1,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_STATUS_UPGRADE": 2,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION":         3,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_ACTION_HINT":    4,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_STORY":          5,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_EXAM_BONUS":     6,
		"SWORD_TRAIN_GAME_SOURCE_TYPE_DIALOGUE":       7,
	}
)

func (x NMABOGNBIPH) Enum() *NMABOGNBIPH {
	p := new(NMABOGNBIPH)
	*p = x
	return p
}

func (x NMABOGNBIPH) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NMABOGNBIPH) Descriptor() protoreflect.EnumDescriptor {
	return file_NMABOGNBIPH_proto_enumTypes[0].Descriptor()
}

func (NMABOGNBIPH) Type() protoreflect.EnumType {
	return &file_NMABOGNBIPH_proto_enumTypes[0]
}

func (x NMABOGNBIPH) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NMABOGNBIPH.Descriptor instead.
func (NMABOGNBIPH) EnumDescriptor() ([]byte, []int) {
	return file_NMABOGNBIPH_proto_rawDescGZIP(), []int{0}
}

var File_NMABOGNBIPH_proto protoreflect.FileDescriptor

var file_NMABOGNBIPH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4d, 0x41, 0x42, 0x4f, 0x47, 0x4e, 0x42, 0x49, 0x50, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xea, 0x02, 0x0a, 0x0b, 0x4e, 0x4d, 0x41, 0x42, 0x4f, 0x47, 0x4e, 0x42,
	0x49, 0x50, 0x48, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41,
	0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x57,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x2f, 0x0a, 0x2b, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x57, 0x4f,
	0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49,
	0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x4e, 0x54, 0x10, 0x04,
	0x12, 0x26, 0x0a, 0x22, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x41, 0x4d, 0x5f, 0x42, 0x4f,
	0x4e, 0x55, 0x53, 0x10, 0x06, 0x12, 0x29, 0x0a, 0x25, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47, 0x55, 0x45, 0x10, 0x07,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NMABOGNBIPH_proto_rawDescOnce sync.Once
	file_NMABOGNBIPH_proto_rawDescData = file_NMABOGNBIPH_proto_rawDesc
)

func file_NMABOGNBIPH_proto_rawDescGZIP() []byte {
	file_NMABOGNBIPH_proto_rawDescOnce.Do(func() {
		file_NMABOGNBIPH_proto_rawDescData = protoimpl.X.CompressGZIP(file_NMABOGNBIPH_proto_rawDescData)
	})
	return file_NMABOGNBIPH_proto_rawDescData
}

var file_NMABOGNBIPH_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_NMABOGNBIPH_proto_goTypes = []interface{}{
	(NMABOGNBIPH)(0), // 0: NMABOGNBIPH
}
var file_NMABOGNBIPH_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NMABOGNBIPH_proto_init() }
func file_NMABOGNBIPH_proto_init() {
	if File_NMABOGNBIPH_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NMABOGNBIPH_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NMABOGNBIPH_proto_goTypes,
		DependencyIndexes: file_NMABOGNBIPH_proto_depIdxs,
		EnumInfos:         file_NMABOGNBIPH_proto_enumTypes,
	}.Build()
	File_NMABOGNBIPH_proto = out.File
	file_NMABOGNBIPH_proto_rawDesc = nil
	file_NMABOGNBIPH_proto_goTypes = nil
	file_NMABOGNBIPH_proto_depIdxs = nil
}
