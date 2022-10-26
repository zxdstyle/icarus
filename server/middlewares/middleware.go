package middlewares

import (
	"context"
	"github.com/zxdstyle/icarus/server/requests"
	"github.com/zxdstyle/icarus/server/responses"
)

type (
	Request interface {
		requests.Request
		Next() error
	}

	FuncMiddleware func(ctx context.Context, req Request) responses.Response
)
