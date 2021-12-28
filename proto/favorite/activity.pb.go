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
	Tags                 []string    `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string    `protobuf:"bytes,18,rep,name=assets,proto3" json:"assets,omitempty"`
	Participants         []*PairInfo `protobuf:"bytes,19,rep,name=participants,proto3" json:"participants,omitempty"`
	Targets              []string    `protobuf:"bytes,20,rep,name=targets,proto3" json:"targets,omitempty"`
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

func (m *ActivityInfo) GetParticipants() []*PairInfo {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *ActivityInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
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
	Assets               []string   `protobuf:"bytes,12,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string   `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	Participants         []string   `protobuf:"bytes,14,rep,name=participants,proto3" json:"participants,omitempty"`
	Targets              []string   `protobuf:"bytes,15,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqActivityAdd) Reset()         { *m = ReqActivityAdd{} }
func (m *ReqActivityAdd) String() string { return proto.CompactTextString(m) }
func (*ReqActivityAdd) ProtoMessage()    {}
func (*ReqActivityAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{1}
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
	return fileDescriptor_3d14c73187376933, []int{2}
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
	return fileDescriptor_3d14c73187376933, []int{3}
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
	return fileDescriptor_3d14c73187376933, []int{4}
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
	return fileDescriptor_3d14c73187376933, []int{5}
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
	proto.RegisterType((*ReqActivityAdd)(nil), "omo.msp.favorite.ReqActivityAdd")
	proto.RegisterType((*ReqActivityUpdate)(nil), "omo.msp.favorite.ReqActivityUpdate")
	proto.RegisterType((*ReplyActivityInfo)(nil), "omo.msp.favorite.ReplyActivityInfo")
	proto.RegisterType((*ReplyActivityList)(nil), "omo.msp.favorite.ReplyActivityList")
	proto.RegisterType((*ReplyPairList)(nil), "omo.msp.favorite.ReplyPairList")
}

func init() {
	proto.RegisterFile("proto/favorite/activity.proto", fileDescriptor_3d14c73187376933)
}

var fileDescriptor_3d14c73187376933 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x8e, 0xdb, 0x36,
	0x10, 0xc7, 0x23, 0x7f, 0xc8, 0xeb, 0xb1, 0xe5, 0x75, 0xd8, 0x45, 0x41, 0x78, 0x9b, 0x46, 0x50,
	0x2e, 0x3e, 0x39, 0xa8, 0x8b, 0x5e, 0x0b, 0x78, 0x51, 0x74, 0x91, 0x22, 0x69, 0x52, 0x6d, 0x03,
	0x14, 0xbd, 0x71, 0x25, 0xae, 0x41, 0xd4, 0x12, 0xb5, 0x14, 0xed, 0xc2, 0x3d, 0xf4, 0x01, 0xfa,
	0x3a, 0xed, 0xb5, 0xaf, 0xd1, 0xe7, 0x29, 0x38, 0x94, 0xbc, 0xb2, 0xd7, 0xda, 0x38, 0xeb, 0xdc,
	0x38, 0xc3, 0xe1, 0x9f, 0xc3, 0xf9, 0x8d, 0x06, 0x82, 0x67, 0x99, 0x92, 0x5a, 0xbe, 0xbc, 0x61,
	0x2b, 0xa9, 0x84, 0xe6, 0x2f, 0x59, 0xa4, 0xc5, 0x4a, 0xe8, 0xf5, 0x04, 0xfd, 0x64, 0x28, 0x13,
	0x39, 0x49, 0xf2, 0x6c, 0x52, 0x06, 0x8c, 0xce, 0x77, 0x0e, 0x44, 0x32, 0x49, 0x64, 0x6a, 0xc3,
	0x83, 0x7f, 0x5b, 0xd0, 0x9f, 0x15, 0x0a, 0xaf, 0xd2, 0x1b, 0x49, 0x86, 0xd0, 0x5c, 0x8a, 0x98,
	0x3a, 0xbe, 0x33, 0xee, 0x86, 0x66, 0x49, 0x06, 0xd0, 0x10, 0x31, 0x6d, 0xf8, 0xce, 0xb8, 0x15,
	0x36, 0x44, 0x4c, 0x08, 0xb4, 0xf4, 0x3a, 0xe3, 0xb4, 0xe9, 0x3b, 0x63, 0x2f, 0xc4, 0x35, 0xa1,
	0xd0, 0x89, 0x14, 0x67, 0x9a, 0xc7, 0xb4, 0xe5, 0x3b, 0xe3, 0x66, 0x58, 0x9a, 0x66, 0x67, 0x99,
	0xc5, 0xb8, 0xd3, 0xb6, 0x3b, 0x85, 0xb9, 0x39, 0x23, 0x15, 0x75, 0xf1, 0xb6, 0xd2, 0x24, 0x23,
	0x38, 0x91, 0x19, 0x57, 0xb8, 0xd5, 0xc1, 0xad, 0x8d, 0x6d, 0x6e, 0x4f, 0x59, 0xc2, 0xe9, 0x09,
	0xfa, 0x71, 0x4d, 0x3e, 0x07, 0x57, 0xf1, 0x84, 0xa9, 0xdf, 0x68, 0x17, 0xbd, 0x85, 0x65, 0x6e,
	0x50, 0xfc, 0x76, 0x29, 0x14, 0xa7, 0x60, 0x6f, 0x28, 0x4c, 0x72, 0x06, 0xed, 0x48, 0xae, 0xb8,
	0xa2, 0x3d, 0xf4, 0x5b, 0xc3, 0x78, 0xe5, 0xef, 0x29, 0x57, 0xb4, 0x6f, 0xbd, 0x68, 0x90, 0x09,
	0xb4, 0x4c, 0xc2, 0xd4, 0xf3, 0x9d, 0x71, 0x6f, 0x3a, 0x9a, 0xec, 0x16, 0x78, 0xf2, 0x1d, 0xd3,
	0xdc, 0xd4, 0x2e, 0xc4, 0x38, 0xf2, 0x15, 0xb4, 0xb3, 0x05, 0x8b, 0x38, 0x1d, 0xe0, 0x81, 0xf3,
	0xfb, 0x07, 0xde, 0x99, 0x6d, 0x3c, 0x61, 0x23, 0xc9, 0x17, 0xd0, 0x95, 0x6a, 0xce, 0x52, 0xf1,
	0x07, 0x57, 0xf4, 0x14, 0x2f, 0xbf, 0x73, 0x98, 0xb4, 0x16, 0x22, 0x11, 0x9a, 0x0e, 0xb1, 0xe2,
	0xd6, 0x40, 0x0c, 0x6c, 0x9e, 0xd3, 0xa7, 0x7e, 0xd3, 0x14, 0xc2, 0xac, 0x4d, 0x21, 0x58, 0x9e,
	0x73, 0x9d, 0x53, 0x82, 0xde, 0xc2, 0x22, 0xdf, 0x42, 0x3f, 0x63, 0x4a, 0x8b, 0x48, 0x64, 0x2c,
	0xd5, 0x39, 0xfd, 0xcc, 0x6f, 0xee, 0x7f, 0xca, 0x3b, 0x26, 0x14, 0x26, 0xb6, 0x15, 0x6f, 0x0a,
	0xa9, 0x99, 0x9a, 0x1b, 0xe1, 0x33, 0x14, 0x2e, 0xcd, 0xe0, 0xef, 0x26, 0x0c, 0x42, 0x7e, 0x5b,
	0xb6, 0xd0, 0x2c, 0x8e, 0x37, 0x84, 0x9c, 0x0a, 0xa1, 0x4d, 0x65, 0x1b, 0xd5, 0xca, 0xde, 0x71,
	0x6b, 0x6e, 0x71, 0xdb, 0xd0, 0x69, 0x55, 0xe9, 0x94, 0x7d, 0x67, 0xda, 0xa8, 0x5d, 0xf4, 0x5d,
	0xb5, 0x53, 0xdc, 0x9d, 0x4e, 0x29, 0xb9, 0x75, 0x3e, 0x96, 0xdb, 0xc9, 0xe3, 0xb8, 0x75, 0x77,
	0xb9, 0x3d, 0xd8, 0x7e, 0x96, 0x68, 0xaf, 0x4a, 0xf4, 0x8e, 0x5e, 0x7f, 0x8b, 0x5e, 0x49, 0xda,
	0xab, 0x90, 0x0e, 0x76, 0x88, 0x0e, 0x70, 0xaf, 0x96, 0xda, 0xe9, 0x36, 0xb5, 0x7f, 0x1a, 0xf0,
	0xb4, 0x42, 0xed, 0x3d, 0x7e, 0x91, 0x7b, 0x3e, 0xfd, 0x12, 0x65, 0x63, 0xef, 0xc7, 0x76, 0x08,
	0xb4, 0x2a, 0xa0, 0x76, 0x0d, 0x20, 0xf7, 0x63, 0x01, 0x75, 0x1e, 0x07, 0xe8, 0xe4, 0x01, 0x40,
	0xdd, 0x1a, 0x40, 0x50, 0x01, 0x14, 0xfc, 0x69, 0xaa, 0x96, 0x2d, 0xd6, 0x5b, 0x03, 0xf3, 0x1b,
	0x70, 0x73, 0xcd, 0xf4, 0x32, 0xc7, 0xc2, 0xf5, 0xa6, 0xcf, 0xee, 0xa7, 0x85, 0x87, 0xae, 0x30,
	0x28, 0x2c, 0x82, 0xc9, 0x14, 0x5a, 0x22, 0xbd, 0x91, 0x58, 0xda, 0xde, 0xf4, 0xcb, 0xfb, 0x87,
	0xaa, 0x97, 0x84, 0x18, 0x1b, 0xfc, 0xe7, 0xec, 0x24, 0xf0, 0x5a, 0xe4, 0xfa, 0xb1, 0x09, 0x9c,
	0x41, 0x5b, 0x4b, 0xcd, 0x16, 0x98, 0x81, 0x17, 0x5a, 0xc3, 0x78, 0x33, 0x36, 0xe7, 0x79, 0x31,
	0xdd, 0xad, 0x61, 0xfa, 0xc0, 0x2c, 0x10, 0xad, 0x17, 0xe2, 0xda, 0xf4, 0x41, 0xba, 0x4c, 0xae,
	0xb9, 0xe5, 0xea, 0x85, 0x85, 0x65, 0x1e, 0xb6, 0x10, 0xb9, 0xa6, 0x2e, 0xce, 0x98, 0x0f, 0x3e,
	0xcc, 0xc4, 0x06, 0x2b, 0xf0, 0x30, 0x45, 0x33, 0x7e, 0x8e, 0x79, 0xd3, 0xa4, 0xb8, 0xbb, 0xf1,
	0xc1, 0xf9, 0x86, 0x71, 0xd3, 0xbf, 0x3a, 0x70, 0x5a, 0xa6, 0x73, 0xc5, 0xd5, 0x4a, 0x44, 0x9c,
	0xfc, 0x04, 0xee, 0x2c, 0x8e, 0xdf, 0xa6, 0x9c, 0xf8, 0xfb, 0x2e, 0xad, 0x8e, 0xba, 0xd1, 0x8b,
	0x9a, 0xb4, 0xaa, 0x4f, 0x0c, 0x9e, 0x90, 0x1f, 0xc1, 0xbd, 0xe4, 0xda, 0x48, 0xee, 0x7d, 0xc7,
	0xed, 0x92, 0xe7, 0xda, 0x84, 0x1e, 0xaa, 0xf7, 0x16, 0x3a, 0x97, 0x5c, 0x63, 0xa1, 0x8e, 0x14,
	0x34, 0x1a, 0xc1, 0x13, 0xf2, 0x1e, 0x7a, 0x97, 0x5c, 0x5f, 0xac, 0xbf, 0x17, 0x0b, 0xcd, 0x15,
	0x79, 0x5e, 0x2b, 0x6a, 0x03, 0x0e, 0x95, 0xfd, 0x05, 0xc0, 0x8e, 0x96, 0x0b, 0x96, 0x73, 0xf2,
	0xe2, 0xc1, 0x72, 0xda, 0xc0, 0x43, 0x2b, 0xf0, 0x43, 0xa9, 0xfc, 0xb3, 0x19, 0x86, 0xf5, 0x45,
	0x30, 0x89, 0x8c, 0xce, 0x6b, 0x34, 0x8b, 0x2c, 0x5f, 0x43, 0xdf, 0x6a, 0xcd, 0xec, 0xb8, 0x3d,
	0x4e, 0xed, 0x0d, 0x78, 0x65, 0x66, 0x38, 0x6b, 0x8f, 0x94, 0x7b, 0x05, 0xdd, 0x90, 0x27, 0x72,
	0xc5, 0x0f, 0xe8, 0x9e, 0x3a, 0xa9, 0xa2, 0x66, 0x6f, 0xa0, 0x3b, 0xcb, 0x32, 0x9e, 0xc6, 0x07,
	0x48, 0x3d, 0xaf, 0x91, 0x2a, 0x3f, 0x50, 0x6c, 0xc2, 0xde, 0xd5, 0xf2, 0x5a, 0x2b, 0x16, 0xe9,
	0x4f, 0x22, 0x78, 0x31, 0xfc, 0x75, 0xb0, 0xfd, 0xa7, 0x7a, 0xed, 0xa2, 0xfd, 0xf5, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x36, 0x7d, 0xc7, 0x81, 0xf3, 0x0a, 0x00, 0x00,
}
