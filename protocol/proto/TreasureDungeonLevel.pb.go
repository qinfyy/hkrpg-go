// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonLevel.proto

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

type TreasureDungeonLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BKFEABHCNKF uint32                       `protobuf:"varint,5,opt,name=BKFEABHCNKF,proto3" json:"BKFEABHCNKF,omitempty"`
	BOEFMHFDNCM []*TreasureDungeonRecordData `protobuf:"bytes,7,rep,name=BOEFMHFDNCM,proto3" json:"BOEFMHFDNCM,omitempty"`
	DKBCFCPKELN bool                         `protobuf:"varint,1449,opt,name=DKBCFCPKELN,proto3" json:"DKBCFCPKELN,omitempty"`
	AvatarList  []*OMAGJFAGNLH               `protobuf:"bytes,262,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	OFEMOANLCIO uint32                       `protobuf:"varint,6,opt,name=OFEMOANLCIO,proto3" json:"OFEMOANLCIO,omitempty"`
	ItemList    []*KBJKOKHJOGF               `protobuf:"bytes,1943,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	JIEHDHEANNK uint32                       `protobuf:"varint,8,opt,name=JIEHDHEANNK,proto3" json:"JIEHDHEANNK,omitempty"`
	PFMACGECCAG []*HEEENFFOBLE               `protobuf:"bytes,1,rep,name=PFMACGECCAG,proto3" json:"PFMACGECCAG,omitempty"`
	NBHGGLCOJCD uint32                       `protobuf:"varint,1227,opt,name=NBHGGLCOJCD,proto3" json:"NBHGGLCOJCD,omitempty"`
	LNGPLFNBEDN bool                         `protobuf:"varint,1590,opt,name=LNGPLFNBEDN,proto3" json:"LNGPLFNBEDN,omitempty"`
	BuffList    []*EELOILJMJKN               `protobuf:"bytes,1448,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	LJJDNLFCPMF bool                         `protobuf:"varint,708,opt,name=LJJDNLFCPMF,proto3" json:"LJJDNLFCPMF,omitempty"`
	BFAOFHNBPNA []*NFJAKOCBDCP               `protobuf:"bytes,955,rep,name=BFAOFHNBPNA,proto3" json:"BFAOFHNBPNA,omitempty"`
	PGOHBDFNLAO uint32                       `protobuf:"varint,9,opt,name=PGOHBDFNLAO,proto3" json:"PGOHBDFNLAO,omitempty"`
	MapId       uint32                       `protobuf:"varint,3,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	BDOEPLHLHNL []*OMAGJFAGNLH               `protobuf:"bytes,387,rep,name=BDOEPLHLHNL,proto3" json:"BDOEPLHLHNL,omitempty"`
	MBJGCOMPGHP uint32                       `protobuf:"varint,14,opt,name=MBJGCOMPGHP,proto3" json:"MBJGCOMPGHP,omitempty"`
	BFHEILPBKNA uint32                       `protobuf:"varint,10,opt,name=BFHEILPBKNA,proto3" json:"BFHEILPBKNA,omitempty"`
}

func (x *TreasureDungeonLevel) Reset() {
	*x = TreasureDungeonLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureDungeonLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureDungeonLevel) ProtoMessage() {}

func (x *TreasureDungeonLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureDungeonLevel.ProtoReflect.Descriptor instead.
func (*TreasureDungeonLevel) Descriptor() ([]byte, []int) {
	return file_TreasureDungeonLevel_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureDungeonLevel) GetBKFEABHCNKF() uint32 {
	if x != nil {
		return x.BKFEABHCNKF
	}
	return 0
}

func (x *TreasureDungeonLevel) GetBOEFMHFDNCM() []*TreasureDungeonRecordData {
	if x != nil {
		return x.BOEFMHFDNCM
	}
	return nil
}

func (x *TreasureDungeonLevel) GetDKBCFCPKELN() bool {
	if x != nil {
		return x.DKBCFCPKELN
	}
	return false
}

func (x *TreasureDungeonLevel) GetAvatarList() []*OMAGJFAGNLH {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetOFEMOANLCIO() uint32 {
	if x != nil {
		return x.OFEMOANLCIO
	}
	return 0
}

func (x *TreasureDungeonLevel) GetItemList() []*KBJKOKHJOGF {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetJIEHDHEANNK() uint32 {
	if x != nil {
		return x.JIEHDHEANNK
	}
	return 0
}

func (x *TreasureDungeonLevel) GetPFMACGECCAG() []*HEEENFFOBLE {
	if x != nil {
		return x.PFMACGECCAG
	}
	return nil
}

func (x *TreasureDungeonLevel) GetNBHGGLCOJCD() uint32 {
	if x != nil {
		return x.NBHGGLCOJCD
	}
	return 0
}

func (x *TreasureDungeonLevel) GetLNGPLFNBEDN() bool {
	if x != nil {
		return x.LNGPLFNBEDN
	}
	return false
}

func (x *TreasureDungeonLevel) GetBuffList() []*EELOILJMJKN {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetLJJDNLFCPMF() bool {
	if x != nil {
		return x.LJJDNLFCPMF
	}
	return false
}

func (x *TreasureDungeonLevel) GetBFAOFHNBPNA() []*NFJAKOCBDCP {
	if x != nil {
		return x.BFAOFHNBPNA
	}
	return nil
}

func (x *TreasureDungeonLevel) GetPGOHBDFNLAO() uint32 {
	if x != nil {
		return x.PGOHBDFNLAO
	}
	return 0
}

func (x *TreasureDungeonLevel) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *TreasureDungeonLevel) GetBDOEPLHLHNL() []*OMAGJFAGNLH {
	if x != nil {
		return x.BDOEPLHLHNL
	}
	return nil
}

func (x *TreasureDungeonLevel) GetMBJGCOMPGHP() uint32 {
	if x != nil {
		return x.MBJGCOMPGHP
	}
	return 0
}

func (x *TreasureDungeonLevel) GetBFHEILPBKNA() uint32 {
	if x != nil {
		return x.BFHEILPBKNA
	}
	return 0
}

var File_TreasureDungeonLevel_proto protoreflect.FileDescriptor

var file_TreasureDungeonLevel_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x45,
	0x45, 0x45, 0x4e, 0x46, 0x46, 0x4f, 0x42, 0x4c, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4e, 0x46, 0x4a, 0x41, 0x4b, 0x4f, 0x43, 0x42, 0x44, 0x43, 0x50, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4d, 0x41, 0x47, 0x4a, 0x46, 0x41, 0x47, 0x4e, 0x4c, 0x48, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x45, 0x4c, 0x4f, 0x49, 0x4c, 0x4a, 0x4d, 0x4a,
	0x4b, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x42, 0x4a, 0x4b, 0x4f, 0x4b,
	0x48, 0x4a, 0x4f, 0x47, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a,
	0x14, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4b, 0x46, 0x45, 0x41, 0x42, 0x48,
	0x43, 0x4e, 0x4b, 0x46, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4b, 0x46, 0x45,
	0x41, 0x42, 0x48, 0x43, 0x4e, 0x4b, 0x46, 0x12, 0x3c, 0x0a, 0x0b, 0x42, 0x4f, 0x45, 0x46, 0x4d,
	0x48, 0x46, 0x44, 0x4e, 0x43, 0x4d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x42, 0x4f, 0x45, 0x46, 0x4d, 0x48,
	0x46, 0x44, 0x4e, 0x43, 0x4d, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x4b, 0x42, 0x43, 0x46, 0x43, 0x50,
	0x4b, 0x45, 0x4c, 0x4e, 0x18, 0xa9, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4b, 0x42,
	0x43, 0x46, 0x43, 0x50, 0x4b, 0x45, 0x4c, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x86, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4f, 0x4d, 0x41, 0x47, 0x4a, 0x46, 0x41, 0x47, 0x4e, 0x4c, 0x48, 0x52, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x46, 0x45, 0x4d,
	0x4f, 0x41, 0x4e, 0x4c, 0x43, 0x49, 0x4f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f,
	0x46, 0x45, 0x4d, 0x4f, 0x41, 0x4e, 0x4c, 0x43, 0x49, 0x4f, 0x12, 0x2a, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x97, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4b, 0x42, 0x4a, 0x4b, 0x4f, 0x4b, 0x48, 0x4a, 0x4f, 0x47, 0x46, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x45, 0x48, 0x44, 0x48,
	0x45, 0x41, 0x4e, 0x4e, 0x4b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x49, 0x45,
	0x48, 0x44, 0x48, 0x45, 0x41, 0x4e, 0x4e, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x46, 0x4d, 0x41,
	0x43, 0x47, 0x45, 0x43, 0x43, 0x41, 0x47, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x48, 0x45, 0x45, 0x45, 0x4e, 0x46, 0x46, 0x4f, 0x42, 0x4c, 0x45, 0x52, 0x0b, 0x50, 0x46, 0x4d,
	0x41, 0x43, 0x47, 0x45, 0x43, 0x43, 0x41, 0x47, 0x12, 0x21, 0x0a, 0x0b, 0x4e, 0x42, 0x48, 0x47,
	0x47, 0x4c, 0x43, 0x4f, 0x4a, 0x43, 0x44, 0x18, 0xcb, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4e, 0x42, 0x48, 0x47, 0x47, 0x4c, 0x43, 0x4f, 0x4a, 0x43, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x4c,
	0x4e, 0x47, 0x50, 0x4c, 0x46, 0x4e, 0x42, 0x45, 0x44, 0x4e, 0x18, 0xb6, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x4c, 0x4e, 0x47, 0x50, 0x4c, 0x46, 0x4e, 0x42, 0x45, 0x44, 0x4e, 0x12, 0x2a,
	0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xa8, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x45, 0x4c, 0x4f, 0x49, 0x4c, 0x4a, 0x4d, 0x4a, 0x4b, 0x4e,
	0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x4a,
	0x4a, 0x44, 0x4e, 0x4c, 0x46, 0x43, 0x50, 0x4d, 0x46, 0x18, 0xc4, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4c, 0x4a, 0x4a, 0x44, 0x4e, 0x4c, 0x46, 0x43, 0x50, 0x4d, 0x46, 0x12, 0x2f, 0x0a,
	0x0b, 0x42, 0x46, 0x41, 0x4f, 0x46, 0x48, 0x4e, 0x42, 0x50, 0x4e, 0x41, 0x18, 0xbb, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x46, 0x4a, 0x41, 0x4b, 0x4f, 0x43, 0x42, 0x44, 0x43,
	0x50, 0x52, 0x0b, 0x42, 0x46, 0x41, 0x4f, 0x46, 0x48, 0x4e, 0x42, 0x50, 0x4e, 0x41, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x47, 0x4f, 0x48, 0x42, 0x44, 0x46, 0x4e, 0x4c, 0x41, 0x4f, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x47, 0x4f, 0x48, 0x42, 0x44, 0x46, 0x4e, 0x4c, 0x41, 0x4f,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x42, 0x44, 0x4f, 0x45, 0x50,
	0x4c, 0x48, 0x4c, 0x48, 0x4e, 0x4c, 0x18, 0x83, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4f, 0x4d, 0x41, 0x47, 0x4a, 0x46, 0x41, 0x47, 0x4e, 0x4c, 0x48, 0x52, 0x0b, 0x42, 0x44, 0x4f,
	0x45, 0x50, 0x4c, 0x48, 0x4c, 0x48, 0x4e, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x42, 0x4a, 0x47,
	0x43, 0x4f, 0x4d, 0x50, 0x47, 0x48, 0x50, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d,
	0x42, 0x4a, 0x47, 0x43, 0x4f, 0x4d, 0x50, 0x47, 0x48, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46,
	0x48, 0x45, 0x49, 0x4c, 0x50, 0x42, 0x4b, 0x4e, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x46, 0x48, 0x45, 0x49, 0x4c, 0x50, 0x42, 0x4b, 0x4e, 0x41, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonLevel_proto_rawDescOnce sync.Once
	file_TreasureDungeonLevel_proto_rawDescData = file_TreasureDungeonLevel_proto_rawDesc
)

func file_TreasureDungeonLevel_proto_rawDescGZIP() []byte {
	file_TreasureDungeonLevel_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonLevel_proto_rawDescData)
	})
	return file_TreasureDungeonLevel_proto_rawDescData
}

var file_TreasureDungeonLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureDungeonLevel_proto_goTypes = []interface{}{
	(*TreasureDungeonLevel)(nil),      // 0: TreasureDungeonLevel
	(*TreasureDungeonRecordData)(nil), // 1: TreasureDungeonRecordData
	(*OMAGJFAGNLH)(nil),               // 2: OMAGJFAGNLH
	(*KBJKOKHJOGF)(nil),               // 3: KBJKOKHJOGF
	(*HEEENFFOBLE)(nil),               // 4: HEEENFFOBLE
	(*EELOILJMJKN)(nil),               // 5: EELOILJMJKN
	(*NFJAKOCBDCP)(nil),               // 6: NFJAKOCBDCP
}
var file_TreasureDungeonLevel_proto_depIdxs = []int32{
	1, // 0: TreasureDungeonLevel.BOEFMHFDNCM:type_name -> TreasureDungeonRecordData
	2, // 1: TreasureDungeonLevel.avatar_list:type_name -> OMAGJFAGNLH
	3, // 2: TreasureDungeonLevel.item_list:type_name -> KBJKOKHJOGF
	4, // 3: TreasureDungeonLevel.PFMACGECCAG:type_name -> HEEENFFOBLE
	5, // 4: TreasureDungeonLevel.buff_list:type_name -> EELOILJMJKN
	6, // 5: TreasureDungeonLevel.BFAOFHNBPNA:type_name -> NFJAKOCBDCP
	2, // 6: TreasureDungeonLevel.BDOEPLHLHNL:type_name -> OMAGJFAGNLH
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_TreasureDungeonLevel_proto_init() }
func file_TreasureDungeonLevel_proto_init() {
	if File_TreasureDungeonLevel_proto != nil {
		return
	}
	file_HEEENFFOBLE_proto_init()
	file_NFJAKOCBDCP_proto_init()
	file_OMAGJFAGNLH_proto_init()
	file_EELOILJMJKN_proto_init()
	file_KBJKOKHJOGF_proto_init()
	file_TreasureDungeonRecordData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TreasureDungeonLevel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureDungeonLevel); i {
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
			RawDescriptor: file_TreasureDungeonLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonLevel_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonLevel_proto_depIdxs,
		MessageInfos:      file_TreasureDungeonLevel_proto_msgTypes,
	}.Build()
	File_TreasureDungeonLevel_proto = out.File
	file_TreasureDungeonLevel_proto_rawDesc = nil
	file_TreasureDungeonLevel_proto_goTypes = nil
	file_TreasureDungeonLevel_proto_depIdxs = nil
}
