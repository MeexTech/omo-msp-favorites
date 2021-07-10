// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/favorite.proto

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

type FavoriteInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Creator              string   `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,9,opt,name=operator,proto3" json:"operator,omitempty"`
	Cover                string   `protobuf:"bytes,10,opt,name=cover,proto3" json:"cover,omitempty"`
	Origin               string   `protobuf:"bytes,11,opt,name=origin,proto3" json:"origin,omitempty"`
	Type                 uint32   `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
	Tags                 []string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	Keys                 []string `protobuf:"bytes,15,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavoriteInfo) Reset()         { *m = FavoriteInfo{} }
func (m *FavoriteInfo) String() string { return proto.CompactTextString(m) }
func (*FavoriteInfo) ProtoMessage()    {}
func (*FavoriteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{0}
}

func (m *FavoriteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteInfo.Unmarshal(m, b)
}
func (m *FavoriteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteInfo.Marshal(b, m, deterministic)
}
func (m *FavoriteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteInfo.Merge(m, src)
}
func (m *FavoriteInfo) XXX_Size() int {
	return xxx_messageInfo_FavoriteInfo.Size(m)
}
func (m *FavoriteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteInfo proto.InternalMessageInfo

func (m *FavoriteInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FavoriteInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FavoriteInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FavoriteInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *FavoriteInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FavoriteInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FavoriteInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *FavoriteInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FavoriteInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *FavoriteInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *FavoriteInfo) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *FavoriteInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FavoriteInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FavoriteInfo) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ReqFavoriteAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Origin               string   `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
	Person               bool     `protobuf:"varint,8,opt,name=person,proto3" json:"person,omitempty"`
	Keys                 []string `protobuf:"bytes,9,rep,name=keys,proto3" json:"keys,omitempty"`
	Tags                 []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteAdd) Reset()         { *m = ReqFavoriteAdd{} }
func (m *ReqFavoriteAdd) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteAdd) ProtoMessage()    {}
func (*ReqFavoriteAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{1}
}

func (m *ReqFavoriteAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteAdd.Unmarshal(m, b)
}
func (m *ReqFavoriteAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteAdd.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteAdd.Merge(m, src)
}
func (m *ReqFavoriteAdd) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteAdd.Size(m)
}
func (m *ReqFavoriteAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteAdd proto.InternalMessageInfo

func (m *ReqFavoriteAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFavoriteAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqFavoriteAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFavoriteAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqFavoriteAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqFavoriteAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteAdd) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ReqFavoriteAdd) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

func (m *ReqFavoriteAdd) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ReqFavoriteAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqFavoriteUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Origin               string   `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	Type                 uint32   `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Person               bool     `protobuf:"varint,8,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteUpdate) Reset()         { *m = ReqFavoriteUpdate{} }
func (m *ReqFavoriteUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteUpdate) ProtoMessage()    {}
func (*ReqFavoriteUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{2}
}

func (m *ReqFavoriteUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteUpdate.Unmarshal(m, b)
}
func (m *ReqFavoriteUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteUpdate.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteUpdate.Merge(m, src)
}
func (m *ReqFavoriteUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteUpdate.Size(m)
}
func (m *ReqFavoriteUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteUpdate proto.InternalMessageInfo

func (m *ReqFavoriteUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqFavoriteUpdate) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

type ReqFavoriteTags struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Person               bool     `protobuf:"varint,3,opt,name=person,proto3" json:"person,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteTags) Reset()         { *m = ReqFavoriteTags{} }
func (m *ReqFavoriteTags) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteTags) ProtoMessage()    {}
func (*ReqFavoriteTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{3}
}

func (m *ReqFavoriteTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteTags.Unmarshal(m, b)
}
func (m *ReqFavoriteTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteTags.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteTags.Merge(m, src)
}
func (m *ReqFavoriteTags) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteTags.Size(m)
}
func (m *ReqFavoriteTags) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteTags.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteTags proto.InternalMessageInfo

func (m *ReqFavoriteTags) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteTags) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteTags) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

func (m *ReqFavoriteTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqFavoriteList struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Person               bool     `protobuf:"varint,3,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteList) Reset()         { *m = ReqFavoriteList{} }
func (m *ReqFavoriteList) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteList) ProtoMessage()    {}
func (*ReqFavoriteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{4}
}

func (m *ReqFavoriteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteList.Unmarshal(m, b)
}
func (m *ReqFavoriteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteList.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteList.Merge(m, src)
}
func (m *ReqFavoriteList) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteList.Size(m)
}
func (m *ReqFavoriteList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteList proto.InternalMessageInfo

func (m *ReqFavoriteList) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqFavoriteList) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqFavoriteList) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

type ReplyFavoriteInfo struct {
	Info                 *FavoriteInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyFavoriteInfo) Reset()         { *m = ReplyFavoriteInfo{} }
func (m *ReplyFavoriteInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteInfo) ProtoMessage()    {}
func (*ReplyFavoriteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{5}
}

func (m *ReplyFavoriteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteInfo.Unmarshal(m, b)
}
func (m *ReplyFavoriteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteInfo.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteInfo.Merge(m, src)
}
func (m *ReplyFavoriteInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteInfo.Size(m)
}
func (m *ReplyFavoriteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteInfo proto.InternalMessageInfo

func (m *ReplyFavoriteInfo) GetInfo() *FavoriteInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyFavoriteInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyFavoriteList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*FavoriteInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyFavoriteList) Reset()         { *m = ReplyFavoriteList{} }
func (m *ReplyFavoriteList) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteList) ProtoMessage()    {}
func (*ReplyFavoriteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{6}
}

func (m *ReplyFavoriteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteList.Unmarshal(m, b)
}
func (m *ReplyFavoriteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteList.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteList.Merge(m, src)
}
func (m *ReplyFavoriteList) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteList.Size(m)
}
func (m *ReplyFavoriteList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteList proto.InternalMessageInfo

func (m *ReplyFavoriteList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFavoriteList) GetList() []*FavoriteInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqFavoriteKeys struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Person               bool     `protobuf:"varint,3,opt,name=person,proto3" json:"person,omitempty"`
	Keys                 []string `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteKeys) Reset()         { *m = ReqFavoriteKeys{} }
func (m *ReqFavoriteKeys) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteKeys) ProtoMessage()    {}
func (*ReqFavoriteKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{7}
}

func (m *ReqFavoriteKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteKeys.Unmarshal(m, b)
}
func (m *ReqFavoriteKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteKeys.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteKeys.Merge(m, src)
}
func (m *ReqFavoriteKeys) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteKeys.Size(m)
}
func (m *ReqFavoriteKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteKeys proto.InternalMessageInfo

func (m *ReqFavoriteKeys) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteKeys) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteKeys) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

func (m *ReqFavoriteKeys) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ReplyFavoriteKeys struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Keys                 []string     `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyFavoriteKeys) Reset()         { *m = ReplyFavoriteKeys{} }
func (m *ReplyFavoriteKeys) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteKeys) ProtoMessage()    {}
func (*ReplyFavoriteKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{8}
}

func (m *ReplyFavoriteKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteKeys.Unmarshal(m, b)
}
func (m *ReplyFavoriteKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteKeys.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteKeys.Merge(m, src)
}
func (m *ReplyFavoriteKeys) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteKeys.Size(m)
}
func (m *ReplyFavoriteKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteKeys proto.InternalMessageInfo

func (m *ReplyFavoriteKeys) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyFavoriteKeys) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ReplyFavoriteKeys) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*FavoriteInfo)(nil), "omo.msp.Favorite.FavoriteInfo")
	proto.RegisterType((*ReqFavoriteAdd)(nil), "omo.msp.Favorite.ReqFavoriteAdd")
	proto.RegisterType((*ReqFavoriteUpdate)(nil), "omo.msp.Favorite.ReqFavoriteUpdate")
	proto.RegisterType((*ReqFavoriteTags)(nil), "omo.msp.Favorite.ReqFavoriteTags")
	proto.RegisterType((*ReqFavoriteList)(nil), "omo.msp.Favorite.ReqFavoriteList")
	proto.RegisterType((*ReplyFavoriteInfo)(nil), "omo.msp.Favorite.ReplyFavoriteInfo")
	proto.RegisterType((*ReplyFavoriteList)(nil), "omo.msp.Favorite.ReplyFavoriteList")
	proto.RegisterType((*ReqFavoriteKeys)(nil), "omo.msp.Favorite.ReqFavoriteKeys")
	proto.RegisterType((*ReplyFavoriteKeys)(nil), "omo.msp.Favorite.ReplyFavoriteKeys")
}

func init() {
	proto.RegisterFile("proto/favorite/favorite.proto", fileDescriptor_bdf9f6d1ce804485)
}

var fileDescriptor_bdf9f6d1ce804485 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x4e, 0xdb, 0x30,
	0x18, 0x25, 0x49, 0x9b, 0xd2, 0xaf, 0xac, 0x30, 0x6b, 0x42, 0x16, 0x88, 0x29, 0x0b, 0x37, 0xbd,
	0xea, 0xa4, 0x4e, 0x7b, 0x00, 0x90, 0x36, 0x84, 0x36, 0x0d, 0x91, 0x32, 0x34, 0xed, 0x2e, 0x34,
	0x06, 0x45, 0x90, 0x38, 0x38, 0x6e, 0xa7, 0xde, 0xec, 0xa9, 0xf6, 0x0c, 0xdb, 0xeb, 0xec, 0x11,
	0x26, 0x7f, 0x49, 0x1a, 0x37, 0x4d, 0x4b, 0xab, 0x71, 0xe7, 0xcf, 0x3e, 0x39, 0xdf, 0x39, 0xc7,
	0x3f, 0x0a, 0x1c, 0x25, 0x82, 0x4b, 0xfe, 0xf6, 0xd6, 0x9f, 0x70, 0x11, 0x4a, 0x36, 0x1b, 0xf4,
	0x71, 0x9e, 0xec, 0xf1, 0x88, 0xf7, 0xa3, 0x34, 0xe9, 0x7f, 0xcc, 0xe7, 0x0f, 0x0e, 0x2b, 0x1f,
	0x8c, 0x78, 0x14, 0xf1, 0x38, 0x83, 0xbb, 0xbf, 0x4d, 0xd8, 0x29, 0x90, 0xe7, 0xf1, 0x2d, 0x27,
	0x7b, 0x60, 0x8d, 0xc3, 0x80, 0x1a, 0x8e, 0xd1, 0x6b, 0x7b, 0x6a, 0x48, 0xba, 0x60, 0x86, 0x01,
	0x35, 0x1d, 0xa3, 0xd7, 0xf0, 0xcc, 0x30, 0x20, 0x14, 0x5a, 0x23, 0xc1, 0x7c, 0xc9, 0x02, 0x6a,
	0x39, 0x46, 0xcf, 0xf2, 0x8a, 0x52, 0xad, 0x8c, 0x93, 0x00, 0x57, 0x1a, 0xd9, 0x4a, 0x5e, 0x12,
	0x02, 0x8d, 0xd8, 0x8f, 0x18, 0x6d, 0x22, 0x2d, 0x8e, 0xc9, 0x2b, 0x68, 0xf2, 0x1f, 0x31, 0x13,
	0xd4, 0xc6, 0xc9, 0xac, 0x20, 0xfb, 0x60, 0x0b, 0x16, 0xf9, 0xe2, 0x9e, 0xb6, 0x70, 0x3a, 0xaf,
	0x66, 0x5d, 0xb9, 0xa0, 0xdb, 0xb8, 0x50, 0x94, 0xe4, 0x00, 0xb6, 0x79, 0xc2, 0x04, 0x2e, 0xb5,
	0x71, 0x69, 0x56, 0xab, 0x1e, 0x23, 0x3e, 0x61, 0x82, 0x42, 0xd6, 0x03, 0x0b, 0xd5, 0x83, 0x8b,
	0xf0, 0x2e, 0x8c, 0x69, 0x27, 0xeb, 0x91, 0x55, 0x4a, 0xa5, 0x9c, 0x26, 0x8c, 0xee, 0x38, 0x46,
	0xef, 0x85, 0x87, 0x63, 0x9c, 0xf3, 0xef, 0x52, 0xda, 0x75, 0x2c, 0xa5, 0x5c, 0x8d, 0xd5, 0xdc,
	0x3d, 0x9b, 0xa6, 0x74, 0x37, 0x9b, 0x53, 0x63, 0xf7, 0xaf, 0x01, 0x5d, 0x8f, 0x3d, 0x16, 0x59,
	0x9e, 0x04, 0xa5, 0x69, 0xa3, 0xce, 0xb4, 0x59, 0x6f, 0xda, 0x9a, 0x33, 0x3d, 0x93, 0xdf, 0xd0,
	0xe5, 0x17, 0x32, 0x55, 0x98, 0xcd, 0x5c, 0xa6, 0x1e, 0x82, 0x5d, 0x09, 0xa1, 0xb4, 0xdb, 0x9a,
	0xb3, 0xbb, 0x0f, 0x76, 0xc2, 0x44, 0xca, 0x63, 0x4c, 0x74, 0xdb, 0xcb, 0xab, 0x99, 0xbd, 0x76,
	0x69, 0x6f, 0x16, 0x03, 0x94, 0x31, 0xb8, 0x7f, 0x0c, 0x78, 0xa9, 0x59, 0xfe, 0x8a, 0x7b, 0x5d,
	0x73, 0x80, 0x8a, 0x1c, 0x4c, 0x2d, 0x87, 0xcd, 0x1c, 0xeb, 0xee, 0x9a, 0x4b, 0xdd, 0xd9, 0xb5,
	0x9b, 0xd9, 0xd2, 0x36, 0x73, 0x89, 0x63, 0xf7, 0x1e, 0x76, 0x35, 0x23, 0x57, 0x6a, 0x8f, 0x17,
	0x6d, 0xe8, 0x22, 0xcc, 0x45, 0x11, 0x39, 0xb1, 0x55, 0x8d, 0x12, 0x63, 0x6b, 0x68, 0xb1, 0x0d,
	0xe7, 0x9a, 0x7d, 0x0e, 0x53, 0x59, 0x9e, 0x0a, 0x43, 0x3f, 0x15, 0x85, 0x03, 0xb3, 0xd6, 0xc1,
	0x5c, 0x23, 0xf7, 0xa7, 0xda, 0x8a, 0xe4, 0x61, 0x3a, 0x77, 0x97, 0x07, 0xd0, 0x08, 0xe3, 0x5b,
	0x8e, 0xac, 0x9d, 0xc1, 0xeb, 0x7e, 0xf5, 0x69, 0xe8, 0xeb, 0x68, 0x0f, 0xb1, 0xe4, 0x3d, 0xd8,
	0xa9, 0xf4, 0xe5, 0x38, 0xc5, 0xb6, 0x9d, 0xc1, 0xd1, 0xe2, 0x57, 0xd8, 0x68, 0x88, 0x20, 0x2f,
	0x07, 0x2f, 0xf4, 0x47, 0x5b, 0x25, 0x97, 0xb1, 0x01, 0x97, 0x92, 0xfd, 0x10, 0xa6, 0x92, 0x9a,
	0x8e, 0xb5, 0x8e, 0x6c, 0x85, 0xad, 0xec, 0xe0, 0x27, 0x75, 0x64, 0x9f, 0x6d, 0x07, 0xf1, 0x32,
	0x34, 0xb4, 0xbb, 0x9e, 0x54, 0xcc, 0x2e, 0x69, 0x57, 0x7c, 0x6a, 0x6a, 0xf7, 0xa8, 0x8c, 0xc4,
	0xda, 0x20, 0x92, 0xc1, 0xaf, 0x16, 0xec, 0x16, 0x80, 0x21, 0x13, 0x93, 0x70, 0xc4, 0xc8, 0x25,
	0xd8, 0x27, 0x41, 0x70, 0x11, 0x33, 0xe2, 0xd4, 0x91, 0xe8, 0x4f, 0xd1, 0xc1, 0xf1, 0x92, 0x36,
	0x7a, 0x92, 0xee, 0x16, 0xf9, 0x02, 0xf6, 0x19, 0x93, 0x8a, 0xb2, 0x56, 0xd7, 0xe3, 0x98, 0xa5,
	0x52, 0x41, 0xd7, 0xe5, 0x1b, 0x42, 0xe7, 0x8c, 0xc9, 0xd3, 0xe9, 0x45, 0x76, 0x25, 0x9f, 0x8b,
	0xb4, 0x75, 0xc6, 0x24, 0x1e, 0xb0, 0x37, 0x2b, 0x8d, 0x2b, 0xc8, 0x93, 0xa4, 0x0a, 0xe4, 0x6e,
	0x91, 0x4b, 0x68, 0xa3, 0x52, 0xa4, 0x5d, 0xae, 0x73, 0x13, 0xca, 0x6f, 0x00, 0xd9, 0x93, 0x78,
	0xea, 0xa7, 0x8c, 0x1c, 0xaf, 0x94, 0x9a, 0x01, 0xd7, 0x4d, 0xe0, 0xba, 0x60, 0xc6, 0x97, 0x6a,
	0x75, 0x08, 0x0a, 0xb2, 0x2e, 0xef, 0x39, 0xb4, 0x3d, 0x16, 0xf1, 0x09, 0x5b, 0xe3, 0x04, 0x1c,
	0x2e, 0xa1, 0xac, 0x4a, 0xc4, 0xbb, 0xb1, 0x5a, 0xa2, 0x82, 0x3c, 0x29, 0x51, 0x81, 0xdc, 0x2d,
	0x72, 0x05, 0x3b, 0x27, 0x49, 0xc2, 0xe2, 0xe0, 0x43, 0x2c, 0x43, 0x39, 0xfd, 0xdf, 0x23, 0x95,
	0xb3, 0x5e, 0x43, 0x77, 0x38, 0xbe, 0x91, 0xc2, 0x1f, 0xc9, 0xe7, 0xe4, 0x3d, 0xdd, 0xfb, 0xde,
	0x9d, 0xff, 0xf9, 0xba, 0xb1, 0xb1, 0x7e, 0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa5, 0x3a,
	0xfb, 0xc6, 0x09, 0x00, 0x00,
}
