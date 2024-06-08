// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JJILCMBDNLL.proto

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

type JJILCMBDNLL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase       string   `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	ACNEBKEPBAN float32  `protobuf:"fixed32,2,opt,name=ACNEBKEPBAN,proto3" json:"ACNEBKEPBAN,omitempty"`
	BJEPMNEDMFH float32  `protobuf:"fixed32,3,opt,name=BJEPMNEDMFH,proto3" json:"BJEPMNEDMFH,omitempty"`
	CHLMOCKFMBN uint32   `protobuf:"varint,4,opt,name=CHLMOCKFMBN,proto3" json:"CHLMOCKFMBN,omitempty"`
	ADEOCAOLNCI uint32   `protobuf:"varint,5,opt,name=ADEOCAOLNCI,proto3" json:"ADEOCAOLNCI,omitempty"`
	MENOJAMFIOD uint32   `protobuf:"varint,6,opt,name=MENOJAMFIOD,proto3" json:"MENOJAMFIOD,omitempty"`
	FPGLNAEBHEA uint32   `protobuf:"varint,7,opt,name=FPGLNAEBHEA,proto3" json:"FPGLNAEBHEA,omitempty"`
	FHBAFMLLNCK uint32   `protobuf:"varint,8,opt,name=FHBAFMLLNCK,proto3" json:"FHBAFMLLNCK,omitempty"`
	KGGPCGELPEJ []uint32 `protobuf:"varint,9,rep,packed,name=KGGPCGELPEJ,proto3" json:"KGGPCGELPEJ,omitempty"`
	NCJHFBOJEGC []uint32 `protobuf:"varint,10,rep,packed,name=NCJHFBOJEGC,proto3" json:"NCJHFBOJEGC,omitempty"`
}

func (x *JJILCMBDNLL) Reset() {
	*x = JJILCMBDNLL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JJILCMBDNLL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JJILCMBDNLL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JJILCMBDNLL) ProtoMessage() {}

func (x *JJILCMBDNLL) ProtoReflect() protoreflect.Message {
	mi := &file_JJILCMBDNLL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JJILCMBDNLL.ProtoReflect.Descriptor instead.
func (*JJILCMBDNLL) Descriptor() ([]byte, []int) {
	return file_JJILCMBDNLL_proto_rawDescGZIP(), []int{0}
}

func (x *JJILCMBDNLL) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *JJILCMBDNLL) GetACNEBKEPBAN() float32 {
	if x != nil {
		return x.ACNEBKEPBAN
	}
	return 0
}

func (x *JJILCMBDNLL) GetBJEPMNEDMFH() float32 {
	if x != nil {
		return x.BJEPMNEDMFH
	}
	return 0
}

func (x *JJILCMBDNLL) GetCHLMOCKFMBN() uint32 {
	if x != nil {
		return x.CHLMOCKFMBN
	}
	return 0
}

func (x *JJILCMBDNLL) GetADEOCAOLNCI() uint32 {
	if x != nil {
		return x.ADEOCAOLNCI
	}
	return 0
}

func (x *JJILCMBDNLL) GetMENOJAMFIOD() uint32 {
	if x != nil {
		return x.MENOJAMFIOD
	}
	return 0
}

func (x *JJILCMBDNLL) GetFPGLNAEBHEA() uint32 {
	if x != nil {
		return x.FPGLNAEBHEA
	}
	return 0
}

func (x *JJILCMBDNLL) GetFHBAFMLLNCK() uint32 {
	if x != nil {
		return x.FHBAFMLLNCK
	}
	return 0
}

func (x *JJILCMBDNLL) GetKGGPCGELPEJ() []uint32 {
	if x != nil {
		return x.KGGPCGELPEJ
	}
	return nil
}

func (x *JJILCMBDNLL) GetNCJHFBOJEGC() []uint32 {
	if x != nil {
		return x.NCJHFBOJEGC
	}
	return nil
}

var File_JJILCMBDNLL_proto protoreflect.FileDescriptor

var file_JJILCMBDNLL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x4a, 0x49, 0x4c, 0x43, 0x4d, 0x42, 0x44, 0x4e, 0x4c, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x0b, 0x4a, 0x4a, 0x49, 0x4c, 0x43, 0x4d, 0x42, 0x44,
	0x4e, 0x4c, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x43, 0x4e,
	0x45, 0x42, 0x4b, 0x45, 0x50, 0x42, 0x41, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x41, 0x43, 0x4e, 0x45, 0x42, 0x4b, 0x45, 0x50, 0x42, 0x41, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x4a, 0x45, 0x50, 0x4d, 0x4e, 0x45, 0x44, 0x4d, 0x46, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x42, 0x4a, 0x45, 0x50, 0x4d, 0x4e, 0x45, 0x44, 0x4d, 0x46, 0x48, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x48, 0x4c, 0x4d, 0x4f, 0x43, 0x4b, 0x46, 0x4d, 0x42, 0x4e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x43, 0x48, 0x4c, 0x4d, 0x4f, 0x43, 0x4b, 0x46, 0x4d, 0x42, 0x4e, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x44, 0x45, 0x4f, 0x43, 0x41, 0x4f, 0x4c, 0x4e, 0x43, 0x49, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x44, 0x45, 0x4f, 0x43, 0x41, 0x4f, 0x4c, 0x4e, 0x43,
	0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x45, 0x4e, 0x4f, 0x4a, 0x41, 0x4d, 0x46, 0x49, 0x4f, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x45, 0x4e, 0x4f, 0x4a, 0x41, 0x4d, 0x46,
	0x49, 0x4f, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x50, 0x47, 0x4c, 0x4e, 0x41, 0x45, 0x42, 0x48,
	0x45, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x50, 0x47, 0x4c, 0x4e, 0x41,
	0x45, 0x42, 0x48, 0x45, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x48, 0x42, 0x41, 0x46, 0x4d, 0x4c,
	0x4c, 0x4e, 0x43, 0x4b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x48, 0x42, 0x41,
	0x46, 0x4d, 0x4c, 0x4c, 0x4e, 0x43, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x47, 0x47, 0x50, 0x43,
	0x47, 0x45, 0x4c, 0x50, 0x45, 0x4a, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x47,
	0x47, 0x50, 0x43, 0x47, 0x45, 0x4c, 0x50, 0x45, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x43, 0x4a,
	0x48, 0x46, 0x42, 0x4f, 0x4a, 0x45, 0x47, 0x43, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x4e, 0x43, 0x4a, 0x48, 0x46, 0x42, 0x4f, 0x4a, 0x45, 0x47, 0x43, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JJILCMBDNLL_proto_rawDescOnce sync.Once
	file_JJILCMBDNLL_proto_rawDescData = file_JJILCMBDNLL_proto_rawDesc
)

func file_JJILCMBDNLL_proto_rawDescGZIP() []byte {
	file_JJILCMBDNLL_proto_rawDescOnce.Do(func() {
		file_JJILCMBDNLL_proto_rawDescData = protoimpl.X.CompressGZIP(file_JJILCMBDNLL_proto_rawDescData)
	})
	return file_JJILCMBDNLL_proto_rawDescData
}

var file_JJILCMBDNLL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JJILCMBDNLL_proto_goTypes = []interface{}{
	(*JJILCMBDNLL)(nil), // 0: JJILCMBDNLL
}
var file_JJILCMBDNLL_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JJILCMBDNLL_proto_init() }
func file_JJILCMBDNLL_proto_init() {
	if File_JJILCMBDNLL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JJILCMBDNLL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JJILCMBDNLL); i {
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
			RawDescriptor: file_JJILCMBDNLL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JJILCMBDNLL_proto_goTypes,
		DependencyIndexes: file_JJILCMBDNLL_proto_depIdxs,
		MessageInfos:      file_JJILCMBDNLL_proto_msgTypes,
	}.Build()
	File_JJILCMBDNLL_proto = out.File
	file_JJILCMBDNLL_proto_rawDesc = nil
	file_JJILCMBDNLL_proto_goTypes = nil
	file_JJILCMBDNLL_proto_depIdxs = nil
}