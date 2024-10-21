// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueAdventureRoomGameplayWolfGunTarget.proto

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

type RogueAdventureRoomGameplayWolfGunTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetNone    *MHJEHLNICFP `protobuf:"bytes,10,opt,name=target_none,json=targetNone,proto3" json:"target_none,omitempty"`
	TargetCoin    *FKPILFBKDLA `protobuf:"bytes,8,opt,name=target_coin,json=targetCoin,proto3" json:"target_coin,omitempty"`
	TargetMiracle *LNKGGAPBPLF `protobuf:"bytes,4,opt,name=target_miracle,json=targetMiracle,proto3" json:"target_miracle,omitempty"`
	TargetRuanmei *JJOBHOEDLHO `protobuf:"bytes,1,opt,name=target_ruanmei,json=targetRuanmei,proto3" json:"target_ruanmei,omitempty"`
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) Reset() {
	*x = RogueAdventureRoomGameplayWolfGunTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueAdventureRoomGameplayWolfGunTarget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueAdventureRoomGameplayWolfGunTarget) ProtoMessage() {}

func (x *RogueAdventureRoomGameplayWolfGunTarget) ProtoReflect() protoreflect.Message {
	mi := &file_RogueAdventureRoomGameplayWolfGunTarget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueAdventureRoomGameplayWolfGunTarget.ProtoReflect.Descriptor instead.
func (*RogueAdventureRoomGameplayWolfGunTarget) Descriptor() ([]byte, []int) {
	return file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescGZIP(), []int{0}
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) GetTargetNone() *MHJEHLNICFP {
	if x != nil {
		return x.TargetNone
	}
	return nil
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) GetTargetCoin() *FKPILFBKDLA {
	if x != nil {
		return x.TargetCoin
	}
	return nil
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) GetTargetMiracle() *LNKGGAPBPLF {
	if x != nil {
		return x.TargetMiracle
	}
	return nil
}

func (x *RogueAdventureRoomGameplayWolfGunTarget) GetTargetRuanmei() *JJOBHOEDLHO {
	if x != nil {
		return x.TargetRuanmei
	}
	return nil
}

var File_RogueAdventureRoomGameplayWolfGunTarget_proto protoreflect.FileDescriptor

var file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x57, 0x6f, 0x6c, 0x66,
	0x47, 0x75, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4d, 0x48, 0x4a, 0x45, 0x48, 0x4c, 0x4e, 0x49, 0x43, 0x46, 0x50, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4e, 0x4b, 0x47, 0x47, 0x41, 0x50, 0x42, 0x50, 0x4c, 0x46, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4b, 0x50, 0x49, 0x4c, 0x46, 0x42, 0x4b, 0x44,
	0x4c, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4a, 0x4f, 0x42, 0x48, 0x4f,
	0x45, 0x44, 0x4c, 0x48, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x27,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x57, 0x6f, 0x6c, 0x66, 0x47, 0x75,
	0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d,
	0x48, 0x4a, 0x45, 0x48, 0x4c, 0x4e, 0x49, 0x43, 0x46, 0x50, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4b,
	0x50, 0x49, 0x4c, 0x46, 0x42, 0x4b, 0x44, 0x4c, 0x41, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x33, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4c, 0x4e, 0x4b, 0x47, 0x47, 0x41, 0x50, 0x42, 0x50, 0x4c, 0x46, 0x52, 0x0d, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x75, 0x61, 0x6e, 0x6d, 0x65, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4a, 0x4f, 0x42, 0x48, 0x4f, 0x45, 0x44, 0x4c, 0x48, 0x4f,
	0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x75, 0x61, 0x6e, 0x6d, 0x65, 0x69, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescOnce sync.Once
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescData = file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDesc
)

func file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescGZIP() []byte {
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescOnce.Do(func() {
		file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescData)
	})
	return file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDescData
}

var file_RogueAdventureRoomGameplayWolfGunTarget_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueAdventureRoomGameplayWolfGunTarget_proto_goTypes = []interface{}{
	(*RogueAdventureRoomGameplayWolfGunTarget)(nil), // 0: RogueAdventureRoomGameplayWolfGunTarget
	(*MHJEHLNICFP)(nil),                             // 1: MHJEHLNICFP
	(*FKPILFBKDLA)(nil),                             // 2: FKPILFBKDLA
	(*LNKGGAPBPLF)(nil),                             // 3: LNKGGAPBPLF
	(*JJOBHOEDLHO)(nil),                             // 4: JJOBHOEDLHO
}
var file_RogueAdventureRoomGameplayWolfGunTarget_proto_depIdxs = []int32{
	1, // 0: RogueAdventureRoomGameplayWolfGunTarget.target_none:type_name -> MHJEHLNICFP
	2, // 1: RogueAdventureRoomGameplayWolfGunTarget.target_coin:type_name -> FKPILFBKDLA
	3, // 2: RogueAdventureRoomGameplayWolfGunTarget.target_miracle:type_name -> LNKGGAPBPLF
	4, // 3: RogueAdventureRoomGameplayWolfGunTarget.target_ruanmei:type_name -> JJOBHOEDLHO
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueAdventureRoomGameplayWolfGunTarget_proto_init() }
func file_RogueAdventureRoomGameplayWolfGunTarget_proto_init() {
	if File_RogueAdventureRoomGameplayWolfGunTarget_proto != nil {
		return
	}
	file_MHJEHLNICFP_proto_init()
	file_LNKGGAPBPLF_proto_init()
	file_FKPILFBKDLA_proto_init()
	file_JJOBHOEDLHO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueAdventureRoomGameplayWolfGunTarget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueAdventureRoomGameplayWolfGunTarget); i {
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
			RawDescriptor: file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueAdventureRoomGameplayWolfGunTarget_proto_goTypes,
		DependencyIndexes: file_RogueAdventureRoomGameplayWolfGunTarget_proto_depIdxs,
		MessageInfos:      file_RogueAdventureRoomGameplayWolfGunTarget_proto_msgTypes,
	}.Build()
	File_RogueAdventureRoomGameplayWolfGunTarget_proto = out.File
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_rawDesc = nil
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_goTypes = nil
	file_RogueAdventureRoomGameplayWolfGunTarget_proto_depIdxs = nil
}
