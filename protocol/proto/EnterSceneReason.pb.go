// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterSceneReason.proto

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

type EnterSceneReason int32

const (
	EnterSceneReason_ENTER_SCENE_REASON_NONE              EnterSceneReason = 0
	EnterSceneReason_ENTER_SCENE_REASON_CHALLENGE_TIMEOUT EnterSceneReason = 1
	EnterSceneReason_ENTER_SCENE_REASON_ROGUE_TIMEOUT     EnterSceneReason = 2
)

// Enum value maps for EnterSceneReason.
var (
	EnterSceneReason_name = map[int32]string{
		0: "ENTER_SCENE_REASON_NONE",
		1: "ENTER_SCENE_REASON_CHALLENGE_TIMEOUT",
		2: "ENTER_SCENE_REASON_ROGUE_TIMEOUT",
	}
	EnterSceneReason_value = map[string]int32{
		"ENTER_SCENE_REASON_NONE":              0,
		"ENTER_SCENE_REASON_CHALLENGE_TIMEOUT": 1,
		"ENTER_SCENE_REASON_ROGUE_TIMEOUT":     2,
	}
)

func (x EnterSceneReason) Enum() *EnterSceneReason {
	p := new(EnterSceneReason)
	*p = x
	return p
}

func (x EnterSceneReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnterSceneReason) Descriptor() protoreflect.EnumDescriptor {
	return file_EnterSceneReason_proto_enumTypes[0].Descriptor()
}

func (EnterSceneReason) Type() protoreflect.EnumType {
	return &file_EnterSceneReason_proto_enumTypes[0]
}

func (x EnterSceneReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnterSceneReason.Descriptor instead.
func (EnterSceneReason) EnumDescriptor() ([]byte, []int) {
	return file_EnterSceneReason_proto_rawDescGZIP(), []int{0}
}

var File_EnterSceneReason_proto protoreflect.FileDescriptor

var file_EnterSceneReason_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7f, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x4e, 0x54,
	0x45, 0x52, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x43, 0x45,
	0x4e, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterSceneReason_proto_rawDescOnce sync.Once
	file_EnterSceneReason_proto_rawDescData = file_EnterSceneReason_proto_rawDesc
)

func file_EnterSceneReason_proto_rawDescGZIP() []byte {
	file_EnterSceneReason_proto_rawDescOnce.Do(func() {
		file_EnterSceneReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterSceneReason_proto_rawDescData)
	})
	return file_EnterSceneReason_proto_rawDescData
}

var file_EnterSceneReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EnterSceneReason_proto_goTypes = []interface{}{
	(EnterSceneReason)(0), // 0: EnterSceneReason
}
var file_EnterSceneReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterSceneReason_proto_init() }
func file_EnterSceneReason_proto_init() {
	if File_EnterSceneReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EnterSceneReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterSceneReason_proto_goTypes,
		DependencyIndexes: file_EnterSceneReason_proto_depIdxs,
		EnumInfos:         file_EnterSceneReason_proto_enumTypes,
	}.Build()
	File_EnterSceneReason_proto = out.File
	file_EnterSceneReason_proto_rawDesc = nil
	file_EnterSceneReason_proto_goTypes = nil
	file_EnterSceneReason_proto_depIdxs = nil
}
