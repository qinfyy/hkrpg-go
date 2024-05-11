// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueHandbookAeon.proto

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

type RogueHandbookAeon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp               uint32   `protobuf:"varint,6,opt,name=exp,proto3" json:"exp,omitempty"`
	Level             uint32   `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	AeonId            uint32   `protobuf:"varint,5,opt,name=aeon_id,json=aeonId,proto3" json:"aeon_id,omitempty"`
	ArchiveUnlockList []uint32 `protobuf:"varint,15,rep,packed,name=archive_unlock_list,json=archiveUnlockList,proto3" json:"archive_unlock_list,omitempty"` // KPFEEKHKANG
	MaxLevel          uint32   `protobuf:"varint,14,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	TakenRewardList   []uint32 `protobuf:"varint,7,rep,packed,name=taken_reward_list,json=takenRewardList,proto3" json:"taken_reward_list,omitempty"` // BODAGGMFDMJ
}

func (x *RogueHandbookAeon) Reset() {
	*x = RogueHandbookAeon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueHandbookAeon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueHandbookAeon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueHandbookAeon) ProtoMessage() {}

func (x *RogueHandbookAeon) ProtoReflect() protoreflect.Message {
	mi := &file_RogueHandbookAeon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueHandbookAeon.ProtoReflect.Descriptor instead.
func (*RogueHandbookAeon) Descriptor() ([]byte, []int) {
	return file_RogueHandbookAeon_proto_rawDescGZIP(), []int{0}
}

func (x *RogueHandbookAeon) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *RogueHandbookAeon) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RogueHandbookAeon) GetAeonId() uint32 {
	if x != nil {
		return x.AeonId
	}
	return 0
}

func (x *RogueHandbookAeon) GetArchiveUnlockList() []uint32 {
	if x != nil {
		return x.ArchiveUnlockList
	}
	return nil
}

func (x *RogueHandbookAeon) GetMaxLevel() uint32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *RogueHandbookAeon) GetTakenRewardList() []uint32 {
	if x != nil {
		return x.TakenRewardList
	}
	return nil
}

var File_RogueHandbookAeon_proto protoreflect.FileDescriptor

var file_RogueHandbookAeon_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x41,
	0x65, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x11, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x41, 0x65, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x65, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x65, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2a, 0x0a,
	0x11, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueHandbookAeon_proto_rawDescOnce sync.Once
	file_RogueHandbookAeon_proto_rawDescData = file_RogueHandbookAeon_proto_rawDesc
)

func file_RogueHandbookAeon_proto_rawDescGZIP() []byte {
	file_RogueHandbookAeon_proto_rawDescOnce.Do(func() {
		file_RogueHandbookAeon_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueHandbookAeon_proto_rawDescData)
	})
	return file_RogueHandbookAeon_proto_rawDescData
}

var file_RogueHandbookAeon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueHandbookAeon_proto_goTypes = []interface{}{
	(*RogueHandbookAeon)(nil), // 0: RogueHandbookAeon
}
var file_RogueHandbookAeon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueHandbookAeon_proto_init() }
func file_RogueHandbookAeon_proto_init() {
	if File_RogueHandbookAeon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueHandbookAeon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueHandbookAeon); i {
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
			RawDescriptor: file_RogueHandbookAeon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueHandbookAeon_proto_goTypes,
		DependencyIndexes: file_RogueHandbookAeon_proto_depIdxs,
		MessageInfos:      file_RogueHandbookAeon_proto_msgTypes,
	}.Build()
	File_RogueHandbookAeon_proto = out.File
	file_RogueHandbookAeon_proto_rawDesc = nil
	file_RogueHandbookAeon_proto_goTypes = nil
	file_RogueHandbookAeon_proto_depIdxs = nil
}
