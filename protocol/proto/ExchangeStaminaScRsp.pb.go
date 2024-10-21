// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ExchangeStaminaScRsp.proto

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

type ExchangeStaminaScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeTimesFieldNumber   uint32      `protobuf:"varint,10,opt,name=ExchangeTimesFieldNumber,proto3" json:"ExchangeTimesFieldNumber,omitempty"`
	LastRecoverTimeFieldNumber int64       `protobuf:"varint,3,opt,name=LastRecoverTimeFieldNumber,proto3" json:"LastRecoverTimeFieldNumber,omitempty"`
	StaminaAddFieldNumber      uint32      `protobuf:"varint,9,opt,name=StaminaAddFieldNumber,proto3" json:"StaminaAddFieldNumber,omitempty"`
	Retcode                    uint32      `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ItemCostListFieldNumber    []*ItemCost `protobuf:"bytes,14,rep,name=ItemCostListFieldNumber,proto3" json:"ItemCostListFieldNumber,omitempty"`
}

func (x *ExchangeStaminaScRsp) Reset() {
	*x = ExchangeStaminaScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExchangeStaminaScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeStaminaScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeStaminaScRsp) ProtoMessage() {}

func (x *ExchangeStaminaScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ExchangeStaminaScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeStaminaScRsp.ProtoReflect.Descriptor instead.
func (*ExchangeStaminaScRsp) Descriptor() ([]byte, []int) {
	return file_ExchangeStaminaScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeStaminaScRsp) GetExchangeTimesFieldNumber() uint32 {
	if x != nil {
		return x.ExchangeTimesFieldNumber
	}
	return 0
}

func (x *ExchangeStaminaScRsp) GetLastRecoverTimeFieldNumber() int64 {
	if x != nil {
		return x.LastRecoverTimeFieldNumber
	}
	return 0
}

func (x *ExchangeStaminaScRsp) GetStaminaAddFieldNumber() uint32 {
	if x != nil {
		return x.StaminaAddFieldNumber
	}
	return 0
}

func (x *ExchangeStaminaScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ExchangeStaminaScRsp) GetItemCostListFieldNumber() []*ItemCost {
	if x != nil {
		return x.ItemCostListFieldNumber
	}
	return nil
}

var File_ExchangeStaminaScRsp_proto protoreflect.FileDescriptor

var file_ExchangeStaminaScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e,
	0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a,
	0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x18, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x1a, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x41, 0x64, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x15, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x41, 0x64, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x43, 0x0a, 0x17, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x17, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExchangeStaminaScRsp_proto_rawDescOnce sync.Once
	file_ExchangeStaminaScRsp_proto_rawDescData = file_ExchangeStaminaScRsp_proto_rawDesc
)

func file_ExchangeStaminaScRsp_proto_rawDescGZIP() []byte {
	file_ExchangeStaminaScRsp_proto_rawDescOnce.Do(func() {
		file_ExchangeStaminaScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExchangeStaminaScRsp_proto_rawDescData)
	})
	return file_ExchangeStaminaScRsp_proto_rawDescData
}

var file_ExchangeStaminaScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExchangeStaminaScRsp_proto_goTypes = []interface{}{
	(*ExchangeStaminaScRsp)(nil), // 0: ExchangeStaminaScRsp
	(*ItemCost)(nil),             // 1: ItemCost
}
var file_ExchangeStaminaScRsp_proto_depIdxs = []int32{
	1, // 0: ExchangeStaminaScRsp.ItemCostListFieldNumber:type_name -> ItemCost
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ExchangeStaminaScRsp_proto_init() }
func file_ExchangeStaminaScRsp_proto_init() {
	if File_ExchangeStaminaScRsp_proto != nil {
		return
	}
	file_ItemCost_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExchangeStaminaScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeStaminaScRsp); i {
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
			RawDescriptor: file_ExchangeStaminaScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExchangeStaminaScRsp_proto_goTypes,
		DependencyIndexes: file_ExchangeStaminaScRsp_proto_depIdxs,
		MessageInfos:      file_ExchangeStaminaScRsp_proto_msgTypes,
	}.Build()
	File_ExchangeStaminaScRsp_proto = out.File
	file_ExchangeStaminaScRsp_proto_rawDesc = nil
	file_ExchangeStaminaScRsp_proto_goTypes = nil
	file_ExchangeStaminaScRsp_proto_depIdxs = nil
}
