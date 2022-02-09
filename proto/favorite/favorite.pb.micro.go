// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/favorite/favorite.proto

package favorite

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FavoriteService service

type FavoriteService interface {
	AddOne(ctx context.Context, in *ReqFavoriteAdd, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	GetByOrigin(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	GetList(ctx context.Context, in *ReqFavoriteList, opts ...client.CallOption) (*ReplyFavoriteList, error)
	GetByList(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyFavoriteList, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyFavoriteList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateBase(ctx context.Context, in *ReqFavoriteUpdate, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	UpdateMeta(ctx context.Context, in *ReqFavoriteMeta, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	UpdateTags(ctx context.Context, in *ReqFavoriteTags, opts ...client.CallOption) (*ReplyFavoriteInfo, error)
	UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateKeys(ctx context.Context, in *ReqFavoriteKeys, opts ...client.CallOption) (*ReplyFavoriteKeys, error)
	AppendKey(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteKeys, error)
	SubtractKey(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteKeys, error)
	UpdateTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error)
	AppendTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error)
	SubtractTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error)
	UpdateTargets(ctx context.Context, in *ReqFavoriteTargets, opts ...client.CallOption) (*ReplyInfo, error)
}

type favoriteService struct {
	c    client.Client
	name string
}

func NewFavoriteService(name string, c client.Client) FavoriteService {
	return &favoriteService{
		c:    c,
		name: name,
	}
}

func (c *favoriteService) AddOne(ctx context.Context, in *ReqFavoriteAdd, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.AddOne", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetOne", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetByOrigin(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetByOrigin", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetList(ctx context.Context, in *ReqFavoriteList, opts ...client.CallOption) (*ReplyFavoriteList, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetList", in)
	out := new(ReplyFavoriteList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetByList(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyFavoriteList, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetByList", in)
	out := new(ReplyFavoriteList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyFavoriteList, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetByFilter", in)
	out := new(ReplyFavoriteList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateBase(ctx context.Context, in *ReqFavoriteUpdate, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateBase", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateMeta(ctx context.Context, in *ReqFavoriteMeta, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateMeta", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateTags(ctx context.Context, in *ReqFavoriteTags, opts ...client.CallOption) (*ReplyFavoriteInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateTags", in)
	out := new(ReplyFavoriteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateKeys(ctx context.Context, in *ReqFavoriteKeys, opts ...client.CallOption) (*ReplyFavoriteKeys, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateKeys", in)
	out := new(ReplyFavoriteKeys)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) AppendKey(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteKeys, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.AppendKey", in)
	out := new(ReplyFavoriteKeys)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) SubtractKey(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyFavoriteKeys, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.SubtractKey", in)
	out := new(ReplyFavoriteKeys)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateTarget", in)
	out := new(ReplyFavoriteTargets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) AppendTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.AppendTarget", in)
	out := new(ReplyFavoriteTargets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) SubtractTarget(ctx context.Context, in *ReqFavoriteTarget, opts ...client.CallOption) (*ReplyFavoriteTargets, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.SubtractTarget", in)
	out := new(ReplyFavoriteTargets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteService) UpdateTargets(ctx context.Context, in *ReqFavoriteTargets, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "FavoriteService.UpdateTargets", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FavoriteService service

type FavoriteServiceHandler interface {
	AddOne(context.Context, *ReqFavoriteAdd, *ReplyFavoriteInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyFavoriteInfo) error
	GetByOrigin(context.Context, *RequestInfo, *ReplyFavoriteInfo) error
	GetList(context.Context, *ReqFavoriteList, *ReplyFavoriteList) error
	GetByList(context.Context, *RequestList, *ReplyFavoriteList) error
	GetByFilter(context.Context, *RequestFilter, *ReplyFavoriteList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateBase(context.Context, *ReqFavoriteUpdate, *ReplyFavoriteInfo) error
	UpdateMeta(context.Context, *ReqFavoriteMeta, *ReplyFavoriteInfo) error
	UpdateTags(context.Context, *ReqFavoriteTags, *ReplyFavoriteInfo) error
	UpdateStatus(context.Context, *RequestState, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	UpdateKeys(context.Context, *ReqFavoriteKeys, *ReplyFavoriteKeys) error
	AppendKey(context.Context, *RequestInfo, *ReplyFavoriteKeys) error
	SubtractKey(context.Context, *RequestInfo, *ReplyFavoriteKeys) error
	UpdateTarget(context.Context, *ReqFavoriteTarget, *ReplyFavoriteTargets) error
	AppendTarget(context.Context, *ReqFavoriteTarget, *ReplyFavoriteTargets) error
	SubtractTarget(context.Context, *ReqFavoriteTarget, *ReplyFavoriteTargets) error
	UpdateTargets(context.Context, *ReqFavoriteTargets, *ReplyInfo) error
}

func RegisterFavoriteServiceHandler(s server.Server, hdlr FavoriteServiceHandler, opts ...server.HandlerOption) error {
	type favoriteService interface {
		AddOne(ctx context.Context, in *ReqFavoriteAdd, out *ReplyFavoriteInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyFavoriteInfo) error
		GetByOrigin(ctx context.Context, in *RequestInfo, out *ReplyFavoriteInfo) error
		GetList(ctx context.Context, in *ReqFavoriteList, out *ReplyFavoriteList) error
		GetByList(ctx context.Context, in *RequestList, out *ReplyFavoriteList) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyFavoriteList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateBase(ctx context.Context, in *ReqFavoriteUpdate, out *ReplyFavoriteInfo) error
		UpdateMeta(ctx context.Context, in *ReqFavoriteMeta, out *ReplyFavoriteInfo) error
		UpdateTags(ctx context.Context, in *ReqFavoriteTags, out *ReplyFavoriteInfo) error
		UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		UpdateKeys(ctx context.Context, in *ReqFavoriteKeys, out *ReplyFavoriteKeys) error
		AppendKey(ctx context.Context, in *RequestInfo, out *ReplyFavoriteKeys) error
		SubtractKey(ctx context.Context, in *RequestInfo, out *ReplyFavoriteKeys) error
		UpdateTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error
		AppendTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error
		SubtractTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error
		UpdateTargets(ctx context.Context, in *ReqFavoriteTargets, out *ReplyInfo) error
	}
	type FavoriteService struct {
		favoriteService
	}
	h := &favoriteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FavoriteService{h}, opts...))
}

type favoriteServiceHandler struct {
	FavoriteServiceHandler
}

func (h *favoriteServiceHandler) AddOne(ctx context.Context, in *ReqFavoriteAdd, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.AddOne(ctx, in, out)
}

func (h *favoriteServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.GetOne(ctx, in, out)
}

func (h *favoriteServiceHandler) GetByOrigin(ctx context.Context, in *RequestInfo, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.GetByOrigin(ctx, in, out)
}

func (h *favoriteServiceHandler) GetList(ctx context.Context, in *ReqFavoriteList, out *ReplyFavoriteList) error {
	return h.FavoriteServiceHandler.GetList(ctx, in, out)
}

func (h *favoriteServiceHandler) GetByList(ctx context.Context, in *RequestList, out *ReplyFavoriteList) error {
	return h.FavoriteServiceHandler.GetByList(ctx, in, out)
}

func (h *favoriteServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyFavoriteList) error {
	return h.FavoriteServiceHandler.GetByFilter(ctx, in, out)
}

func (h *favoriteServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.FavoriteServiceHandler.GetStatistic(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.FavoriteServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateBase(ctx context.Context, in *ReqFavoriteUpdate, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.UpdateBase(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateMeta(ctx context.Context, in *ReqFavoriteMeta, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.UpdateMeta(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateTags(ctx context.Context, in *ReqFavoriteTags, out *ReplyFavoriteInfo) error {
	return h.FavoriteServiceHandler.UpdateTags(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error {
	return h.FavoriteServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *favoriteServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.FavoriteServiceHandler.RemoveOne(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateKeys(ctx context.Context, in *ReqFavoriteKeys, out *ReplyFavoriteKeys) error {
	return h.FavoriteServiceHandler.UpdateKeys(ctx, in, out)
}

func (h *favoriteServiceHandler) AppendKey(ctx context.Context, in *RequestInfo, out *ReplyFavoriteKeys) error {
	return h.FavoriteServiceHandler.AppendKey(ctx, in, out)
}

func (h *favoriteServiceHandler) SubtractKey(ctx context.Context, in *RequestInfo, out *ReplyFavoriteKeys) error {
	return h.FavoriteServiceHandler.SubtractKey(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error {
	return h.FavoriteServiceHandler.UpdateTarget(ctx, in, out)
}

func (h *favoriteServiceHandler) AppendTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error {
	return h.FavoriteServiceHandler.AppendTarget(ctx, in, out)
}

func (h *favoriteServiceHandler) SubtractTarget(ctx context.Context, in *ReqFavoriteTarget, out *ReplyFavoriteTargets) error {
	return h.FavoriteServiceHandler.SubtractTarget(ctx, in, out)
}

func (h *favoriteServiceHandler) UpdateTargets(ctx context.Context, in *ReqFavoriteTargets, out *ReplyInfo) error {
	return h.FavoriteServiceHandler.UpdateTargets(ctx, in, out)
}
