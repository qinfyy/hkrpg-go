// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetLoginActivityCsReq.proto

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

type GetLoginActivityCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoginActivityCsReq) Reset() {
	*x = GetLoginActivityCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetLoginActivityCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginActivityCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginActivityCsReq) ProtoMessage() {}

func (x *GetLoginActivityCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetLoginActivityCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginActivityCsReq.ProtoReflect.Descriptor instead.
func (*GetLoginActivityCsReq) Descriptor() ([]byte, []int) {
	return file_GetLoginActivityCsReq_proto_rawDescGZIP(), []int{0}
}

var File_GetLoginActivityCsReq_proto protoreflect.FileDescriptor

var file_GetLoginActivityCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetLoginActivityCsReq_proto_rawDescOnce sync.Once
	file_GetLoginActivityCsReq_proto_rawDescData = file_GetLoginActivityCsReq_proto_rawDesc
)

func file_GetLoginActivityCsReq_proto_rawDescGZIP() []byte {
	file_GetLoginActivityCsReq_proto_rawDescOnce.Do(func() {
		file_GetLoginActivityCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetLoginActivityCsReq_proto_rawDescData)
	})
	return file_GetLoginActivityCsReq_proto_rawDescData
}

var file_GetLoginActivityCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetLoginActivityCsReq_proto_goTypes = []interface{}{
	(*GetLoginActivityCsReq)(nil), // 0: GetLoginActivityCsReq
}
var file_GetLoginActivityCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetLoginActivityCsReq_proto_init() }
func file_GetLoginActivityCsReq_proto_init() {
	if File_GetLoginActivityCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetLoginActivityCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginActivityCsReq); i {
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
			RawDescriptor: file_GetLoginActivityCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetLoginActivityCsReq_proto_goTypes,
		DependencyIndexes: file_GetLoginActivityCsReq_proto_depIdxs,
		MessageInfos:      file_GetLoginActivityCsReq_proto_msgTypes,
	}.Build()
	File_GetLoginActivityCsReq_proto = out.File
	file_GetLoginActivityCsReq_proto_rawDesc = nil
	file_GetLoginActivityCsReq_proto_goTypes = nil
	file_GetLoginActivityCsReq_proto_depIdxs = nil
}
