// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DailyTaskDataScNotify.proto

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

type DailyTaskDataScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedNum   uint32       `protobuf:"varint,9,opt,name=finished_num,json=finishedNum,proto3" json:"finished_num,omitempty"`
	DailyTaskList []*DailyTask `protobuf:"bytes,14,rep,name=daily_task_list,json=dailyTaskList,proto3" json:"daily_task_list,omitempty"`
}

func (x *DailyTaskDataScNotify) Reset() {
	*x = DailyTaskDataScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyTaskDataScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyTaskDataScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyTaskDataScNotify) ProtoMessage() {}

func (x *DailyTaskDataScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DailyTaskDataScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyTaskDataScNotify.ProtoReflect.Descriptor instead.
func (*DailyTaskDataScNotify) Descriptor() ([]byte, []int) {
	return file_DailyTaskDataScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DailyTaskDataScNotify) GetFinishedNum() uint32 {
	if x != nil {
		return x.FinishedNum
	}
	return 0
}

func (x *DailyTaskDataScNotify) GetDailyTaskList() []*DailyTask {
	if x != nil {
		return x.DailyTaskList
	}
	return nil
}

var File_DailyTaskDataScNotify_proto protoreflect.FileDescriptor

var file_DailyTaskDataScNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x15, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x0f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x0d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DailyTaskDataScNotify_proto_rawDescOnce sync.Once
	file_DailyTaskDataScNotify_proto_rawDescData = file_DailyTaskDataScNotify_proto_rawDesc
)

func file_DailyTaskDataScNotify_proto_rawDescGZIP() []byte {
	file_DailyTaskDataScNotify_proto_rawDescOnce.Do(func() {
		file_DailyTaskDataScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyTaskDataScNotify_proto_rawDescData)
	})
	return file_DailyTaskDataScNotify_proto_rawDescData
}

var file_DailyTaskDataScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyTaskDataScNotify_proto_goTypes = []interface{}{
	(*DailyTaskDataScNotify)(nil), // 0: DailyTaskDataScNotify
	(*DailyTask)(nil),             // 1: DailyTask
}
var file_DailyTaskDataScNotify_proto_depIdxs = []int32{
	1, // 0: DailyTaskDataScNotify.daily_task_list:type_name -> DailyTask
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DailyTaskDataScNotify_proto_init() }
func file_DailyTaskDataScNotify_proto_init() {
	if File_DailyTaskDataScNotify_proto != nil {
		return
	}
	file_DailyTask_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DailyTaskDataScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyTaskDataScNotify); i {
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
			RawDescriptor: file_DailyTaskDataScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyTaskDataScNotify_proto_goTypes,
		DependencyIndexes: file_DailyTaskDataScNotify_proto_depIdxs,
		MessageInfos:      file_DailyTaskDataScNotify_proto_msgTypes,
	}.Build()
	File_DailyTaskDataScNotify_proto = out.File
	file_DailyTaskDataScNotify_proto_rawDesc = nil
	file_DailyTaskDataScNotify_proto_goTypes = nil
	file_DailyTaskDataScNotify_proto_depIdxs = nil
}