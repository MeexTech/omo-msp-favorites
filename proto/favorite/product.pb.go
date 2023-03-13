// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/product.proto

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

type ProductInfo struct {
	Uid                  string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64           `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32           `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64            `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64            `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string           `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string           `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string           `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Menus                string           `protobuf:"bytes,9,opt,name=menus,proto3" json:"menus,omitempty"`
	Entry                string           `protobuf:"bytes,10,opt,name=entry,proto3" json:"entry,omitempty"`
	Key                  string           `protobuf:"bytes,11,opt,name=key,proto3" json:"key,omitempty"`
	Templet              string           `protobuf:"bytes,12,opt,name=templet,proto3" json:"templet,omitempty"`
	Catalogs             string           `protobuf:"bytes,13,opt,name=catalogs,proto3" json:"catalogs,omitempty"`
	Status               uint32           `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	Remark               string           `protobuf:"bytes,15,opt,name=remark,proto3" json:"remark,omitempty"`
	Revises              []string         `protobuf:"bytes,31,rep,name=revises,proto3" json:"revises,omitempty"`
	Effects              []*ProductEffect `protobuf:"bytes,32,rep,name=effects,proto3" json:"effects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProductInfo) Reset()         { *m = ProductInfo{} }
func (m *ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()    {}
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{0}
}

func (m *ProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInfo.Unmarshal(m, b)
}
func (m *ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInfo.Marshal(b, m, deterministic)
}
func (m *ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfo.Merge(m, src)
}
func (m *ProductInfo) XXX_Size() int {
	return xxx_messageInfo_ProductInfo.Size(m)
}
func (m *ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfo proto.InternalMessageInfo

func (m *ProductInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ProductInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ProductInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ProductInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ProductInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ProductInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ProductInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductInfo) GetMenus() string {
	if m != nil {
		return m.Menus
	}
	return ""
}

func (m *ProductInfo) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func (m *ProductInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProductInfo) GetTemplet() string {
	if m != nil {
		return m.Templet
	}
	return ""
}

func (m *ProductInfo) GetCatalogs() string {
	if m != nil {
		return m.Catalogs
	}
	return ""
}

func (m *ProductInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ProductInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ProductInfo) GetRevises() []string {
	if m != nil {
		return m.Revises
	}
	return nil
}

func (m *ProductInfo) GetEffects() []*ProductEffect {
	if m != nil {
		return m.Effects
	}
	return nil
}

type ProductEffect struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Min                  uint32   `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max                  uint32   `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductEffect) Reset()         { *m = ProductEffect{} }
func (m *ProductEffect) String() string { return proto.CompactTextString(m) }
func (*ProductEffect) ProtoMessage()    {}
func (*ProductEffect) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{1}
}

func (m *ProductEffect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductEffect.Unmarshal(m, b)
}
func (m *ProductEffect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductEffect.Marshal(b, m, deterministic)
}
func (m *ProductEffect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductEffect.Merge(m, src)
}
func (m *ProductEffect) XXX_Size() int {
	return xxx_messageInfo_ProductEffect.Size(m)
}
func (m *ProductEffect) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductEffect.DiscardUnknown(m)
}

var xxx_messageInfo_ProductEffect proto.InternalMessageInfo

func (m *ProductEffect) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ProductEffect) GetMin() uint32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *ProductEffect) GetMax() uint32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type ReqProductAdd struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  string           `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Entry                string           `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
	Menus                string           `protobuf:"bytes,4,opt,name=menus,proto3" json:"menus,omitempty"`
	Operator             string           `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32           `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Templet              string           `protobuf:"bytes,7,opt,name=templet,proto3" json:"templet,omitempty"`
	Remark               string           `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Revises              []string         `protobuf:"bytes,17,rep,name=revises,proto3" json:"revises,omitempty"`
	Effects              []*ProductEffect `protobuf:"bytes,18,rep,name=effects,proto3" json:"effects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqProductAdd) Reset()         { *m = ReqProductAdd{} }
func (m *ReqProductAdd) String() string { return proto.CompactTextString(m) }
func (*ReqProductAdd) ProtoMessage()    {}
func (*ReqProductAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{2}
}

func (m *ReqProductAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqProductAdd.Unmarshal(m, b)
}
func (m *ReqProductAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqProductAdd.Marshal(b, m, deterministic)
}
func (m *ReqProductAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqProductAdd.Merge(m, src)
}
func (m *ReqProductAdd) XXX_Size() int {
	return xxx_messageInfo_ReqProductAdd.Size(m)
}
func (m *ReqProductAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqProductAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqProductAdd proto.InternalMessageInfo

func (m *ReqProductAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqProductAdd) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqProductAdd) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func (m *ReqProductAdd) GetMenus() string {
	if m != nil {
		return m.Menus
	}
	return ""
}

func (m *ReqProductAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqProductAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqProductAdd) GetTemplet() string {
	if m != nil {
		return m.Templet
	}
	return ""
}

func (m *ReqProductAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqProductAdd) GetRevises() []string {
	if m != nil {
		return m.Revises
	}
	return nil
}

func (m *ReqProductAdd) GetEffects() []*ProductEffect {
	if m != nil {
		return m.Effects
	}
	return nil
}

type ReqProductUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqProductUpdate) Reset()         { *m = ReqProductUpdate{} }
func (m *ReqProductUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqProductUpdate) ProtoMessage()    {}
func (*ReqProductUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{3}
}

func (m *ReqProductUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqProductUpdate.Unmarshal(m, b)
}
func (m *ReqProductUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqProductUpdate.Marshal(b, m, deterministic)
}
func (m *ReqProductUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqProductUpdate.Merge(m, src)
}
func (m *ReqProductUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqProductUpdate.Size(m)
}
func (m *ReqProductUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqProductUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqProductUpdate proto.InternalMessageInfo

func (m *ReqProductUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqProductUpdate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqProductUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqProductUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqProductUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyProductInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *ProductInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyProductInfo) Reset()         { *m = ReplyProductInfo{} }
func (m *ReplyProductInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyProductInfo) ProtoMessage()    {}
func (*ReplyProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{4}
}

func (m *ReplyProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyProductInfo.Unmarshal(m, b)
}
func (m *ReplyProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyProductInfo.Marshal(b, m, deterministic)
}
func (m *ReplyProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyProductInfo.Merge(m, src)
}
func (m *ReplyProductInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyProductInfo.Size(m)
}
func (m *ReplyProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyProductInfo proto.InternalMessageInfo

func (m *ReplyProductInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyProductInfo) GetInfo() *ProductInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyProductList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*ProductInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyProductList) Reset()         { *m = ReplyProductList{} }
func (m *ReplyProductList) String() string { return proto.CompactTextString(m) }
func (*ReplyProductList) ProtoMessage()    {}
func (*ReplyProductList) Descriptor() ([]byte, []int) {
	return fileDescriptor_26f760d583f07007, []int{5}
}

func (m *ReplyProductList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyProductList.Unmarshal(m, b)
}
func (m *ReplyProductList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyProductList.Marshal(b, m, deterministic)
}
func (m *ReplyProductList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyProductList.Merge(m, src)
}
func (m *ReplyProductList) XXX_Size() int {
	return xxx_messageInfo_ReplyProductList.Size(m)
}
func (m *ReplyProductList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyProductList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyProductList proto.InternalMessageInfo

func (m *ReplyProductList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyProductList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyProductList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyProductList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyProductList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyProductList) GetList() []*ProductInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "omo.msp.favorite.ProductInfo")
	proto.RegisterType((*ProductEffect)(nil), "omo.msp.favorite.ProductEffect")
	proto.RegisterType((*ReqProductAdd)(nil), "omo.msp.favorite.ReqProductAdd")
	proto.RegisterType((*ReqProductUpdate)(nil), "omo.msp.favorite.ReqProductUpdate")
	proto.RegisterType((*ReplyProductInfo)(nil), "omo.msp.favorite.ReplyProductInfo")
	proto.RegisterType((*ReplyProductList)(nil), "omo.msp.favorite.ReplyProductList")
}

func init() {
	proto.RegisterFile("proto/favorite/product.proto", fileDescriptor_26f760d583f07007)
}

var fileDescriptor_26f760d583f07007 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0xdb, 0x4a,
	0x10, 0xc7, 0xb1, 0x93, 0xc0, 0x06, 0xe7, 0xe5, 0xad, 0xd0, 0xd3, 0x2a, 0xbc, 0x27, 0x2c, 0x9f,
	0x72, 0x0a, 0x7a, 0x54, 0x3d, 0xf4, 0x08, 0x52, 0x8b, 0x2a, 0x15, 0x81, 0x4c, 0xdb, 0x43, 0x6f,
	0xc6, 0x9e, 0x20, 0x8b, 0xd8, 0x6b, 0x76, 0xd7, 0xa8, 0x91, 0xaa, 0xde, 0xfa, 0x1d, 0xfa, 0x59,
	0xfa, 0x39, 0xfa, 0x81, 0xaa, 0x9d, 0x5d, 0x3b, 0x0e, 0x24, 0x40, 0x7b, 0x9b, 0xdf, 0xcc, 0x78,
	0xfe, 0xfc, 0x7e, 0xb3, 0x32, 0xf9, 0xb7, 0x14, 0x5c, 0xf1, 0xc3, 0x59, 0x7c, 0xc7, 0x45, 0xa6,
	0xe0, 0xb0, 0x14, 0x3c, 0xad, 0x12, 0x35, 0x45, 0x37, 0x1d, 0xf1, 0x9c, 0x4f, 0x73, 0x59, 0x4e,
	0xeb, 0xf8, 0x78, 0xff, 0x5e, 0x7e, 0xc2, 0xf3, 0x9c, 0x17, 0x26, 0x3d, 0xfc, 0xe1, 0x92, 0xc1,
	0x85, 0x29, 0xf0, 0xb6, 0x98, 0x71, 0x3a, 0x22, 0x6e, 0x95, 0xa5, 0xcc, 0x09, 0x9c, 0xc9, 0x4e,
	0xa4, 0x4d, 0x3a, 0x24, 0x9d, 0x2c, 0x65, 0x9d, 0xc0, 0x99, 0x78, 0x51, 0x27, 0x4b, 0x29, 0x25,
	0x9e, 0x5a, 0x94, 0xc0, 0xdc, 0xc0, 0x99, 0xf8, 0x11, 0xda, 0x94, 0x91, 0x7e, 0x22, 0x20, 0x56,
	0x90, 0x32, 0x2f, 0x70, 0x26, 0x6e, 0x54, 0x43, 0x1d, 0xa9, 0xca, 0x14, 0x23, 0x5d, 0x13, 0xb1,
	0xb0, 0xf9, 0x86, 0x0b, 0xd6, 0xc3, 0x6e, 0x35, 0xa4, 0x63, 0xb2, 0xcd, 0x4b, 0x10, 0x18, 0xea,
	0x63, 0xa8, 0xc1, 0xba, 0x7b, 0x11, 0xe7, 0xc0, 0xb6, 0xd1, 0x8f, 0x36, 0xdd, 0x23, 0xdd, 0x1c,
	0x8a, 0x4a, 0xb2, 0x1d, 0x74, 0x1a, 0xa0, 0xbd, 0x50, 0x28, 0xb1, 0x60, 0xc4, 0x78, 0x11, 0xe8,
	0xfd, 0x6e, 0x60, 0xc1, 0x06, 0x66, 0xbf, 0x1b, 0x58, 0xe8, 0x39, 0x14, 0xe4, 0xe5, 0x1c, 0x14,
	0xdb, 0x35, 0x73, 0x58, 0xa8, 0xe7, 0x48, 0x62, 0x15, 0xcf, 0xf9, 0xb5, 0x64, 0xbe, 0x99, 0xa3,
	0xc6, 0xf4, 0x1f, 0xd2, 0x93, 0x2a, 0x56, 0x95, 0x64, 0x43, 0xe4, 0xc1, 0x22, 0xed, 0x17, 0x90,
	0xc7, 0xe2, 0x86, 0xfd, 0x85, 0x5f, 0x58, 0xa4, 0xbb, 0x08, 0xb8, 0xcb, 0x24, 0x48, 0x76, 0x10,
	0xb8, 0xba, 0x8b, 0x85, 0xf4, 0x15, 0xe9, 0xc3, 0x6c, 0x06, 0x89, 0x92, 0x2c, 0x08, 0xdc, 0xc9,
	0xe0, 0xe8, 0x60, 0x7a, 0x5f, 0xc2, 0xa9, 0x55, 0xe8, 0x35, 0xe6, 0x45, 0x75, 0x7e, 0x78, 0x46,
	0xfc, 0x95, 0x88, 0xee, 0x52, 0xc6, 0x4a, 0x81, 0x28, 0xac, 0x82, 0x35, 0xd4, 0x7b, 0xe7, 0x59,
	0x81, 0x32, 0xfa, 0x91, 0x36, 0xd1, 0x13, 0x7f, 0xb6, 0x32, 0x6a, 0x33, 0xfc, 0xde, 0x21, 0x7e,
	0x04, 0xb7, 0xb6, 0xe4, 0x71, 0x9a, 0x36, 0x6c, 0x3b, 0x2d, 0xb6, 0x2d, 0x83, 0x9d, 0x25, 0x83,
	0x0d, 0xd3, 0x6e, 0x9b, 0xe9, 0x46, 0x15, 0xaf, 0xad, 0x4a, 0x5b, 0xdb, 0xee, 0x43, 0x6d, 0xf1,
	0xb2, 0x7a, 0xab, 0x97, 0x55, 0xab, 0xd3, 0x5f, 0x55, 0x67, 0xc9, 0xf4, 0xf6, 0x26, 0xa6, 0xff,
	0xde, 0xc8, 0x34, 0xfd, 0x4d, 0xa6, 0xbf, 0x92, 0xd1, 0x92, 0x99, 0x0f, 0x78, 0xc1, 0x6b, 0x9e,
	0xca, 0x43, 0x6a, 0x6a, 0x02, 0xdd, 0x16, 0x81, 0xcb, 0xc1, 0xbd, 0x95, 0xc1, 0x1f, 0xa1, 0x26,
	0xfc, 0xa2, 0xfb, 0x97, 0xf3, 0x45, 0xfb, 0xa9, 0xbe, 0x6c, 0x4e, 0x50, 0x8f, 0x30, 0x38, 0xfa,
	0xef, 0xe1, 0x36, 0xf8, 0xcd, 0x25, 0x26, 0x35, 0x17, 0xfa, 0x3f, 0xf1, 0xb2, 0x62, 0xc6, 0x71,
	0xca, 0xb5, 0x1f, 0xb5, 0x7a, 0x44, 0x98, 0x1a, 0xfe, 0x74, 0x56, 0xdb, 0xbf, 0xcb, 0xa4, 0xfa,
	0xd3, 0xf6, 0x7b, 0xa4, 0xab, 0xb8, 0x8a, 0xe7, 0xf6, 0x14, 0x0d, 0xd0, 0xde, 0x32, 0xbe, 0x06,
	0x69, 0xcf, 0xd1, 0x00, 0xcd, 0x9e, 0x36, 0x90, 0x27, 0x3f, 0x42, 0x5b, 0xb3, 0x57, 0x54, 0xf9,
	0x15, 0x18, 0x8e, 0xfc, 0xc8, 0x22, 0xbd, 0xd6, 0x3c, 0x93, 0x8a, 0x8d, 0x50, 0xd9, 0xa7, 0xd6,
	0xd2, 0xa9, 0x47, 0xdf, 0x3c, 0x32, 0xb4, 0xde, 0x4b, 0x10, 0x77, 0x59, 0x02, 0xf4, 0x9c, 0xf4,
	0x8e, 0xd3, 0xf4, 0xbc, 0x00, 0x7a, 0xb0, 0x6e, 0x9d, 0xd6, 0xdb, 0x18, 0x87, 0x1b, 0xf6, 0x6d,
	0xf5, 0x09, 0xb7, 0xe8, 0x19, 0xe9, 0x9d, 0x82, 0xd2, 0x05, 0xd7, 0xf2, 0x73, 0x5b, 0x81, 0xc4,
	0xd4, 0x67, 0x96, 0x7b, 0x4f, 0x06, 0xa7, 0xa0, 0x4e, 0x16, 0x6f, 0xb2, 0xb9, 0x02, 0xb1, 0x61,
	0x48, 0x5d, 0xd3, 0x24, 0x3c, 0x55, 0x55, 0x0b, 0x19, 0x6e, 0xd1, 0x4b, 0xb2, 0x7b, 0x0a, 0x4a,
	0x0b, 0x95, 0x49, 0x95, 0x25, 0x4f, 0x97, 0x0d, 0x1e, 0xd1, 0x1a, 0x4b, 0x84, 0x5b, 0xf4, 0x82,
	0x0c, 0xcd, 0x43, 0x79, 0xc6, 0xb4, 0x26, 0x71, 0xbc, 0xbf, 0xa1, 0xac, 0x5d, 0xfe, 0x23, 0x21,
	0xb6, 0x62, 0x2c, 0x81, 0x86, 0x8f, 0x09, 0x64, 0x0b, 0x3e, 0x8b, 0xd4, 0x93, 0xd1, 0xa7, 0xe1,
	0xea, 0x2f, 0xf2, 0xaa, 0x87, 0xf8, 0xc5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xc8, 0xe8,
	0x7b, 0x6b, 0x07, 0x00, 0x00,
}