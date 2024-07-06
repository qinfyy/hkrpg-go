// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetSceneMapInfoCsReq.proto

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

type GetSceneMapInfoCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CGIIIEKIDJE uint32   `protobuf:"varint,4,opt,name=CGIIIEKIDJE,proto3" json:"CGIIIEKIDJE,omitempty"`
	EntryId     uint32   `protobuf:"varint,9,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	EntryIdList []uint32 `protobuf:"varint,2,rep,packed,name=entry_id_list,json=entryIdList,proto3" json:"entry_id_list,omitempty"`
	BDKGPGBNFNN bool     `protobuf:"varint,7,opt,name=BDKGPGBNFNN,proto3" json:"BDKGPGBNFNN,omitempty"`
}

func (x *GetSceneMapInfoCsReq) Reset() {
	*x = GetSceneMapInfoCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneMapInfoCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneMapInfoCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneMapInfoCsReq) ProtoMessage() {}

func (x *GetSceneMapInfoCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneMapInfoCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneMapInfoCsReq.ProtoReflect.Descriptor instead.
func (*GetSceneMapInfoCsReq) Descriptor() ([]byte, []int) {
	return file_GetSceneMapInfoCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneMapInfoCsReq) GetCGIIIEKIDJE() uint32 {
	if x != nil {
		return x.CGIIIEKIDJE
	}
	return 0
}

func (x *GetSceneMapInfoCsReq) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *GetSceneMapInfoCsReq) GetEntryIdList() []uint32 {
	if x != nil {
		return x.EntryIdList
	}
	return nil
}

func (x *GetSceneMapInfoCsReq) GetBDKGPGBNFNN() bool {
	if x != nil {
		return x.BDKGPGBNFNN
	}
	return false
}

var File_GetSceneMapInfoCsReq_proto protoreflect.FileDescriptor

var file_GetSceneMapInfoCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x47, 0x49, 0x49, 0x49, 0x45, 0x4b,
	0x49, 0x44, 0x4a, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x47, 0x49, 0x49,
	0x49, 0x45, 0x4b, 0x49, 0x44, 0x4a, 0x45, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44, 0x4b, 0x47, 0x50, 0x47,
	0x42, 0x4e, 0x46, 0x4e, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x44, 0x4b,
	0x47, 0x50, 0x47, 0x42, 0x4e, 0x46, 0x4e, 0x4e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSceneMapInfoCsReq_proto_rawDescOnce sync.Once
	file_GetSceneMapInfoCsReq_proto_rawDescData = file_GetSceneMapInfoCsReq_proto_rawDesc
)

func file_GetSceneMapInfoCsReq_proto_rawDescGZIP() []byte {
	file_GetSceneMapInfoCsReq_proto_rawDescOnce.Do(func() {
		file_GetSceneMapInfoCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneMapInfoCsReq_proto_rawDescData)
	})
	return file_GetSceneMapInfoCsReq_proto_rawDescData
}

var file_GetSceneMapInfoCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneMapInfoCsReq_proto_goTypes = []interface{}{
	(*GetSceneMapInfoCsReq)(nil), // 0: GetSceneMapInfoCsReq
}
var file_GetSceneMapInfoCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetSceneMapInfoCsReq_proto_init() }
func file_GetSceneMapInfoCsReq_proto_init() {
	if File_GetSceneMapInfoCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetSceneMapInfoCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneMapInfoCsReq); i {
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
			RawDescriptor: file_GetSceneMapInfoCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneMapInfoCsReq_proto_goTypes,
		DependencyIndexes: file_GetSceneMapInfoCsReq_proto_depIdxs,
		MessageInfos:      file_GetSceneMapInfoCsReq_proto_msgTypes,
	}.Build()
	File_GetSceneMapInfoCsReq_proto = out.File
	file_GetSceneMapInfoCsReq_proto_rawDesc = nil
	file_GetSceneMapInfoCsReq_proto_goTypes = nil
	file_GetSceneMapInfoCsReq_proto_depIdxs = nil
}