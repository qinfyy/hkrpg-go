// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPunkLordDataScRsp.proto

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

type GetPunkLordDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PAECAGPDIFL uint32   `protobuf:"varint,4,opt,name=PAECAGPDIFL,proto3" json:"PAECAGPDIFL,omitempty"`
	EGEENCOKLKJ uint32   `protobuf:"varint,7,opt,name=EGEENCOKLKJ,proto3" json:"EGEENCOKLKJ,omitempty"`
	KICGEHINHHM uint32   `protobuf:"varint,3,opt,name=KICGEHINHHM,proto3" json:"KICGEHINHHM,omitempty"`
	CBFKBIGOONK int64    `protobuf:"varint,12,opt,name=CBFKBIGOONK,proto3" json:"CBFKBIGOONK,omitempty"`
	Retcode     uint32   `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	HLDOMJLOMIA []uint32 `protobuf:"varint,1,rep,packed,name=HLDOMJLOMIA,proto3" json:"HLDOMJLOMIA,omitempty"`
	DKMNNHGNMFD uint32   `protobuf:"varint,10,opt,name=DKMNNHGNMFD,proto3" json:"DKMNNHGNMFD,omitempty"`
	OJMIJDHJCLE uint32   `protobuf:"varint,13,opt,name=OJMIJDHJCLE,proto3" json:"OJMIJDHJCLE,omitempty"`
}

func (x *GetPunkLordDataScRsp) Reset() {
	*x = GetPunkLordDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPunkLordDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPunkLordDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPunkLordDataScRsp) ProtoMessage() {}

func (x *GetPunkLordDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPunkLordDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPunkLordDataScRsp.ProtoReflect.Descriptor instead.
func (*GetPunkLordDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetPunkLordDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPunkLordDataScRsp) GetPAECAGPDIFL() uint32 {
	if x != nil {
		return x.PAECAGPDIFL
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetEGEENCOKLKJ() uint32 {
	if x != nil {
		return x.EGEENCOKLKJ
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetKICGEHINHHM() uint32 {
	if x != nil {
		return x.KICGEHINHHM
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetCBFKBIGOONK() int64 {
	if x != nil {
		return x.CBFKBIGOONK
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetHLDOMJLOMIA() []uint32 {
	if x != nil {
		return x.HLDOMJLOMIA
	}
	return nil
}

func (x *GetPunkLordDataScRsp) GetDKMNNHGNMFD() uint32 {
	if x != nil {
		return x.DKMNNHGNMFD
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetOJMIJDHJCLE() uint32 {
	if x != nil {
		return x.OJMIJDHJCLE
	}
	return 0
}

var File_GetPunkLordDataScRsp_proto protoreflect.FileDescriptor

var file_GetPunkLordDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x41, 0x45, 0x43, 0x41, 0x47, 0x50,
	0x44, 0x49, 0x46, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x41, 0x45, 0x43,
	0x41, 0x47, 0x50, 0x44, 0x49, 0x46, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x47, 0x45, 0x45, 0x4e,
	0x43, 0x4f, 0x4b, 0x4c, 0x4b, 0x4a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x47,
	0x45, 0x45, 0x4e, 0x43, 0x4f, 0x4b, 0x4c, 0x4b, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x43,
	0x47, 0x45, 0x48, 0x49, 0x4e, 0x48, 0x48, 0x4d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4b, 0x49, 0x43, 0x47, 0x45, 0x48, 0x49, 0x4e, 0x48, 0x48, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x42, 0x46, 0x4b, 0x42, 0x49, 0x47, 0x4f, 0x4f, 0x4e, 0x4b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x43, 0x42, 0x46, 0x4b, 0x42, 0x49, 0x47, 0x4f, 0x4f, 0x4e, 0x4b, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4c, 0x44, 0x4f, 0x4d,
	0x4a, 0x4c, 0x4f, 0x4d, 0x49, 0x41, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4c,
	0x44, 0x4f, 0x4d, 0x4a, 0x4c, 0x4f, 0x4d, 0x49, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4b, 0x4d,
	0x4e, 0x4e, 0x48, 0x47, 0x4e, 0x4d, 0x46, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x4b, 0x4d, 0x4e, 0x4e, 0x48, 0x47, 0x4e, 0x4d, 0x46, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x4a, 0x4d, 0x49, 0x4a, 0x44, 0x48, 0x4a, 0x43, 0x4c, 0x45, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4f, 0x4a, 0x4d, 0x49, 0x4a, 0x44, 0x48, 0x4a, 0x43, 0x4c, 0x45, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPunkLordDataScRsp_proto_rawDescOnce sync.Once
	file_GetPunkLordDataScRsp_proto_rawDescData = file_GetPunkLordDataScRsp_proto_rawDesc
)

func file_GetPunkLordDataScRsp_proto_rawDescGZIP() []byte {
	file_GetPunkLordDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetPunkLordDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPunkLordDataScRsp_proto_rawDescData)
	})
	return file_GetPunkLordDataScRsp_proto_rawDescData
}

var file_GetPunkLordDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPunkLordDataScRsp_proto_goTypes = []interface{}{
	(*GetPunkLordDataScRsp)(nil), // 0: GetPunkLordDataScRsp
}
var file_GetPunkLordDataScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPunkLordDataScRsp_proto_init() }
func file_GetPunkLordDataScRsp_proto_init() {
	if File_GetPunkLordDataScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPunkLordDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPunkLordDataScRsp); i {
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
			RawDescriptor: file_GetPunkLordDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPunkLordDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetPunkLordDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetPunkLordDataScRsp_proto_msgTypes,
	}.Build()
	File_GetPunkLordDataScRsp_proto = out.File
	file_GetPunkLordDataScRsp_proto_rawDesc = nil
	file_GetPunkLordDataScRsp_proto_goTypes = nil
	file_GetPunkLordDataScRsp_proto_depIdxs = nil
}