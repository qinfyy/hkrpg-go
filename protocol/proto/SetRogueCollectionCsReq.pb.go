// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetRogueCollectionCsReq.proto

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

type SetRogueCollectionCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MBIMLIIPCDI []RogueCollectionExhibitionOperateType `protobuf:"varint,3,rep,packed,name=MBIMLIIPCDI,proto3,enum=RogueCollectionExhibitionOperateType" json:"MBIMLIIPCDI,omitempty"`
	PIHNJBPPKHO []uint32                               `protobuf:"varint,4,rep,packed,name=PIHNJBPPKHO,proto3" json:"PIHNJBPPKHO,omitempty"`
	PDNOJBIDKKM []uint32                               `protobuf:"varint,10,rep,packed,name=PDNOJBIDKKM,proto3" json:"PDNOJBIDKKM,omitempty"`
}

func (x *SetRogueCollectionCsReq) Reset() {
	*x = SetRogueCollectionCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetRogueCollectionCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRogueCollectionCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRogueCollectionCsReq) ProtoMessage() {}

func (x *SetRogueCollectionCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetRogueCollectionCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRogueCollectionCsReq.ProtoReflect.Descriptor instead.
func (*SetRogueCollectionCsReq) Descriptor() ([]byte, []int) {
	return file_SetRogueCollectionCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetRogueCollectionCsReq) GetMBIMLIIPCDI() []RogueCollectionExhibitionOperateType {
	if x != nil {
		return x.MBIMLIIPCDI
	}
	return nil
}

func (x *SetRogueCollectionCsReq) GetPIHNJBPPKHO() []uint32 {
	if x != nil {
		return x.PIHNJBPPKHO
	}
	return nil
}

func (x *SetRogueCollectionCsReq) GetPDNOJBIDKKM() []uint32 {
	if x != nil {
		return x.PDNOJBIDKKM
	}
	return nil
}

var File_SetRogueCollectionCsReq_proto protoreflect.FileDescriptor

var file_SetRogueCollectionCsReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x17,
	0x53, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x47, 0x0a, 0x0b, 0x4d, 0x42, 0x49, 0x4d, 0x4c,
	0x49, 0x49, 0x50, 0x43, 0x44, 0x49, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x4d, 0x42, 0x49, 0x4d, 0x4c, 0x49, 0x49, 0x50, 0x43, 0x44, 0x49,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x49, 0x48, 0x4e, 0x4a, 0x42, 0x50, 0x50, 0x4b, 0x48, 0x4f, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x49, 0x48, 0x4e, 0x4a, 0x42, 0x50, 0x50, 0x4b,
	0x48, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x44, 0x4e, 0x4f, 0x4a, 0x42, 0x49, 0x44, 0x4b, 0x4b,
	0x4d, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x44, 0x4e, 0x4f, 0x4a, 0x42, 0x49,
	0x44, 0x4b, 0x4b, 0x4d, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetRogueCollectionCsReq_proto_rawDescOnce sync.Once
	file_SetRogueCollectionCsReq_proto_rawDescData = file_SetRogueCollectionCsReq_proto_rawDesc
)

func file_SetRogueCollectionCsReq_proto_rawDescGZIP() []byte {
	file_SetRogueCollectionCsReq_proto_rawDescOnce.Do(func() {
		file_SetRogueCollectionCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetRogueCollectionCsReq_proto_rawDescData)
	})
	return file_SetRogueCollectionCsReq_proto_rawDescData
}

var file_SetRogueCollectionCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetRogueCollectionCsReq_proto_goTypes = []interface{}{
	(*SetRogueCollectionCsReq)(nil),           // 0: SetRogueCollectionCsReq
	(RogueCollectionExhibitionOperateType)(0), // 1: RogueCollectionExhibitionOperateType
}
var file_SetRogueCollectionCsReq_proto_depIdxs = []int32{
	1, // 0: SetRogueCollectionCsReq.MBIMLIIPCDI:type_name -> RogueCollectionExhibitionOperateType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetRogueCollectionCsReq_proto_init() }
func file_SetRogueCollectionCsReq_proto_init() {
	if File_SetRogueCollectionCsReq_proto != nil {
		return
	}
	file_RogueCollectionExhibitionOperateType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetRogueCollectionCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRogueCollectionCsReq); i {
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
			RawDescriptor: file_SetRogueCollectionCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetRogueCollectionCsReq_proto_goTypes,
		DependencyIndexes: file_SetRogueCollectionCsReq_proto_depIdxs,
		MessageInfos:      file_SetRogueCollectionCsReq_proto_msgTypes,
	}.Build()
	File_SetRogueCollectionCsReq_proto = out.File
	file_SetRogueCollectionCsReq_proto_rawDesc = nil
	file_SetRogueCollectionCsReq_proto_goTypes = nil
	file_SetRogueCollectionCsReq_proto_depIdxs = nil
}
