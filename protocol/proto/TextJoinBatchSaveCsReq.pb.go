// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TextJoinBatchSaveCsReq.proto

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

type TextJoinBatchSaveCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextJoinList []*TextJoinInfo `protobuf:"bytes,13,rep,name=text_join_list,json=textJoinList,proto3" json:"text_join_list,omitempty"`
}

func (x *TextJoinBatchSaveCsReq) Reset() {
	*x = TextJoinBatchSaveCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TextJoinBatchSaveCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextJoinBatchSaveCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextJoinBatchSaveCsReq) ProtoMessage() {}

func (x *TextJoinBatchSaveCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_TextJoinBatchSaveCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextJoinBatchSaveCsReq.ProtoReflect.Descriptor instead.
func (*TextJoinBatchSaveCsReq) Descriptor() ([]byte, []int) {
	return file_TextJoinBatchSaveCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *TextJoinBatchSaveCsReq) GetTextJoinList() []*TextJoinInfo {
	if x != nil {
		return x.TextJoinList
	}
	return nil
}

var File_TextJoinBatchSaveCsReq_proto protoreflect.FileDescriptor

var file_TextJoinBatchSaveCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x54, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x61, 0x76, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x54, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x16, 0x54, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a, 0x0e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TextJoinBatchSaveCsReq_proto_rawDescOnce sync.Once
	file_TextJoinBatchSaveCsReq_proto_rawDescData = file_TextJoinBatchSaveCsReq_proto_rawDesc
)

func file_TextJoinBatchSaveCsReq_proto_rawDescGZIP() []byte {
	file_TextJoinBatchSaveCsReq_proto_rawDescOnce.Do(func() {
		file_TextJoinBatchSaveCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TextJoinBatchSaveCsReq_proto_rawDescData)
	})
	return file_TextJoinBatchSaveCsReq_proto_rawDescData
}

var file_TextJoinBatchSaveCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TextJoinBatchSaveCsReq_proto_goTypes = []interface{}{
	(*TextJoinBatchSaveCsReq)(nil), // 0: TextJoinBatchSaveCsReq
	(*TextJoinInfo)(nil),           // 1: TextJoinInfo
}
var file_TextJoinBatchSaveCsReq_proto_depIdxs = []int32{
	1, // 0: TextJoinBatchSaveCsReq.text_join_list:type_name -> TextJoinInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TextJoinBatchSaveCsReq_proto_init() }
func file_TextJoinBatchSaveCsReq_proto_init() {
	if File_TextJoinBatchSaveCsReq_proto != nil {
		return
	}
	file_TextJoinInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TextJoinBatchSaveCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextJoinBatchSaveCsReq); i {
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
			RawDescriptor: file_TextJoinBatchSaveCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TextJoinBatchSaveCsReq_proto_goTypes,
		DependencyIndexes: file_TextJoinBatchSaveCsReq_proto_depIdxs,
		MessageInfos:      file_TextJoinBatchSaveCsReq_proto_msgTypes,
	}.Build()
	File_TextJoinBatchSaveCsReq_proto = out.File
	file_TextJoinBatchSaveCsReq_proto_rawDesc = nil
	file_TextJoinBatchSaveCsReq_proto_goTypes = nil
	file_TextJoinBatchSaveCsReq_proto_depIdxs = nil
}
