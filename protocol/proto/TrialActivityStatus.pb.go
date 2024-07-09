// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TrialActivityStatus.proto

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

type TrialActivityStatus int32

const (
	TrialActivityStatus_TRIAL_ACTIVITY_STATUS_NONE   TrialActivityStatus = 0
	TrialActivityStatus_TRIAL_ACTIVITY_STATUS_FINISH TrialActivityStatus = 1
)

// Enum value maps for TrialActivityStatus.
var (
	TrialActivityStatus_name = map[int32]string{
		0: "TRIAL_ACTIVITY_STATUS_NONE",
		1: "TRIAL_ACTIVITY_STATUS_FINISH",
	}
	TrialActivityStatus_value = map[string]int32{
		"TRIAL_ACTIVITY_STATUS_NONE":   0,
		"TRIAL_ACTIVITY_STATUS_FINISH": 1,
	}
)

func (x TrialActivityStatus) Enum() *TrialActivityStatus {
	p := new(TrialActivityStatus)
	*p = x
	return p
}

func (x TrialActivityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrialActivityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_TrialActivityStatus_proto_enumTypes[0].Descriptor()
}

func (TrialActivityStatus) Type() protoreflect.EnumType {
	return &file_TrialActivityStatus_proto_enumTypes[0]
}

func (x TrialActivityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrialActivityStatus.Descriptor instead.
func (TrialActivityStatus) EnumDescriptor() ([]byte, []int) {
	return file_TrialActivityStatus_proto_rawDescGZIP(), []int{0}
}

var File_TrialActivityStatus_proto protoreflect.FileDescriptor

var file_TrialActivityStatus_proto_rawDesc = []byte{
	0x0a, 0x19, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x57, 0x0a, 0x13, 0x54,
	0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrialActivityStatus_proto_rawDescOnce sync.Once
	file_TrialActivityStatus_proto_rawDescData = file_TrialActivityStatus_proto_rawDesc
)

func file_TrialActivityStatus_proto_rawDescGZIP() []byte {
	file_TrialActivityStatus_proto_rawDescOnce.Do(func() {
		file_TrialActivityStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrialActivityStatus_proto_rawDescData)
	})
	return file_TrialActivityStatus_proto_rawDescData
}

var file_TrialActivityStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TrialActivityStatus_proto_goTypes = []interface{}{
	(TrialActivityStatus)(0), // 0: TrialActivityStatus
}
var file_TrialActivityStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TrialActivityStatus_proto_init() }
func file_TrialActivityStatus_proto_init() {
	if File_TrialActivityStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TrialActivityStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrialActivityStatus_proto_goTypes,
		DependencyIndexes: file_TrialActivityStatus_proto_depIdxs,
		EnumInfos:         file_TrialActivityStatus_proto_enumTypes,
	}.Build()
	File_TrialActivityStatus_proto = out.File
	file_TrialActivityStatus_proto_rawDesc = nil
	file_TrialActivityStatus_proto_goTypes = nil
	file_TrialActivityStatus_proto_depIdxs = nil
}