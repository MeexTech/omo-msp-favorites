// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/activity.proto

package favorite

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActivityInfo struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32      `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64       `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64       `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string      `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string      `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string      `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string      `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Require              string      `protobuf:"bytes,10,opt,name=require,proto3" json:"require,omitempty"`
	Cover                string      `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Owner                string      `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	Date                 *DateInfo   `protobuf:"bytes,13,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo  `protobuf:"bytes,14,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string      `protobuf:"bytes,15,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Limit                uint32      `protobuf:"varint,16,opt,name=limit,proto3" json:"limit,omitempty"`
	Status               uint32      `protobuf:"varint,17,opt,name=status,proto3" json:"status,omitempty"`
	Show                 uint32      `protobuf:"varint,18,opt,name=show,proto3" json:"show,omitempty"`
	Prize                *PrizeInfo  `protobuf:"bytes,19,opt,name=prize,proto3" json:"prize,omitempty"`
	Tags                 []string    `protobuf:"bytes,20,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string    `protobuf:"bytes,21,rep,name=assets,proto3" json:"assets,omitempty"`
	Targets              []string    `protobuf:"bytes,22,rep,name=targets,proto3" json:"targets,omitempty"`
	Participants         []*PairInfo `protobuf:"bytes,23,rep,name=participants,proto3" json:"participants,omitempty"`
	Opuses               []*OpusInfo `protobuf:"bytes,24,rep,name=opuses,proto3" json:"opuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ActivityInfo) Reset()         { *m = ActivityInfo{} }
func (m *ActivityInfo) String() string { return proto.CompactTextString(m) }
func (*ActivityInfo) ProtoMessage()    {}
func (*ActivityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{0}
}

func (m *ActivityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityInfo.Unmarshal(m, b)
}
func (m *ActivityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityInfo.Marshal(b, m, deterministic)
}
func (m *ActivityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityInfo.Merge(m, src)
}
func (m *ActivityInfo) XXX_Size() int {
	return xxx_messageInfo_ActivityInfo.Size(m)
}
func (m *ActivityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityInfo proto.InternalMessageInfo

func (m *ActivityInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ActivityInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActivityInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ActivityInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ActivityInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ActivityInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ActivityInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ActivityInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActivityInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ActivityInfo) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ActivityInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ActivityInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ActivityInfo) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ActivityInfo) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ActivityInfo) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ActivityInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ActivityInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ActivityInfo) GetShow() uint32 {
	if m != nil {
		return m.Show
	}
	return 0
}

func (m *ActivityInfo) GetPrize() *PrizeInfo {
	if m != nil {
		return m.Prize
	}
	return nil
}

func (m *ActivityInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ActivityInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ActivityInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ActivityInfo) GetParticipants() []*PairInfo {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *ActivityInfo) GetOpuses() []*OpusInfo {
	if m != nil {
		return m.Opuses
	}
	return nil
}

type PrizeInfo struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string      `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Ranks                []*RankInfo `protobuf:"bytes,3,rep,name=ranks,proto3" json:"ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PrizeInfo) Reset()         { *m = PrizeInfo{} }
func (m *PrizeInfo) String() string { return proto.CompactTextString(m) }
func (*PrizeInfo) ProtoMessage()    {}
func (*PrizeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{1}
}

func (m *PrizeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrizeInfo.Unmarshal(m, b)
}
func (m *PrizeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrizeInfo.Marshal(b, m, deterministic)
}
func (m *PrizeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrizeInfo.Merge(m, src)
}
func (m *PrizeInfo) XXX_Size() int {
	return xxx_messageInfo_PrizeInfo.Size(m)
}
func (m *PrizeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PrizeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PrizeInfo proto.InternalMessageInfo

func (m *PrizeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrizeInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *PrizeInfo) GetRanks() []*RankInfo {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type RankInfo struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankInfo) Reset()         { *m = RankInfo{} }
func (m *RankInfo) String() string { return proto.CompactTextString(m) }
func (*RankInfo) ProtoMessage()    {}
func (*RankInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{2}
}

func (m *RankInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInfo.Unmarshal(m, b)
}
func (m *RankInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInfo.Marshal(b, m, deterministic)
}
func (m *RankInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInfo.Merge(m, src)
}
func (m *RankInfo) XXX_Size() int {
	return xxx_messageInfo_RankInfo.Size(m)
}
func (m *RankInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankInfo proto.InternalMessageInfo

func (m *RankInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RankInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type OpusInfo struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpusInfo) Reset()         { *m = OpusInfo{} }
func (m *OpusInfo) String() string { return proto.CompactTextString(m) }
func (*OpusInfo) ProtoMessage()    {}
func (*OpusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{3}
}

func (m *OpusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpusInfo.Unmarshal(m, b)
}
func (m *OpusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpusInfo.Marshal(b, m, deterministic)
}
func (m *OpusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpusInfo.Merge(m, src)
}
func (m *OpusInfo) XXX_Size() int {
	return xxx_messageInfo_OpusInfo.Size(m)
}
func (m *OpusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpusInfo proto.InternalMessageInfo

func (m *OpusInfo) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *OpusInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *OpusInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ReqActivityAdd struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string     `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string     `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 int32      `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string     `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,8,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string     `protobuf:"bytes,9,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Require              string     `protobuf:"bytes,10,opt,name=require,proto3" json:"require,omitempty"`
	Limit                uint32     `protobuf:"varint,11,opt,name=limit,proto3" json:"limit,omitempty"`
	Show                 uint32     `protobuf:"varint,12,opt,name=show,proto3" json:"show,omitempty"`
	Prize                *PrizeInfo `protobuf:"bytes,13,opt,name=prize,proto3" json:"prize,omitempty"`
	Assets               []string   `protobuf:"bytes,14,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string   `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Participants         []string   `protobuf:"bytes,16,rep,name=participants,proto3" json:"participants,omitempty"`
	Targets              []string   `protobuf:"bytes,17,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqActivityAdd) Reset()         { *m = ReqActivityAdd{} }
func (m *ReqActivityAdd) String() string { return proto.CompactTextString(m) }
func (*ReqActivityAdd) ProtoMessage()    {}
func (*ReqActivityAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{4}
}

func (m *ReqActivityAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityAdd.Unmarshal(m, b)
}
func (m *ReqActivityAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityAdd.Marshal(b, m, deterministic)
}
func (m *ReqActivityAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityAdd.Merge(m, src)
}
func (m *ReqActivityAdd) XXX_Size() int {
	return xxx_messageInfo_ReqActivityAdd.Size(m)
}
func (m *ReqActivityAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityAdd proto.InternalMessageInfo

func (m *ReqActivityAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqActivityAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqActivityAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqActivityAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqActivityAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityAdd) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqActivityAdd) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqActivityAdd) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ReqActivityAdd) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ReqActivityAdd) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqActivityAdd) GetShow() uint32 {
	if m != nil {
		return m.Show
	}
	return 0
}

func (m *ReqActivityAdd) GetPrize() *PrizeInfo {
	if m != nil {
		return m.Prize
	}
	return nil
}

func (m *ReqActivityAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqActivityAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqActivityAdd) GetParticipants() []string {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *ReqActivityAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqActivityUpdate struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string     `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string     `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,7,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string     `protobuf:"bytes,8,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Require              string     `protobuf:"bytes,9,opt,name=require,proto3" json:"require,omitempty"`
	Limit                uint32     `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqActivityUpdate) Reset()         { *m = ReqActivityUpdate{} }
func (m *ReqActivityUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqActivityUpdate) ProtoMessage()    {}
func (*ReqActivityUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{5}
}

func (m *ReqActivityUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityUpdate.Unmarshal(m, b)
}
func (m *ReqActivityUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityUpdate.Marshal(b, m, deterministic)
}
func (m *ReqActivityUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityUpdate.Merge(m, src)
}
func (m *ReqActivityUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqActivityUpdate.Size(m)
}
func (m *ReqActivityUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityUpdate proto.InternalMessageInfo

func (m *ReqActivityUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqActivityUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqActivityUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityUpdate) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqActivityUpdate) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqActivityUpdate) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ReqActivityUpdate) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ReqActivityUpdate) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReqActivityPrize struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string      `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Operator             string      `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Ranks                []*RankInfo `protobuf:"bytes,5,rep,name=ranks,proto3" json:"ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReqActivityPrize) Reset()         { *m = ReqActivityPrize{} }
func (m *ReqActivityPrize) String() string { return proto.CompactTextString(m) }
func (*ReqActivityPrize) ProtoMessage()    {}
func (*ReqActivityPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{6}
}

func (m *ReqActivityPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityPrize.Unmarshal(m, b)
}
func (m *ReqActivityPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityPrize.Marshal(b, m, deterministic)
}
func (m *ReqActivityPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityPrize.Merge(m, src)
}
func (m *ReqActivityPrize) XXX_Size() int {
	return xxx_messageInfo_ReqActivityPrize.Size(m)
}
func (m *ReqActivityPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityPrize.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityPrize proto.InternalMessageInfo

func (m *ReqActivityPrize) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityPrize) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityPrize) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ReqActivityPrize) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityPrize) GetRanks() []*RankInfo {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type ReqActivityOpus struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Rank                 uint32   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Asset                string   `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqActivityOpus) Reset()         { *m = ReqActivityOpus{} }
func (m *ReqActivityOpus) String() string { return proto.CompactTextString(m) }
func (*ReqActivityOpus) ProtoMessage()    {}
func (*ReqActivityOpus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{7}
}

func (m *ReqActivityOpus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityOpus.Unmarshal(m, b)
}
func (m *ReqActivityOpus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityOpus.Marshal(b, m, deterministic)
}
func (m *ReqActivityOpus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityOpus.Merge(m, src)
}
func (m *ReqActivityOpus) XXX_Size() int {
	return xxx_messageInfo_ReqActivityOpus.Size(m)
}
func (m *ReqActivityOpus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityOpus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityOpus proto.InternalMessageInfo

func (m *ReqActivityOpus) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityOpus) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *ReqActivityOpus) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqActivityOpus) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqActivityOpus) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqActivityState struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqActivityState) Reset()         { *m = ReqActivityState{} }
func (m *ReqActivityState) String() string { return proto.CompactTextString(m) }
func (*ReqActivityState) ProtoMessage()    {}
func (*ReqActivityState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{8}
}

func (m *ReqActivityState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityState.Unmarshal(m, b)
}
func (m *ReqActivityState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityState.Marshal(b, m, deterministic)
}
func (m *ReqActivityState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityState.Merge(m, src)
}
func (m *ReqActivityState) XXX_Size() int {
	return xxx_messageInfo_ReqActivityState.Size(m)
}
func (m *ReqActivityState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityState.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityState proto.InternalMessageInfo

func (m *ReqActivityState) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityState) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityState) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReplyActivityInfo struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *ActivityInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyActivityInfo) Reset()         { *m = ReplyActivityInfo{} }
func (m *ReplyActivityInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyActivityInfo) ProtoMessage()    {}
func (*ReplyActivityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{9}
}

func (m *ReplyActivityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyActivityInfo.Unmarshal(m, b)
}
func (m *ReplyActivityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyActivityInfo.Marshal(b, m, deterministic)
}
func (m *ReplyActivityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyActivityInfo.Merge(m, src)
}
func (m *ReplyActivityInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyActivityInfo.Size(m)
}
func (m *ReplyActivityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyActivityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyActivityInfo proto.InternalMessageInfo

func (m *ReplyActivityInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyActivityInfo) GetInfo() *ActivityInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyActivityList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32          `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32          `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32          `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*ActivityInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyActivityList) Reset()         { *m = ReplyActivityList{} }
func (m *ReplyActivityList) String() string { return proto.CompactTextString(m) }
func (*ReplyActivityList) ProtoMessage()    {}
func (*ReplyActivityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{10}
}

func (m *ReplyActivityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyActivityList.Unmarshal(m, b)
}
func (m *ReplyActivityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyActivityList.Marshal(b, m, deterministic)
}
func (m *ReplyActivityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyActivityList.Merge(m, src)
}
func (m *ReplyActivityList) XXX_Size() int {
	return xxx_messageInfo_ReplyActivityList.Size(m)
}
func (m *ReplyActivityList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyActivityList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyActivityList proto.InternalMessageInfo

func (m *ReplyActivityList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyActivityList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyActivityList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyActivityList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyActivityList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyActivityList) GetList() []*ActivityInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyPairList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*PairInfo  `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyPairList) Reset()         { *m = ReplyPairList{} }
func (m *ReplyPairList) String() string { return proto.CompactTextString(m) }
func (*ReplyPairList) ProtoMessage()    {}
func (*ReplyPairList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{11}
}

func (m *ReplyPairList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyPairList.Unmarshal(m, b)
}
func (m *ReplyPairList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyPairList.Marshal(b, m, deterministic)
}
func (m *ReplyPairList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyPairList.Merge(m, src)
}
func (m *ReplyPairList) XXX_Size() int {
	return xxx_messageInfo_ReplyPairList.Size(m)
}
func (m *ReplyPairList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyPairList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyPairList proto.InternalMessageInfo

func (m *ReplyPairList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyPairList) GetList() []*PairInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivityInfo)(nil), "omo.msp.favorite.ActivityInfo")
	proto.RegisterType((*PrizeInfo)(nil), "omo.msp.favorite.PrizeInfo")
	proto.RegisterType((*RankInfo)(nil), "omo.msp.favorite.RankInfo")
	proto.RegisterType((*OpusInfo)(nil), "omo.msp.favorite.OpusInfo")
	proto.RegisterType((*ReqActivityAdd)(nil), "omo.msp.favorite.ReqActivityAdd")
	proto.RegisterType((*ReqActivityUpdate)(nil), "omo.msp.favorite.ReqActivityUpdate")
	proto.RegisterType((*ReqActivityPrize)(nil), "omo.msp.favorite.ReqActivityPrize")
	proto.RegisterType((*ReqActivityOpus)(nil), "omo.msp.favorite.ReqActivityOpus")
	proto.RegisterType((*ReqActivityState)(nil), "omo.msp.favorite.ReqActivityState")
	proto.RegisterType((*ReplyActivityInfo)(nil), "omo.msp.favorite.ReplyActivityInfo")
	proto.RegisterType((*ReplyActivityList)(nil), "omo.msp.favorite.ReplyActivityList")
	proto.RegisterType((*ReplyPairList)(nil), "omo.msp.favorite.ReplyPairList")
}

func init() {
	proto.RegisterFile("proto/favorite/activity.proto", fileDescriptor_3d14c73187376933)
}

var fileDescriptor_3d14c73187376933 = []byte{
	// 1115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x45, 0x49, 0x96, 0x46, 0x96, 0x2d, 0x6f, 0xdd, 0x74, 0x61, 0x37, 0x8d, 0xca, 0x5c,
	0x7c, 0x52, 0x5a, 0x17, 0xbd, 0x16, 0xb0, 0x51, 0xd4, 0x48, 0xe0, 0xd4, 0x0e, 0xdd, 0x00, 0x41,
	0x6f, 0x6b, 0x71, 0xed, 0x12, 0x96, 0x48, 0x7a, 0xb9, 0x52, 0xea, 0x1c, 0x0a, 0xf4, 0x35, 0xfa,
	0x1c, 0x7d, 0x90, 0xa2, 0x87, 0x3e, 0x4c, 0x4f, 0xc1, 0xcc, 0x92, 0xf4, 0x4a, 0x21, 0x25, 0xd9,
	0xce, 0x6d, 0x67, 0x77, 0x76, 0x7e, 0xbe, 0x99, 0x6f, 0x96, 0x84, 0x27, 0x89, 0x8a, 0x75, 0xfc,
	0xfc, 0x42, 0x4c, 0x63, 0x15, 0x6a, 0xf9, 0x5c, 0x0c, 0x75, 0x38, 0x0d, 0xf5, 0xcd, 0x80, 0xf6,
	0x59, 0x2f, 0x1e, 0xc7, 0x83, 0x71, 0x9a, 0x0c, 0x72, 0x85, 0x9d, 0xdd, 0xb9, 0x0b, 0xc3, 0x78,
	0x3c, 0x8e, 0x23, 0xa3, 0xee, 0xfd, 0xd3, 0x80, 0xf5, 0x83, 0xcc, 0xc2, 0x8b, 0xe8, 0x22, 0x66,
	0x3d, 0x70, 0x27, 0x61, 0xc0, 0x9d, 0xbe, 0xb3, 0xd7, 0xf6, 0x71, 0xc9, 0x36, 0xa0, 0x16, 0x06,
	0xbc, 0xd6, 0x77, 0xf6, 0xea, 0x7e, 0x2d, 0x0c, 0x18, 0x83, 0xba, 0xbe, 0x49, 0x24, 0x77, 0xfb,
	0xce, 0x5e, 0xd7, 0xa7, 0x35, 0xe3, 0xb0, 0x36, 0x54, 0x52, 0x68, 0x19, 0xf0, 0x7a, 0xdf, 0xd9,
	0x73, 0xfd, 0x5c, 0xc4, 0x93, 0x49, 0x12, 0xd0, 0x49, 0xc3, 0x9c, 0x64, 0x62, 0x71, 0x27, 0x56,
	0xbc, 0x49, 0xde, 0x72, 0x91, 0xed, 0x40, 0x2b, 0x4e, 0xa4, 0xa2, 0xa3, 0x35, 0x3a, 0x2a, 0x64,
	0xf4, 0x1e, 0x89, 0xb1, 0xe4, 0x2d, 0xda, 0xa7, 0x35, 0x7b, 0x0c, 0x4d, 0x25, 0xc7, 0x42, 0x5d,
	0xf1, 0x36, 0xed, 0x66, 0x12, 0x7a, 0x50, 0xf2, 0x7a, 0x12, 0x2a, 0xc9, 0xc1, 0x78, 0xc8, 0x44,
	0xb6, 0x0d, 0x8d, 0x61, 0x3c, 0x95, 0x8a, 0x77, 0x68, 0xdf, 0x08, 0xb8, 0x1b, 0xbf, 0x8b, 0xa4,
	0xe2, 0xeb, 0x66, 0x97, 0x04, 0x36, 0x80, 0x3a, 0x06, 0xcc, 0xbb, 0x7d, 0x67, 0xaf, 0xb3, 0xbf,
	0x33, 0x98, 0x07, 0x78, 0xf0, 0xa3, 0xd0, 0x12, 0xb1, 0xf3, 0x49, 0x8f, 0x7d, 0x0b, 0x8d, 0x64,
	0x24, 0x86, 0x92, 0x6f, 0xd0, 0x85, 0xdd, 0x8f, 0x2f, 0x9c, 0xe2, 0x31, 0xdd, 0x30, 0x9a, 0xec,
	0x4b, 0x68, 0xc7, 0xea, 0x52, 0x44, 0xe1, 0x7b, 0xa9, 0xf8, 0x26, 0x39, 0xbf, 0xdd, 0xc0, 0xb0,
	0x46, 0xe1, 0x38, 0xd4, 0xbc, 0x47, 0x88, 0x1b, 0x01, 0x93, 0x4e, 0xb5, 0xd0, 0x93, 0x94, 0x6f,
	0xd1, 0x76, 0x26, 0x21, 0x40, 0xe9, 0x6f, 0xf1, 0x3b, 0xce, 0x4c, 0x79, 0x70, 0x4d, 0x21, 0xa9,
	0xf0, 0xbd, 0xe4, 0x9f, 0x55, 0x86, 0x84, 0xc7, 0x59, 0x48, 0xb8, 0xa4, 0x2a, 0x8b, 0xcb, 0x94,
	0x6f, 0xf7, 0x5d, 0xc4, 0x19, 0xd7, 0xe8, 0x52, 0xa4, 0xa9, 0xd4, 0x29, 0xff, 0x9c, 0x76, 0x33,
	0x09, 0x71, 0xd6, 0x42, 0x5d, 0xe2, 0xc1, 0x63, 0x3a, 0xc8, 0x45, 0xf6, 0x03, 0xac, 0x27, 0x42,
	0xe9, 0x70, 0x18, 0x26, 0x22, 0xd2, 0x29, 0xff, 0xa2, 0xef, 0x96, 0x63, 0x78, 0x2a, 0x42, 0x45,
	0xee, 0x67, 0xf4, 0xd9, 0x3e, 0x34, 0xe3, 0x64, 0x92, 0xca, 0x94, 0xf3, 0xaa, 0x9b, 0x27, 0xc9,
	0x24, 0xa5, 0x9b, 0x99, 0xa6, 0x27, 0xa1, 0x5d, 0x64, 0x53, 0xb4, 0x8b, 0x63, 0xb5, 0x0b, 0x83,
	0x7a, 0x20, 0xd3, 0x21, 0xb5, 0x74, 0xdb, 0xa7, 0x35, 0xfb, 0x06, 0x1a, 0x4a, 0x44, 0x57, 0x29,
	0x77, 0xab, 0xfc, 0xf8, 0x22, 0xba, 0x32, 0x00, 0x91, 0xa2, 0xf7, 0x12, 0x5a, 0xf9, 0x16, 0x56,
	0x28, 0x8c, 0x02, 0xf9, 0x3b, 0xb9, 0xe9, 0xfa, 0x46, 0x28, 0x7c, 0xd7, 0x2c, 0xdf, 0xd4, 0x78,
	0x93, 0x48, 0x67, 0xec, 0x31, 0x82, 0x77, 0x0c, 0xad, 0x3c, 0x0d, 0xbc, 0x85, 0x0e, 0x32, 0x53,
	0xb4, 0xc6, 0x5b, 0x04, 0x75, 0x66, 0xca, 0x08, 0x56, 0xdb, 0xbb, 0x76, 0xdb, 0x7b, 0xff, 0xbb,
	0xb0, 0xe1, 0xcb, 0xeb, 0x9c, 0xd6, 0x07, 0x41, 0x50, 0x0a, 0x43, 0xd1, 0xed, 0x35, 0xbb, 0xdb,
	0x2b, 0x8c, 0xde, 0x32, 0xa6, 0x6e, 0x33, 0x26, 0x9f, 0x05, 0x48, 0xed, 0x46, 0x36, 0x0b, 0x6c,
	0xf6, 0x36, 0xe7, 0xd8, 0x9b, 0x73, 0x69, 0xed, 0xae, 0x5c, 0x6a, 0xdd, 0x8f, 0x4b, 0xed, 0x79,
	0x2e, 0x2d, 0x1c, 0x09, 0x86, 0x65, 0x1d, 0x9b, 0x65, 0x39, 0x9b, 0xd6, 0xcb, 0xd8, 0xd4, 0x5d,
	0x99, 0x4d, 0xb7, 0xcc, 0xd9, 0x98, 0x61, 0x4e, 0xce, 0xb2, 0x4d, 0x8b, 0x65, 0xde, 0x1c, 0x67,
	0x7a, 0x74, 0x36, 0xcb, 0x0b, 0x8b, 0x71, 0x5b, 0x33, 0x8c, 0xf3, 0xfe, 0xae, 0xc1, 0x96, 0x55,
	0xfc, 0x37, 0x34, 0x6c, 0x4b, 0xa6, 0x7a, 0x59, 0x73, 0xde, 0xad, 0xf6, 0x76, 0x9d, 0x1b, 0x15,
	0x75, 0x6e, 0xde, 0xb5, 0xce, 0x6b, 0xf7, 0xab, 0x73, 0x6b, 0x41, 0x9d, 0xdb, 0x15, 0x75, 0x06,
	0xab, 0xce, 0xde, 0x5f, 0x0e, 0xf4, 0x2c, 0xd8, 0xa8, 0x80, 0x2b, 0xa2, 0x96, 0x8f, 0x13, 0xd7,
	0x1a, 0x27, 0x36, 0x36, 0xf5, 0x39, 0x6c, 0x8a, 0x51, 0xd3, 0x58, 0x75, 0xd4, 0xfc, 0xe9, 0xc0,
	0xa6, 0x15, 0x1c, 0x8e, 0x8a, 0xf2, 0xd8, 0x68, 0x70, 0xd4, 0xca, 0x06, 0x87, 0x5b, 0x3e, 0x38,
	0xea, 0x33, 0x75, 0x5e, 0x50, 0x51, 0xef, 0xed, 0x0c, 0x3e, 0x67, 0xba, 0xbc, 0xab, 0x6c, 0x0b,
	0xb5, 0xb9, 0xbc, 0x6f, 0x1f, 0x2c, 0xd7, 0x7e, 0xb0, 0xbc, 0x3f, 0xb0, 0x61, 0x93, 0xd1, 0xcd,
	0xcc, 0x67, 0xc8, 0xf7, 0x85, 0xb2, 0x43, 0x1d, 0xf1, 0xa4, 0x04, 0x25, 0xbc, 0x74, 0x46, 0x4a,
	0xc5, 0xe3, 0xb7, 0x0f, 0xf5, 0x30, 0xba, 0x88, 0xc9, 0x77, 0x67, 0xff, 0xab, 0x8f, 0x2f, 0xd9,
	0x4e, 0x7c, 0xd2, 0xf5, 0xfe, 0x73, 0xe6, 0x02, 0x38, 0x0e, 0x53, 0x7d, 0xdf, 0x00, 0xb6, 0xa1,
	0xa1, 0x63, 0x2d, 0x46, 0x59, 0x15, 0x8c, 0x80, 0xbb, 0x89, 0xb8, 0x94, 0x79, 0xe6, 0x46, 0xc0,
	0x82, 0xe1, 0x82, 0x8a, 0xd0, 0xf5, 0x69, 0x8d, 0x20, 0x45, 0x93, 0xf1, 0xb9, 0x34, 0x05, 0xe8,
	0xfa, 0x99, 0x84, 0x89, 0x8d, 0xc2, 0x54, 0xf3, 0x26, 0xf5, 0xcc, 0xd2, 0xc4, 0x50, 0xd7, 0x9b,
	0x42, 0x97, 0x42, 0xc4, 0xb7, 0xf5, 0x21, 0x39, 0x0d, 0x32, 0xdf, 0xb5, 0xa5, 0x8f, 0x37, 0xe9,
	0xed, 0xff, 0xdb, 0x86, 0xcd, 0xa2, 0x51, 0xa4, 0x9a, 0x86, 0x43, 0xc9, 0x5e, 0x43, 0xf3, 0x20,
	0x08, 0x4e, 0x22, 0xc9, 0xfa, 0x65, 0x4e, 0xed, 0xc7, 0x6a, 0xe7, 0x59, 0x45, 0x58, 0x76, 0x8a,
	0xde, 0x23, 0xf6, 0x33, 0x34, 0x8f, 0xa4, 0x46, 0x93, 0xa5, 0x79, 0x5c, 0x4f, 0x64, 0xaa, 0x51,
	0x75, 0x55, 0x7b, 0x27, 0xb0, 0x76, 0x24, 0x35, 0x01, 0xf5, 0x40, 0x83, 0x68, 0xc3, 0x7b, 0xc4,
	0xde, 0x40, 0xe7, 0x48, 0xea, 0xc3, 0x9b, 0x9f, 0xc2, 0x91, 0x96, 0x8a, 0x3d, 0xad, 0x34, 0x6a,
	0x14, 0x56, 0x35, 0xfb, 0x16, 0xc0, 0x4c, 0xf5, 0x43, 0x91, 0x4a, 0xf6, 0x6c, 0x21, 0x9c, 0x46,
	0x71, 0x55, 0x04, 0x5e, 0xe6, 0x96, 0x7f, 0xc1, 0x77, 0xa8, 0x1a, 0x04, 0x0c, 0x64, 0x67, 0xb7,
	0xc2, 0x66, 0x16, 0xe5, 0x31, 0xac, 0x1b, 0x5b, 0x07, 0xe6, 0xa5, 0x7b, 0x98, 0xb5, 0x57, 0xd0,
	0xcd, 0x23, 0x33, 0x1f, 0x96, 0x0f, 0x33, 0xf7, 0x3a, 0x0f, 0xce, 0x74, 0x3a, 0xf3, 0x16, 0x82,
	0x48, 0xc3, 0xae, 0xd2, 0x64, 0xd1, 0x3d, 0x19, 0x76, 0x67, 0xf8, 0x89, 0xf0, 0x09, 0x0c, 0xbe,
	0x80, 0xb6, 0x2f, 0xc7, 0xf1, 0x54, 0xae, 0xd0, 0xe1, 0x4b, 0x4c, 0xbd, 0x82, 0xf6, 0x41, 0x92,
	0xc8, 0x28, 0x58, 0xc1, 0xd4, 0xd3, 0x0a, 0x53, 0xf9, 0x10, 0xa1, 0x54, 0x3b, 0x67, 0x93, 0x73,
	0xad, 0xc4, 0x50, 0x7f, 0x1a, 0x83, 0xa7, 0x79, 0x39, 0x4e, 0xe8, 0x0b, 0x9e, 0x7d, 0xbd, 0x10,
	0x3d, 0x54, 0x5a, 0x96, 0xf1, 0x29, 0x74, 0x8c, 0x45, 0xf3, 0x90, 0x2f, 0x2e, 0x07, 0xe9, 0x2c,
	0xb1, 0x78, 0xd8, 0xfb, 0x75, 0x63, 0xf6, 0x3f, 0xfa, 0xbc, 0x49, 0xf2, 0x77, 0x1f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x9e, 0x2c, 0x7d, 0x91, 0x0f, 0x00, 0x00,
}
