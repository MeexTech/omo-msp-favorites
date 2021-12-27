// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/article.proto

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

type ArticleInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
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
	Assets               []string `protobuf:"bytes,14,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleInfo) Reset()         { *m = ArticleInfo{} }
func (m *ArticleInfo) String() string { return proto.CompactTextString(m) }
func (*ArticleInfo) ProtoMessage()    {}
func (*ArticleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{0}
}

func (m *ArticleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleInfo.Unmarshal(m, b)
}
func (m *ArticleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleInfo.Marshal(b, m, deterministic)
}
func (m *ArticleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleInfo.Merge(m, src)
}
func (m *ArticleInfo) XXX_Size() int {
	return xxx_messageInfo_ArticleInfo.Size(m)
}
func (m *ArticleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleInfo proto.InternalMessageInfo

func (m *ArticleInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ArticleInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ArticleInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ArticleInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ArticleInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ArticleInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ArticleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleInfo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ArticleInfo) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *ArticleInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ArticleInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ArticleInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ArticleInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqArticleAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subtitle             string   `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Assets               []string `protobuf:"bytes,6,rep,name=assets,proto3" json:"assets,omitempty"`
	Targets              []string `protobuf:"bytes,7,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArticleAdd) Reset()         { *m = ReqArticleAdd{} }
func (m *ReqArticleAdd) String() string { return proto.CompactTextString(m) }
func (*ReqArticleAdd) ProtoMessage()    {}
func (*ReqArticleAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{1}
}

func (m *ReqArticleAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqArticleAdd.Unmarshal(m, b)
}
func (m *ReqArticleAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqArticleAdd.Marshal(b, m, deterministic)
}
func (m *ReqArticleAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArticleAdd.Merge(m, src)
}
func (m *ReqArticleAdd) XXX_Size() int {
	return xxx_messageInfo_ReqArticleAdd.Size(m)
}
func (m *ReqArticleAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArticleAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArticleAdd proto.InternalMessageInfo

func (m *ReqArticleAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqArticleAdd) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *ReqArticleAdd) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ReqArticleAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqArticleAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqArticleAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqArticleAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqArticleUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subtitle             string   `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Assets               []string `protobuf:"bytes,6,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArticleUpdate) Reset()         { *m = ReqArticleUpdate{} }
func (m *ReqArticleUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqArticleUpdate) ProtoMessage()    {}
func (*ReqArticleUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{2}
}

func (m *ReqArticleUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqArticleUpdate.Unmarshal(m, b)
}
func (m *ReqArticleUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqArticleUpdate.Marshal(b, m, deterministic)
}
func (m *ReqArticleUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArticleUpdate.Merge(m, src)
}
func (m *ReqArticleUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqArticleUpdate.Size(m)
}
func (m *ReqArticleUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArticleUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArticleUpdate proto.InternalMessageInfo

func (m *ReqArticleUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqArticleUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqArticleUpdate) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *ReqArticleUpdate) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ReqArticleUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqArticleUpdate) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqArticleList struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Targets              []string `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArticleList) Reset()         { *m = ReqArticleList{} }
func (m *ReqArticleList) String() string { return proto.CompactTextString(m) }
func (*ReqArticleList) ProtoMessage()    {}
func (*ReqArticleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{3}
}

func (m *ReqArticleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqArticleList.Unmarshal(m, b)
}
func (m *ReqArticleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqArticleList.Marshal(b, m, deterministic)
}
func (m *ReqArticleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArticleList.Merge(m, src)
}
func (m *ReqArticleList) XXX_Size() int {
	return xxx_messageInfo_ReqArticleList.Size(m)
}
func (m *ReqArticleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArticleList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArticleList proto.InternalMessageInfo

func (m *ReqArticleList) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqArticleList) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqArticleList) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqArticleAssets struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Assets               []string `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArticleAssets) Reset()         { *m = ReqArticleAssets{} }
func (m *ReqArticleAssets) String() string { return proto.CompactTextString(m) }
func (*ReqArticleAssets) ProtoMessage()    {}
func (*ReqArticleAssets) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{4}
}

func (m *ReqArticleAssets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqArticleAssets.Unmarshal(m, b)
}
func (m *ReqArticleAssets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqArticleAssets.Marshal(b, m, deterministic)
}
func (m *ReqArticleAssets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArticleAssets.Merge(m, src)
}
func (m *ReqArticleAssets) XXX_Size() int {
	return xxx_messageInfo_ReqArticleAssets.Size(m)
}
func (m *ReqArticleAssets) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArticleAssets.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArticleAssets proto.InternalMessageInfo

func (m *ReqArticleAssets) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqArticleAssets) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqArticleState struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArticleState) Reset()         { *m = ReqArticleState{} }
func (m *ReqArticleState) String() string { return proto.CompactTextString(m) }
func (*ReqArticleState) ProtoMessage()    {}
func (*ReqArticleState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{5}
}

func (m *ReqArticleState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqArticleState.Unmarshal(m, b)
}
func (m *ReqArticleState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqArticleState.Marshal(b, m, deterministic)
}
func (m *ReqArticleState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArticleState.Merge(m, src)
}
func (m *ReqArticleState) XXX_Size() int {
	return xxx_messageInfo_ReqArticleState.Size(m)
}
func (m *ReqArticleState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArticleState.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArticleState proto.InternalMessageInfo

func (m *ReqArticleState) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqArticleState) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReplyArticleInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *ArticleInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyArticleInfo) Reset()         { *m = ReplyArticleInfo{} }
func (m *ReplyArticleInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyArticleInfo) ProtoMessage()    {}
func (*ReplyArticleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{6}
}

func (m *ReplyArticleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyArticleInfo.Unmarshal(m, b)
}
func (m *ReplyArticleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyArticleInfo.Marshal(b, m, deterministic)
}
func (m *ReplyArticleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyArticleInfo.Merge(m, src)
}
func (m *ReplyArticleInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyArticleInfo.Size(m)
}
func (m *ReplyArticleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyArticleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyArticleInfo proto.InternalMessageInfo

func (m *ReplyArticleInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyArticleInfo) GetInfo() *ArticleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyArticleList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*ArticleInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyArticleList) Reset()         { *m = ReplyArticleList{} }
func (m *ReplyArticleList) String() string { return proto.CompactTextString(m) }
func (*ReplyArticleList) ProtoMessage()    {}
func (*ReplyArticleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f089b0f7c4866fe, []int{7}
}

func (m *ReplyArticleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyArticleList.Unmarshal(m, b)
}
func (m *ReplyArticleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyArticleList.Marshal(b, m, deterministic)
}
func (m *ReplyArticleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyArticleList.Merge(m, src)
}
func (m *ReplyArticleList) XXX_Size() int {
	return xxx_messageInfo_ReplyArticleList.Size(m)
}
func (m *ReplyArticleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyArticleList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyArticleList proto.InternalMessageInfo

func (m *ReplyArticleList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyArticleList) GetList() []*ArticleInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ArticleInfo)(nil), "omo.msp.favorite.ArticleInfo")
	proto.RegisterType((*ReqArticleAdd)(nil), "omo.msp.favorite.ReqArticleAdd")
	proto.RegisterType((*ReqArticleUpdate)(nil), "omo.msp.favorite.ReqArticleUpdate")
	proto.RegisterType((*ReqArticleList)(nil), "omo.msp.favorite.ReqArticleList")
	proto.RegisterType((*ReqArticleAssets)(nil), "omo.msp.favorite.ReqArticleAssets")
	proto.RegisterType((*ReqArticleState)(nil), "omo.msp.favorite.ReqArticleState")
	proto.RegisterType((*ReplyArticleInfo)(nil), "omo.msp.favorite.ReplyArticleInfo")
	proto.RegisterType((*ReplyArticleList)(nil), "omo.msp.favorite.ReplyArticleList")
}

func init() {
	proto.RegisterFile("proto/favorite/article.proto", fileDescriptor_8f089b0f7c4866fe)
}

var fileDescriptor_8f089b0f7c4866fe = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x7f, 0xe2, 0xb4, 0x93, 0x26, 0x44, 0x2b, 0x84, 0x56, 0x29, 0x08, 0xe3, 0x53, 0x4e,
	0xa9, 0x28, 0xe2, 0x04, 0x97, 0xf4, 0x52, 0x21, 0x81, 0x0a, 0x2e, 0x70, 0xe0, 0xe6, 0xc4, 0x1b,
	0x64, 0x29, 0xf1, 0xba, 0xde, 0x4d, 0x51, 0x24, 0x0e, 0x3c, 0x08, 0xef, 0xc1, 0x73, 0xf0, 0x46,
	0x68, 0x67, 0xfd, 0xb3, 0x0e, 0x4e, 0x52, 0xb8, 0xed, 0xec, 0x7c, 0xfb, 0xcd, 0xf7, 0x8d, 0x67,
	0x64, 0x78, 0x9c, 0xe5, 0x5c, 0xf2, 0xf3, 0x45, 0x74, 0xc7, 0xf3, 0x44, 0xb2, 0xf3, 0x28, 0x97,
	0xc9, 0x7c, 0xc9, 0x26, 0x78, 0x4d, 0x86, 0x7c, 0xc5, 0x27, 0x2b, 0x91, 0x4d, 0xca, 0xfc, 0xe8,
	0x6c, 0x0b, 0x3f, 0xe7, 0xab, 0x15, 0x4f, 0x35, 0x3c, 0xf8, 0x6d, 0x43, 0x6f, 0xaa, 0x09, 0xde,
	0xa4, 0x0b, 0x4e, 0x86, 0xe0, 0xac, 0x93, 0x98, 0x5a, 0xbe, 0x35, 0x3e, 0x09, 0xd5, 0x91, 0x0c,
	0xc0, 0x4e, 0x62, 0x6a, 0xfb, 0xd6, 0xd8, 0x0d, 0xed, 0x24, 0x26, 0x04, 0x5c, 0xb9, 0xc9, 0x18,
	0x75, 0x7c, 0x6b, 0xdc, 0x0f, 0xf1, 0x4c, 0x28, 0x74, 0xe7, 0x39, 0x8b, 0x24, 0x8b, 0xa9, 0xeb,
	0x5b, 0x63, 0x27, 0x2c, 0x43, 0x95, 0x59, 0x67, 0x31, 0x66, 0x3a, 0x3a, 0x53, 0x84, 0xd5, 0x1b,
	0x9e, 0x53, 0x0f, 0xab, 0x95, 0x21, 0x19, 0xc1, 0x31, 0xcf, 0x58, 0x8e, 0xa9, 0x2e, 0xa6, 0xaa,
	0x58, 0x55, 0x4f, 0xa3, 0x15, 0xa3, 0xc7, 0x78, 0x8f, 0x67, 0x75, 0x37, 0xe3, 0xf1, 0x86, 0x9e,
	0xe8, 0x3b, 0x75, 0x56, 0x1c, 0x62, 0x3d, 0x93, 0x89, 0x5c, 0x32, 0x0a, 0x9a, 0xa3, 0x8c, 0xc9,
	0x43, 0xe8, 0xf0, 0x6f, 0x29, 0xcb, 0x69, 0x0f, 0x13, 0x3a, 0x20, 0x8f, 0xc0, 0x13, 0x32, 0x92,
	0x6b, 0x41, 0x4f, 0xd1, 0x59, 0x11, 0x29, 0x9d, 0x32, 0xca, 0xbf, 0x32, 0x29, 0x68, 0xdf, 0x77,
	0x94, 0xce, 0x22, 0x54, 0x2f, 0x22, 0x21, 0x54, 0x62, 0x80, 0x89, 0x22, 0x0a, 0x7e, 0x59, 0xd0,
	0x0f, 0xd9, 0x6d, 0xd1, 0xd6, 0x69, 0x1c, 0x57, 0xaa, 0x2d, 0x43, 0xb5, 0xa9, 0xd0, 0xde, 0x52,
	0x58, 0x3a, 0x72, 0x0c, 0x47, 0x95, 0x6a, 0xd7, 0x54, 0x6d, 0xf6, 0xaa, 0xb3, 0xd5, 0xab, 0x5a,
	0x9f, 0x67, 0xea, 0x33, 0x1d, 0x75, 0x1b, 0x8e, 0x82, 0x9f, 0x16, 0x0c, 0x6b, 0xe5, 0x9f, 0xf0,
	0x4b, 0xb5, 0x8c, 0x44, 0x69, 0xc7, 0xde, 0x61, 0xc7, 0xd9, 0x61, 0xc7, 0x6d, 0x7e, 0xa0, 0x7f,
	0x15, 0x1e, 0x7c, 0x84, 0x41, 0xad, 0xee, 0x6d, 0x22, 0x64, 0xdd, 0x14, 0xcb, 0x6c, 0x4a, 0x39,
	0xa2, 0x76, 0x73, 0x44, 0x4b, 0xd3, 0x4e, 0xd3, 0xf4, 0x6b, 0xd3, 0xf3, 0x54, 0xb7, 0xe8, 0x6f,
	0xcf, 0xb5, 0x26, 0xbb, 0xa1, 0xe9, 0x15, 0x3c, 0xa8, 0x5f, 0xdf, 0xc8, 0xf6, 0x86, 0xd5, 0xb3,
	0x65, 0x9b, 0xb3, 0x15, 0x7c, 0x57, 0xa5, 0xb3, 0xe5, 0xc6, 0xdc, 0xc0, 0x97, 0x15, 0x56, 0x11,
	0xf4, 0x2e, 0x9e, 0x4c, 0xb6, 0x37, 0x7a, 0x82, 0x6f, 0x6e, 0x10, 0x54, 0x8d, 0xe9, 0x73, 0x70,
	0x93, 0x74, 0xc1, 0xb1, 0x40, 0xeb, 0x23, 0xa3, 0x46, 0x88, 0xd0, 0xed, 0xea, 0xd8, 0xd0, 0xff,
	0xaf, 0xbe, 0x4c, 0x84, 0xc4, 0xde, 0x1c, 0xae, 0xae, 0xa0, 0x17, 0x3f, 0x5c, 0x18, 0x94, 0x6d,
	0x63, 0xf9, 0x5d, 0x32, 0x67, 0xe4, 0x1a, 0xbc, 0x69, 0x1c, 0x5f, 0xa7, 0x8c, 0x3c, 0x6d, 0x2b,
	0x6b, 0x6c, 0xd4, 0x28, 0xd8, 0xa1, 0xcb, 0xa8, 0x13, 0x1c, 0x91, 0x77, 0xe0, 0x5d, 0x31, 0xa9,
	0x08, 0x5b, 0x7d, 0xdc, 0xae, 0x99, 0x90, 0x0a, 0x7a, 0x4f, 0xba, 0x0f, 0xd0, 0xbd, 0x62, 0x12,
	0xfb, 0xe4, 0xef, 0x13, 0xa8, 0x10, 0x87, 0x28, 0x15, 0x26, 0x38, 0x22, 0x9f, 0x01, 0xf4, 0x9a,
	0x5d, 0x46, 0x82, 0x91, 0x60, 0x1f, 0xab, 0xc6, 0xdd, 0x5b, 0xea, 0xa9, 0xc6, 0x17, 0x03, 0xbd,
	0x97, 0x59, 0x63, 0x46, 0x67, 0x3b, 0x98, 0x0b, 0xca, 0xf7, 0x25, 0xa5, 0xfe, 0xf6, 0xe4, 0xd9,
	0x3e, 0x4a, 0xdc, 0x84, 0x03, 0x8c, 0x97, 0xc3, 0x2f, 0x83, 0xe6, 0xbf, 0x69, 0xe6, 0x61, 0xfc,
	0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x46, 0x8c, 0x3f, 0xe4, 0x06, 0x00, 0x00,
}
