package middlewares

import (
	"context"
	"github.com/zxdstyle/icarus/server/responses"
	"net/http"
)

func Cors() FuncMiddleware {
	return func(ctx context.Context, req Request) responses.Response {
		req.SetHeader("Access-Control-Allow-Origin", "*")
		req.SetHeader("Access-Control-Allow-Headers", "*")

		if req.Method() == http.MethodOptions {
			return responses.Empty()

		}

		if err := req.Next(); err != nil {
			return responses.Error(err)
		}
		return nil
	}
}
