// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ServerAnnounceNotify.proto

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

type ServerAnnounceNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnnounceDataList []*AnnounceData `protobuf:"bytes,6,rep,name=announce_data_list,json=announceDataList,proto3" json:"announce_data_list,omitempty"`
}

func (x *ServerAnnounceNotify) Reset() {
	*x = ServerAnnounceNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServerAnnounceNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAnnounceNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAnnounceNotify) ProtoMessage() {}

func (x *ServerAnnounceNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ServerAnnounceNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAnnounceNotify.ProtoReflect.Descriptor instead.
func (*ServerAnnounceNotify) Descriptor() ([]byte, []int) {
	return file_ServerAnnounceNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ServerAnnounceNotify) GetAnnounceDataList() []*AnnounceData {
	if x != nil {
		return x.AnnounceDataList
	}
	return nil
}

var File_ServerAnnounceNotify_proto protoreflect.FileDescriptor

var file_ServerAnnounceNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x53, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3b, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServerAnnounceNotify_proto_rawDescOnce sync.Once
	file_ServerAnnounceNotify_proto_rawDescData = file_ServerAnnounceNotify_proto_rawDesc
)

func file_ServerAnnounceNotify_proto_rawDescGZIP() []byte {
	file_ServerAnnounceNotify_proto_rawDescOnce.Do(func() {
		file_ServerAnnounceNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServerAnnounceNotify_proto_rawDescData)
	})
	return file_ServerAnnounceNotify_proto_rawDescData
}

var file_ServerAnnounceNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ServerAnnounceNotify_proto_goTypes = []interface{}{
	(*ServerAnnounceNotify)(nil), // 0: ServerAnnounceNotify
	(*AnnounceData)(nil),         // 1: AnnounceData
}
var file_ServerAnnounceNotify_proto_depIdxs = []int32{
	1, // 0: ServerAnnounceNotify.announce_data_list:type_name -> AnnounceData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ServerAnnounceNotify_proto_init() }
func file_ServerAnnounceNotify_proto_init() {
	if File_ServerAnnounceNotify_proto != nil {
		return
	}
	file_AnnounceData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ServerAnnounceNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAnnounceNotify); i {
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
			RawDescriptor: file_ServerAnnounceNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ServerAnnounceNotify_proto_goTypes,
		DependencyIndexes: file_ServerAnnounceNotify_proto_depIdxs,
		MessageInfos:      file_ServerAnnounceNotify_proto_msgTypes,
	}.Build()
	File_ServerAnnounceNotify_proto = out.File
	file_ServerAnnounceNotify_proto_rawDesc = nil
	file_ServerAnnounceNotify_proto_goTypes = nil
	file_ServerAnnounceNotify_proto_depIdxs = nil
}