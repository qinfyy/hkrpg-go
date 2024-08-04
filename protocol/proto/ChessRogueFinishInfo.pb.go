// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueFinishInfo.proto

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

type ChessRogueFinishInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuitReason            ChessRogueQuitReason `protobuf:"varint,14,opt,name=quit_reason,json=quitReason,proto3,enum=ChessRogueQuitReason" json:"quit_reason,omitempty"`
	LEBECEFMDJN           uint32               `protobuf:"varint,1129,opt,name=LEBECEFMDJN,proto3" json:"LEBECEFMDJN,omitempty"`
	RogueLineup           *LineupInfo          `protobuf:"bytes,10,opt,name=rogue_lineup,json=rogueLineup,proto3" json:"rogue_lineup,omitempty"`
	RogueBuffInfo         *ChessRogueBuff      `protobuf:"bytes,7,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	RogueSubMode          uint32               `protobuf:"varint,862,opt,name=rogue_sub_mode,json=rogueSubMode,proto3" json:"rogue_sub_mode,omitempty"`
	HHMFIDKFNNI           []uint32             `protobuf:"varint,5,rep,packed,name=HHMFIDKFNNI,proto3" json:"HHMFIDKFNNI,omitempty"`
	EndAreaId             uint32               `protobuf:"varint,1,opt,name=end_area_id,json=endAreaId,proto3" json:"end_area_id,omitempty"`
	CNCAOLEDBDI           uint32               `protobuf:"varint,1705,opt,name=CNCAOLEDBDI,proto3" json:"CNCAOLEDBDI,omitempty"`
	LastLayerId           uint32               `protobuf:"varint,2,opt,name=last_layer_id,json=lastLayerId,proto3" json:"last_layer_id,omitempty"`
	EPGJCMNBIPJ           *ItemList            `protobuf:"bytes,3,opt,name=EPGJCMNBIPJ,proto3" json:"EPGJCMNBIPJ,omitempty"`
	DifficultyLevel       uint32               `protobuf:"varint,11,opt,name=difficulty_level,json=difficultyLevel,proto3" json:"difficulty_level,omitempty"`
	ScoreId               uint32               `protobuf:"varint,468,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	GameMiracleInfo       *ChessRogueMiracle   `protobuf:"bytes,12,opt,name=game_miracle_info,json=gameMiracleInfo,proto3" json:"game_miracle_info,omitempty"`
	DODPBNFKKEL           uint32               `protobuf:"varint,6,opt,name=DODPBNFKKEL,proto3" json:"DODPBNFKKEL,omitempty"`
	OLMBPLAIMLP           uint32               `protobuf:"varint,15,opt,name=OLMBPLAIMLP,proto3" json:"OLMBPLAIMLP,omitempty"`
	BIOHIBDDDFG           uint32               `protobuf:"varint,912,opt,name=BIOHIBDDDFG,proto3" json:"BIOHIBDDDFG,omitempty"`
	IsFinish              bool                 `protobuf:"varint,4,opt,name=is_finish,json=isFinish,proto3" json:"is_finish,omitempty"`
	NKGKDMFHGFJ           *JPFECHLHHEN         `protobuf:"bytes,1250,opt,name=NKGKDMFHGFJ,proto3" json:"NKGKDMFHGFJ,omitempty"`
	ChessRogueMainStoryId uint32               `protobuf:"varint,9,opt,name=chess_rogue_main_story_id,json=chessRogueMainStoryId,proto3" json:"chess_rogue_main_story_id,omitempty"`
}

func (x *ChessRogueFinishInfo) Reset() {
	*x = ChessRogueFinishInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueFinishInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueFinishInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueFinishInfo) ProtoMessage() {}

func (x *ChessRogueFinishInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueFinishInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueFinishInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueFinishInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueFinishInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueFinishInfo) GetQuitReason() ChessRogueQuitReason {
	if x != nil {
		return x.QuitReason
	}
	return ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_NONE
}

func (x *ChessRogueFinishInfo) GetLEBECEFMDJN() uint32 {
	if x != nil {
		return x.LEBECEFMDJN
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetRogueLineup() *LineupInfo {
	if x != nil {
		return x.RogueLineup
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetRogueBuffInfo() *ChessRogueBuff {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetRogueSubMode() uint32 {
	if x != nil {
		return x.RogueSubMode
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetHHMFIDKFNNI() []uint32 {
	if x != nil {
		return x.HHMFIDKFNNI
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetEndAreaId() uint32 {
	if x != nil {
		return x.EndAreaId
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetCNCAOLEDBDI() uint32 {
	if x != nil {
		return x.CNCAOLEDBDI
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetLastLayerId() uint32 {
	if x != nil {
		return x.LastLayerId
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetEPGJCMNBIPJ() *ItemList {
	if x != nil {
		return x.EPGJCMNBIPJ
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetDifficultyLevel() uint32 {
	if x != nil {
		return x.DifficultyLevel
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetGameMiracleInfo() *ChessRogueMiracle {
	if x != nil {
		return x.GameMiracleInfo
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetDODPBNFKKEL() uint32 {
	if x != nil {
		return x.DODPBNFKKEL
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetOLMBPLAIMLP() uint32 {
	if x != nil {
		return x.OLMBPLAIMLP
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetBIOHIBDDDFG() uint32 {
	if x != nil {
		return x.BIOHIBDDDFG
	}
	return 0
}

func (x *ChessRogueFinishInfo) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

func (x *ChessRogueFinishInfo) GetNKGKDMFHGFJ() *JPFECHLHHEN {
	if x != nil {
		return x.NKGKDMFHGFJ
	}
	return nil
}

func (x *ChessRogueFinishInfo) GetChessRogueMainStoryId() uint32 {
	if x != nil {
		return x.ChessRogueMainStoryId
	}
	return 0
}

var File_ChessRogueFinishInfo_proto protoreflect.FileDescriptor

var file_ChessRogueFinishInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x50, 0x46, 0x45, 0x43, 0x48, 0x4c, 0x48, 0x48, 0x45, 0x4e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x06, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x52, 0x0a, 0x71, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0b, 0x4c, 0x45, 0x42, 0x45, 0x43, 0x45, 0x46, 0x4d, 0x44, 0x4a, 0x4e, 0x18, 0xe9, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x45, 0x42, 0x45, 0x43, 0x45, 0x46, 0x4d, 0x44, 0x4a, 0x4e,
	0x12, 0x2e, 0x0a, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x12, 0x37, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0xde, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x75, 0x62, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x48, 0x48, 0x4d, 0x46, 0x49, 0x44, 0x4b, 0x46, 0x4e, 0x4e, 0x49, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x48, 0x4d, 0x46, 0x49, 0x44, 0x4b, 0x46, 0x4e,
	0x4e, 0x49, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x4e, 0x43, 0x41, 0x4f, 0x4c, 0x45, 0x44, 0x42, 0x44,
	0x49, 0x18, 0xa9, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4e, 0x43, 0x41, 0x4f, 0x4c,
	0x45, 0x44, 0x42, 0x44, 0x49, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b, 0x45, 0x50, 0x47,
	0x4a, 0x43, 0x4d, 0x4e, 0x42, 0x49, 0x50, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x45, 0x50, 0x47, 0x4a, 0x43,
	0x4d, 0x4e, 0x42, 0x49, 0x50, 0x4a, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xd4, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a,
	0x11, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x0f, 0x67, 0x61,
	0x6d, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x4f, 0x44, 0x50, 0x42, 0x4e, 0x46, 0x4b, 0x4b, 0x45, 0x4c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4f, 0x44, 0x50, 0x42, 0x4e, 0x46, 0x4b, 0x4b, 0x45, 0x4c, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4c, 0x4d, 0x42, 0x50, 0x4c, 0x41, 0x49, 0x4d, 0x4c, 0x50, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x4d, 0x42, 0x50, 0x4c, 0x41, 0x49, 0x4d, 0x4c,
	0x50, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x49, 0x4f, 0x48, 0x49, 0x42, 0x44, 0x44, 0x44, 0x46, 0x47,
	0x18, 0x90, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x49, 0x4f, 0x48, 0x49, 0x42, 0x44,
	0x44, 0x44, 0x46, 0x47, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x12, 0x2f, 0x0a, 0x0b, 0x4e, 0x4b, 0x47, 0x4b, 0x44, 0x4d, 0x46, 0x48, 0x47, 0x46, 0x4a,
	0x18, 0xe2, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x50, 0x46, 0x45, 0x43, 0x48,
	0x4c, 0x48, 0x48, 0x45, 0x4e, 0x52, 0x0b, 0x4e, 0x4b, 0x47, 0x4b, 0x44, 0x4d, 0x46, 0x48, 0x47,
	0x46, 0x4a, 0x12, 0x38, 0x0a, 0x19, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueFinishInfo_proto_rawDescOnce sync.Once
	file_ChessRogueFinishInfo_proto_rawDescData = file_ChessRogueFinishInfo_proto_rawDesc
)

func file_ChessRogueFinishInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueFinishInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueFinishInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueFinishInfo_proto_rawDescData)
	})
	return file_ChessRogueFinishInfo_proto_rawDescData
}

var file_ChessRogueFinishInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueFinishInfo_proto_goTypes = []interface{}{
	(*ChessRogueFinishInfo)(nil), // 0: ChessRogueFinishInfo
	(ChessRogueQuitReason)(0),    // 1: ChessRogueQuitReason
	(*LineupInfo)(nil),           // 2: LineupInfo
	(*ChessRogueBuff)(nil),       // 3: ChessRogueBuff
	(*ItemList)(nil),             // 4: ItemList
	(*ChessRogueMiracle)(nil),    // 5: ChessRogueMiracle
	(*JPFECHLHHEN)(nil),          // 6: JPFECHLHHEN
}
var file_ChessRogueFinishInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueFinishInfo.quit_reason:type_name -> ChessRogueQuitReason
	2, // 1: ChessRogueFinishInfo.rogue_lineup:type_name -> LineupInfo
	3, // 2: ChessRogueFinishInfo.rogue_buff_info:type_name -> ChessRogueBuff
	4, // 3: ChessRogueFinishInfo.EPGJCMNBIPJ:type_name -> ItemList
	5, // 4: ChessRogueFinishInfo.game_miracle_info:type_name -> ChessRogueMiracle
	6, // 5: ChessRogueFinishInfo.NKGKDMFHGFJ:type_name -> JPFECHLHHEN
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ChessRogueFinishInfo_proto_init() }
func file_ChessRogueFinishInfo_proto_init() {
	if File_ChessRogueFinishInfo_proto != nil {
		return
	}
	file_ChessRogueBuff_proto_init()
	file_JPFECHLHHEN_proto_init()
	file_ItemList_proto_init()
	file_ChessRogueQuitReason_proto_init()
	file_LineupInfo_proto_init()
	file_ChessRogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueFinishInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueFinishInfo); i {
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
			RawDescriptor: file_ChessRogueFinishInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueFinishInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueFinishInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueFinishInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueFinishInfo_proto = out.File
	file_ChessRogueFinishInfo_proto_rawDesc = nil
	file_ChessRogueFinishInfo_proto_goTypes = nil
	file_ChessRogueFinishInfo_proto_depIdxs = nil
}
