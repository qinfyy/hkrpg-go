// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueHandbookMiracle.proto

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

type RogueHandbookMiracle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleId  uint32 `protobuf:"varint,3,opt,name=miracle_id,json=miracleId,proto3" json:"miracle_id,omitempty"`
	IsUnlocked bool   `protobuf:"varint,14,opt,name=is_unlocked,json=isUnlocked,proto3" json:"is_unlocked,omitempty"`
}

func (x *RogueHandbookMiracle) Reset() {
	*x = RogueHandbookMiracle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueHandbookMiracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueHandbookMiracle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueHandbookMiracle) ProtoMessage() {}

func (x *RogueHandbookMiracle) ProtoReflect() protoreflect.Message {
	mi := &file_RogueHandbookMiracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueHandbookMiracle.ProtoReflect.Descriptor instead.
func (*RogueHandbookMiracle) Descriptor() ([]byte, []int) {
	return file_RogueHandbookMiracle_proto_rawDescGZIP(), []int{0}
}

func (x *RogueHandbookMiracle) GetMiracleId() uint32 {
	if x != nil {
		return x.MiracleId
	}
	return 0
}

func (x *RogueHandbookMiracle) GetIsUnlocked() bool {
	if x != nil {
		return x.IsUnlocked
	}
	return false
}

var File_RogueHandbookMiracle_proto protoreflect.FileDescriptor

var file_RogueHandbookMiracle_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x4d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueHandbookMiracle_proto_rawDescOnce sync.Once
	file_RogueHandbookMiracle_proto_rawDescData = file_RogueHandbookMiracle_proto_rawDesc
)

func file_RogueHandbookMiracle_proto_rawDescGZIP() []byte {
	file_RogueHandbookMiracle_proto_rawDescOnce.Do(func() {
		file_RogueHandbookMiracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueHandbookMiracle_proto_rawDescData)
	})
	return file_RogueHandbookMiracle_proto_rawDescData
}

var file_RogueHandbookMiracle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueHandbookMiracle_proto_goTypes = []interface{}{
	(*RogueHandbookMiracle)(nil), // 0: RogueHandbookMiracle
}
var file_RogueHandbookMiracle_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueHandbookMiracle_proto_init() }
func file_RogueHandbookMiracle_proto_init() {
	if File_RogueHandbookMiracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueHandbookMiracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueHandbookMiracle); i {
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
			RawDescriptor: file_RogueHandbookMiracle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueHandbookMiracle_proto_goTypes,
		DependencyIndexes: file_RogueHandbookMiracle_proto_depIdxs,
		MessageInfos:      file_RogueHandbookMiracle_proto_msgTypes,
	}.Build()
	File_RogueHandbookMiracle_proto = out.File
	file_RogueHandbookMiracle_proto_rawDesc = nil
	file_RogueHandbookMiracle_proto_goTypes = nil
	file_RogueHandbookMiracle_proto_depIdxs = nil
}
