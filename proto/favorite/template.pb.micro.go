// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/favorite/template.proto

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

// Client API for ProductTemplateService service

type ProductTemplateService interface {
	AddOne(ctx context.Context, in *ReqProductTemplateAdd, opts ...client.CallOption) (*ReplyProductTemplateInfo, error)
	Import(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyProductTemplateList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateBase(ctx context.Context, in *ReqProductTemplateUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateGraph(ctx context.Context, in *ReqProductTemplateGraph, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateReferences(ctx context.Context, in *ReqProductTemplateReference, opts ...client.CallOption) (*ReplyProductTemplateRefs, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetReleaseOn(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateResult, error)
	GetReleaseList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyProductTemplateResults, error)
}

type productTemplateService struct {
	c    client.Client
	name string
}

func NewProductTemplateService(name string, c client.Client) ProductTemplateService {
	return &productTemplateService{
		c:    c,
		name: name,
	}
}

func (c *productTemplateService) AddOne(ctx context.Context, in *ReqProductTemplateAdd, opts ...client.CallOption) (*ReplyProductTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.AddOne", in)
	out := new(ReplyProductTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) Import(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.Import", in)
	out := new(ReplyProductTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.GetOne", in)
	out := new(ReplyProductTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyProductTemplateList, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.GetByFilter", in)
	out := new(ReplyProductTemplateList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) UpdateBase(ctx context.Context, in *ReqProductTemplateUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) UpdateGraph(ctx context.Context, in *ReqProductTemplateGraph, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.UpdateGraph", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) UpdateReferences(ctx context.Context, in *ReqProductTemplateReference, opts ...client.CallOption) (*ReplyProductTemplateRefs, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.UpdateReferences", in)
	out := new(ReplyProductTemplateRefs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) GetReleaseOn(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyProductTemplateResult, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.GetReleaseOn", in)
	out := new(ReplyProductTemplateResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productTemplateService) GetReleaseList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyProductTemplateResults, error) {
	req := c.c.NewRequest(c.name, "ProductTemplateService.GetReleaseList", in)
	out := new(ReplyProductTemplateResults)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductTemplateService service

type ProductTemplateServiceHandler interface {
	AddOne(context.Context, *ReqProductTemplateAdd, *ReplyProductTemplateInfo) error
	Import(context.Context, *RequestInfo, *ReplyProductTemplateInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyProductTemplateInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyProductTemplateList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateBase(context.Context, *ReqProductTemplateUpdate, *ReplyInfo) error
	UpdateGraph(context.Context, *ReqProductTemplateGraph, *ReplyInfo) error
	UpdateStatus(context.Context, *RequestState, *ReplyInfo) error
	UpdateReferences(context.Context, *ReqProductTemplateReference, *ReplyProductTemplateRefs) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetReleaseOn(context.Context, *RequestInfo, *ReplyProductTemplateResult) error
	GetReleaseList(context.Context, *RequestFilter, *ReplyProductTemplateResults) error
}

func RegisterProductTemplateServiceHandler(s server.Server, hdlr ProductTemplateServiceHandler, opts ...server.HandlerOption) error {
	type productTemplateService interface {
		AddOne(ctx context.Context, in *ReqProductTemplateAdd, out *ReplyProductTemplateInfo) error
		Import(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyProductTemplateList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateBase(ctx context.Context, in *ReqProductTemplateUpdate, out *ReplyInfo) error
		UpdateGraph(ctx context.Context, in *ReqProductTemplateGraph, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error
		UpdateReferences(ctx context.Context, in *ReqProductTemplateReference, out *ReplyProductTemplateRefs) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetReleaseOn(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateResult) error
		GetReleaseList(ctx context.Context, in *RequestFilter, out *ReplyProductTemplateResults) error
	}
	type ProductTemplateService struct {
		productTemplateService
	}
	h := &productTemplateServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductTemplateService{h}, opts...))
}

type productTemplateServiceHandler struct {
	ProductTemplateServiceHandler
}

func (h *productTemplateServiceHandler) AddOne(ctx context.Context, in *ReqProductTemplateAdd, out *ReplyProductTemplateInfo) error {
	return h.ProductTemplateServiceHandler.AddOne(ctx, in, out)
}

func (h *productTemplateServiceHandler) Import(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateInfo) error {
	return h.ProductTemplateServiceHandler.Import(ctx, in, out)
}

func (h *productTemplateServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateInfo) error {
	return h.ProductTemplateServiceHandler.GetOne(ctx, in, out)
}

func (h *productTemplateServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyProductTemplateList) error {
	return h.ProductTemplateServiceHandler.GetByFilter(ctx, in, out)
}

func (h *productTemplateServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.ProductTemplateServiceHandler.GetStatistic(ctx, in, out)
}

func (h *productTemplateServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.ProductTemplateServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *productTemplateServiceHandler) UpdateBase(ctx context.Context, in *ReqProductTemplateUpdate, out *ReplyInfo) error {
	return h.ProductTemplateServiceHandler.UpdateBase(ctx, in, out)
}

func (h *productTemplateServiceHandler) UpdateGraph(ctx context.Context, in *ReqProductTemplateGraph, out *ReplyInfo) error {
	return h.ProductTemplateServiceHandler.UpdateGraph(ctx, in, out)
}

func (h *productTemplateServiceHandler) UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error {
	return h.ProductTemplateServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *productTemplateServiceHandler) UpdateReferences(ctx context.Context, in *ReqProductTemplateReference, out *ReplyProductTemplateRefs) error {
	return h.ProductTemplateServiceHandler.UpdateReferences(ctx, in, out)
}

func (h *productTemplateServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.ProductTemplateServiceHandler.RemoveOne(ctx, in, out)
}

func (h *productTemplateServiceHandler) GetReleaseOn(ctx context.Context, in *RequestInfo, out *ReplyProductTemplateResult) error {
	return h.ProductTemplateServiceHandler.GetReleaseOn(ctx, in, out)
}

func (h *productTemplateServiceHandler) GetReleaseList(ctx context.Context, in *RequestFilter, out *ReplyProductTemplateResults) error {
	return h.ProductTemplateServiceHandler.GetReleaseList(ctx, in, out)
}
