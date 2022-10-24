package handler

import (
	"context"
	"github.com/zxdstyle/icarus/server/requests"
	"github.com/zxdstyle/icarus/server/responses"
)

type (
	FuncHandler func(ctx context.Context, req requests.Request) responses.Response

	ResourceHandler interface {
		List(ctx context.Context, req requests.Request) responses.Response
		Show(ctx context.Context, req requests.Request) responses.Response
		Create(ctx context.Context, req requests.Request) responses.Response
		Update(ctx context.Context, req requests.Request) responses.Response
		Destroy(ctx context.Context, req requests.Request) responses.Response
	}
)
