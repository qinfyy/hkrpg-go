// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FODJFBNFPJC.proto

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

type FODJFBNFPJC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IMGILNNKGEK bool           `protobuf:"varint,9,opt,name=IMGILNNKGEK,proto3" json:"IMGILNNKGEK,omitempty"`
	GGPLKBFBPNP uint32         `protobuf:"varint,8,opt,name=GGPLKBFBPNP,proto3" json:"GGPLKBFBPNP,omitempty"`
	BuffList    []*HKMCMFCONNN `protobuf:"bytes,744,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	MKKIBDNBNFE bool           `protobuf:"varint,14,opt,name=MKKIBDNBNFE,proto3" json:"MKKIBDNBNFE,omitempty"`
	EFKNEAIPGLF bool           `protobuf:"varint,5,opt,name=EFKNEAIPGLF,proto3" json:"EFKNEAIPGLF,omitempty"`
	LMCBFDAFMPP bool           `protobuf:"varint,15,opt,name=LMCBFDAFMPP,proto3" json:"LMCBFDAFMPP,omitempty"`
	LGOOOMGJKHH uint32         `protobuf:"varint,10,opt,name=LGOOOMGJKHH,proto3" json:"LGOOOMGJKHH,omitempty"`
	HKOGCPNHNCA uint32         `protobuf:"varint,11,opt,name=HKOGCPNHNCA,proto3" json:"HKOGCPNHNCA,omitempty"`
}

func (x *FODJFBNFPJC) Reset() {
	*x = FODJFBNFPJC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FODJFBNFPJC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FODJFBNFPJC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FODJFBNFPJC) ProtoMessage() {}

func (x *FODJFBNFPJC) ProtoReflect() protoreflect.Message {
	mi := &file_FODJFBNFPJC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FODJFBNFPJC.ProtoReflect.Descriptor instead.
func (*FODJFBNFPJC) Descriptor() ([]byte, []int) {
	return file_FODJFBNFPJC_proto_rawDescGZIP(), []int{0}
}

func (x *FODJFBNFPJC) GetIMGILNNKGEK() bool {
	if x != nil {
		return x.IMGILNNKGEK
	}
	return false
}

func (x *FODJFBNFPJC) GetGGPLKBFBPNP() uint32 {
	if x != nil {
		return x.GGPLKBFBPNP
	}
	return 0
}

func (x *FODJFBNFPJC) GetBuffList() []*HKMCMFCONNN {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *FODJFBNFPJC) GetMKKIBDNBNFE() bool {
	if x != nil {
		return x.MKKIBDNBNFE
	}
	return false
}

func (x *FODJFBNFPJC) GetEFKNEAIPGLF() bool {
	if x != nil {
		return x.EFKNEAIPGLF
	}
	return false
}

func (x *FODJFBNFPJC) GetLMCBFDAFMPP() bool {
	if x != nil {
		return x.LMCBFDAFMPP
	}
	return false
}

func (x *FODJFBNFPJC) GetLGOOOMGJKHH() uint32 {
	if x != nil {
		return x.LGOOOMGJKHH
	}
	return 0
}

func (x *FODJFBNFPJC) GetHKOGCPNHNCA() uint32 {
	if x != nil {
		return x.HKOGCPNHNCA
	}
	return 0
}

var File_FODJFBNFPJC_proto protoreflect.FileDescriptor

var file_FODJFBNFPJC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x4f, 0x44, 0x4a, 0x46, 0x42, 0x4e, 0x46, 0x50, 0x4a, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4b, 0x4d, 0x43, 0x4d, 0x46, 0x43, 0x4f, 0x4e, 0x4e, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x46, 0x4f, 0x44, 0x4a, 0x46,
	0x42, 0x4e, 0x46, 0x50, 0x4a, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4d, 0x47, 0x49, 0x4c, 0x4e,
	0x4e, 0x4b, 0x47, 0x45, 0x4b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x4d, 0x47,
	0x49, 0x4c, 0x4e, 0x4e, 0x4b, 0x47, 0x45, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x47, 0x50, 0x4c,
	0x4b, 0x42, 0x46, 0x42, 0x50, 0x4e, 0x50, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47,
	0x47, 0x50, 0x4c, 0x4b, 0x42, 0x46, 0x42, 0x50, 0x4e, 0x50, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x75,
	0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xe8, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x48, 0x4b, 0x4d, 0x43, 0x4d, 0x46, 0x43, 0x4f, 0x4e, 0x4e, 0x4e, 0x52, 0x08, 0x62, 0x75,
	0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4b, 0x4b, 0x49, 0x42, 0x44,
	0x4e, 0x42, 0x4e, 0x46, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x4b, 0x4b,
	0x49, 0x42, 0x44, 0x4e, 0x42, 0x4e, 0x46, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x4b, 0x4e,
	0x45, 0x41, 0x49, 0x50, 0x47, 0x4c, 0x46, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45,
	0x46, 0x4b, 0x4e, 0x45, 0x41, 0x49, 0x50, 0x47, 0x4c, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4d,
	0x43, 0x42, 0x46, 0x44, 0x41, 0x46, 0x4d, 0x50, 0x50, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x4c, 0x4d, 0x43, 0x42, 0x46, 0x44, 0x41, 0x46, 0x4d, 0x50, 0x50, 0x12, 0x20, 0x0a, 0x0b,
	0x4c, 0x47, 0x4f, 0x4f, 0x4f, 0x4d, 0x47, 0x4a, 0x4b, 0x48, 0x48, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4c, 0x47, 0x4f, 0x4f, 0x4f, 0x4d, 0x47, 0x4a, 0x4b, 0x48, 0x48, 0x12, 0x20,
	0x0a, 0x0b, 0x48, 0x4b, 0x4f, 0x47, 0x43, 0x50, 0x4e, 0x48, 0x4e, 0x43, 0x41, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4b, 0x4f, 0x47, 0x43, 0x50, 0x4e, 0x48, 0x4e, 0x43, 0x41,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_FODJFBNFPJC_proto_rawDescOnce sync.Once
	file_FODJFBNFPJC_proto_rawDescData = file_FODJFBNFPJC_proto_rawDesc
)

func file_FODJFBNFPJC_proto_rawDescGZIP() []byte {
	file_FODJFBNFPJC_proto_rawDescOnce.Do(func() {
		file_FODJFBNFPJC_proto_rawDescData = protoimpl.X.CompressGZIP(file_FODJFBNFPJC_proto_rawDescData)
	})
	return file_FODJFBNFPJC_proto_rawDescData
}

var file_FODJFBNFPJC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FODJFBNFPJC_proto_goTypes = []interface{}{
	(*FODJFBNFPJC)(nil), // 0: FODJFBNFPJC
	(*HKMCMFCONNN)(nil), // 1: HKMCMFCONNN
}
var file_FODJFBNFPJC_proto_depIdxs = []int32{
	1, // 0: FODJFBNFPJC.buff_list:type_name -> HKMCMFCONNN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FODJFBNFPJC_proto_init() }
func file_FODJFBNFPJC_proto_init() {
	if File_FODJFBNFPJC_proto != nil {
		return
	}
	file_HKMCMFCONNN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FODJFBNFPJC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FODJFBNFPJC); i {
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
			RawDescriptor: file_FODJFBNFPJC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FODJFBNFPJC_proto_goTypes,
		DependencyIndexes: file_FODJFBNFPJC_proto_depIdxs,
		MessageInfos:      file_FODJFBNFPJC_proto_msgTypes,
	}.Build()
	File_FODJFBNFPJC_proto = out.File
	file_FODJFBNFPJC_proto_rawDesc = nil
	file_FODJFBNFPJC_proto_goTypes = nil
	file_FODJFBNFPJC_proto_depIdxs = nil
}