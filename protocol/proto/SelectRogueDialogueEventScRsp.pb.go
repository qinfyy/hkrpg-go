// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SelectRogueDialogueEventScRsp.proto

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

type SelectRogueDialogueEventScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogueEventId uint32              `protobuf:"varint,4,opt,name=dialogue_event_id,json=dialogueEventId,proto3" json:"dialogue_event_id,omitempty"`
	DialogueResult  *DialogueResult     `protobuf:"bytes,8,opt,name=dialogue_result,json=dialogueResult,proto3" json:"dialogue_result,omitempty"`
	EventData       *RogueDialogueEvent `protobuf:"bytes,10,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	Retcode         uint32              `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SelectRogueDialogueEventScRsp) Reset() {
	*x = SelectRogueDialogueEventScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SelectRogueDialogueEventScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectRogueDialogueEventScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectRogueDialogueEventScRsp) ProtoMessage() {}

func (x *SelectRogueDialogueEventScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SelectRogueDialogueEventScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectRogueDialogueEventScRsp.ProtoReflect.Descriptor instead.
func (*SelectRogueDialogueEventScRsp) Descriptor() ([]byte, []int) {
	return file_SelectRogueDialogueEventScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SelectRogueDialogueEventScRsp) GetDialogueEventId() uint32 {
	if x != nil {
		return x.DialogueEventId
	}
	return 0
}

func (x *SelectRogueDialogueEventScRsp) GetDialogueResult() *DialogueResult {
	if x != nil {
		return x.DialogueResult
	}
	return nil
}

func (x *SelectRogueDialogueEventScRsp) GetEventData() *RogueDialogueEvent {
	if x != nil {
		return x.EventData
	}
	return nil
}

func (x *SelectRogueDialogueEventScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SelectRogueDialogueEventScRsp_proto protoreflect.FileDescriptor

var file_SelectRogueDialogueEventScRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x1d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SelectRogueDialogueEventScRsp_proto_rawDescOnce sync.Once
	file_SelectRogueDialogueEventScRsp_proto_rawDescData = file_SelectRogueDialogueEventScRsp_proto_rawDesc
)

func file_SelectRogueDialogueEventScRsp_proto_rawDescGZIP() []byte {
	file_SelectRogueDialogueEventScRsp_proto_rawDescOnce.Do(func() {
		file_SelectRogueDialogueEventScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SelectRogueDialogueEventScRsp_proto_rawDescData)
	})
	return file_SelectRogueDialogueEventScRsp_proto_rawDescData
}

var file_SelectRogueDialogueEventScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SelectRogueDialogueEventScRsp_proto_goTypes = []interface{}{
	(*SelectRogueDialogueEventScRsp)(nil), // 0: SelectRogueDialogueEventScRsp
	(*DialogueResult)(nil),                // 1: DialogueResult
	(*RogueDialogueEvent)(nil),            // 2: RogueDialogueEvent
}
var file_SelectRogueDialogueEventScRsp_proto_depIdxs = []int32{
	1, // 0: SelectRogueDialogueEventScRsp.dialogue_result:type_name -> DialogueResult
	2, // 1: SelectRogueDialogueEventScRsp.event_data:type_name -> RogueDialogueEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SelectRogueDialogueEventScRsp_proto_init() }
func file_SelectRogueDialogueEventScRsp_proto_init() {
	if File_SelectRogueDialogueEventScRsp_proto != nil {
		return
	}
	file_RogueDialogueEvent_proto_init()
	file_DialogueResult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SelectRogueDialogueEventScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectRogueDialogueEventScRsp); i {
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
			RawDescriptor: file_SelectRogueDialogueEventScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SelectRogueDialogueEventScRsp_proto_goTypes,
		DependencyIndexes: file_SelectRogueDialogueEventScRsp_proto_depIdxs,
		MessageInfos:      file_SelectRogueDialogueEventScRsp_proto_msgTypes,
	}.Build()
	File_SelectRogueDialogueEventScRsp_proto = out.File
	file_SelectRogueDialogueEventScRsp_proto_rawDesc = nil
	file_SelectRogueDialogueEventScRsp_proto_goTypes = nil
	file_SelectRogueDialogueEventScRsp_proto_depIdxs = nil
}
