// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JHPGIPFHPJM.proto

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

type JHPGIPFHPJM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JKBOLMLOEOE LobbyCharacterType   `protobuf:"varint,1,opt,name=JKBOLMLOEOE,proto3,enum=LobbyCharacterType" json:"JKBOLMLOEOE,omitempty"`
	Status      LobbyCharacterStatus `protobuf:"varint,2,opt,name=status,proto3,enum=LobbyCharacterStatus" json:"status,omitempty"`
}

func (x *JHPGIPFHPJM) Reset() {
	*x = JHPGIPFHPJM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JHPGIPFHPJM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JHPGIPFHPJM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JHPGIPFHPJM) ProtoMessage() {}

func (x *JHPGIPFHPJM) ProtoReflect() protoreflect.Message {
	mi := &file_JHPGIPFHPJM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JHPGIPFHPJM.ProtoReflect.Descriptor instead.
func (*JHPGIPFHPJM) Descriptor() ([]byte, []int) {
	return file_JHPGIPFHPJM_proto_rawDescGZIP(), []int{0}
}

func (x *JHPGIPFHPJM) GetJKBOLMLOEOE() LobbyCharacterType {
	if x != nil {
		return x.JKBOLMLOEOE
	}
	return LobbyCharacterType_LobbyCharacter_None
}

func (x *JHPGIPFHPJM) GetStatus() LobbyCharacterStatus {
	if x != nil {
		return x.Status
	}
	return LobbyCharacterStatus_LobbyCharacterStatus_None
}

var File_JHPGIPFHPJM_proto protoreflect.FileDescriptor

var file_JHPGIPFHPJM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x48, 0x50, 0x47, 0x49, 0x50, 0x46, 0x48, 0x50, 0x4a, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0b, 0x4a, 0x48, 0x50,
	0x47, 0x49, 0x50, 0x46, 0x48, 0x50, 0x4a, 0x4d, 0x12, 0x35, 0x0a, 0x0b, 0x4a, 0x4b, 0x42, 0x4f,
	0x4c, 0x4d, 0x4c, 0x4f, 0x45, 0x4f, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x4a, 0x4b, 0x42, 0x4f, 0x4c, 0x4d, 0x4c, 0x4f, 0x45, 0x4f, 0x45, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JHPGIPFHPJM_proto_rawDescOnce sync.Once
	file_JHPGIPFHPJM_proto_rawDescData = file_JHPGIPFHPJM_proto_rawDesc
)

func file_JHPGIPFHPJM_proto_rawDescGZIP() []byte {
	file_JHPGIPFHPJM_proto_rawDescOnce.Do(func() {
		file_JHPGIPFHPJM_proto_rawDescData = protoimpl.X.CompressGZIP(file_JHPGIPFHPJM_proto_rawDescData)
	})
	return file_JHPGIPFHPJM_proto_rawDescData
}

var file_JHPGIPFHPJM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JHPGIPFHPJM_proto_goTypes = []interface{}{
	(*JHPGIPFHPJM)(nil),       // 0: JHPGIPFHPJM
	(LobbyCharacterType)(0),   // 1: LobbyCharacterType
	(LobbyCharacterStatus)(0), // 2: LobbyCharacterStatus
}
var file_JHPGIPFHPJM_proto_depIdxs = []int32{
	1, // 0: JHPGIPFHPJM.JKBOLMLOEOE:type_name -> LobbyCharacterType
	2, // 1: JHPGIPFHPJM.status:type_name -> LobbyCharacterStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_JHPGIPFHPJM_proto_init() }
func file_JHPGIPFHPJM_proto_init() {
	if File_JHPGIPFHPJM_proto != nil {
		return
	}
	file_LobbyCharacterStatus_proto_init()
	file_LobbyCharacterType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JHPGIPFHPJM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JHPGIPFHPJM); i {
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
			RawDescriptor: file_JHPGIPFHPJM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JHPGIPFHPJM_proto_goTypes,
		DependencyIndexes: file_JHPGIPFHPJM_proto_depIdxs,
		MessageInfos:      file_JHPGIPFHPJM_proto_msgTypes,
	}.Build()
	File_JHPGIPFHPJM_proto = out.File
	file_JHPGIPFHPJM_proto_rawDesc = nil
	file_JHPGIPFHPJM_proto_goTypes = nil
	file_JHPGIPFHPJM_proto_depIdxs = nil
}
