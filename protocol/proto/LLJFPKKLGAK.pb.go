// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LLJFPKKLGAK.proto

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

type LLJFPKKLGAK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OLMCGHBDKIE *LANBPBBAJLM                `protobuf:"bytes,1,opt,name=OLMCGHBDKIE,proto3" json:"OLMCGHBDKIE,omitempty"`
	PJALJFEJGDD *RogueSyncContextBoardEvent `protobuf:"bytes,9,opt,name=PJALJFEJGDD,proto3" json:"PJALJFEJGDD,omitempty"`
}

func (x *LLJFPKKLGAK) Reset() {
	*x = LLJFPKKLGAK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLJFPKKLGAK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LLJFPKKLGAK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LLJFPKKLGAK) ProtoMessage() {}

func (x *LLJFPKKLGAK) ProtoReflect() protoreflect.Message {
	mi := &file_LLJFPKKLGAK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LLJFPKKLGAK.ProtoReflect.Descriptor instead.
func (*LLJFPKKLGAK) Descriptor() ([]byte, []int) {
	return file_LLJFPKKLGAK_proto_rawDescGZIP(), []int{0}
}

func (x *LLJFPKKLGAK) GetOLMCGHBDKIE() *LANBPBBAJLM {
	if x != nil {
		return x.OLMCGHBDKIE
	}
	return nil
}

func (x *LLJFPKKLGAK) GetPJALJFEJGDD() *RogueSyncContextBoardEvent {
	if x != nil {
		return x.PJALJFEJGDD
	}
	return nil
}

var File_LLJFPKKLGAK_proto protoreflect.FileDescriptor

var file_LLJFPKKLGAK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x4c, 0x4a, 0x46, 0x50, 0x4b, 0x4b, 0x4c, 0x47, 0x41, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x41, 0x4e, 0x42, 0x50, 0x42, 0x42, 0x41, 0x4a, 0x4c, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x79, 0x6e,
	0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0b, 0x4c, 0x4c, 0x4a, 0x46,
	0x50, 0x4b, 0x4b, 0x4c, 0x47, 0x41, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4c, 0x4d, 0x43, 0x47,
	0x48, 0x42, 0x44, 0x4b, 0x49, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c,
	0x41, 0x4e, 0x42, 0x50, 0x42, 0x42, 0x41, 0x4a, 0x4c, 0x4d, 0x52, 0x0b, 0x4f, 0x4c, 0x4d, 0x43,
	0x47, 0x48, 0x42, 0x44, 0x4b, 0x49, 0x45, 0x12, 0x3d, 0x0a, 0x0b, 0x50, 0x4a, 0x41, 0x4c, 0x4a,
	0x46, 0x45, 0x4a, 0x47, 0x44, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x50, 0x4a, 0x41, 0x4c, 0x4a,
	0x46, 0x45, 0x4a, 0x47, 0x44, 0x44, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LLJFPKKLGAK_proto_rawDescOnce sync.Once
	file_LLJFPKKLGAK_proto_rawDescData = file_LLJFPKKLGAK_proto_rawDesc
)

func file_LLJFPKKLGAK_proto_rawDescGZIP() []byte {
	file_LLJFPKKLGAK_proto_rawDescOnce.Do(func() {
		file_LLJFPKKLGAK_proto_rawDescData = protoimpl.X.CompressGZIP(file_LLJFPKKLGAK_proto_rawDescData)
	})
	return file_LLJFPKKLGAK_proto_rawDescData
}

var file_LLJFPKKLGAK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LLJFPKKLGAK_proto_goTypes = []interface{}{
	(*LLJFPKKLGAK)(nil),                // 0: LLJFPKKLGAK
	(*LANBPBBAJLM)(nil),                // 1: LANBPBBAJLM
	(*RogueSyncContextBoardEvent)(nil), // 2: RogueSyncContextBoardEvent
}
var file_LLJFPKKLGAK_proto_depIdxs = []int32{
	1, // 0: LLJFPKKLGAK.OLMCGHBDKIE:type_name -> LANBPBBAJLM
	2, // 1: LLJFPKKLGAK.PJALJFEJGDD:type_name -> RogueSyncContextBoardEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LLJFPKKLGAK_proto_init() }
func file_LLJFPKKLGAK_proto_init() {
	if File_LLJFPKKLGAK_proto != nil {
		return
	}
	file_LANBPBBAJLM_proto_init()
	file_RogueSyncContextBoardEvent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LLJFPKKLGAK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LLJFPKKLGAK); i {
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
			RawDescriptor: file_LLJFPKKLGAK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LLJFPKKLGAK_proto_goTypes,
		DependencyIndexes: file_LLJFPKKLGAK_proto_depIdxs,
		MessageInfos:      file_LLJFPKKLGAK_proto_msgTypes,
	}.Build()
	File_LLJFPKKLGAK_proto = out.File
	file_LLJFPKKLGAK_proto_rawDesc = nil
	file_LLJFPKKLGAK_proto_goTypes = nil
	file_LLJFPKKLGAK_proto_depIdxs = nil
}
