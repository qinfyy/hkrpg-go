// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleRogueCommonPendingActionCsReq.proto

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

type HandleRogueCommonPendingActionCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueLocation uint32 `protobuf:"varint,6,opt,name=queue_location,json=queueLocation,proto3" json:"queue_location,omitempty"`
	// Types that are assignable to Action:
	//
	//	*HandleRogueCommonPendingActionCsReq_BuffSelectResult
	//	*HandleRogueCommonPendingActionCsReq_CLOMAFJKKOP
	//	*HandleRogueCommonPendingActionCsReq_EBFFLNAPJHJ
	//	*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult
	//	*HandleRogueCommonPendingActionCsReq_NLHEFFPICMN
	//	*HandleRogueCommonPendingActionCsReq_MiracleSelectResult
	//	*HandleRogueCommonPendingActionCsReq_GDLGGEDGJKN
	//	*HandleRogueCommonPendingActionCsReq_FKOJLMBOKGJ
	//	*HandleRogueCommonPendingActionCsReq_MLDGGJIBFCO
	//	*HandleRogueCommonPendingActionCsReq_LDKODDLEAKP
	//	*HandleRogueCommonPendingActionCsReq_JLMOPLPFPOP
	//	*HandleRogueCommonPendingActionCsReq_KPAHINHGNJI
	//	*HandleRogueCommonPendingActionCsReq_BonusSelectResult
	//	*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult
	//	*HandleRogueCommonPendingActionCsReq_GLLFBCGDMJA
	//	*HandleRogueCommonPendingActionCsReq_KEPJDAKKDOJ
	//	*HandleRogueCommonPendingActionCsReq_IOJKNOLKFLD
	//	*HandleRogueCommonPendingActionCsReq_NANBLBHDPPG
	Action isHandleRogueCommonPendingActionCsReq_Action `protobuf_oneof:"action"`
}

func (x *HandleRogueCommonPendingActionCsReq) Reset() {
	*x = HandleRogueCommonPendingActionCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleRogueCommonPendingActionCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleRogueCommonPendingActionCsReq) ProtoMessage() {}

func (x *HandleRogueCommonPendingActionCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleRogueCommonPendingActionCsReq.ProtoReflect.Descriptor instead.
func (*HandleRogueCommonPendingActionCsReq) Descriptor() ([]byte, []int) {
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *HandleRogueCommonPendingActionCsReq) GetQueueLocation() uint32 {
	if x != nil {
		return x.QueueLocation
	}
	return 0
}

func (m *HandleRogueCommonPendingActionCsReq) GetAction() isHandleRogueCommonPendingActionCsReq_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffSelectResult() *RogueBuffSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BuffSelectResult); ok {
		return x.BuffSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetCLOMAFJKKOP() *FCGNPINDCJO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_CLOMAFJKKOP); ok {
		return x.CLOMAFJKKOP
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetEBFFLNAPJHJ() *KKHCONJBMHD {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_EBFFLNAPJHJ); ok {
		return x.EBFFLNAPJHJ
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffRerollSelectResult() *RogueBuffRerollResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult); ok {
		return x.BuffRerollSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetNLHEFFPICMN() *BEKMDCIMHHG {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_NLHEFFPICMN); ok {
		return x.NLHEFFPICMN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMiracleSelectResult() *RogueMiracleSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult); ok {
		return x.MiracleSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetGDLGGEDGJKN() *JILKCLLHDHL {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_GDLGGEDGJKN); ok {
		return x.GDLGGEDGJKN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetFKOJLMBOKGJ() *GBKOOIIODFC {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_FKOJLMBOKGJ); ok {
		return x.FKOJLMBOKGJ
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMLDGGJIBFCO() *BJMEIIHOINJ {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MLDGGJIBFCO); ok {
		return x.MLDGGJIBFCO
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetLDKODDLEAKP() *DKDLBLJHAMB {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_LDKODDLEAKP); ok {
		return x.LDKODDLEAKP
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetJLMOPLPFPOP() *NLFOBIBCPEO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_JLMOPLPFPOP); ok {
		return x.JLMOPLPFPOP
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetKPAHINHGNJI() *ADMOCCBHGIE {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_KPAHINHGNJI); ok {
		return x.KPAHINHGNJI
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBonusSelectResult() *RogueBonusSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BonusSelectResult); ok {
		return x.BonusSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetRogueTournFormulaResult() *RogueTournFormulaResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult); ok {
		return x.RogueTournFormulaResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetGLLFBCGDMJA() *DELGOOPEKOM {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_GLLFBCGDMJA); ok {
		return x.GLLFBCGDMJA
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetKEPJDAKKDOJ() *EOAHHEAOFFK {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_KEPJDAKKDOJ); ok {
		return x.KEPJDAKKDOJ
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetIOJKNOLKFLD() *OHOPHFCMOKH {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_IOJKNOLKFLD); ok {
		return x.IOJKNOLKFLD
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetNANBLBHDPPG() *IIJJHNNNKCF {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_NANBLBHDPPG); ok {
		return x.NANBLBHDPPG
	}
	return nil
}

type isHandleRogueCommonPendingActionCsReq_Action interface {
	isHandleRogueCommonPendingActionCsReq_Action()
}

type HandleRogueCommonPendingActionCsReq_BuffSelectResult struct {
	BuffSelectResult *RogueBuffSelectResult `protobuf:"bytes,309,opt,name=buff_select_result,json=buffSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_CLOMAFJKKOP struct {
	CLOMAFJKKOP *FCGNPINDCJO `protobuf:"bytes,806,opt,name=CLOMAFJKKOP,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_EBFFLNAPJHJ struct {
	EBFFLNAPJHJ *KKHCONJBMHD `protobuf:"bytes,1291,opt,name=EBFFLNAPJHJ,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult struct {
	BuffRerollSelectResult *RogueBuffRerollResult `protobuf:"bytes,1008,opt,name=buff_reroll_select_result,json=buffRerollSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_NLHEFFPICMN struct {
	NLHEFFPICMN *BEKMDCIMHHG `protobuf:"bytes,1974,opt,name=NLHEFFPICMN,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MiracleSelectResult struct {
	MiracleSelectResult *RogueMiracleSelectResult `protobuf:"bytes,1833,opt,name=miracle_select_result,json=miracleSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_GDLGGEDGJKN struct {
	GDLGGEDGJKN *JILKCLLHDHL `protobuf:"bytes,165,opt,name=GDLGGEDGJKN,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_FKOJLMBOKGJ struct {
	FKOJLMBOKGJ *GBKOOIIODFC `protobuf:"bytes,1639,opt,name=FKOJLMBOKGJ,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MLDGGJIBFCO struct {
	MLDGGJIBFCO *BJMEIIHOINJ `protobuf:"bytes,1902,opt,name=MLDGGJIBFCO,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_LDKODDLEAKP struct {
	LDKODDLEAKP *DKDLBLJHAMB `protobuf:"bytes,1740,opt,name=LDKODDLEAKP,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_JLMOPLPFPOP struct {
	JLMOPLPFPOP *NLFOBIBCPEO `protobuf:"bytes,356,opt,name=JLMOPLPFPOP,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_KPAHINHGNJI struct {
	KPAHINHGNJI *ADMOCCBHGIE `protobuf:"bytes,814,opt,name=KPAHINHGNJI,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BonusSelectResult struct {
	BonusSelectResult *RogueBonusSelectResult `protobuf:"bytes,897,opt,name=bonus_select_result,json=bonusSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult struct {
	RogueTournFormulaResult *RogueTournFormulaResult `protobuf:"bytes,1752,opt,name=rogue_tourn_formula_result,json=rogueTournFormulaResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_GLLFBCGDMJA struct {
	GLLFBCGDMJA *DELGOOPEKOM `protobuf:"bytes,1134,opt,name=GLLFBCGDMJA,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_KEPJDAKKDOJ struct {
	KEPJDAKKDOJ *EOAHHEAOFFK `protobuf:"bytes,1441,opt,name=KEPJDAKKDOJ,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_IOJKNOLKFLD struct {
	IOJKNOLKFLD *OHOPHFCMOKH `protobuf:"bytes,1006,opt,name=IOJKNOLKFLD,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_NANBLBHDPPG struct {
	NANBLBHDPPG *IIJJHNNNKCF `protobuf:"bytes,10380,opt,name=NANBLBHDPPG,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionCsReq_BuffSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_CLOMAFJKKOP) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_EBFFLNAPJHJ) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_NLHEFFPICMN) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MiracleSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_GDLGGEDGJKN) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_FKOJLMBOKGJ) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MLDGGJIBFCO) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_LDKODDLEAKP) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_JLMOPLPFPOP) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_KPAHINHGNJI) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BonusSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_GLLFBCGDMJA) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_KEPJDAKKDOJ) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_IOJKNOLKFLD) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_NANBLBHDPPG) isHandleRogueCommonPendingActionCsReq_Action() {
}

var File_HandleRogueCommonPendingActionCsReq_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4f, 0x41,
	0x48, 0x48, 0x45, 0x41, 0x4f, 0x46, 0x46, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4a, 0x49, 0x4c, 0x4b, 0x43, 0x4c, 0x4c, 0x48, 0x44, 0x48, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4e, 0x4c, 0x46, 0x4f, 0x42, 0x49, 0x42, 0x43, 0x50, 0x45, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x45, 0x4c, 0x47, 0x4f, 0x4f, 0x50, 0x45, 0x4b, 0x4f,
	0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72,
	0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4b, 0x4b, 0x48, 0x43, 0x4f, 0x4e, 0x4a, 0x42, 0x4d, 0x48, 0x44, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4b, 0x44, 0x4c, 0x42, 0x4c, 0x4a, 0x48, 0x41, 0x4d, 0x42, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x48, 0x4f, 0x50, 0x48, 0x46, 0x43, 0x4d, 0x4f,
	0x4b, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x49, 0x4a, 0x4a, 0x48, 0x4e,
	0x4e, 0x4e, 0x4b, 0x43, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x42, 0x4b,
	0x4f, 0x4f, 0x49, 0x49, 0x4f, 0x44, 0x46, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x41, 0x44, 0x4d, 0x4f, 0x43, 0x43, 0x42, 0x48, 0x47, 0x49, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x42, 0x45, 0x4b, 0x4d, 0x44, 0x43, 0x49, 0x4d, 0x48, 0x48, 0x47, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x43, 0x47, 0x4e, 0x50, 0x49, 0x4e, 0x44, 0x43, 0x4a,
	0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x4a, 0x4d, 0x45, 0x49, 0x49, 0x48,
	0x4f, 0x49, 0x4e, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x09, 0x0a, 0x23, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x12, 0x62, 0x75, 0x66, 0x66,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xb5,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52,
	0x10, 0x62, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x31, 0x0a, 0x0b, 0x43, 0x4c, 0x4f, 0x4d, 0x41, 0x46, 0x4a, 0x4b, 0x4b, 0x4f, 0x50,
	0x18, 0xa6, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x43, 0x47, 0x4e, 0x50, 0x49,
	0x4e, 0x44, 0x43, 0x4a, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x43, 0x4c, 0x4f, 0x4d, 0x41, 0x46, 0x4a,
	0x4b, 0x4b, 0x4f, 0x50, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x42, 0x46, 0x46, 0x4c, 0x4e, 0x41, 0x50,
	0x4a, 0x48, 0x4a, 0x18, 0x8b, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4b, 0x48,
	0x43, 0x4f, 0x4e, 0x4a, 0x42, 0x4d, 0x48, 0x44, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x42, 0x46, 0x46,
	0x4c, 0x4e, 0x41, 0x50, 0x4a, 0x48, 0x4a, 0x12, 0x54, 0x0a, 0x19, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x72, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0xf0, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x16, 0x62, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c,
	0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a,
	0x0b, 0x4e, 0x4c, 0x48, 0x45, 0x46, 0x46, 0x50, 0x49, 0x43, 0x4d, 0x4e, 0x18, 0xb6, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x45, 0x4b, 0x4d, 0x44, 0x43, 0x49, 0x4d, 0x48, 0x48,
	0x47, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x4c, 0x48, 0x45, 0x46, 0x46, 0x50, 0x49, 0x43, 0x4d, 0x4e,
	0x12, 0x50, 0x0a, 0x15, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xa9, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x13, 0x6d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x44, 0x4c, 0x47, 0x47, 0x45, 0x44, 0x47, 0x4a, 0x4b,
	0x4e, 0x18, 0xa5, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x49, 0x4c, 0x4b, 0x43,
	0x4c, 0x4c, 0x48, 0x44, 0x48, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x44, 0x4c, 0x47, 0x47, 0x45,
	0x44, 0x47, 0x4a, 0x4b, 0x4e, 0x12, 0x31, 0x0a, 0x0b, 0x46, 0x4b, 0x4f, 0x4a, 0x4c, 0x4d, 0x42,
	0x4f, 0x4b, 0x47, 0x4a, 0x18, 0xe7, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x42,
	0x4b, 0x4f, 0x4f, 0x49, 0x49, 0x4f, 0x44, 0x46, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x4b, 0x4f,
	0x4a, 0x4c, 0x4d, 0x42, 0x4f, 0x4b, 0x47, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x4d, 0x4c, 0x44, 0x47,
	0x47, 0x4a, 0x49, 0x42, 0x46, 0x43, 0x4f, 0x18, 0xee, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x42, 0x4a, 0x4d, 0x45, 0x49, 0x49, 0x48, 0x4f, 0x49, 0x4e, 0x4a, 0x48, 0x00, 0x52, 0x0b,
	0x4d, 0x4c, 0x44, 0x47, 0x47, 0x4a, 0x49, 0x42, 0x46, 0x43, 0x4f, 0x12, 0x31, 0x0a, 0x0b, 0x4c,
	0x44, 0x4b, 0x4f, 0x44, 0x44, 0x4c, 0x45, 0x41, 0x4b, 0x50, 0x18, 0xcc, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4b, 0x44, 0x4c, 0x42, 0x4c, 0x4a, 0x48, 0x41, 0x4d, 0x42, 0x48,
	0x00, 0x52, 0x0b, 0x4c, 0x44, 0x4b, 0x4f, 0x44, 0x44, 0x4c, 0x45, 0x41, 0x4b, 0x50, 0x12, 0x31,
	0x0a, 0x0b, 0x4a, 0x4c, 0x4d, 0x4f, 0x50, 0x4c, 0x50, 0x46, 0x50, 0x4f, 0x50, 0x18, 0xe4, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4c, 0x46, 0x4f, 0x42, 0x49, 0x42, 0x43, 0x50,
	0x45, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x4a, 0x4c, 0x4d, 0x4f, 0x50, 0x4c, 0x50, 0x46, 0x50, 0x4f,
	0x50, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x50, 0x41, 0x48, 0x49, 0x4e, 0x48, 0x47, 0x4e, 0x4a, 0x49,
	0x18, 0xae, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x44, 0x4d, 0x4f, 0x43, 0x43,
	0x42, 0x48, 0x47, 0x49, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x50, 0x41, 0x48, 0x49, 0x4e, 0x48,
	0x47, 0x4e, 0x4a, 0x49, 0x12, 0x4a, 0x0a, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x81, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x11, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x58, 0x0a, 0x1a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xd8,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48,
	0x00, 0x52, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x4c,
	0x4c, 0x46, 0x42, 0x43, 0x47, 0x44, 0x4d, 0x4a, 0x41, 0x18, 0xee, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x45, 0x4c, 0x47, 0x4f, 0x4f, 0x50, 0x45, 0x4b, 0x4f, 0x4d, 0x48, 0x00,
	0x52, 0x0b, 0x47, 0x4c, 0x4c, 0x46, 0x42, 0x43, 0x47, 0x44, 0x4d, 0x4a, 0x41, 0x12, 0x31, 0x0a,
	0x0b, 0x4b, 0x45, 0x50, 0x4a, 0x44, 0x41, 0x4b, 0x4b, 0x44, 0x4f, 0x4a, 0x18, 0xa1, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4f, 0x41, 0x48, 0x48, 0x45, 0x41, 0x4f, 0x46, 0x46,
	0x4b, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x45, 0x50, 0x4a, 0x44, 0x41, 0x4b, 0x4b, 0x44, 0x4f, 0x4a,
	0x12, 0x31, 0x0a, 0x0b, 0x49, 0x4f, 0x4a, 0x4b, 0x4e, 0x4f, 0x4c, 0x4b, 0x46, 0x4c, 0x44, 0x18,
	0xee, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x48, 0x4f, 0x50, 0x48, 0x46, 0x43,
	0x4d, 0x4f, 0x4b, 0x48, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x4f, 0x4a, 0x4b, 0x4e, 0x4f, 0x4c, 0x4b,
	0x46, 0x4c, 0x44, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x41, 0x4e, 0x42, 0x4c, 0x42, 0x48, 0x44, 0x50,
	0x50, 0x47, 0x18, 0x8c, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x49, 0x4a, 0x4a,
	0x48, 0x4e, 0x4e, 0x4e, 0x4b, 0x43, 0x46, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x41, 0x4e, 0x42, 0x4c,
	0x42, 0x48, 0x44, 0x50, 0x50, 0x47, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce sync.Once
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = file_HandleRogueCommonPendingActionCsReq_proto_rawDesc
)

func file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP() []byte {
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce.Do(func() {
		file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleRogueCommonPendingActionCsReq_proto_rawDescData)
	})
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescData
}

var file_HandleRogueCommonPendingActionCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleRogueCommonPendingActionCsReq_proto_goTypes = []interface{}{
	(*HandleRogueCommonPendingActionCsReq)(nil), // 0: HandleRogueCommonPendingActionCsReq
	(*RogueBuffSelectResult)(nil),               // 1: RogueBuffSelectResult
	(*FCGNPINDCJO)(nil),                         // 2: FCGNPINDCJO
	(*KKHCONJBMHD)(nil),                         // 3: KKHCONJBMHD
	(*RogueBuffRerollResult)(nil),               // 4: RogueBuffRerollResult
	(*BEKMDCIMHHG)(nil),                         // 5: BEKMDCIMHHG
	(*RogueMiracleSelectResult)(nil),            // 6: RogueMiracleSelectResult
	(*JILKCLLHDHL)(nil),                         // 7: JILKCLLHDHL
	(*GBKOOIIODFC)(nil),                         // 8: GBKOOIIODFC
	(*BJMEIIHOINJ)(nil),                         // 9: BJMEIIHOINJ
	(*DKDLBLJHAMB)(nil),                         // 10: DKDLBLJHAMB
	(*NLFOBIBCPEO)(nil),                         // 11: NLFOBIBCPEO
	(*ADMOCCBHGIE)(nil),                         // 12: ADMOCCBHGIE
	(*RogueBonusSelectResult)(nil),              // 13: RogueBonusSelectResult
	(*RogueTournFormulaResult)(nil),             // 14: RogueTournFormulaResult
	(*DELGOOPEKOM)(nil),                         // 15: DELGOOPEKOM
	(*EOAHHEAOFFK)(nil),                         // 16: EOAHHEAOFFK
	(*OHOPHFCMOKH)(nil),                         // 17: OHOPHFCMOKH
	(*IIJJHNNNKCF)(nil),                         // 18: IIJJHNNNKCF
}
var file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = []int32{
	1,  // 0: HandleRogueCommonPendingActionCsReq.buff_select_result:type_name -> RogueBuffSelectResult
	2,  // 1: HandleRogueCommonPendingActionCsReq.CLOMAFJKKOP:type_name -> FCGNPINDCJO
	3,  // 2: HandleRogueCommonPendingActionCsReq.EBFFLNAPJHJ:type_name -> KKHCONJBMHD
	4,  // 3: HandleRogueCommonPendingActionCsReq.buff_reroll_select_result:type_name -> RogueBuffRerollResult
	5,  // 4: HandleRogueCommonPendingActionCsReq.NLHEFFPICMN:type_name -> BEKMDCIMHHG
	6,  // 5: HandleRogueCommonPendingActionCsReq.miracle_select_result:type_name -> RogueMiracleSelectResult
	7,  // 6: HandleRogueCommonPendingActionCsReq.GDLGGEDGJKN:type_name -> JILKCLLHDHL
	8,  // 7: HandleRogueCommonPendingActionCsReq.FKOJLMBOKGJ:type_name -> GBKOOIIODFC
	9,  // 8: HandleRogueCommonPendingActionCsReq.MLDGGJIBFCO:type_name -> BJMEIIHOINJ
	10, // 9: HandleRogueCommonPendingActionCsReq.LDKODDLEAKP:type_name -> DKDLBLJHAMB
	11, // 10: HandleRogueCommonPendingActionCsReq.JLMOPLPFPOP:type_name -> NLFOBIBCPEO
	12, // 11: HandleRogueCommonPendingActionCsReq.KPAHINHGNJI:type_name -> ADMOCCBHGIE
	13, // 12: HandleRogueCommonPendingActionCsReq.bonus_select_result:type_name -> RogueBonusSelectResult
	14, // 13: HandleRogueCommonPendingActionCsReq.rogue_tourn_formula_result:type_name -> RogueTournFormulaResult
	15, // 14: HandleRogueCommonPendingActionCsReq.GLLFBCGDMJA:type_name -> DELGOOPEKOM
	16, // 15: HandleRogueCommonPendingActionCsReq.KEPJDAKKDOJ:type_name -> EOAHHEAOFFK
	17, // 16: HandleRogueCommonPendingActionCsReq.IOJKNOLKFLD:type_name -> OHOPHFCMOKH
	18, // 17: HandleRogueCommonPendingActionCsReq.NANBLBHDPPG:type_name -> IIJJHNNNKCF
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionCsReq_proto_init() }
func file_HandleRogueCommonPendingActionCsReq_proto_init() {
	if File_HandleRogueCommonPendingActionCsReq_proto != nil {
		return
	}
	file_EOAHHEAOFFK_proto_init()
	file_JILKCLLHDHL_proto_init()
	file_NLFOBIBCPEO_proto_init()
	file_DELGOOPEKOM_proto_init()
	file_RogueBuffSelectResult_proto_init()
	file_RogueBonusSelectResult_proto_init()
	file_RogueBuffRerollResult_proto_init()
	file_KKHCONJBMHD_proto_init()
	file_RogueMiracleSelectResult_proto_init()
	file_DKDLBLJHAMB_proto_init()
	file_OHOPHFCMOKH_proto_init()
	file_IIJJHNNNKCF_proto_init()
	file_GBKOOIIODFC_proto_init()
	file_ADMOCCBHGIE_proto_init()
	file_BEKMDCIMHHG_proto_init()
	file_FCGNPINDCJO_proto_init()
	file_BJMEIIHOINJ_proto_init()
	file_RogueTournFormulaResult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleRogueCommonPendingActionCsReq); i {
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
	file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HandleRogueCommonPendingActionCsReq_BuffSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_CLOMAFJKKOP)(nil),
		(*HandleRogueCommonPendingActionCsReq_EBFFLNAPJHJ)(nil),
		(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_NLHEFFPICMN)(nil),
		(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_GDLGGEDGJKN)(nil),
		(*HandleRogueCommonPendingActionCsReq_FKOJLMBOKGJ)(nil),
		(*HandleRogueCommonPendingActionCsReq_MLDGGJIBFCO)(nil),
		(*HandleRogueCommonPendingActionCsReq_LDKODDLEAKP)(nil),
		(*HandleRogueCommonPendingActionCsReq_JLMOPLPFPOP)(nil),
		(*HandleRogueCommonPendingActionCsReq_KPAHINHGNJI)(nil),
		(*HandleRogueCommonPendingActionCsReq_BonusSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_GLLFBCGDMJA)(nil),
		(*HandleRogueCommonPendingActionCsReq_KEPJDAKKDOJ)(nil),
		(*HandleRogueCommonPendingActionCsReq_IOJKNOLKFLD)(nil),
		(*HandleRogueCommonPendingActionCsReq_NANBLBHDPPG)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HandleRogueCommonPendingActionCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleRogueCommonPendingActionCsReq_proto_goTypes,
		DependencyIndexes: file_HandleRogueCommonPendingActionCsReq_proto_depIdxs,
		MessageInfos:      file_HandleRogueCommonPendingActionCsReq_proto_msgTypes,
	}.Build()
	File_HandleRogueCommonPendingActionCsReq_proto = out.File
	file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = nil
	file_HandleRogueCommonPendingActionCsReq_proto_goTypes = nil
	file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = nil
}
