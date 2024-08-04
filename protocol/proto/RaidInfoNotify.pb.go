// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RaidInfoNotify.proto

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

type RaidInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RaidId         uint32            `protobuf:"varint,4,opt,name=raid_id,json=raidId,proto3" json:"raid_id,omitempty"`
	Status         RaidStatus        `protobuf:"varint,2,opt,name=status,proto3,enum=RaidStatus" json:"status,omitempty"`
	WorldLevel     uint32            `protobuf:"varint,1,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	RaidTargetInfo []*RaidTargetInfo `protobuf:"bytes,10,rep,name=raid_target_info,json=raidTargetInfo,proto3" json:"raid_target_info,omitempty"`
	ItemList       *ItemList         `protobuf:"bytes,11,opt,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	RaidFinishTime uint64            `protobuf:"varint,5,opt,name=raid_finish_time,json=raidFinishTime,proto3" json:"raid_finish_time,omitempty"`
}

func (x *RaidInfoNotify) Reset() {
	*x = RaidInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RaidInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaidInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaidInfoNotify) ProtoMessage() {}

func (x *RaidInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RaidInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaidInfoNotify.ProtoReflect.Descriptor instead.
func (*RaidInfoNotify) Descriptor() ([]byte, []int) {
	return file_RaidInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RaidInfoNotify) GetRaidId() uint32 {
	if x != nil {
		return x.RaidId
	}
	return 0
}

func (x *RaidInfoNotify) GetStatus() RaidStatus {
	if x != nil {
		return x.Status
	}
	return RaidStatus_RAID_STATUS_NONE
}

func (x *RaidInfoNotify) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *RaidInfoNotify) GetRaidTargetInfo() []*RaidTargetInfo {
	if x != nil {
		return x.RaidTargetInfo
	}
	return nil
}

func (x *RaidInfoNotify) GetItemList() *ItemList {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *RaidInfoNotify) GetRaidFinishTime() uint64 {
	if x != nil {
		return x.RaidFinishTime
	}
	return 0
}

var File_RaidInfoNotify_proto protoreflect.FileDescriptor

var file_RaidInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x14, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x61, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x52, 0x61, 0x69, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc,
	0x01, 0x0a, 0x0e, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x69,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x39, 0x0a, 0x10, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x61, 0x69,
	0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x72, 0x61, 0x69,
	0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72,
	0x61, 0x69, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RaidInfoNotify_proto_rawDescOnce sync.Once
	file_RaidInfoNotify_proto_rawDescData = file_RaidInfoNotify_proto_rawDesc
)

func file_RaidInfoNotify_proto_rawDescGZIP() []byte {
	file_RaidInfoNotify_proto_rawDescOnce.Do(func() {
		file_RaidInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RaidInfoNotify_proto_rawDescData)
	})
	return file_RaidInfoNotify_proto_rawDescData
}

var file_RaidInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RaidInfoNotify_proto_goTypes = []interface{}{
	(*RaidInfoNotify)(nil), // 0: RaidInfoNotify
	(RaidStatus)(0),        // 1: RaidStatus
	(*RaidTargetInfo)(nil), // 2: RaidTargetInfo
	(*ItemList)(nil),       // 3: ItemList
}
var file_RaidInfoNotify_proto_depIdxs = []int32{
	1, // 0: RaidInfoNotify.status:type_name -> RaidStatus
	2, // 1: RaidInfoNotify.raid_target_info:type_name -> RaidTargetInfo
	3, // 2: RaidInfoNotify.item_list:type_name -> ItemList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RaidInfoNotify_proto_init() }
func file_RaidInfoNotify_proto_init() {
	if File_RaidInfoNotify_proto != nil {
		return
	}
	file_ItemList_proto_init()
	file_RaidStatus_proto_init()
	file_RaidTargetInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RaidInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaidInfoNotify); i {
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
			RawDescriptor: file_RaidInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RaidInfoNotify_proto_goTypes,
		DependencyIndexes: file_RaidInfoNotify_proto_depIdxs,
		MessageInfos:      file_RaidInfoNotify_proto_msgTypes,
	}.Build()
	File_RaidInfoNotify_proto = out.File
	file_RaidInfoNotify_proto_rawDesc = nil
	file_RaidInfoNotify_proto_goTypes = nil
	file_RaidInfoNotify_proto_depIdxs = nil
}
