// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpdatePsnSettingsInfoCsReq.proto

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

type UpdatePsnSettingsInfoCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CDBHDDDBGIP:
	//
	//	*UpdatePsnSettingsInfoCsReq_PNJCIINPILL
	//	*UpdatePsnSettingsInfoCsReq_DDHKNBKPEJO
	CDBHDDDBGIP isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP `protobuf_oneof:"CDBHDDDBGIP"`
}

func (x *UpdatePsnSettingsInfoCsReq) Reset() {
	*x = UpdatePsnSettingsInfoCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdatePsnSettingsInfoCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePsnSettingsInfoCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePsnSettingsInfoCsReq) ProtoMessage() {}

func (x *UpdatePsnSettingsInfoCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_UpdatePsnSettingsInfoCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePsnSettingsInfoCsReq.ProtoReflect.Descriptor instead.
func (*UpdatePsnSettingsInfoCsReq) Descriptor() ([]byte, []int) {
	return file_UpdatePsnSettingsInfoCsReq_proto_rawDescGZIP(), []int{0}
}

func (m *UpdatePsnSettingsInfoCsReq) GetCDBHDDDBGIP() isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP {
	if m != nil {
		return m.CDBHDDDBGIP
	}
	return nil
}

func (x *UpdatePsnSettingsInfoCsReq) GetPNJCIINPILL() *EKJOBFKOMIE {
	if x, ok := x.GetCDBHDDDBGIP().(*UpdatePsnSettingsInfoCsReq_PNJCIINPILL); ok {
		return x.PNJCIINPILL
	}
	return nil
}

func (x *UpdatePsnSettingsInfoCsReq) GetDDHKNBKPEJO() *KFBHJBGHHDO {
	if x, ok := x.GetCDBHDDDBGIP().(*UpdatePsnSettingsInfoCsReq_DDHKNBKPEJO); ok {
		return x.DDHKNBKPEJO
	}
	return nil
}

type isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP interface {
	isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP()
}

type UpdatePsnSettingsInfoCsReq_PNJCIINPILL struct {
	PNJCIINPILL *EKJOBFKOMIE `protobuf:"bytes,741,opt,name=PNJCIINPILL,proto3,oneof"`
}

type UpdatePsnSettingsInfoCsReq_DDHKNBKPEJO struct {
	DDHKNBKPEJO *KFBHJBGHHDO `protobuf:"bytes,1350,opt,name=DDHKNBKPEJO,proto3,oneof"`
}

func (*UpdatePsnSettingsInfoCsReq_PNJCIINPILL) isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP() {}

func (*UpdatePsnSettingsInfoCsReq_DDHKNBKPEJO) isUpdatePsnSettingsInfoCsReq_CDBHDDDBGIP() {}

var File_UpdatePsnSettingsInfoCsReq_proto protoreflect.FileDescriptor

var file_UpdatePsnSettingsInfoCsReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x73, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x46, 0x42, 0x48, 0x4a, 0x42, 0x47, 0x48, 0x48, 0x44, 0x4f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4b, 0x4a, 0x4f, 0x42, 0x46, 0x4b, 0x4f, 0x4d,
	0x49, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x73, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x0b, 0x50, 0x4e, 0x4a, 0x43, 0x49,
	0x49, 0x4e, 0x50, 0x49, 0x4c, 0x4c, 0x18, 0xe5, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x45, 0x4b, 0x4a, 0x4f, 0x42, 0x46, 0x4b, 0x4f, 0x4d, 0x49, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x50,
	0x4e, 0x4a, 0x43, 0x49, 0x49, 0x4e, 0x50, 0x49, 0x4c, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x44,
	0x48, 0x4b, 0x4e, 0x42, 0x4b, 0x50, 0x45, 0x4a, 0x4f, 0x18, 0xc6, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4b, 0x46, 0x42, 0x48, 0x4a, 0x42, 0x47, 0x48, 0x48, 0x44, 0x4f, 0x48, 0x00,
	0x52, 0x0b, 0x44, 0x44, 0x48, 0x4b, 0x4e, 0x42, 0x4b, 0x50, 0x45, 0x4a, 0x4f, 0x42, 0x0d, 0x0a,
	0x0b, 0x43, 0x44, 0x42, 0x48, 0x44, 0x44, 0x44, 0x42, 0x47, 0x49, 0x50, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdatePsnSettingsInfoCsReq_proto_rawDescOnce sync.Once
	file_UpdatePsnSettingsInfoCsReq_proto_rawDescData = file_UpdatePsnSettingsInfoCsReq_proto_rawDesc
)

func file_UpdatePsnSettingsInfoCsReq_proto_rawDescGZIP() []byte {
	file_UpdatePsnSettingsInfoCsReq_proto_rawDescOnce.Do(func() {
		file_UpdatePsnSettingsInfoCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdatePsnSettingsInfoCsReq_proto_rawDescData)
	})
	return file_UpdatePsnSettingsInfoCsReq_proto_rawDescData
}

var file_UpdatePsnSettingsInfoCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdatePsnSettingsInfoCsReq_proto_goTypes = []interface{}{
	(*UpdatePsnSettingsInfoCsReq)(nil), // 0: UpdatePsnSettingsInfoCsReq
	(*EKJOBFKOMIE)(nil),                // 1: EKJOBFKOMIE
	(*KFBHJBGHHDO)(nil),                // 2: KFBHJBGHHDO
}
var file_UpdatePsnSettingsInfoCsReq_proto_depIdxs = []int32{
	1, // 0: UpdatePsnSettingsInfoCsReq.PNJCIINPILL:type_name -> EKJOBFKOMIE
	2, // 1: UpdatePsnSettingsInfoCsReq.DDHKNBKPEJO:type_name -> KFBHJBGHHDO
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_UpdatePsnSettingsInfoCsReq_proto_init() }
func file_UpdatePsnSettingsInfoCsReq_proto_init() {
	if File_UpdatePsnSettingsInfoCsReq_proto != nil {
		return
	}
	file_KFBHJBGHHDO_proto_init()
	file_EKJOBFKOMIE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpdatePsnSettingsInfoCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePsnSettingsInfoCsReq); i {
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
	file_UpdatePsnSettingsInfoCsReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UpdatePsnSettingsInfoCsReq_PNJCIINPILL)(nil),
		(*UpdatePsnSettingsInfoCsReq_DDHKNBKPEJO)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_UpdatePsnSettingsInfoCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdatePsnSettingsInfoCsReq_proto_goTypes,
		DependencyIndexes: file_UpdatePsnSettingsInfoCsReq_proto_depIdxs,
		MessageInfos:      file_UpdatePsnSettingsInfoCsReq_proto_msgTypes,
	}.Build()
	File_UpdatePsnSettingsInfoCsReq_proto = out.File
	file_UpdatePsnSettingsInfoCsReq_proto_rawDesc = nil
	file_UpdatePsnSettingsInfoCsReq_proto_goTypes = nil
	file_UpdatePsnSettingsInfoCsReq_proto_depIdxs = nil
}
