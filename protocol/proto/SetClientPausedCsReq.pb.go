// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetClientPausedCsReq.proto

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

type SetClientPausedCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paused bool `protobuf:"varint,5,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *SetClientPausedCsReq) Reset() {
	*x = SetClientPausedCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetClientPausedCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetClientPausedCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetClientPausedCsReq) ProtoMessage() {}

func (x *SetClientPausedCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetClientPausedCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetClientPausedCsReq.ProtoReflect.Descriptor instead.
func (*SetClientPausedCsReq) Descriptor() ([]byte, []int) {
	return file_SetClientPausedCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetClientPausedCsReq) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

var File_SetClientPausedCsReq_proto protoreflect.FileDescriptor

var file_SetClientPausedCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65,
	0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x14,
	0x53, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetClientPausedCsReq_proto_rawDescOnce sync.Once
	file_SetClientPausedCsReq_proto_rawDescData = file_SetClientPausedCsReq_proto_rawDesc
)

func file_SetClientPausedCsReq_proto_rawDescGZIP() []byte {
	file_SetClientPausedCsReq_proto_rawDescOnce.Do(func() {
		file_SetClientPausedCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetClientPausedCsReq_proto_rawDescData)
	})
	return file_SetClientPausedCsReq_proto_rawDescData
}

var file_SetClientPausedCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetClientPausedCsReq_proto_goTypes = []interface{}{
	(*SetClientPausedCsReq)(nil), // 0: SetClientPausedCsReq
}
var file_SetClientPausedCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetClientPausedCsReq_proto_init() }
func file_SetClientPausedCsReq_proto_init() {
	if File_SetClientPausedCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetClientPausedCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetClientPausedCsReq); i {
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
			RawDescriptor: file_SetClientPausedCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetClientPausedCsReq_proto_goTypes,
		DependencyIndexes: file_SetClientPausedCsReq_proto_depIdxs,
		MessageInfos:      file_SetClientPausedCsReq_proto_msgTypes,
	}.Build()
	File_SetClientPausedCsReq_proto = out.File
	file_SetClientPausedCsReq_proto_rawDesc = nil
	file_SetClientPausedCsReq_proto_goTypes = nil
	file_SetClientPausedCsReq_proto_depIdxs = nil
}
