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
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32     `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64      `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64      `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string     `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string     `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string     `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string     `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Require              string     `protobuf:"bytes,10,opt,name=require,proto3" json:"require,omitempty"`
	Cover                string     `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Owner                string     `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,13,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,14,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string     `protobuf:"bytes,15,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Tags                 []string   `protobuf:"bytes,16,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string   `protobuf:"bytes,17,rep,name=assets,proto3" json:"assets,omitempty"`
	Participants         []string   `protobuf:"bytes,18,rep,name=participants,proto3" json:"participants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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

func (m *ActivityInfo) GetParticipants() []string {
	if m != nil {
		return m.Participants
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
	Assets               []string   `protobuf:"bytes,11,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string   `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	Participants         []string   `protobuf:"bytes,13,rep,name=participants,proto3" json:"participants,omitempty"`
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
	List                 []*ActivityInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
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

func (m *ReplyActivityList) GetList() []*ActivityInfo {
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
}

func init() {
	proto.RegisterFile("proto/favorite/activity.proto", fileDescriptor_3d14c73187376933)
}

var fileDescriptor_3d14c73187376933 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0xb1, 0x9d, 0x38, 0xf1, 0x49, 0x08, 0x61, 0x74, 0x75, 0x35, 0x0a, 0x97, 0x2b, 0x2b,
	0x6c, 0xb2, 0x0a, 0x6a, 0xaa, 0x3e, 0x40, 0x50, 0x25, 0x44, 0x8b, 0x4a, 0x6b, 0x5a, 0xa9, 0xea,
	0x6e, 0xb0, 0x07, 0x34, 0x2a, 0xf6, 0x98, 0xf1, 0x24, 0x15, 0x5d, 0x74, 0xd3, 0xf7, 0xe8, 0x93,
	0xf5, 0x25, 0xfa, 0x06, 0xd5, 0x1c, 0xdb, 0xc1, 0xf9, 0x42, 0x81, 0xec, 0x7c, 0xbe, 0xcf, 0xfc,
	0xfe, 0x33, 0x51, 0xe0, 0x30, 0x55, 0x52, 0xcb, 0xe3, 0x6b, 0x36, 0x95, 0x4a, 0x68, 0x7e, 0xcc,
	0x42, 0x2d, 0xa6, 0x42, 0xdf, 0x0f, 0xd1, 0x4f, 0xba, 0x32, 0x96, 0xc3, 0x38, 0x4b, 0x87, 0x65,
	0x42, 0xef, 0x60, 0xa1, 0x20, 0x94, 0x71, 0x2c, 0x93, 0x3c, 0xbd, 0xff, 0xc7, 0x81, 0xf6, 0xb8,
	0xe8, 0x70, 0x96, 0x5c, 0x4b, 0xd2, 0x05, 0x67, 0x22, 0x22, 0x6a, 0xf9, 0xd6, 0xc0, 0x0b, 0xcc,
	0x27, 0xe9, 0x80, 0x2d, 0x22, 0x6a, 0xfb, 0xd6, 0xa0, 0x16, 0xd8, 0x22, 0x22, 0x04, 0x6a, 0xfa,
	0x3e, 0xe5, 0xd4, 0xf1, 0xad, 0xc1, 0x6e, 0x80, 0xdf, 0x84, 0x42, 0x23, 0x54, 0x9c, 0x69, 0x1e,
	0xd1, 0x9a, 0x6f, 0x0d, 0x9c, 0xa0, 0x34, 0x4d, 0x64, 0x92, 0x46, 0x18, 0xa9, 0xe7, 0x91, 0xc2,
	0x9c, 0xd5, 0x48, 0x45, 0x5d, 0x9c, 0x56, 0x9a, 0xa4, 0x07, 0x4d, 0x99, 0x72, 0x85, 0xa1, 0x06,
	0x86, 0x66, 0xb6, 0x99, 0x9e, 0xb0, 0x98, 0xd3, 0x26, 0xfa, 0xf1, 0x9b, 0xfc, 0x0b, 0xae, 0xe2,
	0x31, 0x53, 0x5f, 0xa9, 0x87, 0xde, 0xc2, 0x32, 0x13, 0x14, 0xbf, 0x9b, 0x08, 0xc5, 0x29, 0xe4,
	0x13, 0x0a, 0x93, 0xfc, 0x03, 0xf5, 0x50, 0x4e, 0xb9, 0xa2, 0x2d, 0xf4, 0xe7, 0x86, 0xf1, 0xca,
	0x6f, 0x09, 0x57, 0xb4, 0x9d, 0x7b, 0xd1, 0x20, 0x43, 0xa8, 0x99, 0x85, 0xe9, 0xae, 0x6f, 0x0d,
	0x5a, 0xa3, 0xde, 0x70, 0x11, 0xf0, 0xf0, 0x35, 0xd3, 0xdc, 0xb0, 0x0b, 0x30, 0x8f, 0xbc, 0x80,
	0x7a, 0x7a, 0xcb, 0x42, 0x4e, 0x3b, 0x58, 0x70, 0xb0, 0x5c, 0xf0, 0xde, 0x84, 0xb1, 0x22, 0xcf,
	0x24, 0xff, 0x81, 0x27, 0xd5, 0x0d, 0x4b, 0xc4, 0x77, 0xae, 0xe8, 0x1e, 0x0e, 0x7f, 0x70, 0x20,
	0x70, 0x76, 0x93, 0xd1, 0xae, 0xef, 0x98, 0x23, 0x9b, 0x6f, 0x73, 0x64, 0x96, 0x65, 0x5c, 0x67,
	0x74, 0x1f, 0xbd, 0x85, 0x45, 0xfa, 0xd0, 0x4e, 0x99, 0xd2, 0x22, 0x14, 0x29, 0x4b, 0x74, 0x46,
	0x09, 0x46, 0xe7, 0x7c, 0xfd, 0x9f, 0x0e, 0x74, 0x02, 0x7e, 0x57, 0xca, 0x3e, 0x8e, 0xa2, 0x19,
	0x55, 0xab, 0x42, 0x75, 0x46, 0xc3, 0xae, 0xd2, 0x78, 0x60, 0xed, 0xcc, 0xb1, 0x9e, 0x11, 0xad,
	0x55, 0x89, 0x96, 0x77, 0xc5, 0x48, 0x5f, 0x2f, 0xee, 0x4a, 0x55, 0x5d, 0x77, 0x41, 0xdd, 0x92,
	0x75, 0xe3, 0xa9, 0xac, 0x9b, 0xcf, 0x63, 0xed, 0x2d, 0xb2, 0x5e, 0x7f, 0x65, 0x1e, 0x88, 0xb7,
	0xe6, 0x88, 0x97, 0xea, 0xb4, 0x2b, 0xea, 0x2c, 0xaa, 0xb0, 0xbb, 0x42, 0x85, 0x5f, 0x36, 0xec,
	0x57, 0x54, 0xf8, 0x84, 0xaf, 0x62, 0xc5, 0xf3, 0x2b, 0xa5, 0xb1, 0x57, 0x5e, 0xf8, 0x4d, 0x44,
	0xa8, 0x02, 0xaf, 0xaf, 0x01, 0xee, 0x3e, 0x15, 0x78, 0xe3, 0x79, 0xc0, 0x9b, 0x8f, 0x00, 0xf7,
	0xe6, 0x80, 0xf7, 0x7f, 0x18, 0x3e, 0xe9, 0xed, 0xfd, 0xdc, 0xcf, 0xd3, 0x2b, 0x70, 0x33, 0xcd,
	0xf4, 0x24, 0x43, 0x44, 0xad, 0xd1, 0xe1, 0xf2, 0x02, 0x58, 0x74, 0x89, 0x49, 0x41, 0x91, 0x4c,
	0x46, 0x50, 0x13, 0xc9, 0xb5, 0x44, 0x88, 0xad, 0xd1, 0xff, 0xcb, 0x45, 0xd5, 0x21, 0x01, 0xe6,
	0x2e, 0xcd, 0x3f, 0x17, 0x99, 0xde, 0x62, 0xfe, 0xad, 0xc8, 0x34, 0xb5, 0x7d, 0x67, 0x93, 0xf9,
	0x26, 0x77, 0xf4, 0xbb, 0x0e, 0x7b, 0xa5, 0xfb, 0x92, 0xab, 0xa9, 0x08, 0x39, 0xf9, 0x00, 0xee,
	0x38, 0x8a, 0x2e, 0x12, 0x4e, 0xfc, 0x55, 0x83, 0xab, 0x6f, 0xba, 0x77, 0xb4, 0x66, 0xb5, 0xea,
	0xa8, 0xfe, 0x0e, 0x79, 0x07, 0xee, 0x29, 0xd7, 0xa6, 0xe5, 0xca, 0xb3, 0xdc, 0x4d, 0x78, 0xa6,
	0x4d, 0xea, 0xa6, 0xfd, 0x2e, 0xa0, 0x71, 0xca, 0x35, 0xc2, 0xda, 0xb2, 0xa1, 0xe9, 0xd1, 0xdf,
	0x21, 0x9f, 0x01, 0xf2, 0xc7, 0x71, 0xc2, 0x32, 0x4e, 0x8e, 0x1e, 0x3d, 0x77, 0x9e, 0xb8, 0xe9,
	0xaa, 0x6f, 0xca, 0xce, 0x1f, 0xcd, 0xa3, 0x5d, 0xbf, 0xad, 0x59, 0xa4, 0x77, 0xb0, 0xa6, 0x67,
	0xb1, 0xe5, 0x39, 0xb4, 0xf3, 0x5e, 0xe3, 0xfc, 0x67, 0x61, 0xbb, 0x6e, 0x67, 0xe0, 0x05, 0x3c,
	0x96, 0x53, 0xbe, 0x81, 0x2e, 0xeb, 0x5a, 0x15, 0x87, 0x3c, 0x03, 0x6f, 0x9c, 0xa6, 0x3c, 0x89,
	0xb6, 0x68, 0x55, 0x6c, 0xf5, 0x16, 0x5a, 0x97, 0x93, 0x2b, 0xad, 0x58, 0xa8, 0xb7, 0x6e, 0x76,
	0xd2, 0xfd, 0xd2, 0x99, 0xff, 0x63, 0x72, 0xe5, 0xa2, 0xfd, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x16, 0x92, 0x13, 0xa0, 0xe2, 0x08, 0x00, 0x00,
}
