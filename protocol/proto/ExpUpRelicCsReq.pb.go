// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ExpUpRelicCsReq.proto

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

type ExpUpRelicCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelicUniqueId uint32        `protobuf:"varint,11,opt,name=relic_unique_id,json=relicUniqueId,proto3" json:"relic_unique_id,omitempty"`
	ItemCostList  *ItemCostList `protobuf:"bytes,3,opt,name=item_cost_list,json=itemCostList,proto3" json:"item_cost_list,omitempty"`
}

func (x *ExpUpRelicCsReq) Reset() {
	*x = ExpUpRelicCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExpUpRelicCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpUpRelicCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpUpRelicCsReq) ProtoMessage() {}

func (x *ExpUpRelicCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ExpUpRelicCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpUpRelicCsReq.ProtoReflect.Descriptor instead.
func (*ExpUpRelicCsReq) Descriptor() ([]byte, []int) {
	return file_ExpUpRelicCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ExpUpRelicCsReq) GetRelicUniqueId() uint32 {
	if x != nil {
		return x.RelicUniqueId
	}
	return 0
}

func (x *ExpUpRelicCsReq) GetItemCostList() *ItemCostList {
	if x != nil {
		return x.ItemCostList
	}
	return nil
}

var File_ExpUpRelicCsReq_proto protoreflect.FileDescriptor

var file_ExpUpRelicCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x78, 0x70, 0x55, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0f, 0x45,
	0x78, 0x70, 0x55, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x26,
	0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x55, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x69,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExpUpRelicCsReq_proto_rawDescOnce sync.Once
	file_ExpUpRelicCsReq_proto_rawDescData = file_ExpUpRelicCsReq_proto_rawDesc
)

func file_ExpUpRelicCsReq_proto_rawDescGZIP() []byte {
	file_ExpUpRelicCsReq_proto_rawDescOnce.Do(func() {
		file_ExpUpRelicCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExpUpRelicCsReq_proto_rawDescData)
	})
	return file_ExpUpRelicCsReq_proto_rawDescData
}

var file_ExpUpRelicCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExpUpRelicCsReq_proto_goTypes = []interface{}{
	(*ExpUpRelicCsReq)(nil), // 0: ExpUpRelicCsReq
	(*ItemCostList)(nil),    // 1: ItemCostList
}
var file_ExpUpRelicCsReq_proto_depIdxs = []int32{
	1, // 0: ExpUpRelicCsReq.item_cost_list:type_name -> ItemCostList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ExpUpRelicCsReq_proto_init() }
func file_ExpUpRelicCsReq_proto_init() {
	if File_ExpUpRelicCsReq_proto != nil {
		return
	}
	file_ItemCostList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExpUpRelicCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpUpRelicCsReq); i {
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
			RawDescriptor: file_ExpUpRelicCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExpUpRelicCsReq_proto_goTypes,
		DependencyIndexes: file_ExpUpRelicCsReq_proto_depIdxs,
		MessageInfos:      file_ExpUpRelicCsReq_proto_msgTypes,
	}.Build()
	File_ExpUpRelicCsReq_proto = out.File
	file_ExpUpRelicCsReq_proto_rawDesc = nil
	file_ExpUpRelicCsReq_proto_goTypes = nil
	file_ExpUpRelicCsReq_proto_depIdxs = nil
}
