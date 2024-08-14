// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: IOBLHLMIDAE.proto

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

type IOBLHLMIDAE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	COENPLNMMOH       *FMBLGECBIBP       `protobuf:"bytes,6,opt,name=COENPLNMMOH,proto3" json:"COENPLNMMOH,omitempty"`
	OGAAFEIKNON       *PECNIJEAIIB       `protobuf:"bytes,15,opt,name=OGAAFEIKNON,proto3" json:"OGAAFEIKNON,omitempty"`
	RogueTournCurInfo *RogueTournCurInfo `protobuf:"bytes,5,opt,name=rogue_tourn_cur_info,json=rogueTournCurInfo,proto3" json:"rogue_tourn_cur_info,omitempty"`
	RogueLineupInfo   *LineupInfo        `protobuf:"bytes,11,opt,name=rogue_lineup_info,json=rogueLineupInfo,proto3" json:"rogue_lineup_info,omitempty"`
	BFIGBKDHIGJ       *CJIAHGOOHAD       `protobuf:"bytes,3,opt,name=BFIGBKDHIGJ,proto3" json:"BFIGBKDHIGJ,omitempty"`
}

func (x *IOBLHLMIDAE) Reset() {
	*x = IOBLHLMIDAE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IOBLHLMIDAE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOBLHLMIDAE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOBLHLMIDAE) ProtoMessage() {}

func (x *IOBLHLMIDAE) ProtoReflect() protoreflect.Message {
	mi := &file_IOBLHLMIDAE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOBLHLMIDAE.ProtoReflect.Descriptor instead.
func (*IOBLHLMIDAE) Descriptor() ([]byte, []int) {
	return file_IOBLHLMIDAE_proto_rawDescGZIP(), []int{0}
}

func (x *IOBLHLMIDAE) GetCOENPLNMMOH() *FMBLGECBIBP {
	if x != nil {
		return x.COENPLNMMOH
	}
	return nil
}

func (x *IOBLHLMIDAE) GetOGAAFEIKNON() *PECNIJEAIIB {
	if x != nil {
		return x.OGAAFEIKNON
	}
	return nil
}

func (x *IOBLHLMIDAE) GetRogueTournCurInfo() *RogueTournCurInfo {
	if x != nil {
		return x.RogueTournCurInfo
	}
	return nil
}

func (x *IOBLHLMIDAE) GetRogueLineupInfo() *LineupInfo {
	if x != nil {
		return x.RogueLineupInfo
	}
	return nil
}

func (x *IOBLHLMIDAE) GetBFIGBKDHIGJ() *CJIAHGOOHAD {
	if x != nil {
		return x.BFIGBKDHIGJ
	}
	return nil
}

var File_IOBLHLMIDAE_proto protoreflect.FileDescriptor

var file_IOBLHLMIDAE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x4f, 0x42, 0x4c, 0x48, 0x4c, 0x4d, 0x49, 0x44, 0x41, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4a, 0x49, 0x41, 0x48, 0x47, 0x4f, 0x4f, 0x48, 0x41, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4d, 0x42, 0x4c, 0x47, 0x45, 0x43, 0x42,
	0x49, 0x42, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x45, 0x43, 0x4e, 0x49, 0x4a, 0x45, 0x41, 0x49, 0x49,
	0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0b, 0x49, 0x4f, 0x42, 0x4c,
	0x48, 0x4c, 0x4d, 0x49, 0x44, 0x41, 0x45, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x4f, 0x45, 0x4e, 0x50,
	0x4c, 0x4e, 0x4d, 0x4d, 0x4f, 0x48, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46,
	0x4d, 0x42, 0x4c, 0x47, 0x45, 0x43, 0x42, 0x49, 0x42, 0x50, 0x52, 0x0b, 0x43, 0x4f, 0x45, 0x4e,
	0x50, 0x4c, 0x4e, 0x4d, 0x4d, 0x4f, 0x48, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x47, 0x41, 0x41, 0x46,
	0x45, 0x49, 0x4b, 0x4e, 0x4f, 0x4e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50,
	0x45, 0x43, 0x4e, 0x49, 0x4a, 0x45, 0x41, 0x49, 0x49, 0x42, 0x52, 0x0b, 0x4f, 0x47, 0x41, 0x41,
	0x46, 0x45, 0x49, 0x4b, 0x4e, 0x4f, 0x4e, 0x12, 0x43, 0x0a, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x11,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x46, 0x49, 0x47, 0x42, 0x4b, 0x44,
	0x48, 0x49, 0x47, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4a, 0x49,
	0x41, 0x48, 0x47, 0x4f, 0x4f, 0x48, 0x41, 0x44, 0x52, 0x0b, 0x42, 0x46, 0x49, 0x47, 0x42, 0x4b,
	0x44, 0x48, 0x49, 0x47, 0x4a, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IOBLHLMIDAE_proto_rawDescOnce sync.Once
	file_IOBLHLMIDAE_proto_rawDescData = file_IOBLHLMIDAE_proto_rawDesc
)

func file_IOBLHLMIDAE_proto_rawDescGZIP() []byte {
	file_IOBLHLMIDAE_proto_rawDescOnce.Do(func() {
		file_IOBLHLMIDAE_proto_rawDescData = protoimpl.X.CompressGZIP(file_IOBLHLMIDAE_proto_rawDescData)
	})
	return file_IOBLHLMIDAE_proto_rawDescData
}

var file_IOBLHLMIDAE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IOBLHLMIDAE_proto_goTypes = []interface{}{
	(*IOBLHLMIDAE)(nil),       // 0: IOBLHLMIDAE
	(*FMBLGECBIBP)(nil),       // 1: FMBLGECBIBP
	(*PECNIJEAIIB)(nil),       // 2: PECNIJEAIIB
	(*RogueTournCurInfo)(nil), // 3: RogueTournCurInfo
	(*LineupInfo)(nil),        // 4: LineupInfo
	(*CJIAHGOOHAD)(nil),       // 5: CJIAHGOOHAD
}
var file_IOBLHLMIDAE_proto_depIdxs = []int32{
	1, // 0: IOBLHLMIDAE.COENPLNMMOH:type_name -> FMBLGECBIBP
	2, // 1: IOBLHLMIDAE.OGAAFEIKNON:type_name -> PECNIJEAIIB
	3, // 2: IOBLHLMIDAE.rogue_tourn_cur_info:type_name -> RogueTournCurInfo
	4, // 3: IOBLHLMIDAE.rogue_lineup_info:type_name -> LineupInfo
	5, // 4: IOBLHLMIDAE.BFIGBKDHIGJ:type_name -> CJIAHGOOHAD
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_IOBLHLMIDAE_proto_init() }
func file_IOBLHLMIDAE_proto_init() {
	if File_IOBLHLMIDAE_proto != nil {
		return
	}
	file_CJIAHGOOHAD_proto_init()
	file_FMBLGECBIBP_proto_init()
	file_LineupInfo_proto_init()
	file_RogueTournCurInfo_proto_init()
	file_PECNIJEAIIB_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IOBLHLMIDAE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOBLHLMIDAE); i {
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
			RawDescriptor: file_IOBLHLMIDAE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IOBLHLMIDAE_proto_goTypes,
		DependencyIndexes: file_IOBLHLMIDAE_proto_depIdxs,
		MessageInfos:      file_IOBLHLMIDAE_proto_msgTypes,
	}.Build()
	File_IOBLHLMIDAE_proto = out.File
	file_IOBLHLMIDAE_proto_rawDesc = nil
	file_IOBLHLMIDAE_proto_goTypes = nil
	file_IOBLHLMIDAE_proto_depIdxs = nil
}