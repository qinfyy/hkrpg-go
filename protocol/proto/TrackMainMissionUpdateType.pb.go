// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TrackMainMissionUpdateType.proto

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

type TrackMainMissionUpdateType int32

const (
	TrackMainMissionUpdateType_TRACK_MAIN_MISSION_UPDATE_NONE         TrackMainMissionUpdateType = 0
	TrackMainMissionUpdateType_TRACK_MAIN_MISSION_UPDATE_AUTO         TrackMainMissionUpdateType = 1
	TrackMainMissionUpdateType_TRACK_MAIN_MISSION_UPDATE_MANUAL       TrackMainMissionUpdateType = 2
	TrackMainMissionUpdateType_TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT TrackMainMissionUpdateType = 3
)

// Enum value maps for TrackMainMissionUpdateType.
var (
	TrackMainMissionUpdateType_name = map[int32]string{
		0: "TRACK_MAIN_MISSION_UPDATE_NONE",
		1: "TRACK_MAIN_MISSION_UPDATE_AUTO",
		2: "TRACK_MAIN_MISSION_UPDATE_MANUAL",
		3: "TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT",
	}
	TrackMainMissionUpdateType_value = map[string]int32{
		"TRACK_MAIN_MISSION_UPDATE_NONE":         0,
		"TRACK_MAIN_MISSION_UPDATE_AUTO":         1,
		"TRACK_MAIN_MISSION_UPDATE_MANUAL":       2,
		"TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT": 3,
	}
)

func (x TrackMainMissionUpdateType) Enum() *TrackMainMissionUpdateType {
	p := new(TrackMainMissionUpdateType)
	*p = x
	return p
}

func (x TrackMainMissionUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackMainMissionUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_TrackMainMissionUpdateType_proto_enumTypes[0].Descriptor()
}

func (TrackMainMissionUpdateType) Type() protoreflect.EnumType {
	return &file_TrackMainMissionUpdateType_proto_enumTypes[0]
}

func (x TrackMainMissionUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackMainMissionUpdateType.Descriptor instead.
func (TrackMainMissionUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_TrackMainMissionUpdateType_proto_rawDescGZIP(), []int{0}
}

var File_TrackMainMissionUpdateType_proto protoreflect.FileDescriptor

var file_TrackMainMissionUpdateType_proto_rawDesc = []byte{
	0x0a, 0x20, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xb6, 0x01, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41,
	0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12,
	0x2a, 0x0a, 0x26, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrackMainMissionUpdateType_proto_rawDescOnce sync.Once
	file_TrackMainMissionUpdateType_proto_rawDescData = file_TrackMainMissionUpdateType_proto_rawDesc
)

func file_TrackMainMissionUpdateType_proto_rawDescGZIP() []byte {
	file_TrackMainMissionUpdateType_proto_rawDescOnce.Do(func() {
		file_TrackMainMissionUpdateType_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrackMainMissionUpdateType_proto_rawDescData)
	})
	return file_TrackMainMissionUpdateType_proto_rawDescData
}

var file_TrackMainMissionUpdateType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TrackMainMissionUpdateType_proto_goTypes = []interface{}{
	(TrackMainMissionUpdateType)(0), // 0: TrackMainMissionUpdateType
}
var file_TrackMainMissionUpdateType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TrackMainMissionUpdateType_proto_init() }
func file_TrackMainMissionUpdateType_proto_init() {
	if File_TrackMainMissionUpdateType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TrackMainMissionUpdateType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrackMainMissionUpdateType_proto_goTypes,
		DependencyIndexes: file_TrackMainMissionUpdateType_proto_depIdxs,
		EnumInfos:         file_TrackMainMissionUpdateType_proto_enumTypes,
	}.Build()
	File_TrackMainMissionUpdateType_proto = out.File
	file_TrackMainMissionUpdateType_proto_rawDesc = nil
	file_TrackMainMissionUpdateType_proto_goTypes = nil
	file_TrackMainMissionUpdateType_proto_depIdxs = nil
}