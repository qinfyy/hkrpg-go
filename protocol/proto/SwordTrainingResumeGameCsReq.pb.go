// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingResumeGameCsReq.proto

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

type SwordTrainingResumeGameCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameStoryLineId uint32 `protobuf:"varint,8,opt,name=game_story_line_id,json=gameStoryLineId,proto3" json:"game_story_line_id,omitempty"`
}

func (x *SwordTrainingResumeGameCsReq) Reset() {
	*x = SwordTrainingResumeGameCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingResumeGameCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingResumeGameCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingResumeGameCsReq) ProtoMessage() {}

func (x *SwordTrainingResumeGameCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingResumeGameCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingResumeGameCsReq.ProtoReflect.Descriptor instead.
func (*SwordTrainingResumeGameCsReq) Descriptor() ([]byte, []int) {
	return file_SwordTrainingResumeGameCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingResumeGameCsReq) GetGameStoryLineId() uint32 {
	if x != nil {
		return x.GameStoryLineId
	}
	return 0
}

var File_SwordTrainingResumeGameCsReq_proto protoreflect.FileDescriptor

var file_SwordTrainingResumeGameCsReq_proto_rawDesc = []byte{
	0x0a, 0x22, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x1c, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingResumeGameCsReq_proto_rawDescOnce sync.Once
	file_SwordTrainingResumeGameCsReq_proto_rawDescData = file_SwordTrainingResumeGameCsReq_proto_rawDesc
)

func file_SwordTrainingResumeGameCsReq_proto_rawDescGZIP() []byte {
	file_SwordTrainingResumeGameCsReq_proto_rawDescOnce.Do(func() {
		file_SwordTrainingResumeGameCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingResumeGameCsReq_proto_rawDescData)
	})
	return file_SwordTrainingResumeGameCsReq_proto_rawDescData
}

var file_SwordTrainingResumeGameCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingResumeGameCsReq_proto_goTypes = []interface{}{
	(*SwordTrainingResumeGameCsReq)(nil), // 0: SwordTrainingResumeGameCsReq
}
var file_SwordTrainingResumeGameCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SwordTrainingResumeGameCsReq_proto_init() }
func file_SwordTrainingResumeGameCsReq_proto_init() {
	if File_SwordTrainingResumeGameCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingResumeGameCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingResumeGameCsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SwordTrainingResumeGameCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingResumeGameCsReq_proto_goTypes,
		DependencyIndexes: file_SwordTrainingResumeGameCsReq_proto_depIdxs,
		MessageInfos:      file_SwordTrainingResumeGameCsReq_proto_msgTypes,
	}.Build()
	File_SwordTrainingResumeGameCsReq_proto = out.File
	file_SwordTrainingResumeGameCsReq_proto_rawDesc = nil
	file_SwordTrainingResumeGameCsReq_proto_goTypes = nil
	file_SwordTrainingResumeGameCsReq_proto_depIdxs = nil
}
