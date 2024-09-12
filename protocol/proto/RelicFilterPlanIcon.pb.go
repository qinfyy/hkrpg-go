// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RelicFilterPlanIcon.proto

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

type RelicFilterPlanIcon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvatarIconFieldNumber bool   `protobuf:"varint,6,opt,name=IsAvatarIconFieldNumber,proto3" json:"IsAvatarIconFieldNumber,omitempty"`
	IconIdFieldNumber       uint32 `protobuf:"varint,3,opt,name=IconIdFieldNumber,proto3" json:"IconIdFieldNumber,omitempty"`
}

func (x *RelicFilterPlanIcon) Reset() {
	*x = RelicFilterPlanIcon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RelicFilterPlanIcon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelicFilterPlanIcon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicFilterPlanIcon) ProtoMessage() {}

func (x *RelicFilterPlanIcon) ProtoReflect() protoreflect.Message {
	mi := &file_RelicFilterPlanIcon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelicFilterPlanIcon.ProtoReflect.Descriptor instead.
func (*RelicFilterPlanIcon) Descriptor() ([]byte, []int) {
	return file_RelicFilterPlanIcon_proto_rawDescGZIP(), []int{0}
}

func (x *RelicFilterPlanIcon) GetIsAvatarIconFieldNumber() bool {
	if x != nil {
		return x.IsAvatarIconFieldNumber
	}
	return false
}

func (x *RelicFilterPlanIcon) GetIconIdFieldNumber() uint32 {
	if x != nil {
		return x.IconIdFieldNumber
	}
	return 0
}

var File_RelicFilterPlanIcon_proto protoreflect.FileDescriptor

var file_RelicFilterPlanIcon_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x13, 0x52,
	0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x63,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x49, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x63,
	0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x17, 0x49, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x63, 0x6f,
	0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11,
	0x49, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x49, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RelicFilterPlanIcon_proto_rawDescOnce sync.Once
	file_RelicFilterPlanIcon_proto_rawDescData = file_RelicFilterPlanIcon_proto_rawDesc
)

func file_RelicFilterPlanIcon_proto_rawDescGZIP() []byte {
	file_RelicFilterPlanIcon_proto_rawDescOnce.Do(func() {
		file_RelicFilterPlanIcon_proto_rawDescData = protoimpl.X.CompressGZIP(file_RelicFilterPlanIcon_proto_rawDescData)
	})
	return file_RelicFilterPlanIcon_proto_rawDescData
}

var file_RelicFilterPlanIcon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RelicFilterPlanIcon_proto_goTypes = []interface{}{
	(*RelicFilterPlanIcon)(nil), // 0: RelicFilterPlanIcon
}
var file_RelicFilterPlanIcon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RelicFilterPlanIcon_proto_init() }
func file_RelicFilterPlanIcon_proto_init() {
	if File_RelicFilterPlanIcon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RelicFilterPlanIcon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelicFilterPlanIcon); i {
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
			RawDescriptor: file_RelicFilterPlanIcon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RelicFilterPlanIcon_proto_goTypes,
		DependencyIndexes: file_RelicFilterPlanIcon_proto_depIdxs,
		MessageInfos:      file_RelicFilterPlanIcon_proto_msgTypes,
	}.Build()
	File_RelicFilterPlanIcon_proto = out.File
	file_RelicFilterPlanIcon_proto_rawDesc = nil
	file_RelicFilterPlanIcon_proto_goTypes = nil
	file_RelicFilterPlanIcon_proto_depIdxs = nil
}