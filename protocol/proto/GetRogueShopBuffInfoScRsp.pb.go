// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetRogueShopBuffInfoScRsp.proto

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

type GetRogueShopBuffInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueBuffInfo *IPBJHDCMIKO  `protobuf:"bytes,6,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	CMIJHIFKNPO   bool          `protobuf:"varint,14,opt,name=CMIJHIFKNPO,proto3" json:"CMIJHIFKNPO,omitempty"`
	KNPFHBLMPGI   *ItemCostData `protobuf:"bytes,7,opt,name=KNPFHBLMPGI,proto3" json:"KNPFHBLMPGI,omitempty"`
	Retcode       uint32        `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetRogueShopBuffInfoScRsp) Reset() {
	*x = GetRogueShopBuffInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRogueShopBuffInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRogueShopBuffInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRogueShopBuffInfoScRsp) ProtoMessage() {}

func (x *GetRogueShopBuffInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRogueShopBuffInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRogueShopBuffInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetRogueShopBuffInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetRogueShopBuffInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRogueShopBuffInfoScRsp) GetRogueBuffInfo() *IPBJHDCMIKO {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *GetRogueShopBuffInfoScRsp) GetCMIJHIFKNPO() bool {
	if x != nil {
		return x.CMIJHIFKNPO
	}
	return false
}

func (x *GetRogueShopBuffInfoScRsp) GetKNPFHBLMPGI() *ItemCostData {
	if x != nil {
		return x.KNPFHBLMPGI
	}
	return nil
}

func (x *GetRogueShopBuffInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetRogueShopBuffInfoScRsp_proto protoreflect.FileDescriptor

var file_GetRogueShopBuffInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75,
	0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x50, 0x42, 0x4a, 0x48, 0x44, 0x43, 0x4d, 0x49,
	0x4b, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f,
	0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x50, 0x42, 0x4a, 0x48, 0x44, 0x43, 0x4d, 0x49, 0x4b, 0x4f, 0x52, 0x0d, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x4d, 0x49, 0x4a, 0x48, 0x49, 0x46, 0x4b, 0x4e, 0x50, 0x4f, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x43, 0x4d, 0x49, 0x4a, 0x48, 0x49, 0x46, 0x4b, 0x4e, 0x50, 0x4f, 0x12, 0x2f,
	0x0a, 0x0b, 0x4b, 0x4e, 0x50, 0x46, 0x48, 0x42, 0x4c, 0x4d, 0x50, 0x47, 0x49, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x4b, 0x4e, 0x50, 0x46, 0x48, 0x42, 0x4c, 0x4d, 0x50, 0x47, 0x49, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetRogueShopBuffInfoScRsp_proto_rawDescOnce sync.Once
	file_GetRogueShopBuffInfoScRsp_proto_rawDescData = file_GetRogueShopBuffInfoScRsp_proto_rawDesc
)

func file_GetRogueShopBuffInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetRogueShopBuffInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetRogueShopBuffInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRogueShopBuffInfoScRsp_proto_rawDescData)
	})
	return file_GetRogueShopBuffInfoScRsp_proto_rawDescData
}

var file_GetRogueShopBuffInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRogueShopBuffInfoScRsp_proto_goTypes = []interface{}{
	(*GetRogueShopBuffInfoScRsp)(nil), // 0: GetRogueShopBuffInfoScRsp
	(*IPBJHDCMIKO)(nil),               // 1: IPBJHDCMIKO
	(*ItemCostData)(nil),              // 2: ItemCostData
}
var file_GetRogueShopBuffInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetRogueShopBuffInfoScRsp.rogue_buff_info:type_name -> IPBJHDCMIKO
	2, // 1: GetRogueShopBuffInfoScRsp.KNPFHBLMPGI:type_name -> ItemCostData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetRogueShopBuffInfoScRsp_proto_init() }
func file_GetRogueShopBuffInfoScRsp_proto_init() {
	if File_GetRogueShopBuffInfoScRsp_proto != nil {
		return
	}
	file_ItemCostData_proto_init()
	file_IPBJHDCMIKO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetRogueShopBuffInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRogueShopBuffInfoScRsp); i {
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
			RawDescriptor: file_GetRogueShopBuffInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRogueShopBuffInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetRogueShopBuffInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetRogueShopBuffInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetRogueShopBuffInfoScRsp_proto = out.File
	file_GetRogueShopBuffInfoScRsp_proto_rawDesc = nil
	file_GetRogueShopBuffInfoScRsp_proto_goTypes = nil
	file_GetRogueShopBuffInfoScRsp_proto_depIdxs = nil
}