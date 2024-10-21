// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneEnterStageCsReq.proto

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

type SceneEnterStageCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NNNCABEMAEB  bool         `protobuf:"varint,3,opt,name=NNNCABEMAEB,proto3" json:"NNNCABEMAEB,omitempty"`
	EventId      uint32       `protobuf:"varint,14,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	RebattleType RebattleType `protobuf:"varint,2,opt,name=rebattle_type,json=rebattleType,proto3,enum=RebattleType" json:"rebattle_type,omitempty"`
}

func (x *SceneEnterStageCsReq) Reset() {
	*x = SceneEnterStageCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEnterStageCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEnterStageCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEnterStageCsReq) ProtoMessage() {}

func (x *SceneEnterStageCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEnterStageCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEnterStageCsReq.ProtoReflect.Descriptor instead.
func (*SceneEnterStageCsReq) Descriptor() ([]byte, []int) {
	return file_SceneEnterStageCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEnterStageCsReq) GetNNNCABEMAEB() bool {
	if x != nil {
		return x.NNNCABEMAEB
	}
	return false
}

func (x *SceneEnterStageCsReq) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *SceneEnterStageCsReq) GetRebattleType() RebattleType {
	if x != nil {
		return x.RebattleType
	}
	return RebattleType_REBATTLE_TYPE_NONE
}

var File_SceneEnterStageCsReq_proto protoreflect.FileDescriptor

var file_SceneEnterStageCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x52, 0x65,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x01, 0x0a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4e, 0x4e,
	0x43, 0x41, 0x42, 0x45, 0x4d, 0x41, 0x45, 0x42, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4e, 0x4e, 0x4e, 0x43, 0x41, 0x42, 0x45, 0x4d, 0x41, 0x45, 0x42, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0d, 0x72, 0x65, 0x62, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x52, 0x65, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SceneEnterStageCsReq_proto_rawDescOnce sync.Once
	file_SceneEnterStageCsReq_proto_rawDescData = file_SceneEnterStageCsReq_proto_rawDesc
)

func file_SceneEnterStageCsReq_proto_rawDescGZIP() []byte {
	file_SceneEnterStageCsReq_proto_rawDescOnce.Do(func() {
		file_SceneEnterStageCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEnterStageCsReq_proto_rawDescData)
	})
	return file_SceneEnterStageCsReq_proto_rawDescData
}

var file_SceneEnterStageCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEnterStageCsReq_proto_goTypes = []interface{}{
	(*SceneEnterStageCsReq)(nil), // 0: SceneEnterStageCsReq
	(RebattleType)(0),            // 1: RebattleType
}
var file_SceneEnterStageCsReq_proto_depIdxs = []int32{
	1, // 0: SceneEnterStageCsReq.rebattle_type:type_name -> RebattleType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEnterStageCsReq_proto_init() }
func file_SceneEnterStageCsReq_proto_init() {
	if File_SceneEnterStageCsReq_proto != nil {
		return
	}
	file_RebattleType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEnterStageCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEnterStageCsReq); i {
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
			RawDescriptor: file_SceneEnterStageCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEnterStageCsReq_proto_goTypes,
		DependencyIndexes: file_SceneEnterStageCsReq_proto_depIdxs,
		MessageInfos:      file_SceneEnterStageCsReq_proto_msgTypes,
	}.Build()
	File_SceneEnterStageCsReq_proto = out.File
	file_SceneEnterStageCsReq_proto_rawDesc = nil
	file_SceneEnterStageCsReq_proto_goTypes = nil
	file_SceneEnterStageCsReq_proto_depIdxs = nil
}
