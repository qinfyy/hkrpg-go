// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GIOGIPJLONO.proto

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

type GIOGIPJLONO int32

const (
	GIOGIPJLONO_ROGUE_TOURN_SETTLE_REASON_NONE      GIOGIPJLONO = 0
	GIOGIPJLONO_ROGUE_TOURN_SETTLE_REASON_WIN       GIOGIPJLONO = 1
	GIOGIPJLONO_ROGUE_TOURN_SETTLE_REASON_FAIL      GIOGIPJLONO = 2
	GIOGIPJLONO_ROGUE_TOURN_SETTLE_REASON_INTERRUPT GIOGIPJLONO = 3
)

// Enum value maps for GIOGIPJLONO.
var (
	GIOGIPJLONO_name = map[int32]string{
		0: "ROGUE_TOURN_SETTLE_REASON_NONE",
		1: "ROGUE_TOURN_SETTLE_REASON_WIN",
		2: "ROGUE_TOURN_SETTLE_REASON_FAIL",
		3: "ROGUE_TOURN_SETTLE_REASON_INTERRUPT",
	}
	GIOGIPJLONO_value = map[string]int32{
		"ROGUE_TOURN_SETTLE_REASON_NONE":      0,
		"ROGUE_TOURN_SETTLE_REASON_WIN":       1,
		"ROGUE_TOURN_SETTLE_REASON_FAIL":      2,
		"ROGUE_TOURN_SETTLE_REASON_INTERRUPT": 3,
	}
)

func (x GIOGIPJLONO) Enum() *GIOGIPJLONO {
	p := new(GIOGIPJLONO)
	*p = x
	return p
}

func (x GIOGIPJLONO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GIOGIPJLONO) Descriptor() protoreflect.EnumDescriptor {
	return file_GIOGIPJLONO_proto_enumTypes[0].Descriptor()
}

func (GIOGIPJLONO) Type() protoreflect.EnumType {
	return &file_GIOGIPJLONO_proto_enumTypes[0]
}

func (x GIOGIPJLONO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GIOGIPJLONO.Descriptor instead.
func (GIOGIPJLONO) EnumDescriptor() ([]byte, []int) {
	return file_GIOGIPJLONO_proto_rawDescGZIP(), []int{0}
}

var File_GIOGIPJLONO_proto protoreflect.FileDescriptor

var file_GIOGIPJLONO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x49, 0x4f, 0x47, 0x49, 0x50, 0x4a, 0x4c, 0x4f, 0x4e, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xa1, 0x01, 0x0a, 0x0b, 0x47, 0x49, 0x4f, 0x47, 0x49, 0x50, 0x4a, 0x4c,
	0x4f, 0x4e, 0x4f, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55,
	0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f,
	0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x27,
	0x0a, 0x23, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x52, 0x55, 0x50, 0x54, 0x10, 0x03, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GIOGIPJLONO_proto_rawDescOnce sync.Once
	file_GIOGIPJLONO_proto_rawDescData = file_GIOGIPJLONO_proto_rawDesc
)

func file_GIOGIPJLONO_proto_rawDescGZIP() []byte {
	file_GIOGIPJLONO_proto_rawDescOnce.Do(func() {
		file_GIOGIPJLONO_proto_rawDescData = protoimpl.X.CompressGZIP(file_GIOGIPJLONO_proto_rawDescData)
	})
	return file_GIOGIPJLONO_proto_rawDescData
}

var file_GIOGIPJLONO_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GIOGIPJLONO_proto_goTypes = []interface{}{
	(GIOGIPJLONO)(0), // 0: GIOGIPJLONO
}
var file_GIOGIPJLONO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GIOGIPJLONO_proto_init() }
func file_GIOGIPJLONO_proto_init() {
	if File_GIOGIPJLONO_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GIOGIPJLONO_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GIOGIPJLONO_proto_goTypes,
		DependencyIndexes: file_GIOGIPJLONO_proto_depIdxs,
		EnumInfos:         file_GIOGIPJLONO_proto_enumTypes,
	}.Build()
	File_GIOGIPJLONO_proto = out.File
	file_GIOGIPJLONO_proto_rawDesc = nil
	file_GIOGIPJLONO_proto_goTypes = nil
	file_GIOGIPJLONO_proto_depIdxs = nil
}