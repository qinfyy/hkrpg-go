// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AnnounceData.proto

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

type AnnounceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndTime     int64  `protobuf:"varint,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ConfigId    uint32 `protobuf:"varint,11,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	LAMFOFKABEO uint32 `protobuf:"varint,10,opt,name=LAMFOFKABEO,proto3" json:"LAMFOFKABEO,omitempty"`
	BCPBPCDKAGD string `protobuf:"bytes,1,opt,name=BCPBPCDKAGD,proto3" json:"BCPBPCDKAGD,omitempty"`
	NAFKMMGJCND string `protobuf:"bytes,13,opt,name=NAFKMMGJCND,proto3" json:"NAFKMMGJCND,omitempty"`
	BeginTime   int64  `protobuf:"varint,15,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	CHJPFPLHJBJ string `protobuf:"bytes,4,opt,name=CHJPFPLHJBJ,proto3" json:"CHJPFPLHJBJ,omitempty"`
	DKKPPNNJIFO uint32 `protobuf:"varint,3,opt,name=DKKPPNNJIFO,proto3" json:"DKKPPNNJIFO,omitempty"`
	HLNPMNBKBBC bool   `protobuf:"varint,6,opt,name=HLNPMNBKBBC,proto3" json:"HLNPMNBKBBC,omitempty"`
}

func (x *AnnounceData) Reset() {
	*x = AnnounceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AnnounceData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceData) ProtoMessage() {}

func (x *AnnounceData) ProtoReflect() protoreflect.Message {
	mi := &file_AnnounceData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceData.ProtoReflect.Descriptor instead.
func (*AnnounceData) Descriptor() ([]byte, []int) {
	return file_AnnounceData_proto_rawDescGZIP(), []int{0}
}

func (x *AnnounceData) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *AnnounceData) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *AnnounceData) GetLAMFOFKABEO() uint32 {
	if x != nil {
		return x.LAMFOFKABEO
	}
	return 0
}

func (x *AnnounceData) GetBCPBPCDKAGD() string {
	if x != nil {
		return x.BCPBPCDKAGD
	}
	return ""
}

func (x *AnnounceData) GetNAFKMMGJCND() string {
	if x != nil {
		return x.NAFKMMGJCND
	}
	return ""
}

func (x *AnnounceData) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *AnnounceData) GetCHJPFPLHJBJ() string {
	if x != nil {
		return x.CHJPFPLHJBJ
	}
	return ""
}

func (x *AnnounceData) GetDKKPPNNJIFO() uint32 {
	if x != nil {
		return x.DKKPPNNJIFO
	}
	return 0
}

func (x *AnnounceData) GetHLNPMNBKBBC() bool {
	if x != nil {
		return x.HLNPMNBKBBC
	}
	return false
}

var File_AnnounceData_proto protoreflect.FileDescriptor

var file_AnnounceData_proto_rawDesc = []byte{
	0x0a, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x4c, 0x41, 0x4d, 0x46, 0x4f, 0x46, 0x4b, 0x41, 0x42, 0x45, 0x4f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x41, 0x4d, 0x46, 0x4f, 0x46, 0x4b, 0x41, 0x42, 0x45, 0x4f, 0x12,
	0x20, 0x0a, 0x0b, 0x42, 0x43, 0x50, 0x42, 0x50, 0x43, 0x44, 0x4b, 0x41, 0x47, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x43, 0x50, 0x42, 0x50, 0x43, 0x44, 0x4b, 0x41, 0x47,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x41, 0x46, 0x4b, 0x4d, 0x4d, 0x47, 0x4a, 0x43, 0x4e, 0x44,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x41, 0x46, 0x4b, 0x4d, 0x4d, 0x47, 0x4a,
	0x43, 0x4e, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x48, 0x4a, 0x50, 0x46, 0x50, 0x4c, 0x48, 0x4a, 0x42,
	0x4a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x48, 0x4a, 0x50, 0x46, 0x50, 0x4c,
	0x48, 0x4a, 0x42, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4b, 0x4b, 0x50, 0x50, 0x4e, 0x4e, 0x4a,
	0x49, 0x46, 0x4f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4b, 0x4b, 0x50, 0x50,
	0x4e, 0x4e, 0x4a, 0x49, 0x46, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4c, 0x4e, 0x50, 0x4d, 0x4e,
	0x42, 0x4b, 0x42, 0x42, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x4c, 0x4e,
	0x50, 0x4d, 0x4e, 0x42, 0x4b, 0x42, 0x42, 0x43, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AnnounceData_proto_rawDescOnce sync.Once
	file_AnnounceData_proto_rawDescData = file_AnnounceData_proto_rawDesc
)

func file_AnnounceData_proto_rawDescGZIP() []byte {
	file_AnnounceData_proto_rawDescOnce.Do(func() {
		file_AnnounceData_proto_rawDescData = protoimpl.X.CompressGZIP(file_AnnounceData_proto_rawDescData)
	})
	return file_AnnounceData_proto_rawDescData
}

var file_AnnounceData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AnnounceData_proto_goTypes = []interface{}{
	(*AnnounceData)(nil), // 0: AnnounceData
}
var file_AnnounceData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AnnounceData_proto_init() }
func file_AnnounceData_proto_init() {
	if File_AnnounceData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AnnounceData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceData); i {
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
			RawDescriptor: file_AnnounceData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AnnounceData_proto_goTypes,
		DependencyIndexes: file_AnnounceData_proto_depIdxs,
		MessageInfos:      file_AnnounceData_proto_msgTypes,
	}.Build()
	File_AnnounceData_proto = out.File
	file_AnnounceData_proto_rawDesc = nil
	file_AnnounceData_proto_goTypes = nil
	file_AnnounceData_proto_depIdxs = nil
}