// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordMonsterKilledNotify.proto

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

type PunkLordMonsterKilledNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ICJAMBFGJFJ *KilledPunkLordMonsterInfo `protobuf:"bytes,1,opt,name=ICJAMBFGJFJ,proto3" json:"ICJAMBFGJFJ,omitempty"`
}

func (x *PunkLordMonsterKilledNotify) Reset() {
	*x = PunkLordMonsterKilledNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PunkLordMonsterKilledNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunkLordMonsterKilledNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunkLordMonsterKilledNotify) ProtoMessage() {}

func (x *PunkLordMonsterKilledNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PunkLordMonsterKilledNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunkLordMonsterKilledNotify.ProtoReflect.Descriptor instead.
func (*PunkLordMonsterKilledNotify) Descriptor() ([]byte, []int) {
	return file_PunkLordMonsterKilledNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PunkLordMonsterKilledNotify) GetICJAMBFGJFJ() *KilledPunkLordMonsterInfo {
	if x != nil {
		return x.ICJAMBFGJFJ
	}
	return nil
}

var File_PunkLordMonsterKilledNotify_proto protoreflect.FileDescriptor

var file_PunkLordMonsterKilledNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x75, 0x6e, 0x6b, 0x4c,
	0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x1b, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x49, 0x43, 0x4a, 0x41, 0x4d, 0x42, 0x46, 0x47, 0x4a,
	0x46, 0x4a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x49, 0x43, 0x4a, 0x41, 0x4d, 0x42, 0x46, 0x47, 0x4a, 0x46,
	0x4a, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PunkLordMonsterKilledNotify_proto_rawDescOnce sync.Once
	file_PunkLordMonsterKilledNotify_proto_rawDescData = file_PunkLordMonsterKilledNotify_proto_rawDesc
)

func file_PunkLordMonsterKilledNotify_proto_rawDescGZIP() []byte {
	file_PunkLordMonsterKilledNotify_proto_rawDescOnce.Do(func() {
		file_PunkLordMonsterKilledNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordMonsterKilledNotify_proto_rawDescData)
	})
	return file_PunkLordMonsterKilledNotify_proto_rawDescData
}

var file_PunkLordMonsterKilledNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PunkLordMonsterKilledNotify_proto_goTypes = []interface{}{
	(*PunkLordMonsterKilledNotify)(nil), // 0: PunkLordMonsterKilledNotify
	(*KilledPunkLordMonsterInfo)(nil),   // 1: KilledPunkLordMonsterInfo
}
var file_PunkLordMonsterKilledNotify_proto_depIdxs = []int32{
	1, // 0: PunkLordMonsterKilledNotify.ICJAMBFGJFJ:type_name -> KilledPunkLordMonsterInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PunkLordMonsterKilledNotify_proto_init() }
func file_PunkLordMonsterKilledNotify_proto_init() {
	if File_PunkLordMonsterKilledNotify_proto != nil {
		return
	}
	file_KilledPunkLordMonsterInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PunkLordMonsterKilledNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunkLordMonsterKilledNotify); i {
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
			RawDescriptor: file_PunkLordMonsterKilledNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordMonsterKilledNotify_proto_goTypes,
		DependencyIndexes: file_PunkLordMonsterKilledNotify_proto_depIdxs,
		MessageInfos:      file_PunkLordMonsterKilledNotify_proto_msgTypes,
	}.Build()
	File_PunkLordMonsterKilledNotify_proto = out.File
	file_PunkLordMonsterKilledNotify_proto_rawDesc = nil
	file_PunkLordMonsterKilledNotify_proto_goTypes = nil
	file_PunkLordMonsterKilledNotify_proto_depIdxs = nil
}
