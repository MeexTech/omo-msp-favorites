// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/words.proto

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

type WordsInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Words                string   `protobuf:"bytes,7,opt,name=words,proto3" json:"words,omitempty"`
	Creator              string   `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,9,opt,name=operator,proto3" json:"operator,omitempty"`
	Target               string   `protobuf:"bytes,10,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
	Weight               uint32   `protobuf:"varint,13,opt,name=weight,proto3" json:"weight,omitempty"`
	Asset                string   `protobuf:"bytes,14,opt,name=asset,proto3" json:"asset,omitempty"`
	Quote                string   `protobuf:"bytes,15,opt,name=quote,proto3" json:"quote,omitempty"`
	Device               string   `protobuf:"bytes,16,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordsInfo) Reset()         { *m = WordsInfo{} }
func (m *WordsInfo) String() string { return proto.CompactTextString(m) }
func (*WordsInfo) ProtoMessage()    {}
func (*WordsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b7eebfc6949eed, []int{0}
}

func (m *WordsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordsInfo.Unmarshal(m, b)
}
func (m *WordsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordsInfo.Marshal(b, m, deterministic)
}
func (m *WordsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordsInfo.Merge(m, src)
}
func (m *WordsInfo) XXX_Size() int {
	return xxx_messageInfo_WordsInfo.Size(m)
}
func (m *WordsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WordsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WordsInfo proto.InternalMessageInfo

func (m *WordsInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WordsInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WordsInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *WordsInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *WordsInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WordsInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *WordsInfo) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func (m *WordsInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WordsInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *WordsInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *WordsInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *WordsInfo) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *WordsInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *WordsInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *WordsInfo) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type ReqWordsAdd struct {
	Words                string   `protobuf:"bytes,1,opt,name=words,proto3" json:"words,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Target               string   `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Quote                string   `protobuf:"bytes,6,opt,name=quote,proto3" json:"quote,omitempty"`
	Asset                string   `protobuf:"bytes,7,opt,name=asset,proto3" json:"asset,omitempty"`
	Device               string   `protobuf:"bytes,8,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqWordsAdd) Reset()         { *m = ReqWordsAdd{} }
func (m *ReqWordsAdd) String() string { return proto.CompactTextString(m) }
func (*ReqWordsAdd) ProtoMessage()    {}
func (*ReqWordsAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b7eebfc6949eed, []int{1}
}

func (m *ReqWordsAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqWordsAdd.Unmarshal(m, b)
}
func (m *ReqWordsAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqWordsAdd.Marshal(b, m, deterministic)
}
func (m *ReqWordsAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqWordsAdd.Merge(m, src)
}
func (m *ReqWordsAdd) XXX_Size() int {
	return xxx_messageInfo_ReqWordsAdd.Size(m)
}
func (m *ReqWordsAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqWordsAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqWordsAdd proto.InternalMessageInfo

func (m *ReqWordsAdd) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func (m *ReqWordsAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqWordsAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqWordsAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqWordsAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqWordsAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqWordsAdd) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqWordsAdd) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type ReplyWordsInfo struct {
	Info                 *WordsInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyWordsInfo) Reset()         { *m = ReplyWordsInfo{} }
func (m *ReplyWordsInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyWordsInfo) ProtoMessage()    {}
func (*ReplyWordsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b7eebfc6949eed, []int{2}
}

func (m *ReplyWordsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyWordsInfo.Unmarshal(m, b)
}
func (m *ReplyWordsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyWordsInfo.Marshal(b, m, deterministic)
}
func (m *ReplyWordsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyWordsInfo.Merge(m, src)
}
func (m *ReplyWordsInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyWordsInfo.Size(m)
}
func (m *ReplyWordsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyWordsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyWordsInfo proto.InternalMessageInfo

func (m *ReplyWordsInfo) GetInfo() *WordsInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyWordsInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyWordsList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32       `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32       `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*WordsInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyWordsList) Reset()         { *m = ReplyWordsList{} }
func (m *ReplyWordsList) String() string { return proto.CompactTextString(m) }
func (*ReplyWordsList) ProtoMessage()    {}
func (*ReplyWordsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b7eebfc6949eed, []int{3}
}

func (m *ReplyWordsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyWordsList.Unmarshal(m, b)
}
func (m *ReplyWordsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyWordsList.Marshal(b, m, deterministic)
}
func (m *ReplyWordsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyWordsList.Merge(m, src)
}
func (m *ReplyWordsList) XXX_Size() int {
	return xxx_messageInfo_ReplyWordsList.Size(m)
}
func (m *ReplyWordsList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyWordsList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyWordsList proto.InternalMessageInfo

func (m *ReplyWordsList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyWordsList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyWordsList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyWordsList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyWordsList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyWordsList) GetList() []*WordsInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*WordsInfo)(nil), "omo.msp.favorite.WordsInfo")
	proto.RegisterType((*ReqWordsAdd)(nil), "omo.msp.favorite.ReqWordsAdd")
	proto.RegisterType((*ReplyWordsInfo)(nil), "omo.msp.favorite.ReplyWordsInfo")
	proto.RegisterType((*ReplyWordsList)(nil), "omo.msp.favorite.ReplyWordsList")
}

func init() {
	proto.RegisterFile("proto/favorite/words.proto", fileDescriptor_62b7eebfc6949eed)
}

var fileDescriptor_62b7eebfc6949eed = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0xa6, 0x14, 0x0a, 0x0c, 0x3f, 0x92, 0x89, 0x31, 0x13, 0x88, 0xb1, 0xe9, 0x15, 0x57, 0x90,
	0x60, 0x7c, 0x80, 0xdd, 0x0b, 0xc9, 0x46, 0x13, 0xcd, 0x10, 0x63, 0xe2, 0x5d, 0x97, 0x1e, 0x70,
	0x12, 0xda, 0xe9, 0x76, 0xa6, 0x20, 0x0f, 0xe8, 0xbd, 0x6f, 0xe0, 0x1b, 0xf8, 0x0c, 0x66, 0xce,
	0x94, 0xd2, 0x5d, 0x97, 0x15, 0xbd, 0x3b, 0xdf, 0x39, 0x5f, 0xbf, 0x73, 0xbe, 0xaf, 0x4d, 0xc9,
	0x28, 0xcd, 0xa4, 0x96, 0xb3, 0x75, 0xb8, 0x93, 0x99, 0xd0, 0x30, 0xdb, 0xcb, 0x2c, 0x52, 0x53,
	0x6c, 0xd2, 0xa1, 0x8c, 0xe5, 0x34, 0x56, 0xe9, 0xf4, 0x38, 0x1d, 0x8d, 0x1f, 0xb0, 0x57, 0x32,
	0x8e, 0x65, 0x62, 0xe9, 0xc1, 0xcf, 0x3a, 0xe9, 0x7c, 0x36, 0x8f, 0xdf, 0x24, 0x6b, 0x49, 0x87,
	0xc4, 0xcd, 0x45, 0xc4, 0x1c, 0xdf, 0x99, 0x74, 0xb8, 0x29, 0xe9, 0x80, 0xd4, 0x45, 0xc4, 0xea,
	0xbe, 0x33, 0x69, 0xf0, 0xba, 0x88, 0x28, 0x23, 0xad, 0x55, 0x06, 0xa1, 0x86, 0x88, 0xb9, 0xbe,
	0x33, 0x71, 0xf9, 0x11, 0x9a, 0x49, 0x9e, 0x46, 0x38, 0x69, 0xd8, 0x49, 0x01, 0x29, 0x25, 0x8d,
	0x24, 0x8c, 0x81, 0x35, 0x51, 0x16, 0x6b, 0xfa, 0x9c, 0x34, 0xe5, 0x3e, 0x81, 0x8c, 0x79, 0xd8,
	0xb4, 0xc0, 0x74, 0xd1, 0x0b, 0x6b, 0xd9, 0x2e, 0x82, 0x72, 0xa7, 0xcc, 0x58, 0x1b, 0xfb, 0x47,
	0x48, 0x47, 0xa4, 0x2d, 0x53, 0xc8, 0x70, 0xd4, 0xc1, 0x51, 0x89, 0xe9, 0x0b, 0xe2, 0xe9, 0x30,
	0xdb, 0x80, 0x66, 0x04, 0x27, 0x05, 0x32, 0xd7, 0xe8, 0x43, 0x0a, 0xac, 0xe7, 0x3b, 0x93, 0x3e,
	0xc7, 0xda, 0x70, 0xf7, 0x20, 0x36, 0x5f, 0x35, 0xeb, 0x63, 0xb7, 0x40, 0xe6, 0x9e, 0x50, 0x29,
	0xd0, 0x6c, 0x60, 0xef, 0x41, 0x60, 0xba, 0x77, 0xb9, 0xd4, 0xc0, 0x9e, 0xd9, 0x2e, 0x02, 0xa3,
	0x11, 0xc1, 0x4e, 0xac, 0x80, 0x0d, 0xed, 0x3e, 0x8b, 0x82, 0xef, 0x0e, 0xe9, 0x72, 0xb8, 0xc3,
	0x90, 0xaf, 0xa2, 0xe8, 0xe4, 0xd1, 0xa9, 0x7a, 0x2c, 0xf3, 0xa8, 0x57, 0xf3, 0xa8, 0xfa, 0x73,
	0xcf, 0xfa, 0x6b, 0x3c, 0xea, 0xaf, 0x59, 0xf1, 0x57, 0x5e, 0xec, 0x55, 0x2f, 0x2e, 0xdd, 0xb5,
	0xaa, 0xee, 0x4e, 0x3e, 0xda, 0xf7, 0x7c, 0x7c, 0x23, 0x03, 0x0e, 0xe9, 0xf6, 0x70, 0xfa, 0x5a,
	0x66, 0xa4, 0x21, 0x92, 0xb5, 0x44, 0x23, 0xdd, 0xf9, 0x78, 0xfa, 0xf0, 0xcb, 0x9b, 0x96, 0x54,
	0x8e, 0x44, 0xfa, 0x86, 0x78, 0x4a, 0x87, 0x3a, 0x57, 0xe8, 0xb2, 0x3b, 0x7f, 0xf9, 0xe7, 0x23,
	0xb8, 0x62, 0x89, 0x24, 0x5e, 0x90, 0x83, 0x1f, 0x4e, 0x75, 0xf5, 0x7b, 0xa1, 0x74, 0x45, 0xc9,
	0xf9, 0x07, 0x25, 0xe3, 0x58, 0x4b, 0x1d, 0x6e, 0x71, 0x7f, 0x9f, 0x5b, 0x60, 0xba, 0x69, 0xb8,
	0x01, 0x85, 0x11, 0xf7, 0xb9, 0x05, 0x26, 0x47, 0x53, 0x60, 0xba, 0x7d, 0x8e, 0xb5, 0xc9, 0x26,
	0xc9, 0xe3, 0x5b, 0xc8, 0x8a, 0x74, 0x0b, 0x64, 0x92, 0xd8, 0x0a, 0xa5, 0x99, 0xe7, 0xbb, 0x7f,
	0x4d, 0xc2, 0x10, 0xe7, 0xbf, 0x5c, 0xd2, 0xc3, 0xde, 0x12, 0x32, 0x93, 0x2e, 0x7d, 0x47, 0xbc,
	0xab, 0x28, 0xfa, 0x90, 0x00, 0x7d, 0xd4, 0x4a, 0xf9, 0xf9, 0x8c, 0xfc, 0x33, 0x4e, 0xcb, 0x0d,
	0x41, 0xcd, 0x88, 0x2d, 0x40, 0x9f, 0x17, 0xcb, 0x41, 0x69, 0x43, 0xbd, 0x48, 0x8c, 0x93, 0xee,
	0x02, 0xf4, 0xf5, 0xe1, 0xad, 0xd8, 0x6a, 0xc8, 0xe8, 0xab, 0xb3, 0x8a, 0x96, 0xf0, 0xb4, 0xa6,
	0x79, 0x79, 0x41, 0x8d, 0x2e, 0x49, 0x6f, 0x01, 0xda, 0xbc, 0x1c, 0xa1, 0xb4, 0x58, 0xfd, 0xbf,
	0x68, 0x29, 0x11, 0xd4, 0xe8, 0x47, 0x32, 0xf8, 0x84, 0x7f, 0x9c, 0x0b, 0x6e, 0xb5, 0xc4, 0xd1,
	0xf8, 0x8c, 0x6c, 0x61, 0xfd, 0x86, 0x74, 0x38, 0xc4, 0x72, 0x07, 0x17, 0x44, 0xf9, 0xb4, 0xd4,
	0xf5, 0xf0, 0xcb, 0xe0, 0xfe, 0x6f, 0xf8, 0xd6, 0x43, 0xfc, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x64, 0x7d, 0xe6, 0xcd, 0x05, 0x00, 0x00,
}
