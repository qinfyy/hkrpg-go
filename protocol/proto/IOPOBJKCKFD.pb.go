// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: IOPOBJKCKFD.proto

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

type IOPOBJKCKFD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AAKIGJLJDHJ uint32 `protobuf:"varint,5,opt,name=AAKIGJLJDHJ,proto3" json:"AAKIGJLJDHJ,omitempty"`
	ConfigId    uint32 `protobuf:"varint,4,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	LLKDIMANDCH uint32 `protobuf:"varint,9,opt,name=LLKDIMANDCH,proto3" json:"LLKDIMANDCH,omitempty"`
	// Types that are assignable to BHDFLALIMCI:
	//
	//	*IOPOBJKCKFD_HPNKHMPIDKI
	//	*IOPOBJKCKFD_EBPMCIHJLNP
	//	*IOPOBJKCKFD_IFFFOCLOCGP
	BHDFLALIMCI isIOPOBJKCKFD_BHDFLALIMCI `protobuf_oneof:"BHDFLALIMCI"`
}

func (x *IOPOBJKCKFD) Reset() {
	*x = IOPOBJKCKFD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IOPOBJKCKFD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOPOBJKCKFD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOPOBJKCKFD) ProtoMessage() {}

func (x *IOPOBJKCKFD) ProtoReflect() protoreflect.Message {
	mi := &file_IOPOBJKCKFD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOPOBJKCKFD.ProtoReflect.Descriptor instead.
func (*IOPOBJKCKFD) Descriptor() ([]byte, []int) {
	return file_IOPOBJKCKFD_proto_rawDescGZIP(), []int{0}
}

func (x *IOPOBJKCKFD) GetAAKIGJLJDHJ() uint32 {
	if x != nil {
		return x.AAKIGJLJDHJ
	}
	return 0
}

func (x *IOPOBJKCKFD) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *IOPOBJKCKFD) GetLLKDIMANDCH() uint32 {
	if x != nil {
		return x.LLKDIMANDCH
	}
	return 0
}

func (m *IOPOBJKCKFD) GetBHDFLALIMCI() isIOPOBJKCKFD_BHDFLALIMCI {
	if m != nil {
		return m.BHDFLALIMCI
	}
	return nil
}

func (x *IOPOBJKCKFD) GetHPNKHMPIDKI() *OPIMJFDHBEE {
	if x, ok := x.GetBHDFLALIMCI().(*IOPOBJKCKFD_HPNKHMPIDKI); ok {
		return x.HPNKHMPIDKI
	}
	return nil
}

func (x *IOPOBJKCKFD) GetEBPMCIHJLNP() *IAGKOOPHPBM {
	if x, ok := x.GetBHDFLALIMCI().(*IOPOBJKCKFD_EBPMCIHJLNP); ok {
		return x.EBPMCIHJLNP
	}
	return nil
}

func (x *IOPOBJKCKFD) GetIFFFOCLOCGP() *EJMONFLFDJK {
	if x, ok := x.GetBHDFLALIMCI().(*IOPOBJKCKFD_IFFFOCLOCGP); ok {
		return x.IFFFOCLOCGP
	}
	return nil
}

type isIOPOBJKCKFD_BHDFLALIMCI interface {
	isIOPOBJKCKFD_BHDFLALIMCI()
}

type IOPOBJKCKFD_HPNKHMPIDKI struct {
	HPNKHMPIDKI *OPIMJFDHBEE `protobuf:"bytes,14,opt,name=HPNKHMPIDKI,proto3,oneof"`
}

type IOPOBJKCKFD_EBPMCIHJLNP struct {
	EBPMCIHJLNP *IAGKOOPHPBM `protobuf:"bytes,6,opt,name=EBPMCIHJLNP,proto3,oneof"`
}

type IOPOBJKCKFD_IFFFOCLOCGP struct {
	IFFFOCLOCGP *EJMONFLFDJK `protobuf:"bytes,11,opt,name=IFFFOCLOCGP,proto3,oneof"`
}

func (*IOPOBJKCKFD_HPNKHMPIDKI) isIOPOBJKCKFD_BHDFLALIMCI() {}

func (*IOPOBJKCKFD_EBPMCIHJLNP) isIOPOBJKCKFD_BHDFLALIMCI() {}

func (*IOPOBJKCKFD_IFFFOCLOCGP) isIOPOBJKCKFD_BHDFLALIMCI() {}

var File_IOPOBJKCKFD_proto protoreflect.FileDescriptor

var file_IOPOBJKCKFD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x4f, 0x50, 0x4f, 0x42, 0x4a, 0x4b, 0x43, 0x4b, 0x46, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x50, 0x49, 0x4d, 0x4a, 0x46, 0x44, 0x48, 0x42, 0x45, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4a, 0x4d, 0x4f, 0x4e, 0x46, 0x4c, 0x46,
	0x44, 0x4a, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x41, 0x47, 0x4b, 0x4f,
	0x4f, 0x50, 0x48, 0x50, 0x42, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a,
	0x0b, 0x49, 0x4f, 0x50, 0x4f, 0x42, 0x4a, 0x4b, 0x43, 0x4b, 0x46, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x41, 0x4b, 0x49, 0x47, 0x4a, 0x4c, 0x4a, 0x44, 0x48, 0x4a, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x41, 0x41, 0x4b, 0x49, 0x47, 0x4a, 0x4c, 0x4a, 0x44, 0x48, 0x4a, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4c,
	0x4c, 0x4b, 0x44, 0x49, 0x4d, 0x41, 0x4e, 0x44, 0x43, 0x48, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4c, 0x4c, 0x4b, 0x44, 0x49, 0x4d, 0x41, 0x4e, 0x44, 0x43, 0x48, 0x12, 0x30, 0x0a,
	0x0b, 0x48, 0x50, 0x4e, 0x4b, 0x48, 0x4d, 0x50, 0x49, 0x44, 0x4b, 0x49, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x50, 0x49, 0x4d, 0x4a, 0x46, 0x44, 0x48, 0x42, 0x45, 0x45,
	0x48, 0x00, 0x52, 0x0b, 0x48, 0x50, 0x4e, 0x4b, 0x48, 0x4d, 0x50, 0x49, 0x44, 0x4b, 0x49, 0x12,
	0x30, 0x0a, 0x0b, 0x45, 0x42, 0x50, 0x4d, 0x43, 0x49, 0x48, 0x4a, 0x4c, 0x4e, 0x50, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x41, 0x47, 0x4b, 0x4f, 0x4f, 0x50, 0x48, 0x50,
	0x42, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x42, 0x50, 0x4d, 0x43, 0x49, 0x48, 0x4a, 0x4c, 0x4e,
	0x50, 0x12, 0x30, 0x0a, 0x0b, 0x49, 0x46, 0x46, 0x46, 0x4f, 0x43, 0x4c, 0x4f, 0x43, 0x47, 0x50,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4a, 0x4d, 0x4f, 0x4e, 0x46, 0x4c,
	0x46, 0x44, 0x4a, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x46, 0x46, 0x46, 0x4f, 0x43, 0x4c, 0x4f,
	0x43, 0x47, 0x50, 0x42, 0x0d, 0x0a, 0x0b, 0x42, 0x48, 0x44, 0x46, 0x4c, 0x41, 0x4c, 0x49, 0x4d,
	0x43, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IOPOBJKCKFD_proto_rawDescOnce sync.Once
	file_IOPOBJKCKFD_proto_rawDescData = file_IOPOBJKCKFD_proto_rawDesc
)

func file_IOPOBJKCKFD_proto_rawDescGZIP() []byte {
	file_IOPOBJKCKFD_proto_rawDescOnce.Do(func() {
		file_IOPOBJKCKFD_proto_rawDescData = protoimpl.X.CompressGZIP(file_IOPOBJKCKFD_proto_rawDescData)
	})
	return file_IOPOBJKCKFD_proto_rawDescData
}

var file_IOPOBJKCKFD_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IOPOBJKCKFD_proto_goTypes = []interface{}{
	(*IOPOBJKCKFD)(nil), // 0: IOPOBJKCKFD
	(*OPIMJFDHBEE)(nil), // 1: OPIMJFDHBEE
	(*IAGKOOPHPBM)(nil), // 2: IAGKOOPHPBM
	(*EJMONFLFDJK)(nil), // 3: EJMONFLFDJK
}
var file_IOPOBJKCKFD_proto_depIdxs = []int32{
	1, // 0: IOPOBJKCKFD.HPNKHMPIDKI:type_name -> OPIMJFDHBEE
	2, // 1: IOPOBJKCKFD.EBPMCIHJLNP:type_name -> IAGKOOPHPBM
	3, // 2: IOPOBJKCKFD.IFFFOCLOCGP:type_name -> EJMONFLFDJK
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_IOPOBJKCKFD_proto_init() }
func file_IOPOBJKCKFD_proto_init() {
	if File_IOPOBJKCKFD_proto != nil {
		return
	}
	file_OPIMJFDHBEE_proto_init()
	file_EJMONFLFDJK_proto_init()
	file_IAGKOOPHPBM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IOPOBJKCKFD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOPOBJKCKFD); i {
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
	file_IOPOBJKCKFD_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IOPOBJKCKFD_HPNKHMPIDKI)(nil),
		(*IOPOBJKCKFD_EBPMCIHJLNP)(nil),
		(*IOPOBJKCKFD_IFFFOCLOCGP)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_IOPOBJKCKFD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IOPOBJKCKFD_proto_goTypes,
		DependencyIndexes: file_IOPOBJKCKFD_proto_depIdxs,
		MessageInfos:      file_IOPOBJKCKFD_proto_msgTypes,
	}.Build()
	File_IOPOBJKCKFD_proto = out.File
	file_IOPOBJKCKFD_proto_rawDesc = nil
	file_IOPOBJKCKFD_proto_goTypes = nil
	file_IOPOBJKCKFD_proto_depIdxs = nil
}
