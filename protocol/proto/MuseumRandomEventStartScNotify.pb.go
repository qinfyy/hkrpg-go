// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MuseumRandomEventStartScNotify.proto

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

type MuseumRandomEventStartScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PABJLEADHDE `protobuf:"bytes,15,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *MuseumRandomEventStartScNotify) Reset() {
	*x = MuseumRandomEventStartScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MuseumRandomEventStartScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuseumRandomEventStartScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuseumRandomEventStartScNotify) ProtoMessage() {}

func (x *MuseumRandomEventStartScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MuseumRandomEventStartScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuseumRandomEventStartScNotify.ProtoReflect.Descriptor instead.
func (*MuseumRandomEventStartScNotify) Descriptor() ([]byte, []int) {
	return file_MuseumRandomEventStartScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MuseumRandomEventStartScNotify) GetInfo() *PABJLEADHDE {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_MuseumRandomEventStartScNotify_proto protoreflect.FileDescriptor

var file_MuseumRandomEventStartScNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x41, 0x42, 0x4a, 0x4c, 0x45, 0x41, 0x44,
	0x48, 0x44, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x1e, 0x4d, 0x75, 0x73,
	0x65, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x41, 0x42, 0x4a,
	0x4c, 0x45, 0x41, 0x44, 0x48, 0x44, 0x45, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MuseumRandomEventStartScNotify_proto_rawDescOnce sync.Once
	file_MuseumRandomEventStartScNotify_proto_rawDescData = file_MuseumRandomEventStartScNotify_proto_rawDesc
)

func file_MuseumRandomEventStartScNotify_proto_rawDescGZIP() []byte {
	file_MuseumRandomEventStartScNotify_proto_rawDescOnce.Do(func() {
		file_MuseumRandomEventStartScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MuseumRandomEventStartScNotify_proto_rawDescData)
	})
	return file_MuseumRandomEventStartScNotify_proto_rawDescData
}

var file_MuseumRandomEventStartScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MuseumRandomEventStartScNotify_proto_goTypes = []interface{}{
	(*MuseumRandomEventStartScNotify)(nil), // 0: MuseumRandomEventStartScNotify
	(*PABJLEADHDE)(nil),                    // 1: PABJLEADHDE
}
var file_MuseumRandomEventStartScNotify_proto_depIdxs = []int32{
	1, // 0: MuseumRandomEventStartScNotify.info:type_name -> PABJLEADHDE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MuseumRandomEventStartScNotify_proto_init() }
func file_MuseumRandomEventStartScNotify_proto_init() {
	if File_MuseumRandomEventStartScNotify_proto != nil {
		return
	}
	file_PABJLEADHDE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MuseumRandomEventStartScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuseumRandomEventStartScNotify); i {
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
			RawDescriptor: file_MuseumRandomEventStartScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MuseumRandomEventStartScNotify_proto_goTypes,
		DependencyIndexes: file_MuseumRandomEventStartScNotify_proto_depIdxs,
		MessageInfos:      file_MuseumRandomEventStartScNotify_proto_msgTypes,
	}.Build()
	File_MuseumRandomEventStartScNotify_proto = out.File
	file_MuseumRandomEventStartScNotify_proto_rawDesc = nil
	file_MuseumRandomEventStartScNotify_proto_goTypes = nil
	file_MuseumRandomEventStartScNotify_proto_depIdxs = nil
}
