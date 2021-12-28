// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/favorite/article.proto

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

// Client API for ArticleService service

type ArticleService interface {
	AddOne(ctx context.Context, in *ReqArticleAdd, opts ...client.CallOption) (*ReplyArticleInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyArticleInfo, error)
	GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyArticleList, error)
	UpdateBase(ctx context.Context, in *ReqArticleUpdate, opts ...client.CallOption) (*ReplyArticleInfo, error)
	UpdateAssets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *ReqArticleState, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateTargets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
}

type articleService struct {
	c    client.Client
	name string
}

func NewArticleService(name string, c client.Client) ArticleService {
	return &articleService{
		c:    c,
		name: name,
	}
}

func (c *articleService) AddOne(ctx context.Context, in *ReqArticleAdd, opts ...client.CallOption) (*ReplyArticleInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.AddOne", in)
	out := new(ReplyArticleInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyArticleInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.GetOne", in)
	out := new(ReplyArticleInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyArticleList, error) {
	req := c.c.NewRequest(c.name, "ArticleService.GetList", in)
	out := new(ReplyArticleList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) UpdateBase(ctx context.Context, in *ReqArticleUpdate, opts ...client.CallOption) (*ReplyArticleInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.UpdateBase", in)
	out := new(ReplyArticleInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) UpdateAssets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.UpdateAssets", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) UpdateStatus(ctx context.Context, in *ReqArticleState, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.UpdateTags", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) UpdateTargets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ArticleService.UpdateTargets", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ArticleService service

type ArticleServiceHandler interface {
	AddOne(context.Context, *ReqArticleAdd, *ReplyArticleInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyArticleInfo) error
	GetList(context.Context, *RequestFilter, *ReplyArticleList) error
	UpdateBase(context.Context, *ReqArticleUpdate, *ReplyArticleInfo) error
	UpdateAssets(context.Context, *RequestList, *ReplyInfo) error
	UpdateStatus(context.Context, *ReqArticleState, *ReplyInfo) error
	UpdateTags(context.Context, *RequestList, *ReplyInfo) error
	UpdateTargets(context.Context, *RequestList, *ReplyInfo) error
}

func RegisterArticleServiceHandler(s server.Server, hdlr ArticleServiceHandler, opts ...server.HandlerOption) error {
	type articleService interface {
		AddOne(ctx context.Context, in *ReqArticleAdd, out *ReplyArticleInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyArticleInfo) error
		GetList(ctx context.Context, in *RequestFilter, out *ReplyArticleList) error
		UpdateBase(ctx context.Context, in *ReqArticleUpdate, out *ReplyArticleInfo) error
		UpdateAssets(ctx context.Context, in *RequestList, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *ReqArticleState, out *ReplyInfo) error
		UpdateTags(ctx context.Context, in *RequestList, out *ReplyInfo) error
		UpdateTargets(ctx context.Context, in *RequestList, out *ReplyInfo) error
	}
	type ArticleService struct {
		articleService
	}
	h := &articleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ArticleService{h}, opts...))
}

type articleServiceHandler struct {
	ArticleServiceHandler
}

func (h *articleServiceHandler) AddOne(ctx context.Context, in *ReqArticleAdd, out *ReplyArticleInfo) error {
	return h.ArticleServiceHandler.AddOne(ctx, in, out)
}

func (h *articleServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyArticleInfo) error {
	return h.ArticleServiceHandler.GetOne(ctx, in, out)
}

func (h *articleServiceHandler) GetList(ctx context.Context, in *RequestFilter, out *ReplyArticleList) error {
	return h.ArticleServiceHandler.GetList(ctx, in, out)
}

func (h *articleServiceHandler) UpdateBase(ctx context.Context, in *ReqArticleUpdate, out *ReplyArticleInfo) error {
	return h.ArticleServiceHandler.UpdateBase(ctx, in, out)
}

func (h *articleServiceHandler) UpdateAssets(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.ArticleServiceHandler.UpdateAssets(ctx, in, out)
}

func (h *articleServiceHandler) UpdateStatus(ctx context.Context, in *ReqArticleState, out *ReplyInfo) error {
	return h.ArticleServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *articleServiceHandler) UpdateTags(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.ArticleServiceHandler.UpdateTags(ctx, in, out)
}

func (h *articleServiceHandler) UpdateTargets(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.ArticleServiceHandler.UpdateTargets(ctx, in, out)
}
