// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EFEEMPHMFKI.proto

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

type EFEEMPHMFKI int32

const (
	EFEEMPHMFKI_MATCH3_PLAYER_STATE_ALIVE EFEEMPHMFKI = 0
	EFEEMPHMFKI_MATCH3_PLAYER_STATE_DYING EFEEMPHMFKI = 1
	EFEEMPHMFKI_MATCH3_PLAYER_STATE_DEAD  EFEEMPHMFKI = 2
	EFEEMPHMFKI_MATCH3_PLAYER_STATE_LEAVE EFEEMPHMFKI = 3
)

// Enum value maps for EFEEMPHMFKI.
var (
	EFEEMPHMFKI_name = map[int32]string{
		0: "MATCH3_PLAYER_STATE_ALIVE",
		1: "MATCH3_PLAYER_STATE_DYING",
		2: "MATCH3_PLAYER_STATE_DEAD",
		3: "MATCH3_PLAYER_STATE_LEAVE",
	}
	EFEEMPHMFKI_value = map[string]int32{
		"MATCH3_PLAYER_STATE_ALIVE": 0,
		"MATCH3_PLAYER_STATE_DYING": 1,
		"MATCH3_PLAYER_STATE_DEAD":  2,
		"MATCH3_PLAYER_STATE_LEAVE": 3,
	}
)

func (x EFEEMPHMFKI) Enum() *EFEEMPHMFKI {
	p := new(EFEEMPHMFKI)
	*p = x
	return p
}

func (x EFEEMPHMFKI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EFEEMPHMFKI) Descriptor() protoreflect.EnumDescriptor {
	return file_EFEEMPHMFKI_proto_enumTypes[0].Descriptor()
}

func (EFEEMPHMFKI) Type() protoreflect.EnumType {
	return &file_EFEEMPHMFKI_proto_enumTypes[0]
}

func (x EFEEMPHMFKI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EFEEMPHMFKI.Descriptor instead.
func (EFEEMPHMFKI) EnumDescriptor() ([]byte, []int) {
	return file_EFEEMPHMFKI_proto_rawDescGZIP(), []int{0}
}

var File_EFEEMPHMFKI_proto protoreflect.FileDescriptor

var file_EFEEMPHMFKI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x46, 0x45, 0x45, 0x4d, 0x50, 0x48, 0x4d, 0x46, 0x4b, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x88, 0x01, 0x0a, 0x0b, 0x45, 0x46, 0x45, 0x45, 0x4d, 0x50, 0x48, 0x4d,
	0x46, 0x4b, 0x49, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4c, 0x49, 0x56, 0x45,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x59, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12,
	0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x03, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EFEEMPHMFKI_proto_rawDescOnce sync.Once
	file_EFEEMPHMFKI_proto_rawDescData = file_EFEEMPHMFKI_proto_rawDesc
)

func file_EFEEMPHMFKI_proto_rawDescGZIP() []byte {
	file_EFEEMPHMFKI_proto_rawDescOnce.Do(func() {
		file_EFEEMPHMFKI_proto_rawDescData = protoimpl.X.CompressGZIP(file_EFEEMPHMFKI_proto_rawDescData)
	})
	return file_EFEEMPHMFKI_proto_rawDescData
}

var file_EFEEMPHMFKI_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EFEEMPHMFKI_proto_goTypes = []interface{}{
	(EFEEMPHMFKI)(0), // 0: EFEEMPHMFKI
}
var file_EFEEMPHMFKI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EFEEMPHMFKI_proto_init() }
func file_EFEEMPHMFKI_proto_init() {
	if File_EFEEMPHMFKI_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EFEEMPHMFKI_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EFEEMPHMFKI_proto_goTypes,
		DependencyIndexes: file_EFEEMPHMFKI_proto_depIdxs,
		EnumInfos:         file_EFEEMPHMFKI_proto_enumTypes,
	}.Build()
	File_EFEEMPHMFKI_proto = out.File
	file_EFEEMPHMFKI_proto_rawDesc = nil
	file_EFEEMPHMFKI_proto_goTypes = nil
	file_EFEEMPHMFKI_proto_depIdxs = nil
}