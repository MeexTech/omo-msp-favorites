// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/notice.proto

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

type NoticeInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Body                 string   `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	Subtitle             string   `protobuf:"bytes,10,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Owner                string   `protobuf:"bytes,11,opt,name=owner,proto3" json:"owner,omitempty"`
	Status               uint32   `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Targets              []string `protobuf:"bytes,13,rep,name=targets,proto3" json:"targets,omitempty"`
	Tags                 []string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoticeInfo) Reset()         { *m = NoticeInfo{} }
func (m *NoticeInfo) String() string { return proto.CompactTextString(m) }
func (*NoticeInfo) ProtoMessage()    {}
func (*NoticeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{0}
}

func (m *NoticeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoticeInfo.Unmarshal(m, b)
}
func (m *NoticeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoticeInfo.Marshal(b, m, deterministic)
}
func (m *NoticeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoticeInfo.Merge(m, src)
}
func (m *NoticeInfo) XXX_Size() int {
	return xxx_messageInfo_NoticeInfo.Size(m)
}
func (m *NoticeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NoticeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NoticeInfo proto.InternalMessageInfo

func (m *NoticeInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *NoticeInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NoticeInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *NoticeInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *NoticeInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NoticeInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *NoticeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NoticeInfo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *NoticeInfo) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *NoticeInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NoticeInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *NoticeInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *NoticeInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqNoticeAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subtitle             string   `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string `protobuf:"bytes,6,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNoticeAdd) Reset()         { *m = ReqNoticeAdd{} }
func (m *ReqNoticeAdd) String() string { return proto.CompactTextString(m) }
func (*ReqNoticeAdd) ProtoMessage()    {}
func (*ReqNoticeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{1}
}

func (m *ReqNoticeAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNoticeAdd.Unmarshal(m, b)
}
func (m *ReqNoticeAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNoticeAdd.Marshal(b, m, deterministic)
}
func (m *ReqNoticeAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNoticeAdd.Merge(m, src)
}
func (m *ReqNoticeAdd) XXX_Size() int {
	return xxx_messageInfo_ReqNoticeAdd.Size(m)
}
func (m *ReqNoticeAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNoticeAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNoticeAdd proto.InternalMessageInfo

func (m *ReqNoticeAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqNoticeAdd) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *ReqNoticeAdd) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ReqNoticeAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqNoticeAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqNoticeAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqNoticeAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqNoticeUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subtitle             string   `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Targets              []string `protobuf:"bytes,6,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNoticeUpdate) Reset()         { *m = ReqNoticeUpdate{} }
func (m *ReqNoticeUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqNoticeUpdate) ProtoMessage()    {}
func (*ReqNoticeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{2}
}

func (m *ReqNoticeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNoticeUpdate.Unmarshal(m, b)
}
func (m *ReqNoticeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNoticeUpdate.Marshal(b, m, deterministic)
}
func (m *ReqNoticeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNoticeUpdate.Merge(m, src)
}
func (m *ReqNoticeUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqNoticeUpdate.Size(m)
}
func (m *ReqNoticeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNoticeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNoticeUpdate proto.InternalMessageInfo

func (m *ReqNoticeUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqNoticeUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqNoticeUpdate) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *ReqNoticeUpdate) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ReqNoticeUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqNoticeUpdate) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqNoticeState struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNoticeState) Reset()         { *m = ReqNoticeState{} }
func (m *ReqNoticeState) String() string { return proto.CompactTextString(m) }
func (*ReqNoticeState) ProtoMessage()    {}
func (*ReqNoticeState) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{3}
}

func (m *ReqNoticeState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNoticeState.Unmarshal(m, b)
}
func (m *ReqNoticeState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNoticeState.Marshal(b, m, deterministic)
}
func (m *ReqNoticeState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNoticeState.Merge(m, src)
}
func (m *ReqNoticeState) XXX_Size() int {
	return xxx_messageInfo_ReqNoticeState.Size(m)
}
func (m *ReqNoticeState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNoticeState.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNoticeState proto.InternalMessageInfo

func (m *ReqNoticeState) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqNoticeState) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqNoticeState) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyNoticeInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *NoticeInfo  `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyNoticeInfo) Reset()         { *m = ReplyNoticeInfo{} }
func (m *ReplyNoticeInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyNoticeInfo) ProtoMessage()    {}
func (*ReplyNoticeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{4}
}

func (m *ReplyNoticeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyNoticeInfo.Unmarshal(m, b)
}
func (m *ReplyNoticeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyNoticeInfo.Marshal(b, m, deterministic)
}
func (m *ReplyNoticeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyNoticeInfo.Merge(m, src)
}
func (m *ReplyNoticeInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyNoticeInfo.Size(m)
}
func (m *ReplyNoticeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyNoticeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyNoticeInfo proto.InternalMessageInfo

func (m *ReplyNoticeInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyNoticeInfo) GetInfo() *NoticeInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyNoticeList struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32        `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32        `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32        `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*NoticeInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyNoticeList) Reset()         { *m = ReplyNoticeList{} }
func (m *ReplyNoticeList) String() string { return proto.CompactTextString(m) }
func (*ReplyNoticeList) ProtoMessage()    {}
func (*ReplyNoticeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_949156a11f363448, []int{5}
}

func (m *ReplyNoticeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyNoticeList.Unmarshal(m, b)
}
func (m *ReplyNoticeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyNoticeList.Marshal(b, m, deterministic)
}
func (m *ReplyNoticeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyNoticeList.Merge(m, src)
}
func (m *ReplyNoticeList) XXX_Size() int {
	return xxx_messageInfo_ReplyNoticeList.Size(m)
}
func (m *ReplyNoticeList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyNoticeList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyNoticeList proto.InternalMessageInfo

func (m *ReplyNoticeList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyNoticeList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyNoticeList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyNoticeList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyNoticeList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyNoticeList) GetList() []*NoticeInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*NoticeInfo)(nil), "omo.msp.favorite.NoticeInfo")
	proto.RegisterType((*ReqNoticeAdd)(nil), "omo.msp.favorite.ReqNoticeAdd")
	proto.RegisterType((*ReqNoticeUpdate)(nil), "omo.msp.favorite.ReqNoticeUpdate")
	proto.RegisterType((*ReqNoticeState)(nil), "omo.msp.favorite.ReqNoticeState")
	proto.RegisterType((*ReplyNoticeInfo)(nil), "omo.msp.favorite.ReplyNoticeInfo")
	proto.RegisterType((*ReplyNoticeList)(nil), "omo.msp.favorite.ReplyNoticeList")
}

func init() {
	proto.RegisterFile("proto/favorite/notice.proto", fileDescriptor_949156a11f363448)
}

var fileDescriptor_949156a11f363448 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x4e, 0xe2, 0xc0, 0x25, 0xce, 0x87, 0x46, 0xe8, 0xd3, 0x28, 0xf4, 0xc7, 0x78, 0x95,
	0x55, 0xa8, 0xa8, 0xfa, 0x00, 0xb0, 0x28, 0xa2, 0x82, 0x22, 0x19, 0xda, 0x45, 0x77, 0x4e, 0x3c,
	0xa0, 0x91, 0x62, 0x8f, 0xf1, 0x8c, 0xa9, 0xe8, 0xbb, 0xf4, 0x25, 0xba, 0xea, 0x73, 0xf4, 0x05,
	0xfa, 0x2a, 0xd5, 0xdc, 0xb1, 0x1d, 0x1b, 0xe2, 0x24, 0x6a, 0x77, 0x73, 0x7f, 0xe6, 0xdc, 0x73,
	0xee, 0xf1, 0x24, 0x70, 0x90, 0x66, 0x42, 0x89, 0xa3, 0xdb, 0xf0, 0x41, 0x64, 0x5c, 0xb1, 0xa3,
	0x44, 0x28, 0x3e, 0x63, 0x13, 0xcc, 0x92, 0x3d, 0x11, 0x8b, 0x49, 0x2c, 0xd3, 0x49, 0x59, 0x1e,
	0x3d, 0x6d, 0x9f, 0x89, 0x38, 0x16, 0x89, 0x69, 0xf7, 0x7f, 0xda, 0x00, 0x1f, 0xf1, 0xfe, 0x79,
	0x72, 0x2b, 0xc8, 0x1e, 0x74, 0x72, 0x1e, 0x51, 0xcb, 0xb3, 0xc6, 0x3b, 0x81, 0x3e, 0x92, 0x21,
	0xd8, 0x3c, 0xa2, 0xb6, 0x67, 0x8d, 0xbb, 0x81, 0xcd, 0x23, 0x42, 0xa1, 0x3f, 0xcb, 0x58, 0xa8,
	0x58, 0x44, 0xbb, 0x9e, 0x35, 0xee, 0x04, 0x65, 0xa8, 0x2b, 0x79, 0x1a, 0x61, 0xa5, 0x67, 0x2a,
	0x45, 0x58, 0xdd, 0x11, 0x19, 0x75, 0x10, 0xb9, 0x0c, 0xc9, 0x08, 0xb6, 0x45, 0xca, 0x32, 0x2c,
	0xf5, 0xb1, 0x54, 0xc5, 0x84, 0x40, 0x37, 0x09, 0x63, 0x46, 0xb7, 0x31, 0x8f, 0x67, 0x9d, 0x9b,
	0x8a, 0xe8, 0x91, 0xee, 0x98, 0x9c, 0x3e, 0x6b, 0x0c, 0x99, 0x4f, 0x15, 0x57, 0x73, 0x46, 0xc1,
	0x60, 0x94, 0x31, 0xd9, 0x87, 0x9e, 0xf8, 0x9a, 0xb0, 0x8c, 0xee, 0x62, 0xc1, 0x04, 0xe4, 0x7f,
	0x70, 0xa4, 0x0a, 0x55, 0x2e, 0xe9, 0xc0, 0xb3, 0xc6, 0x6e, 0x50, 0x44, 0x9a, 0xa7, 0x0a, 0xb3,
	0x3b, 0xa6, 0x24, 0x75, 0xbd, 0x8e, 0xe6, 0x59, 0x84, 0x7a, 0xae, 0x0a, 0xef, 0x24, 0x1d, 0x62,
	0x1a, 0xcf, 0xfe, 0x0f, 0x0b, 0x06, 0x01, 0xbb, 0x37, 0xdb, 0x3b, 0x89, 0xa2, 0x8a, 0xb0, 0x55,
	0x23, 0x5c, 0x27, 0x67, 0x3f, 0x21, 0x57, 0x8a, 0xe9, 0xd4, 0xc4, 0x54, 0x84, 0xbb, 0x75, 0xc2,
	0xf5, 0x35, 0xf5, 0x9e, 0xaf, 0x09, 0xa9, 0xf5, 0x17, 0xd4, 0xea, 0x42, 0x9c, 0x86, 0x10, 0xff,
	0xbb, 0x05, 0xff, 0x55, 0xa4, 0x3f, 0xa1, 0x3f, 0x4b, 0x4c, 0x2f, 0x95, 0xd8, 0x2d, 0x4a, 0x3a,
	0x2d, 0x4a, 0xba, 0x4d, 0x5b, 0x5a, 0x39, 0xb7, 0xf3, 0xfb, 0x0c, 0xc3, 0x8a, 0xde, 0xb5, 0x5a,
	0xce, 0x6e, 0x61, 0x9f, 0xdd, 0xb0, 0xaf, 0x3e, 0xb1, 0xd3, 0x9c, 0xe8, 0x7f, 0xd3, 0xb2, 0xd3,
	0xf9, 0x63, 0xed, 0x5b, 0x7f, 0x57, 0xc1, 0x68, 0xec, 0xdd, 0xe3, 0x97, 0x93, 0xa7, 0x4f, 0x67,
	0x82, 0x57, 0xae, 0xb1, 0xa9, 0x9a, 0xf2, 0x06, 0xba, 0x3c, 0xb9, 0x15, 0x38, 0x7b, 0xf7, 0xf8,
	0xc5, 0xf3, 0x4b, 0x8b, 0x11, 0x01, 0x76, 0xfa, 0xbf, 0xac, 0xc6, 0xf0, 0x0b, 0x2e, 0xd5, 0xdf,
	0x0e, 0xdf, 0x87, 0x9e, 0x12, 0x2a, 0x9c, 0x17, 0xca, 0x4d, 0xa0, 0xb3, 0x69, 0x78, 0xc7, 0x24,
	0xaa, 0x76, 0x03, 0x13, 0x68, 0x53, 0xf4, 0x01, 0x4d, 0x71, 0x03, 0x3c, 0xeb, 0xd5, 0x25, 0x79,
	0x3c, 0x65, 0xc6, 0x12, 0x37, 0x28, 0x22, 0x2d, 0x6a, 0xce, 0xa5, 0x42, 0x37, 0xd6, 0x8a, 0xd2,
	0x9d, 0xc7, 0xbf, 0x7b, 0xe0, 0x16, 0x36, 0xb1, 0xec, 0x81, 0xcf, 0x18, 0xb9, 0x04, 0xe7, 0x24,
	0x8a, 0xae, 0x12, 0x46, 0x5e, 0x2d, 0x13, 0xb3, 0x78, 0x28, 0xa3, 0xc3, 0x16, 0xb1, 0x8b, 0x21,
	0xfe, 0x16, 0xb9, 0x00, 0xe7, 0x8c, 0x29, 0x0d, 0xb7, 0x74, 0x37, 0xf7, 0x39, 0x93, 0x4a, 0xb7,
	0x6e, 0x86, 0x76, 0x05, 0xfd, 0x33, 0xa6, 0x70, 0xf5, 0xaf, 0x5b, 0xe1, 0xde, 0xf3, 0xb9, 0x62,
	0xd9, 0x1a, 0x40, 0x8d, 0xe1, 0x6f, 0x91, 0x6b, 0x18, 0x9c, 0x31, 0xa5, 0xed, 0xe1, 0x52, 0xf1,
	0xd9, 0x7a, 0x54, 0x6f, 0x85, 0xc3, 0x08, 0xe1, 0x6f, 0x91, 0x1b, 0x00, 0xf3, 0x26, 0x4f, 0x43,
	0xc9, 0xc8, 0xe1, 0x8a, 0x35, 0x9a, 0xb6, 0x4d, 0xb5, 0x0f, 0x4c, 0xbb, 0xf9, 0x98, 0x88, 0xb7,
	0x02, 0x17, 0xdf, 0xdc, 0xe8, 0xa0, 0x05, 0xb6, 0x00, 0xfc, 0x50, 0xd2, 0xbc, 0xd1, 0x3f, 0x36,
	0xed, 0xf6, 0xe8, 0x55, 0xad, 0xc3, 0xba, 0x04, 0xb7, 0xc4, 0x32, 0x3f, 0xb5, 0xff, 0x06, 0x77,
	0x0e, 0x3b, 0x01, 0x8b, 0xc5, 0x03, 0xdb, 0xe0, 0xc3, 0x59, 0x0d, 0x75, 0xba, 0xf7, 0x65, 0xd8,
	0xfc, 0xe7, 0x9c, 0x3a, 0x18, 0xbf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x50, 0x48, 0x93, 0x5f,
	0x81, 0x07, 0x00, 0x00,
}
