// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DeathSource.proto

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

type DeathSource int32

const (
	DeathSource_UNKNOWN          DeathSource = 0
	DeathSource_KILLED_BY_OTHERS DeathSource = 1
	DeathSource_KILLED_BY_SELF   DeathSource = 2
	DeathSource_ESCAPE           DeathSource = 3
)

// Enum value maps for DeathSource.
var (
	DeathSource_name = map[int32]string{
		0: "UNKNOWN",
		1: "KILLED_BY_OTHERS",
		2: "KILLED_BY_SELF",
		3: "ESCAPE",
	}
	DeathSource_value = map[string]int32{
		"UNKNOWN":          0,
		"KILLED_BY_OTHERS": 1,
		"KILLED_BY_SELF":   2,
		"ESCAPE":           3,
	}
)

func (x DeathSource) Enum() *DeathSource {
	p := new(DeathSource)
	*p = x
	return p
}

func (x DeathSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeathSource) Descriptor() protoreflect.EnumDescriptor {
	return file_DeathSource_proto_enumTypes[0].Descriptor()
}

func (DeathSource) Type() protoreflect.EnumType {
	return &file_DeathSource_proto_enumTypes[0]
}

func (x DeathSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeathSource.Descriptor instead.
func (DeathSource) EnumDescriptor() ([]byte, []int) {
	return file_DeathSource_proto_rawDescGZIP(), []int{0}
}

var File_DeathSource_proto protoreflect.FileDescriptor

var file_DeathSource_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x65, 0x61, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x50, 0x0a, 0x0b, 0x44, 0x65, 0x61, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x5f,
	0x42, 0x59, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x53, 0x43,
	0x41, 0x50, 0x45, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DeathSource_proto_rawDescOnce sync.Once
	file_DeathSource_proto_rawDescData = file_DeathSource_proto_rawDesc
)

func file_DeathSource_proto_rawDescGZIP() []byte {
	file_DeathSource_proto_rawDescOnce.Do(func() {
		file_DeathSource_proto_rawDescData = protoimpl.X.CompressGZIP(file_DeathSource_proto_rawDescData)
	})
	return file_DeathSource_proto_rawDescData
}

var file_DeathSource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DeathSource_proto_goTypes = []interface{}{
	(DeathSource)(0), // 0: DeathSource
}
var file_DeathSource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DeathSource_proto_init() }
func file_DeathSource_proto_init() {
	if File_DeathSource_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DeathSource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DeathSource_proto_goTypes,
		DependencyIndexes: file_DeathSource_proto_depIdxs,
		EnumInfos:         file_DeathSource_proto_enumTypes,
	}.Build()
	File_DeathSource_proto = out.File
	file_DeathSource_proto_rawDesc = nil
	file_DeathSource_proto_goTypes = nil
	file_DeathSource_proto_depIdxs = nil
}
