// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BLBPCMEMJNC.proto

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

type BLBPCMEMJNC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GKPLGLEGHFE *PNKHECJAAOI `protobuf:"bytes,5,opt,name=GKPLGLEGHFE,proto3" json:"GKPLGLEGHFE,omitempty"`
	CAJLMBEDJJJ *EPHHCCKOMLK `protobuf:"bytes,10,opt,name=CAJLMBEDJJJ,proto3" json:"CAJLMBEDJJJ,omitempty"`
	ACCJAFLAOIB *NFEKEDIAPGE `protobuf:"bytes,4,opt,name=ACCJAFLAOIB,proto3" json:"ACCJAFLAOIB,omitempty"`
	ACPIMJOIKNP *INFPPJCEMHA `protobuf:"bytes,3,opt,name=ACPIMJOIKNP,proto3" json:"ACPIMJOIKNP,omitempty"`
	IOCJKOLDMHD *MIHPPHICNIH `protobuf:"bytes,14,opt,name=IOCJKOLDMHD,proto3" json:"IOCJKOLDMHD,omitempty"`
	GKHGEPLBOJP *ONJHNHIKEOC `protobuf:"bytes,13,opt,name=GKHGEPLBOJP,proto3" json:"GKHGEPLBOJP,omitempty"`
	JKPHMMLAPCC *DNOKDEMKPLI `protobuf:"bytes,6,opt,name=JKPHMMLAPCC,proto3" json:"JKPHMMLAPCC,omitempty"`
	CMNPPLPLGMP *CBGNJFGBGEE `protobuf:"bytes,1,opt,name=CMNPPLPLGMP,proto3" json:"CMNPPLPLGMP,omitempty"`
}

func (x *BLBPCMEMJNC) Reset() {
	*x = BLBPCMEMJNC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BLBPCMEMJNC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLBPCMEMJNC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLBPCMEMJNC) ProtoMessage() {}

func (x *BLBPCMEMJNC) ProtoReflect() protoreflect.Message {
	mi := &file_BLBPCMEMJNC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLBPCMEMJNC.ProtoReflect.Descriptor instead.
func (*BLBPCMEMJNC) Descriptor() ([]byte, []int) {
	return file_BLBPCMEMJNC_proto_rawDescGZIP(), []int{0}
}

func (x *BLBPCMEMJNC) GetGKPLGLEGHFE() *PNKHECJAAOI {
	if x != nil {
		return x.GKPLGLEGHFE
	}
	return nil
}

func (x *BLBPCMEMJNC) GetCAJLMBEDJJJ() *EPHHCCKOMLK {
	if x != nil {
		return x.CAJLMBEDJJJ
	}
	return nil
}

func (x *BLBPCMEMJNC) GetACCJAFLAOIB() *NFEKEDIAPGE {
	if x != nil {
		return x.ACCJAFLAOIB
	}
	return nil
}

func (x *BLBPCMEMJNC) GetACPIMJOIKNP() *INFPPJCEMHA {
	if x != nil {
		return x.ACPIMJOIKNP
	}
	return nil
}

func (x *BLBPCMEMJNC) GetIOCJKOLDMHD() *MIHPPHICNIH {
	if x != nil {
		return x.IOCJKOLDMHD
	}
	return nil
}

func (x *BLBPCMEMJNC) GetGKHGEPLBOJP() *ONJHNHIKEOC {
	if x != nil {
		return x.GKHGEPLBOJP
	}
	return nil
}

func (x *BLBPCMEMJNC) GetJKPHMMLAPCC() *DNOKDEMKPLI {
	if x != nil {
		return x.JKPHMMLAPCC
	}
	return nil
}

func (x *BLBPCMEMJNC) GetCMNPPLPLGMP() *CBGNJFGBGEE {
	if x != nil {
		return x.CMNPPLPLGMP
	}
	return nil
}

var File_BLBPCMEMJNC_proto protoreflect.FileDescriptor

var file_BLBPCMEMJNC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x4c, 0x42, 0x50, 0x43, 0x4d, 0x45, 0x4d, 0x4a, 0x4e, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x49, 0x48, 0x50, 0x50, 0x48, 0x49, 0x43, 0x4e, 0x49, 0x48,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x42, 0x47, 0x4e, 0x4a, 0x46, 0x47, 0x42,
	0x47, 0x45, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x50, 0x48, 0x48, 0x43,
	0x43, 0x4b, 0x4f, 0x4d, 0x4c, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4e,
	0x46, 0x50, 0x50, 0x4a, 0x43, 0x45, 0x4d, 0x48, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x50, 0x4e, 0x4b, 0x48, 0x45, 0x43, 0x4a, 0x41, 0x41, 0x4f, 0x49, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4e, 0x4a, 0x48, 0x4e, 0x48, 0x49, 0x4b, 0x45, 0x4f, 0x43, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4e, 0x4f, 0x4b, 0x44, 0x45, 0x4d, 0x4b, 0x50,
	0x4c, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x46, 0x45, 0x4b, 0x45, 0x44,
	0x49, 0x41, 0x50, 0x47, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x0b,
	0x42, 0x4c, 0x42, 0x50, 0x43, 0x4d, 0x45, 0x4d, 0x4a, 0x4e, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x4b, 0x50, 0x4c, 0x47, 0x4c, 0x45, 0x47, 0x48, 0x46, 0x45, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x50, 0x4e, 0x4b, 0x48, 0x45, 0x43, 0x4a, 0x41, 0x41, 0x4f, 0x49, 0x52, 0x0b,
	0x47, 0x4b, 0x50, 0x4c, 0x47, 0x4c, 0x45, 0x47, 0x48, 0x46, 0x45, 0x12, 0x2e, 0x0a, 0x0b, 0x43,
	0x41, 0x4a, 0x4c, 0x4d, 0x42, 0x45, 0x44, 0x4a, 0x4a, 0x4a, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x45, 0x50, 0x48, 0x48, 0x43, 0x43, 0x4b, 0x4f, 0x4d, 0x4c, 0x4b, 0x52, 0x0b,
	0x43, 0x41, 0x4a, 0x4c, 0x4d, 0x42, 0x45, 0x44, 0x4a, 0x4a, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x41,
	0x43, 0x43, 0x4a, 0x41, 0x46, 0x4c, 0x41, 0x4f, 0x49, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4e, 0x46, 0x45, 0x4b, 0x45, 0x44, 0x49, 0x41, 0x50, 0x47, 0x45, 0x52, 0x0b,
	0x41, 0x43, 0x43, 0x4a, 0x41, 0x46, 0x4c, 0x41, 0x4f, 0x49, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x41,
	0x43, 0x50, 0x49, 0x4d, 0x4a, 0x4f, 0x49, 0x4b, 0x4e, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x4e, 0x46, 0x50, 0x50, 0x4a, 0x43, 0x45, 0x4d, 0x48, 0x41, 0x52, 0x0b,
	0x41, 0x43, 0x50, 0x49, 0x4d, 0x4a, 0x4f, 0x49, 0x4b, 0x4e, 0x50, 0x12, 0x2e, 0x0a, 0x0b, 0x49,
	0x4f, 0x43, 0x4a, 0x4b, 0x4f, 0x4c, 0x44, 0x4d, 0x48, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4d, 0x49, 0x48, 0x50, 0x50, 0x48, 0x49, 0x43, 0x4e, 0x49, 0x48, 0x52, 0x0b,
	0x49, 0x4f, 0x43, 0x4a, 0x4b, 0x4f, 0x4c, 0x44, 0x4d, 0x48, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x4b, 0x48, 0x47, 0x45, 0x50, 0x4c, 0x42, 0x4f, 0x4a, 0x50, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4f, 0x4e, 0x4a, 0x48, 0x4e, 0x48, 0x49, 0x4b, 0x45, 0x4f, 0x43, 0x52, 0x0b,
	0x47, 0x4b, 0x48, 0x47, 0x45, 0x50, 0x4c, 0x42, 0x4f, 0x4a, 0x50, 0x12, 0x2e, 0x0a, 0x0b, 0x4a,
	0x4b, 0x50, 0x48, 0x4d, 0x4d, 0x4c, 0x41, 0x50, 0x43, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x4e, 0x4f, 0x4b, 0x44, 0x45, 0x4d, 0x4b, 0x50, 0x4c, 0x49, 0x52, 0x0b,
	0x4a, 0x4b, 0x50, 0x48, 0x4d, 0x4d, 0x4c, 0x41, 0x50, 0x43, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x43,
	0x4d, 0x4e, 0x50, 0x50, 0x4c, 0x50, 0x4c, 0x47, 0x4d, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x43, 0x42, 0x47, 0x4e, 0x4a, 0x46, 0x47, 0x42, 0x47, 0x45, 0x45, 0x52, 0x0b,
	0x43, 0x4d, 0x4e, 0x50, 0x50, 0x4c, 0x50, 0x4c, 0x47, 0x4d, 0x50, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_BLBPCMEMJNC_proto_rawDescOnce sync.Once
	file_BLBPCMEMJNC_proto_rawDescData = file_BLBPCMEMJNC_proto_rawDesc
)

func file_BLBPCMEMJNC_proto_rawDescGZIP() []byte {
	file_BLBPCMEMJNC_proto_rawDescOnce.Do(func() {
		file_BLBPCMEMJNC_proto_rawDescData = protoimpl.X.CompressGZIP(file_BLBPCMEMJNC_proto_rawDescData)
	})
	return file_BLBPCMEMJNC_proto_rawDescData
}

var file_BLBPCMEMJNC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BLBPCMEMJNC_proto_goTypes = []interface{}{
	(*BLBPCMEMJNC)(nil), // 0: BLBPCMEMJNC
	(*PNKHECJAAOI)(nil), // 1: PNKHECJAAOI
	(*EPHHCCKOMLK)(nil), // 2: EPHHCCKOMLK
	(*NFEKEDIAPGE)(nil), // 3: NFEKEDIAPGE
	(*INFPPJCEMHA)(nil), // 4: INFPPJCEMHA
	(*MIHPPHICNIH)(nil), // 5: MIHPPHICNIH
	(*ONJHNHIKEOC)(nil), // 6: ONJHNHIKEOC
	(*DNOKDEMKPLI)(nil), // 7: DNOKDEMKPLI
	(*CBGNJFGBGEE)(nil), // 8: CBGNJFGBGEE
}
var file_BLBPCMEMJNC_proto_depIdxs = []int32{
	1, // 0: BLBPCMEMJNC.GKPLGLEGHFE:type_name -> PNKHECJAAOI
	2, // 1: BLBPCMEMJNC.CAJLMBEDJJJ:type_name -> EPHHCCKOMLK
	3, // 2: BLBPCMEMJNC.ACCJAFLAOIB:type_name -> NFEKEDIAPGE
	4, // 3: BLBPCMEMJNC.ACPIMJOIKNP:type_name -> INFPPJCEMHA
	5, // 4: BLBPCMEMJNC.IOCJKOLDMHD:type_name -> MIHPPHICNIH
	6, // 5: BLBPCMEMJNC.GKHGEPLBOJP:type_name -> ONJHNHIKEOC
	7, // 6: BLBPCMEMJNC.JKPHMMLAPCC:type_name -> DNOKDEMKPLI
	8, // 7: BLBPCMEMJNC.CMNPPLPLGMP:type_name -> CBGNJFGBGEE
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_BLBPCMEMJNC_proto_init() }
func file_BLBPCMEMJNC_proto_init() {
	if File_BLBPCMEMJNC_proto != nil {
		return
	}
	file_MIHPPHICNIH_proto_init()
	file_CBGNJFGBGEE_proto_init()
	file_EPHHCCKOMLK_proto_init()
	file_INFPPJCEMHA_proto_init()
	file_PNKHECJAAOI_proto_init()
	file_ONJHNHIKEOC_proto_init()
	file_DNOKDEMKPLI_proto_init()
	file_NFEKEDIAPGE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BLBPCMEMJNC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLBPCMEMJNC); i {
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
			RawDescriptor: file_BLBPCMEMJNC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BLBPCMEMJNC_proto_goTypes,
		DependencyIndexes: file_BLBPCMEMJNC_proto_depIdxs,
		MessageInfos:      file_BLBPCMEMJNC_proto_msgTypes,
	}.Build()
	File_BLBPCMEMJNC_proto = out.File
	file_BLBPCMEMJNC_proto_rawDesc = nil
	file_BLBPCMEMJNC_proto_goTypes = nil
	file_BLBPCMEMJNC_proto_depIdxs = nil
}