// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTalentStatus.proto

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

type RogueTalentStatus int32

const (
	RogueTalentStatus_ROGUE_TALENT_STATUS_LOCK   RogueTalentStatus = 0
	RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK RogueTalentStatus = 1
	RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE RogueTalentStatus = 2
)

// Enum value maps for RogueTalentStatus.
var (
	RogueTalentStatus_name = map[int32]string{
		0: "ROGUE_TALENT_STATUS_LOCK",
		1: "ROGUE_TALENT_STATUS_UNLOCK",
		2: "ROGUE_TALENT_STATUS_ENABLE",
	}
	RogueTalentStatus_value = map[string]int32{
		"ROGUE_TALENT_STATUS_LOCK":   0,
		"ROGUE_TALENT_STATUS_UNLOCK": 1,
		"ROGUE_TALENT_STATUS_ENABLE": 2,
	}
)

func (x RogueTalentStatus) Enum() *RogueTalentStatus {
	p := new(RogueTalentStatus)
	*p = x
	return p
}

func (x RogueTalentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueTalentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueTalentStatus_proto_enumTypes[0].Descriptor()
}

func (RogueTalentStatus) Type() protoreflect.EnumType {
	return &file_RogueTalentStatus_proto_enumTypes[0]
}

func (x RogueTalentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueTalentStatus.Descriptor instead.
func (RogueTalentStatus) EnumDescriptor() ([]byte, []int) {
	return file_RogueTalentStatus_proto_rawDescGZIP(), []int{0}
}

var File_RogueTalentStatus_proto protoreflect.FileDescriptor

var file_RogueTalentStatus_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x71, 0x0a, 0x11, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x18, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x41, 0x4c, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x41, 0x4c, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x41, 0x4c, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTalentStatus_proto_rawDescOnce sync.Once
	file_RogueTalentStatus_proto_rawDescData = file_RogueTalentStatus_proto_rawDesc
)

func file_RogueTalentStatus_proto_rawDescGZIP() []byte {
	file_RogueTalentStatus_proto_rawDescOnce.Do(func() {
		file_RogueTalentStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTalentStatus_proto_rawDescData)
	})
	return file_RogueTalentStatus_proto_rawDescData
}

var file_RogueTalentStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueTalentStatus_proto_goTypes = []interface{}{
	(RogueTalentStatus)(0), // 0: RogueTalentStatus
}
var file_RogueTalentStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueTalentStatus_proto_init() }
func file_RogueTalentStatus_proto_init() {
	if File_RogueTalentStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueTalentStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTalentStatus_proto_goTypes,
		DependencyIndexes: file_RogueTalentStatus_proto_depIdxs,
		EnumInfos:         file_RogueTalentStatus_proto_enumTypes,
	}.Build()
	File_RogueTalentStatus_proto = out.File
	file_RogueTalentStatus_proto_rawDesc = nil
	file_RogueTalentStatus_proto_goTypes = nil
	file_RogueTalentStatus_proto_depIdxs = nil
}
