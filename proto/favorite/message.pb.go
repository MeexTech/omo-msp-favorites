// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/message.proto

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

type MessageInfo struct {
	Uid                  string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              uint64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Date                 *DateInfo `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Remark               string    `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Organizer            string    `protobuf:"bytes,7,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Type                 uint32    `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	Tags                 []string  `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string  `protobuf:"bytes,10,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{0}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MessageInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MessageInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageInfo) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *MessageInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *MessageInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MessageInfo) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *MessageInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MessageInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReplyMessages struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Count                uint32         `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	List                 []*MessageInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyMessages) Reset()         { *m = ReplyMessages{} }
func (m *ReplyMessages) String() string { return proto.CompactTextString(m) }
func (*ReplyMessages) ProtoMessage()    {}
func (*ReplyMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{1}
}

func (m *ReplyMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMessages.Unmarshal(m, b)
}
func (m *ReplyMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMessages.Marshal(b, m, deterministic)
}
func (m *ReplyMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMessages.Merge(m, src)
}
func (m *ReplyMessages) XXX_Size() int {
	return xxx_messageInfo_ReplyMessages.Size(m)
}
func (m *ReplyMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMessages.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMessages proto.InternalMessageInfo

func (m *ReplyMessages) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMessages) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMessages) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyMessages) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyMessages) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyMessages) GetList() []*MessageInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageInfo)(nil), "omo.msp.favorite.MessageInfo")
	proto.RegisterType((*ReplyMessages)(nil), "omo.msp.favorite.ReplyMessages")
}

func init() {
	proto.RegisterFile("proto/favorite/message.proto", fileDescriptor_12a2cdcb179ecaf9)
}

var fileDescriptor_12a2cdcb179ecaf9 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x95, 0x42, 0xa9, 0x1d, 0x42, 0xd3, 0x4c, 0x8c, 0x99, 0xd4, 0x9a, 0x92, 0xae, 0x58, 0xd1,
	0x58, 0xe3, 0x0f, 0x34, 0x46, 0xe3, 0xc2, 0x85, 0xd3, 0x9d, 0xbb, 0x91, 0xde, 0x12, 0x22, 0x30,
	0x38, 0x73, 0xa9, 0xa9, 0xbf, 0xe0, 0xc7, 0xf9, 0x4b, 0x66, 0x2e, 0x10, 0x7d, 0x7d, 0x7d, 0xbb,
	0x73, 0xce, 0x3d, 0x5c, 0xce, 0x3d, 0xc0, 0xd6, 0xad, 0xd1, 0xa8, 0x77, 0x67, 0x75, 0xd1, 0xa6,
	0x44, 0xd8, 0xd5, 0x60, 0xad, 0x2a, 0x20, 0x23, 0x99, 0x2f, 0x75, 0xad, 0xb3, 0xda, 0xb6, 0xd9,
	0x38, 0x5f, 0xbd, 0xba, 0xf1, 0xe7, 0xba, 0xae, 0x75, 0xd3, 0xdb, 0xb7, 0xbf, 0x27, 0x2c, 0xfa,
	0xdc, 0x2f, 0xf8, 0xd4, 0x9c, 0x35, 0x5f, 0x32, 0xbf, 0x2b, 0x4f, 0xc2, 0x4b, 0xbc, 0x74, 0x2e,
	0x1d, 0xe4, 0x82, 0xcd, 0x72, 0x03, 0x0a, 0xe1, 0x24, 0x26, 0x89, 0x97, 0x06, 0x72, 0xa4, 0x9c,
	0xb3, 0xa0, 0x51, 0x35, 0x08, 0x9f, 0xcc, 0x84, 0x79, 0xc6, 0x82, 0x93, 0x42, 0x10, 0x41, 0xe2,
	0xa5, 0xd1, 0x7e, 0x95, 0xdd, 0xa6, 0xc9, 0xde, 0x2b, 0xa4, 0x37, 0x49, 0xf2, 0xf1, 0x97, 0x2c,
	0x34, 0x50, 0x2b, 0xf3, 0x5d, 0x4c, 0x69, 0xcb, 0xc0, 0xf8, 0x0b, 0x36, 0xd5, 0x3f, 0x1b, 0x30,
	0x22, 0x24, 0xb9, 0x27, 0x7c, 0xcd, 0xe6, 0xda, 0x14, 0xaa, 0x29, 0x7f, 0x81, 0x11, 0x33, 0x9a,
	0xfc, 0x13, 0x5c, 0x1e, 0xbc, 0xb6, 0x20, 0x9e, 0x27, 0x5e, 0x1a, 0x4b, 0xc2, 0xa4, 0xa9, 0xc2,
	0x8a, 0x79, 0xe2, 0xbb, 0x8c, 0x0e, 0xbb, 0x8b, 0x50, 0x99, 0x02, 0xd0, 0x0a, 0x46, 0xf2, 0x48,
	0xb7, 0x7f, 0x3c, 0x16, 0x4b, 0x68, 0xab, 0xeb, 0x50, 0x89, 0xe5, 0xef, 0x58, 0x68, 0x51, 0x61,
	0x67, 0xa9, 0x92, 0x68, 0xff, 0xfa, 0xf1, 0x45, 0xf4, 0xc0, 0x91, 0x4c, 0x72, 0x30, 0xbb, 0xf8,
	0xa8, 0x51, 0x55, 0x54, 0x59, 0x2c, 0x7b, 0xe2, 0xd4, 0xd6, 0x6d, 0xa5, 0xc6, 0x62, 0xd9, 0x13,
	0x17, 0xd1, 0x01, 0xaa, 0x2c, 0x96, 0x84, 0x9d, 0x33, 0xd7, 0x5d, 0x83, 0xd4, 0x4a, 0x2c, 0x7b,
	0xc2, 0xdf, 0xb0, 0xa0, 0x2a, 0x2d, 0x8a, 0x30, 0xf1, 0xef, 0x47, 0xf9, 0xef, 0x4b, 0x4a, 0xb2,
	0xee, 0x73, 0xb6, 0x18, 0xc4, 0x23, 0x98, 0x4b, 0x99, 0x03, 0xff, 0xc2, 0xa2, 0x8f, 0x80, 0x87,
	0xeb, 0x87, 0xb2, 0x42, 0x30, 0x7c, 0x73, 0xef, 0xa0, 0x1f, 0x1d, 0x58, 0xec, 0x0d, 0xab, 0xcd,
	0x13, 0x17, 0x8f, 0x15, 0x6d, 0x9f, 0x1d, 0x96, 0x5f, 0x17, 0x0f, 0xff, 0xb1, 0x6f, 0x21, 0xf1,
	0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x77, 0xeb, 0xd4, 0xac, 0x02, 0x00, 0x00,
}
