// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CBILELFJAOM.proto

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

type CBILELFJAOM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasTakenReward bool   `protobuf:"varint,6,opt,name=has_taken_reward,json=hasTakenReward,proto3" json:"has_taken_reward,omitempty"`
	KEJFKHKPCNE    uint32 `protobuf:"varint,14,opt,name=KEJFKHKPCNE,proto3" json:"KEJFKHKPCNE,omitempty"`
}

func (x *CBILELFJAOM) Reset() {
	*x = CBILELFJAOM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBILELFJAOM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBILELFJAOM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBILELFJAOM) ProtoMessage() {}

func (x *CBILELFJAOM) ProtoReflect() protoreflect.Message {
	mi := &file_CBILELFJAOM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBILELFJAOM.ProtoReflect.Descriptor instead.
func (*CBILELFJAOM) Descriptor() ([]byte, []int) {
	return file_CBILELFJAOM_proto_rawDescGZIP(), []int{0}
}

func (x *CBILELFJAOM) GetHasTakenReward() bool {
	if x != nil {
		return x.HasTakenReward
	}
	return false
}

func (x *CBILELFJAOM) GetKEJFKHKPCNE() uint32 {
	if x != nil {
		return x.KEJFKHKPCNE
	}
	return 0
}

var File_CBILELFJAOM_proto protoreflect.FileDescriptor

var file_CBILELFJAOM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x42, 0x49, 0x4c, 0x45, 0x4c, 0x46, 0x4a, 0x41, 0x4f, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b, 0x43, 0x42, 0x49, 0x4c, 0x45, 0x4c, 0x46, 0x4a, 0x41,
	0x4f, 0x4d, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61,
	0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x4b, 0x45, 0x4a, 0x46, 0x4b, 0x48, 0x4b, 0x50, 0x43, 0x4e, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4b, 0x45, 0x4a, 0x46, 0x4b, 0x48, 0x4b, 0x50, 0x43, 0x4e, 0x45, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CBILELFJAOM_proto_rawDescOnce sync.Once
	file_CBILELFJAOM_proto_rawDescData = file_CBILELFJAOM_proto_rawDesc
)

func file_CBILELFJAOM_proto_rawDescGZIP() []byte {
	file_CBILELFJAOM_proto_rawDescOnce.Do(func() {
		file_CBILELFJAOM_proto_rawDescData = protoimpl.X.CompressGZIP(file_CBILELFJAOM_proto_rawDescData)
	})
	return file_CBILELFJAOM_proto_rawDescData
}

var file_CBILELFJAOM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CBILELFJAOM_proto_goTypes = []interface{}{
	(*CBILELFJAOM)(nil), // 0: CBILELFJAOM
}
var file_CBILELFJAOM_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CBILELFJAOM_proto_init() }
func file_CBILELFJAOM_proto_init() {
	if File_CBILELFJAOM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CBILELFJAOM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBILELFJAOM); i {
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
			RawDescriptor: file_CBILELFJAOM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CBILELFJAOM_proto_goTypes,
		DependencyIndexes: file_CBILELFJAOM_proto_depIdxs,
		MessageInfos:      file_CBILELFJAOM_proto_msgTypes,
	}.Build()
	File_CBILELFJAOM_proto = out.File
	file_CBILELFJAOM_proto_rawDesc = nil
	file_CBILELFJAOM_proto_goTypes = nil
	file_CBILELFJAOM_proto_depIdxs = nil
}
