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

	QueueLocation uint32 `protobuf:"varint,13,opt,name=queue_location,json=queueLocation,proto3" json:"queue_location,omitempty"`
	// Types that are assignable to Action:
	//
	//	*HandleRogueCommonPendingActionCsReq_BuffSelectResult
	//	*HandleRogueCommonPendingActionCsReq_NGMGGIKPPML
	//	*HandleRogueCommonPendingActionCsReq_HNPCNPHPOFO
	//	*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult
	//	*HandleRogueCommonPendingActionCsReq_APEMNEOMEJC
	//	*HandleRogueCommonPendingActionCsReq_MiracleSelectResult
	//	*HandleRogueCommonPendingActionCsReq_OCIDABLHJIH
	//	*HandleRogueCommonPendingActionCsReq_EIKDKBPMJFI
	//	*HandleRogueCommonPendingActionCsReq_HLLIGDDNBDN
	//	*HandleRogueCommonPendingActionCsReq_EMBGMEGJKFI
	//	*HandleRogueCommonPendingActionCsReq_FEDGABPAIOM
	//	*HandleRogueCommonPendingActionCsReq_NDGBENNMADD
	//	*HandleRogueCommonPendingActionCsReq_BonusSelectResult
	//	*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult
	//	*HandleRogueCommonPendingActionCsReq_HEHGEDKECJE
	//	*HandleRogueCommonPendingActionCsReq_CEELCMLKJGD
	//	*HandleRogueCommonPendingActionCsReq_MEBCEBMIGOI
	//	*HandleRogueCommonPendingActionCsReq_NGPNHBFLPEL
	//	*HandleRogueCommonPendingActionCsReq_IEEDBGFDBAL
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

func (x *HandleRogueCommonPendingActionCsReq) GetNGMGGIKPPML() *FIJOIHPLGFI {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_NGMGGIKPPML); ok {
		return x.NGMGGIKPPML
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetHNPCNPHPOFO() *LHMGPKEOGHE {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_HNPCNPHPOFO); ok {
		return x.HNPCNPHPOFO
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffRerollSelectResult() *RogueBuffRerollResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult); ok {
		return x.BuffRerollSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetAPEMNEOMEJC() *NBJGOGBHOKM {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_APEMNEOMEJC); ok {
		return x.APEMNEOMEJC
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMiracleSelectResult() *RogueMiracleSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult); ok {
		return x.MiracleSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetOCIDABLHJIH() *ENOLGAIDDGI {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_OCIDABLHJIH); ok {
		return x.OCIDABLHJIH
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetEIKDKBPMJFI() *OKAOEPBDLKG {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_EIKDKBPMJFI); ok {
		return x.EIKDKBPMJFI
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetHLLIGDDNBDN() *BDLJNOIIOOH {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_HLLIGDDNBDN); ok {
		return x.HLLIGDDNBDN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetEMBGMEGJKFI() *GNDKAOLNAIC {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_EMBGMEGJKFI); ok {
		return x.EMBGMEGJKFI
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetFEDGABPAIOM() *JOGGLAHDIHP {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_FEDGABPAIOM); ok {
		return x.FEDGABPAIOM
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetNDGBENNMADD() *PNHMJIKAAMK {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_NDGBENNMADD); ok {
		return x.NDGBENNMADD
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

func (x *HandleRogueCommonPendingActionCsReq) GetHEHGEDKECJE() *JHDEAOONPLE {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_HEHGEDKECJE); ok {
		return x.HEHGEDKECJE
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetCEELCMLKJGD() *HDCJNNPGEID {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_CEELCMLKJGD); ok {
		return x.CEELCMLKJGD
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMEBCEBMIGOI() *KEADKPNDPML {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MEBCEBMIGOI); ok {
		return x.MEBCEBMIGOI
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetNGPNHBFLPEL() *PALLDJBAIEK {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_NGPNHBFLPEL); ok {
		return x.NGPNHBFLPEL
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetIEEDBGFDBAL() *JDPLMLDAKEC {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_IEEDBGFDBAL); ok {
		return x.IEEDBGFDBAL
	}
	return nil
}

type isHandleRogueCommonPendingActionCsReq_Action interface {
	isHandleRogueCommonPendingActionCsReq_Action()
}

type HandleRogueCommonPendingActionCsReq_BuffSelectResult struct {
	BuffSelectResult *RogueBuffSelectResult `protobuf:"bytes,737,opt,name=buff_select_result,json=buffSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_NGMGGIKPPML struct {
	NGMGGIKPPML *FIJOIHPLGFI `protobuf:"bytes,481,opt,name=NGMGGIKPPML,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_HNPCNPHPOFO struct {
	HNPCNPHPOFO *LHMGPKEOGHE `protobuf:"bytes,1202,opt,name=HNPCNPHPOFO,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult struct {
	BuffRerollSelectResult *RogueBuffRerollResult `protobuf:"bytes,1753,opt,name=buff_reroll_select_result,json=buffRerollSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_APEMNEOMEJC struct {
	APEMNEOMEJC *NBJGOGBHOKM `protobuf:"bytes,226,opt,name=APEMNEOMEJC,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MiracleSelectResult struct {
	MiracleSelectResult *RogueMiracleSelectResult `protobuf:"bytes,1596,opt,name=miracle_select_result,json=miracleSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_OCIDABLHJIH struct {
	OCIDABLHJIH *ENOLGAIDDGI `protobuf:"bytes,1019,opt,name=OCIDABLHJIH,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_EIKDKBPMJFI struct {
	EIKDKBPMJFI *OKAOEPBDLKG `protobuf:"bytes,735,opt,name=EIKDKBPMJFI,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_HLLIGDDNBDN struct {
	HLLIGDDNBDN *BDLJNOIIOOH `protobuf:"bytes,935,opt,name=HLLIGDDNBDN,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_EMBGMEGJKFI struct {
	EMBGMEGJKFI *GNDKAOLNAIC `protobuf:"bytes,252,opt,name=EMBGMEGJKFI,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_FEDGABPAIOM struct {
	FEDGABPAIOM *JOGGLAHDIHP `protobuf:"bytes,2011,opt,name=FEDGABPAIOM,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_NDGBENNMADD struct {
	NDGBENNMADD *PNHMJIKAAMK `protobuf:"bytes,76,opt,name=NDGBENNMADD,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BonusSelectResult struct {
	BonusSelectResult *RogueBonusSelectResult `protobuf:"bytes,1105,opt,name=bonus_select_result,json=bonusSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult struct {
	RogueTournFormulaResult *RogueTournFormulaResult `protobuf:"bytes,908,opt,name=rogue_tourn_formula_result,json=rogueTournFormulaResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_HEHGEDKECJE struct {
	HEHGEDKECJE *JHDEAOONPLE `protobuf:"bytes,1349,opt,name=HEHGEDKECJE,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_CEELCMLKJGD struct {
	CEELCMLKJGD *HDCJNNPGEID `protobuf:"bytes,564,opt,name=CEELCMLKJGD,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MEBCEBMIGOI struct {
	MEBCEBMIGOI *KEADKPNDPML `protobuf:"bytes,1450,opt,name=MEBCEBMIGOI,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_NGPNHBFLPEL struct {
	NGPNHBFLPEL *PALLDJBAIEK `protobuf:"bytes,379,opt,name=NGPNHBFLPEL,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_IEEDBGFDBAL struct {
	IEEDBGFDBAL *JDPLMLDAKEC `protobuf:"bytes,210810,opt,name=IEEDBGFDBAL,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionCsReq_BuffSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_NGMGGIKPPML) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_HNPCNPHPOFO) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_APEMNEOMEJC) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MiracleSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_OCIDABLHJIH) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_EIKDKBPMJFI) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_HLLIGDDNBDN) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_EMBGMEGJKFI) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_FEDGABPAIOM) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_NDGBENNMADD) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BonusSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_HEHGEDKECJE) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_CEELCMLKJGD) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MEBCEBMIGOI) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_NGPNHBFLPEL) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_IEEDBGFDBAL) isHandleRogueCommonPendingActionCsReq_Action() {
}

var File_HandleRogueCommonPendingActionCsReq_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x42, 0x4a,
	0x47, 0x4f, 0x47, 0x42, 0x48, 0x4f, 0x4b, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4a, 0x44, 0x50, 0x4c, 0x4d, 0x4c, 0x44, 0x41, 0x4b, 0x45, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x48, 0x44, 0x43, 0x4a, 0x4e, 0x4e, 0x50, 0x47, 0x45, 0x49, 0x44, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4b, 0x41, 0x4f, 0x45, 0x50, 0x42, 0x44, 0x4c, 0x4b,
	0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x48, 0x44, 0x45, 0x41, 0x4f, 0x4f,
	0x4e, 0x50, 0x4c, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x45, 0x41, 0x44,
	0x4b, 0x50, 0x4e, 0x44, 0x50, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50,
	0x4e, 0x48, 0x4d, 0x4a, 0x49, 0x4b, 0x41, 0x41, 0x4d, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x47, 0x4e, 0x44, 0x4b, 0x41, 0x4f, 0x4c, 0x4e, 0x41, 0x49, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4e, 0x4f, 0x4c, 0x47, 0x41, 0x49, 0x44, 0x44, 0x47, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x41, 0x4c, 0x4c, 0x44, 0x4a, 0x42, 0x41,
	0x49, 0x45, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4c, 0x48, 0x4d, 0x47, 0x50, 0x4b, 0x45, 0x4f, 0x47, 0x48, 0x45, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x49, 0x4a, 0x4f, 0x49, 0x48, 0x50, 0x4c, 0x47, 0x46, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4f, 0x47, 0x47, 0x4c, 0x41, 0x48, 0x44,
	0x49, 0x48, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x44, 0x4c, 0x4a, 0x4e,
	0x4f, 0x49, 0x49, 0x4f, 0x4f, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x09, 0x0a,
	0x23, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x12, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0xe1, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x62, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x47, 0x4d, 0x47, 0x47, 0x49, 0x4b, 0x50,
	0x50, 0x4d, 0x4c, 0x18, 0xe1, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x49, 0x4a,
	0x4f, 0x49, 0x48, 0x50, 0x4c, 0x47, 0x46, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x47, 0x4d, 0x47,
	0x47, 0x49, 0x4b, 0x50, 0x50, 0x4d, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x48, 0x4e, 0x50, 0x43, 0x4e,
	0x50, 0x48, 0x50, 0x4f, 0x46, 0x4f, 0x18, 0xb2, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4c, 0x48, 0x4d, 0x47, 0x50, 0x4b, 0x45, 0x4f, 0x47, 0x48, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x48,
	0x4e, 0x50, 0x43, 0x4e, 0x50, 0x48, 0x50, 0x4f, 0x46, 0x4f, 0x12, 0x54, 0x0a, 0x19, 0x62, 0x75,
	0x66, 0x66, 0x5f, 0x72, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xd9, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x16, 0x62, 0x75, 0x66, 0x66, 0x52, 0x65,
	0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x31, 0x0a, 0x0b, 0x41, 0x50, 0x45, 0x4d, 0x4e, 0x45, 0x4f, 0x4d, 0x45, 0x4a, 0x43, 0x18,
	0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x42, 0x4a, 0x47, 0x4f, 0x47, 0x42,
	0x48, 0x4f, 0x4b, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x50, 0x45, 0x4d, 0x4e, 0x45, 0x4f, 0x4d,
	0x45, 0x4a, 0x43, 0x12, 0x50, 0x0a, 0x15, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xbc, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x13, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x43, 0x49, 0x44, 0x41, 0x42, 0x4c,
	0x48, 0x4a, 0x49, 0x48, 0x18, 0xfb, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4e,
	0x4f, 0x4c, 0x47, 0x41, 0x49, 0x44, 0x44, 0x47, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x43, 0x49,
	0x44, 0x41, 0x42, 0x4c, 0x48, 0x4a, 0x49, 0x48, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x49, 0x4b, 0x44,
	0x4b, 0x42, 0x50, 0x4d, 0x4a, 0x46, 0x49, 0x18, 0xdf, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4f, 0x4b, 0x41, 0x4f, 0x45, 0x50, 0x42, 0x44, 0x4c, 0x4b, 0x47, 0x48, 0x00, 0x52, 0x0b,
	0x45, 0x49, 0x4b, 0x44, 0x4b, 0x42, 0x50, 0x4d, 0x4a, 0x46, 0x49, 0x12, 0x31, 0x0a, 0x0b, 0x48,
	0x4c, 0x4c, 0x49, 0x47, 0x44, 0x44, 0x4e, 0x42, 0x44, 0x4e, 0x18, 0xa7, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x44, 0x4c, 0x4a, 0x4e, 0x4f, 0x49, 0x49, 0x4f, 0x4f, 0x48, 0x48,
	0x00, 0x52, 0x0b, 0x48, 0x4c, 0x4c, 0x49, 0x47, 0x44, 0x44, 0x4e, 0x42, 0x44, 0x4e, 0x12, 0x31,
	0x0a, 0x0b, 0x45, 0x4d, 0x42, 0x47, 0x4d, 0x45, 0x47, 0x4a, 0x4b, 0x46, 0x49, 0x18, 0xfc, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4e, 0x44, 0x4b, 0x41, 0x4f, 0x4c, 0x4e, 0x41,
	0x49, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x4d, 0x42, 0x47, 0x4d, 0x45, 0x47, 0x4a, 0x4b, 0x46,
	0x49, 0x12, 0x31, 0x0a, 0x0b, 0x46, 0x45, 0x44, 0x47, 0x41, 0x42, 0x50, 0x41, 0x49, 0x4f, 0x4d,
	0x18, 0xdb, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4f, 0x47, 0x47, 0x4c, 0x41,
	0x48, 0x44, 0x49, 0x48, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x45, 0x44, 0x47, 0x41, 0x42, 0x50,
	0x41, 0x49, 0x4f, 0x4d, 0x12, 0x30, 0x0a, 0x0b, 0x4e, 0x44, 0x47, 0x42, 0x45, 0x4e, 0x4e, 0x4d,
	0x41, 0x44, 0x44, 0x18, 0x4c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4e, 0x48, 0x4d,
	0x4a, 0x49, 0x4b, 0x41, 0x41, 0x4d, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x44, 0x47, 0x42, 0x45,
	0x4e, 0x4e, 0x4d, 0x41, 0x44, 0x44, 0x12, 0x4a, 0x0a, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xd1, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52,
	0x11, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x58, 0x0a, 0x1a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x8c, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x48, 0x00, 0x52, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b,
	0x48, 0x45, 0x48, 0x47, 0x45, 0x44, 0x4b, 0x45, 0x43, 0x4a, 0x45, 0x18, 0xc5, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x48, 0x44, 0x45, 0x41, 0x4f, 0x4f, 0x4e, 0x50, 0x4c, 0x45,
	0x48, 0x00, 0x52, 0x0b, 0x48, 0x45, 0x48, 0x47, 0x45, 0x44, 0x4b, 0x45, 0x43, 0x4a, 0x45, 0x12,
	0x31, 0x0a, 0x0b, 0x43, 0x45, 0x45, 0x4c, 0x43, 0x4d, 0x4c, 0x4b, 0x4a, 0x47, 0x44, 0x18, 0xb4,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x44, 0x43, 0x4a, 0x4e, 0x4e, 0x50, 0x47,
	0x45, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0b, 0x43, 0x45, 0x45, 0x4c, 0x43, 0x4d, 0x4c, 0x4b, 0x4a,
	0x47, 0x44, 0x12, 0x31, 0x0a, 0x0b, 0x4d, 0x45, 0x42, 0x43, 0x45, 0x42, 0x4d, 0x49, 0x47, 0x4f,
	0x49, 0x18, 0xaa, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x45, 0x41, 0x44, 0x4b,
	0x50, 0x4e, 0x44, 0x50, 0x4d, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x4d, 0x45, 0x42, 0x43, 0x45, 0x42,
	0x4d, 0x49, 0x47, 0x4f, 0x49, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x47, 0x50, 0x4e, 0x48, 0x42, 0x46,
	0x4c, 0x50, 0x45, 0x4c, 0x18, 0xfb, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x41,
	0x4c, 0x4c, 0x44, 0x4a, 0x42, 0x41, 0x49, 0x45, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x47, 0x50,
	0x4e, 0x48, 0x42, 0x46, 0x4c, 0x50, 0x45, 0x4c, 0x12, 0x32, 0x0a, 0x0b, 0x49, 0x45, 0x45, 0x44,
	0x42, 0x47, 0x46, 0x44, 0x42, 0x41, 0x4c, 0x18, 0xfa, 0xee, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4a, 0x44, 0x50, 0x4c, 0x4d, 0x4c, 0x44, 0x41, 0x4b, 0x45, 0x43, 0x48, 0x00, 0x52,
	0x0b, 0x49, 0x45, 0x45, 0x44, 0x42, 0x47, 0x46, 0x44, 0x42, 0x41, 0x4c, 0x42, 0x08, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*FIJOIHPLGFI)(nil),                         // 2: FIJOIHPLGFI
	(*LHMGPKEOGHE)(nil),                         // 3: LHMGPKEOGHE
	(*RogueBuffRerollResult)(nil),               // 4: RogueBuffRerollResult
	(*NBJGOGBHOKM)(nil),                         // 5: NBJGOGBHOKM
	(*RogueMiracleSelectResult)(nil),            // 6: RogueMiracleSelectResult
	(*ENOLGAIDDGI)(nil),                         // 7: ENOLGAIDDGI
	(*OKAOEPBDLKG)(nil),                         // 8: OKAOEPBDLKG
	(*BDLJNOIIOOH)(nil),                         // 9: BDLJNOIIOOH
	(*GNDKAOLNAIC)(nil),                         // 10: GNDKAOLNAIC
	(*JOGGLAHDIHP)(nil),                         // 11: JOGGLAHDIHP
	(*PNHMJIKAAMK)(nil),                         // 12: PNHMJIKAAMK
	(*RogueBonusSelectResult)(nil),              // 13: RogueBonusSelectResult
	(*RogueTournFormulaResult)(nil),             // 14: RogueTournFormulaResult
	(*JHDEAOONPLE)(nil),                         // 15: JHDEAOONPLE
	(*HDCJNNPGEID)(nil),                         // 16: HDCJNNPGEID
	(*KEADKPNDPML)(nil),                         // 17: KEADKPNDPML
	(*PALLDJBAIEK)(nil),                         // 18: PALLDJBAIEK
	(*JDPLMLDAKEC)(nil),                         // 19: JDPLMLDAKEC
}
var file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = []int32{
	1,  // 0: HandleRogueCommonPendingActionCsReq.buff_select_result:type_name -> RogueBuffSelectResult
	2,  // 1: HandleRogueCommonPendingActionCsReq.NGMGGIKPPML:type_name -> FIJOIHPLGFI
	3,  // 2: HandleRogueCommonPendingActionCsReq.HNPCNPHPOFO:type_name -> LHMGPKEOGHE
	4,  // 3: HandleRogueCommonPendingActionCsReq.buff_reroll_select_result:type_name -> RogueBuffRerollResult
	5,  // 4: HandleRogueCommonPendingActionCsReq.APEMNEOMEJC:type_name -> NBJGOGBHOKM
	6,  // 5: HandleRogueCommonPendingActionCsReq.miracle_select_result:type_name -> RogueMiracleSelectResult
	7,  // 6: HandleRogueCommonPendingActionCsReq.OCIDABLHJIH:type_name -> ENOLGAIDDGI
	8,  // 7: HandleRogueCommonPendingActionCsReq.EIKDKBPMJFI:type_name -> OKAOEPBDLKG
	9,  // 8: HandleRogueCommonPendingActionCsReq.HLLIGDDNBDN:type_name -> BDLJNOIIOOH
	10, // 9: HandleRogueCommonPendingActionCsReq.EMBGMEGJKFI:type_name -> GNDKAOLNAIC
	11, // 10: HandleRogueCommonPendingActionCsReq.FEDGABPAIOM:type_name -> JOGGLAHDIHP
	12, // 11: HandleRogueCommonPendingActionCsReq.NDGBENNMADD:type_name -> PNHMJIKAAMK
	13, // 12: HandleRogueCommonPendingActionCsReq.bonus_select_result:type_name -> RogueBonusSelectResult
	14, // 13: HandleRogueCommonPendingActionCsReq.rogue_tourn_formula_result:type_name -> RogueTournFormulaResult
	15, // 14: HandleRogueCommonPendingActionCsReq.HEHGEDKECJE:type_name -> JHDEAOONPLE
	16, // 15: HandleRogueCommonPendingActionCsReq.CEELCMLKJGD:type_name -> HDCJNNPGEID
	17, // 16: HandleRogueCommonPendingActionCsReq.MEBCEBMIGOI:type_name -> KEADKPNDPML
	18, // 17: HandleRogueCommonPendingActionCsReq.NGPNHBFLPEL:type_name -> PALLDJBAIEK
	19, // 18: HandleRogueCommonPendingActionCsReq.IEEDBGFDBAL:type_name -> JDPLMLDAKEC
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionCsReq_proto_init() }
func file_HandleRogueCommonPendingActionCsReq_proto_init() {
	if File_HandleRogueCommonPendingActionCsReq_proto != nil {
		return
	}
	file_NBJGOGBHOKM_proto_init()
	file_JDPLMLDAKEC_proto_init()
	file_HDCJNNPGEID_proto_init()
	file_OKAOEPBDLKG_proto_init()
	file_JHDEAOONPLE_proto_init()
	file_KEADKPNDPML_proto_init()
	file_PNHMJIKAAMK_proto_init()
	file_RogueMiracleSelectResult_proto_init()
	file_GNDKAOLNAIC_proto_init()
	file_ENOLGAIDDGI_proto_init()
	file_PALLDJBAIEK_proto_init()
	file_RogueBonusSelectResult_proto_init()
	file_RogueBuffRerollResult_proto_init()
	file_RogueBuffSelectResult_proto_init()
	file_LHMGPKEOGHE_proto_init()
	file_RogueTournFormulaResult_proto_init()
	file_FIJOIHPLGFI_proto_init()
	file_JOGGLAHDIHP_proto_init()
	file_BDLJNOIIOOH_proto_init()
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
		(*HandleRogueCommonPendingActionCsReq_NGMGGIKPPML)(nil),
		(*HandleRogueCommonPendingActionCsReq_HNPCNPHPOFO)(nil),
		(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_APEMNEOMEJC)(nil),
		(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_OCIDABLHJIH)(nil),
		(*HandleRogueCommonPendingActionCsReq_EIKDKBPMJFI)(nil),
		(*HandleRogueCommonPendingActionCsReq_HLLIGDDNBDN)(nil),
		(*HandleRogueCommonPendingActionCsReq_EMBGMEGJKFI)(nil),
		(*HandleRogueCommonPendingActionCsReq_FEDGABPAIOM)(nil),
		(*HandleRogueCommonPendingActionCsReq_NDGBENNMADD)(nil),
		(*HandleRogueCommonPendingActionCsReq_BonusSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_HEHGEDKECJE)(nil),
		(*HandleRogueCommonPendingActionCsReq_CEELCMLKJGD)(nil),
		(*HandleRogueCommonPendingActionCsReq_MEBCEBMIGOI)(nil),
		(*HandleRogueCommonPendingActionCsReq_NGPNHBFLPEL)(nil),
		(*HandleRogueCommonPendingActionCsReq_IEEDBGFDBAL)(nil),
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
