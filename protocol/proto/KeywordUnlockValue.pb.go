// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KeywordUnlockValue.proto

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

type KeywordUnlockValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeywordUnlockMap map[uint32]bool `protobuf:"bytes,3,rep,name=keyword_unlock_map,json=keywordUnlockMap,proto3" json:"keyword_unlock_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *KeywordUnlockValue) Reset() {
	*x = KeywordUnlockValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeywordUnlockValue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordUnlockValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordUnlockValue) ProtoMessage() {}

func (x *KeywordUnlockValue) ProtoReflect() protoreflect.Message {
	mi := &file_KeywordUnlockValue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordUnlockValue.ProtoReflect.Descriptor instead.
func (*KeywordUnlockValue) Descriptor() ([]byte, []int) {
	return file_KeywordUnlockValue_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordUnlockValue) GetKeywordUnlockMap() map[uint32]bool {
	if x != nil {
		return x.KeywordUnlockMap
	}
	return nil
}

var File_KeywordUnlockValue_proto protoreflect.FileDescriptor

var file_KeywordUnlockValue_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x12, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x57, 0x0a, 0x12, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x70, 0x1a, 0x43, 0x0a, 0x15, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KeywordUnlockValue_proto_rawDescOnce sync.Once
	file_KeywordUnlockValue_proto_rawDescData = file_KeywordUnlockValue_proto_rawDesc
)

func file_KeywordUnlockValue_proto_rawDescGZIP() []byte {
	file_KeywordUnlockValue_proto_rawDescOnce.Do(func() {
		file_KeywordUnlockValue_proto_rawDescData = protoimpl.X.CompressGZIP(file_KeywordUnlockValue_proto_rawDescData)
	})
	return file_KeywordUnlockValue_proto_rawDescData
}

var file_KeywordUnlockValue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_KeywordUnlockValue_proto_goTypes = []interface{}{
	(*KeywordUnlockValue)(nil), // 0: KeywordUnlockValue
	nil,                        // 1: KeywordUnlockValue.KeywordUnlockMapEntry
}
var file_KeywordUnlockValue_proto_depIdxs = []int32{
	1, // 0: KeywordUnlockValue.keyword_unlock_map:type_name -> KeywordUnlockValue.KeywordUnlockMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KeywordUnlockValue_proto_init() }
func file_KeywordUnlockValue_proto_init() {
	if File_KeywordUnlockValue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KeywordUnlockValue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordUnlockValue); i {
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
			RawDescriptor: file_KeywordUnlockValue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KeywordUnlockValue_proto_goTypes,
		DependencyIndexes: file_KeywordUnlockValue_proto_depIdxs,
		MessageInfos:      file_KeywordUnlockValue_proto_msgTypes,
	}.Build()
	File_KeywordUnlockValue_proto = out.File
	file_KeywordUnlockValue_proto_rawDesc = nil
	file_KeywordUnlockValue_proto_goTypes = nil
	file_KeywordUnlockValue_proto_depIdxs = nil
}
