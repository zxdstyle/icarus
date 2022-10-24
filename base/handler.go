package base

import (
	"context"
	"github.com/zxdstyle/liey/pkg/server/requests"
	"github.com/zxdstyle/liey/pkg/server/responses"
)

type (
	Handler[V logicModel] struct {
		logic Logic[V]
	}
)

func NewHandler[V logicModel](l Logic[V]) Handler[V] {
	return Handler[V]{
		logic: l,
	}
}

func (h Handler[V]) List(ctx context.Context, req requests.Request) responses.Response {
	var data []V
	resp := responses.Success(&data)
	if err := h.logic.List(ctx, req, resp); err != nil {
		return responses.Error(err)
	}
	return resp
}

func (h Handler[V]) Show(ctx context.Context, req requests.Request) responses.Response {
	primaryKey := req.GetResourceID()
	var mo V
	if err := h.logic.Show(ctx, primaryKey, req, &mo); err != nil {
		return responses.Error(err)
	}
	return responses.Success(mo)
}

func (h Handler[V]) Create(ctx context.Context, req requests.Request) responses.Response {
	var data V
	if err := req.Validate(&data); err != nil {
		return responses.Error(err)
	}
	if err := h.logic.Create(ctx, &data); err != nil {
		return responses.Error(err)
	}
	return responses.Success(data)
}

func (h Handler[V]) Update(ctx context.Context, req requests.Request) responses.Response {
	panic("implement me")
}

func (h Handler[V]) Destroy(ctx context.Context, req requests.Request) responses.Response {
	panic("implement me")
}
