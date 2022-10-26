package middlewares

import (
	"context"
	"github.com/zxdstyle/icarus/server/responses"
)

func Cors() FuncMiddleware {
	return func(ctx context.Context, req Request) responses.Response {

		if err := req.Next(); err != nil {
			return responses.Error(err)
		}
		return nil
	}
}
