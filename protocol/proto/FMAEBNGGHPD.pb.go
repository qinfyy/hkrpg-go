// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FMAEBNGGHPD.proto

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

type FMAEBNGGHPD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LFBDOEBGLMD uint32       `protobuf:"varint,2,opt,name=LFBDOEBGLMD,proto3" json:"LFBDOEBGLMD,omitempty"`
	MIPJMLKPFDF *FCKIFNDEJLC `protobuf:"bytes,8,opt,name=MIPJMLKPFDF,proto3" json:"MIPJMLKPFDF,omitempty"`
	JPHCENNAMPP *POCKOFMNDAG `protobuf:"bytes,7,opt,name=JPHCENNAMPP,proto3" json:"JPHCENNAMPP,omitempty"`
	JPIKKFGJMCE *FLIJBNILBLK `protobuf:"bytes,9,opt,name=JPIKKFGJMCE,proto3" json:"JPIKKFGJMCE,omitempty"`
	GLAMIJLHCLK *HDKJDNOGIJC `protobuf:"bytes,14,opt,name=GLAMIJLHCLK,proto3" json:"GLAMIJLHCLK,omitempty"`
	MBJIIJLMNHH *KDAJLDOONEP `protobuf:"bytes,1,opt,name=MBJIIJLMNHH,proto3" json:"MBJIIJLMNHH,omitempty"`
	PLCHKGEBGGO *NFKBEABDAPM `protobuf:"bytes,11,opt,name=PLCHKGEBGGO,proto3" json:"PLCHKGEBGGO,omitempty"`
	KJEJOFFOMKF *HOFAEKPEDLN `protobuf:"bytes,12,opt,name=KJEJOFFOMKF,proto3" json:"KJEJOFFOMKF,omitempty"`
	OIBEKLKMHEC *GIGDEINNDJO `protobuf:"bytes,10,opt,name=OIBEKLKMHEC,proto3" json:"OIBEKLKMHEC,omitempty"`
}

func (x *FMAEBNGGHPD) Reset() {
	*x = FMAEBNGGHPD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FMAEBNGGHPD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FMAEBNGGHPD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FMAEBNGGHPD) ProtoMessage() {}

func (x *FMAEBNGGHPD) ProtoReflect() protoreflect.Message {
	mi := &file_FMAEBNGGHPD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FMAEBNGGHPD.ProtoReflect.Descriptor instead.
func (*FMAEBNGGHPD) Descriptor() ([]byte, []int) {
	return file_FMAEBNGGHPD_proto_rawDescGZIP(), []int{0}
}

func (x *FMAEBNGGHPD) GetLFBDOEBGLMD() uint32 {
	if x != nil {
		return x.LFBDOEBGLMD
	}
	return 0
}

func (x *FMAEBNGGHPD) GetMIPJMLKPFDF() *FCKIFNDEJLC {
	if x != nil {
		return x.MIPJMLKPFDF
	}
	return nil
}

func (x *FMAEBNGGHPD) GetJPHCENNAMPP() *POCKOFMNDAG {
	if x != nil {
		return x.JPHCENNAMPP
	}
	return nil
}

func (x *FMAEBNGGHPD) GetJPIKKFGJMCE() *FLIJBNILBLK {
	if x != nil {
		return x.JPIKKFGJMCE
	}
	return nil
}

func (x *FMAEBNGGHPD) GetGLAMIJLHCLK() *HDKJDNOGIJC {
	if x != nil {
		return x.GLAMIJLHCLK
	}
	return nil
}

func (x *FMAEBNGGHPD) GetMBJIIJLMNHH() *KDAJLDOONEP {
	if x != nil {
		return x.MBJIIJLMNHH
	}
	return nil
}

func (x *FMAEBNGGHPD) GetPLCHKGEBGGO() *NFKBEABDAPM {
	if x != nil {
		return x.PLCHKGEBGGO
	}
	return nil
}

func (x *FMAEBNGGHPD) GetKJEJOFFOMKF() *HOFAEKPEDLN {
	if x != nil {
		return x.KJEJOFFOMKF
	}
	return nil
}

func (x *FMAEBNGGHPD) GetOIBEKLKMHEC() *GIGDEINNDJO {
	if x != nil {
		return x.OIBEKLKMHEC
	}
	return nil
}

var File_FMAEBNGGHPD_proto protoreflect.FileDescriptor

var file_FMAEBNGGHPD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x4d, 0x41, 0x45, 0x42, 0x4e, 0x47, 0x47, 0x48, 0x50, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x46, 0x4b, 0x42, 0x45, 0x41, 0x42, 0x44, 0x41, 0x50, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x44, 0x4b, 0x4a, 0x44, 0x4e, 0x4f, 0x47,
	0x49, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x49, 0x47, 0x44, 0x45,
	0x49, 0x4e, 0x4e, 0x44, 0x4a, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x43,
	0x4b, 0x49, 0x46, 0x4e, 0x44, 0x45, 0x4a, 0x4c, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x46, 0x4c, 0x49, 0x4a, 0x42, 0x4e, 0x49, 0x4c, 0x42, 0x4c, 0x4b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x44, 0x41, 0x4a, 0x4c, 0x44, 0x4f, 0x4f, 0x4e, 0x45, 0x50, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4f, 0x43, 0x4b, 0x4f, 0x46, 0x4d, 0x4e, 0x44,
	0x41, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4f, 0x46, 0x41, 0x45, 0x4b,
	0x50, 0x45, 0x44, 0x4c, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x0b,
	0x46, 0x4d, 0x41, 0x45, 0x42, 0x4e, 0x47, 0x47, 0x48, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4c,
	0x46, 0x42, 0x44, 0x4f, 0x45, 0x42, 0x47, 0x4c, 0x4d, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4c, 0x46, 0x42, 0x44, 0x4f, 0x45, 0x42, 0x47, 0x4c, 0x4d, 0x44, 0x12, 0x2e, 0x0a,
	0x0b, 0x4d, 0x49, 0x50, 0x4a, 0x4d, 0x4c, 0x4b, 0x50, 0x46, 0x44, 0x46, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x43, 0x4b, 0x49, 0x46, 0x4e, 0x44, 0x45, 0x4a, 0x4c, 0x43,
	0x52, 0x0b, 0x4d, 0x49, 0x50, 0x4a, 0x4d, 0x4c, 0x4b, 0x50, 0x46, 0x44, 0x46, 0x12, 0x2e, 0x0a,
	0x0b, 0x4a, 0x50, 0x48, 0x43, 0x45, 0x4e, 0x4e, 0x41, 0x4d, 0x50, 0x50, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4f, 0x43, 0x4b, 0x4f, 0x46, 0x4d, 0x4e, 0x44, 0x41, 0x47,
	0x52, 0x0b, 0x4a, 0x50, 0x48, 0x43, 0x45, 0x4e, 0x4e, 0x41, 0x4d, 0x50, 0x50, 0x12, 0x2e, 0x0a,
	0x0b, 0x4a, 0x50, 0x49, 0x4b, 0x4b, 0x46, 0x47, 0x4a, 0x4d, 0x43, 0x45, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4c, 0x49, 0x4a, 0x42, 0x4e, 0x49, 0x4c, 0x42, 0x4c, 0x4b,
	0x52, 0x0b, 0x4a, 0x50, 0x49, 0x4b, 0x4b, 0x46, 0x47, 0x4a, 0x4d, 0x43, 0x45, 0x12, 0x2e, 0x0a,
	0x0b, 0x47, 0x4c, 0x41, 0x4d, 0x49, 0x4a, 0x4c, 0x48, 0x43, 0x4c, 0x4b, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x44, 0x4b, 0x4a, 0x44, 0x4e, 0x4f, 0x47, 0x49, 0x4a, 0x43,
	0x52, 0x0b, 0x47, 0x4c, 0x41, 0x4d, 0x49, 0x4a, 0x4c, 0x48, 0x43, 0x4c, 0x4b, 0x12, 0x2e, 0x0a,
	0x0b, 0x4d, 0x42, 0x4a, 0x49, 0x49, 0x4a, 0x4c, 0x4d, 0x4e, 0x48, 0x48, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x44, 0x41, 0x4a, 0x4c, 0x44, 0x4f, 0x4f, 0x4e, 0x45, 0x50,
	0x52, 0x0b, 0x4d, 0x42, 0x4a, 0x49, 0x49, 0x4a, 0x4c, 0x4d, 0x4e, 0x48, 0x48, 0x12, 0x2e, 0x0a,
	0x0b, 0x50, 0x4c, 0x43, 0x48, 0x4b, 0x47, 0x45, 0x42, 0x47, 0x47, 0x4f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x46, 0x4b, 0x42, 0x45, 0x41, 0x42, 0x44, 0x41, 0x50, 0x4d,
	0x52, 0x0b, 0x50, 0x4c, 0x43, 0x48, 0x4b, 0x47, 0x45, 0x42, 0x47, 0x47, 0x4f, 0x12, 0x2e, 0x0a,
	0x0b, 0x4b, 0x4a, 0x45, 0x4a, 0x4f, 0x46, 0x46, 0x4f, 0x4d, 0x4b, 0x46, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4f, 0x46, 0x41, 0x45, 0x4b, 0x50, 0x45, 0x44, 0x4c, 0x4e,
	0x52, 0x0b, 0x4b, 0x4a, 0x45, 0x4a, 0x4f, 0x46, 0x46, 0x4f, 0x4d, 0x4b, 0x46, 0x12, 0x2e, 0x0a,
	0x0b, 0x4f, 0x49, 0x42, 0x45, 0x4b, 0x4c, 0x4b, 0x4d, 0x48, 0x45, 0x43, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x49, 0x47, 0x44, 0x45, 0x49, 0x4e, 0x4e, 0x44, 0x4a, 0x4f,
	0x52, 0x0b, 0x4f, 0x49, 0x42, 0x45, 0x4b, 0x4c, 0x4b, 0x4d, 0x48, 0x45, 0x43, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FMAEBNGGHPD_proto_rawDescOnce sync.Once
	file_FMAEBNGGHPD_proto_rawDescData = file_FMAEBNGGHPD_proto_rawDesc
)

func file_FMAEBNGGHPD_proto_rawDescGZIP() []byte {
	file_FMAEBNGGHPD_proto_rawDescOnce.Do(func() {
		file_FMAEBNGGHPD_proto_rawDescData = protoimpl.X.CompressGZIP(file_FMAEBNGGHPD_proto_rawDescData)
	})
	return file_FMAEBNGGHPD_proto_rawDescData
}

var file_FMAEBNGGHPD_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FMAEBNGGHPD_proto_goTypes = []interface{}{
	(*FMAEBNGGHPD)(nil), // 0: FMAEBNGGHPD
	(*FCKIFNDEJLC)(nil), // 1: FCKIFNDEJLC
	(*POCKOFMNDAG)(nil), // 2: POCKOFMNDAG
	(*FLIJBNILBLK)(nil), // 3: FLIJBNILBLK
	(*HDKJDNOGIJC)(nil), // 4: HDKJDNOGIJC
	(*KDAJLDOONEP)(nil), // 5: KDAJLDOONEP
	(*NFKBEABDAPM)(nil), // 6: NFKBEABDAPM
	(*HOFAEKPEDLN)(nil), // 7: HOFAEKPEDLN
	(*GIGDEINNDJO)(nil), // 8: GIGDEINNDJO
}
var file_FMAEBNGGHPD_proto_depIdxs = []int32{
	1, // 0: FMAEBNGGHPD.MIPJMLKPFDF:type_name -> FCKIFNDEJLC
	2, // 1: FMAEBNGGHPD.JPHCENNAMPP:type_name -> POCKOFMNDAG
	3, // 2: FMAEBNGGHPD.JPIKKFGJMCE:type_name -> FLIJBNILBLK
	4, // 3: FMAEBNGGHPD.GLAMIJLHCLK:type_name -> HDKJDNOGIJC
	5, // 4: FMAEBNGGHPD.MBJIIJLMNHH:type_name -> KDAJLDOONEP
	6, // 5: FMAEBNGGHPD.PLCHKGEBGGO:type_name -> NFKBEABDAPM
	7, // 6: FMAEBNGGHPD.KJEJOFFOMKF:type_name -> HOFAEKPEDLN
	8, // 7: FMAEBNGGHPD.OIBEKLKMHEC:type_name -> GIGDEINNDJO
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_FMAEBNGGHPD_proto_init() }
func file_FMAEBNGGHPD_proto_init() {
	if File_FMAEBNGGHPD_proto != nil {
		return
	}
	file_NFKBEABDAPM_proto_init()
	file_HDKJDNOGIJC_proto_init()
	file_GIGDEINNDJO_proto_init()
	file_FCKIFNDEJLC_proto_init()
	file_FLIJBNILBLK_proto_init()
	file_KDAJLDOONEP_proto_init()
	file_POCKOFMNDAG_proto_init()
	file_HOFAEKPEDLN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FMAEBNGGHPD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FMAEBNGGHPD); i {
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
			RawDescriptor: file_FMAEBNGGHPD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FMAEBNGGHPD_proto_goTypes,
		DependencyIndexes: file_FMAEBNGGHPD_proto_depIdxs,
		MessageInfos:      file_FMAEBNGGHPD_proto_msgTypes,
	}.Build()
	File_FMAEBNGGHPD_proto = out.File
	file_FMAEBNGGHPD_proto_rawDesc = nil
	file_FMAEBNGGHPD_proto_goTypes = nil
	file_FMAEBNGGHPD_proto_depIdxs = nil
}
