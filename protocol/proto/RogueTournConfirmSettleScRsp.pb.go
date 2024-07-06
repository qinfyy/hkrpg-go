// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournConfirmSettleScRsp.proto

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

type RogueTournConfirmSettleScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PLIEBDLJLBD            *ItemList               `protobuf:"bytes,3,opt,name=PLIEBDLJLBD,proto3" json:"PLIEBDLJLBD,omitempty"`
	ALFKCICLLMG            *ItemList               `protobuf:"bytes,13,opt,name=ALFKCICLLMG,proto3" json:"ALFKCICLLMG,omitempty"`
	DHKCMMIDLEB            *RogueTournSaveList     `protobuf:"bytes,8,opt,name=DHKCMMIDLEB,proto3" json:"DHKCMMIDLEB,omitempty"`
	RogueTournCurSceneInfo *RogueTournCurSceneInfo `protobuf:"bytes,7,opt,name=rogue_tourn_cur_scene_info,json=rogueTournCurSceneInfo,proto3" json:"rogue_tourn_cur_scene_info,omitempty"`
	Retcode                uint32                  `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *RogueTournConfirmSettleScRsp) Reset() {
	*x = RogueTournConfirmSettleScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournConfirmSettleScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournConfirmSettleScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournConfirmSettleScRsp) ProtoMessage() {}

func (x *RogueTournConfirmSettleScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournConfirmSettleScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournConfirmSettleScRsp.ProtoReflect.Descriptor instead.
func (*RogueTournConfirmSettleScRsp) Descriptor() ([]byte, []int) {
	return file_RogueTournConfirmSettleScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournConfirmSettleScRsp) GetPLIEBDLJLBD() *ItemList {
	if x != nil {
		return x.PLIEBDLJLBD
	}
	return nil
}

func (x *RogueTournConfirmSettleScRsp) GetALFKCICLLMG() *ItemList {
	if x != nil {
		return x.ALFKCICLLMG
	}
	return nil
}

func (x *RogueTournConfirmSettleScRsp) GetDHKCMMIDLEB() *RogueTournSaveList {
	if x != nil {
		return x.DHKCMMIDLEB
	}
	return nil
}

func (x *RogueTournConfirmSettleScRsp) GetRogueTournCurSceneInfo() *RogueTournCurSceneInfo {
	if x != nil {
		return x.RogueTournCurSceneInfo
	}
	return nil
}

func (x *RogueTournConfirmSettleScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_RogueTournConfirmSettleScRsp_proto protoreflect.FileDescriptor

var file_RogueTournConfirmSettleScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x43, 0x75, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x61,
	0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a,
	0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2b, 0x0a,
	0x0b, 0x50, 0x4c, 0x49, 0x45, 0x42, 0x44, 0x4c, 0x4a, 0x4c, 0x42, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x50,
	0x4c, 0x49, 0x45, 0x42, 0x44, 0x4c, 0x4a, 0x4c, 0x42, 0x44, 0x12, 0x2b, 0x0a, 0x0b, 0x41, 0x4c,
	0x46, 0x4b, 0x43, 0x49, 0x43, 0x4c, 0x4c, 0x4d, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x41, 0x4c, 0x46, 0x4b,
	0x43, 0x49, 0x43, 0x4c, 0x4c, 0x4d, 0x47, 0x12, 0x35, 0x0a, 0x0b, 0x44, 0x48, 0x4b, 0x43, 0x4d,
	0x4d, 0x49, 0x44, 0x4c, 0x45, 0x42, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x0b, 0x44, 0x48, 0x4b, 0x43, 0x4d, 0x4d, 0x49, 0x44, 0x4c, 0x45, 0x42, 0x12, 0x53,
	0x0a, 0x1a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x75,
	0x72, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43,
	0x75, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x16, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournConfirmSettleScRsp_proto_rawDescOnce sync.Once
	file_RogueTournConfirmSettleScRsp_proto_rawDescData = file_RogueTournConfirmSettleScRsp_proto_rawDesc
)

func file_RogueTournConfirmSettleScRsp_proto_rawDescGZIP() []byte {
	file_RogueTournConfirmSettleScRsp_proto_rawDescOnce.Do(func() {
		file_RogueTournConfirmSettleScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournConfirmSettleScRsp_proto_rawDescData)
	})
	return file_RogueTournConfirmSettleScRsp_proto_rawDescData
}

var file_RogueTournConfirmSettleScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournConfirmSettleScRsp_proto_goTypes = []interface{}{
	(*RogueTournConfirmSettleScRsp)(nil), // 0: RogueTournConfirmSettleScRsp
	(*ItemList)(nil),                     // 1: ItemList
	(*RogueTournSaveList)(nil),           // 2: RogueTournSaveList
	(*RogueTournCurSceneInfo)(nil),       // 3: RogueTournCurSceneInfo
}
var file_RogueTournConfirmSettleScRsp_proto_depIdxs = []int32{
	1, // 0: RogueTournConfirmSettleScRsp.PLIEBDLJLBD:type_name -> ItemList
	1, // 1: RogueTournConfirmSettleScRsp.ALFKCICLLMG:type_name -> ItemList
	2, // 2: RogueTournConfirmSettleScRsp.DHKCMMIDLEB:type_name -> RogueTournSaveList
	3, // 3: RogueTournConfirmSettleScRsp.rogue_tourn_cur_scene_info:type_name -> RogueTournCurSceneInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueTournConfirmSettleScRsp_proto_init() }
func file_RogueTournConfirmSettleScRsp_proto_init() {
	if File_RogueTournConfirmSettleScRsp_proto != nil {
		return
	}
	file_RogueTournCurSceneInfo_proto_init()
	file_RogueTournSaveList_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournConfirmSettleScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournConfirmSettleScRsp); i {
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
			RawDescriptor: file_RogueTournConfirmSettleScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournConfirmSettleScRsp_proto_goTypes,
		DependencyIndexes: file_RogueTournConfirmSettleScRsp_proto_depIdxs,
		MessageInfos:      file_RogueTournConfirmSettleScRsp_proto_msgTypes,
	}.Build()
	File_RogueTournConfirmSettleScRsp_proto = out.File
	file_RogueTournConfirmSettleScRsp_proto_rawDesc = nil
	file_RogueTournConfirmSettleScRsp_proto_goTypes = nil
	file_RogueTournConfirmSettleScRsp_proto_depIdxs = nil
}