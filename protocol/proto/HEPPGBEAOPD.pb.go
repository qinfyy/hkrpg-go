// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HEPPGBEAOPD.proto

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

type HEPPGBEAOPD int32

const (
	HEPPGBEAOPD_TRAIN_VISITOR_REGISTER_GET_TYPE_NONE   HEPPGBEAOPD = 0
	HEPPGBEAOPD_TRAIN_VISITOR_REGISTER_GET_TYPE_AUTO   HEPPGBEAOPD = 1
	HEPPGBEAOPD_TRAIN_VISITOR_REGISTER_GET_TYPE_MANUAL HEPPGBEAOPD = 2
)

// Enum value maps for HEPPGBEAOPD.
var (
	HEPPGBEAOPD_name = map[int32]string{
		0: "TRAIN_VISITOR_REGISTER_GET_TYPE_NONE",
		1: "TRAIN_VISITOR_REGISTER_GET_TYPE_AUTO",
		2: "TRAIN_VISITOR_REGISTER_GET_TYPE_MANUAL",
	}
	HEPPGBEAOPD_value = map[string]int32{
		"TRAIN_VISITOR_REGISTER_GET_TYPE_NONE":   0,
		"TRAIN_VISITOR_REGISTER_GET_TYPE_AUTO":   1,
		"TRAIN_VISITOR_REGISTER_GET_TYPE_MANUAL": 2,
	}
)

func (x HEPPGBEAOPD) Enum() *HEPPGBEAOPD {
	p := new(HEPPGBEAOPD)
	*p = x
	return p
}

func (x HEPPGBEAOPD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HEPPGBEAOPD) Descriptor() protoreflect.EnumDescriptor {
	return file_HEPPGBEAOPD_proto_enumTypes[0].Descriptor()
}

func (HEPPGBEAOPD) Type() protoreflect.EnumType {
	return &file_HEPPGBEAOPD_proto_enumTypes[0]
}

func (x HEPPGBEAOPD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HEPPGBEAOPD.Descriptor instead.
func (HEPPGBEAOPD) EnumDescriptor() ([]byte, []int) {
	return file_HEPPGBEAOPD_proto_rawDescGZIP(), []int{0}
}

var File_HEPPGBEAOPD_proto protoreflect.FileDescriptor

var file_HEPPGBEAOPD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x45, 0x50, 0x50, 0x47, 0x42, 0x45, 0x41, 0x4f, 0x50, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x8d, 0x01, 0x0a, 0x0b, 0x48, 0x45, 0x50, 0x50, 0x47, 0x42, 0x45, 0x41,
	0x4f, 0x50, 0x44, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x56, 0x49, 0x53,
	0x49, 0x54, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x28, 0x0a,
	0x24, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x52, 0x41, 0x49, 0x4e,
	0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HEPPGBEAOPD_proto_rawDescOnce sync.Once
	file_HEPPGBEAOPD_proto_rawDescData = file_HEPPGBEAOPD_proto_rawDesc
)

func file_HEPPGBEAOPD_proto_rawDescGZIP() []byte {
	file_HEPPGBEAOPD_proto_rawDescOnce.Do(func() {
		file_HEPPGBEAOPD_proto_rawDescData = protoimpl.X.CompressGZIP(file_HEPPGBEAOPD_proto_rawDescData)
	})
	return file_HEPPGBEAOPD_proto_rawDescData
}

var file_HEPPGBEAOPD_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HEPPGBEAOPD_proto_goTypes = []interface{}{
	(HEPPGBEAOPD)(0), // 0: HEPPGBEAOPD
}
var file_HEPPGBEAOPD_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HEPPGBEAOPD_proto_init() }
func file_HEPPGBEAOPD_proto_init() {
	if File_HEPPGBEAOPD_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HEPPGBEAOPD_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HEPPGBEAOPD_proto_goTypes,
		DependencyIndexes: file_HEPPGBEAOPD_proto_depIdxs,
		EnumInfos:         file_HEPPGBEAOPD_proto_enumTypes,
	}.Build()
	File_HEPPGBEAOPD_proto = out.File
	file_HEPPGBEAOPD_proto_rawDesc = nil
	file_HEPPGBEAOPD_proto_goTypes = nil
	file_HEPPGBEAOPD_proto_depIdxs = nil
}
