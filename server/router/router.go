package router

import (
	"github.com/zxdstyle/icarus/server/handler"
	"github.com/zxdstyle/icarus/server/middlewares"
)

type Router interface {
	Group(prefix string) Router
	Use(handler middlewares.FuncMiddleware) Router
	GET(path string, handler handler.FuncHandler) Router
	POST(path string, handler handler.FuncHandler) Router
	PUT(path string, handler handler.FuncHandler) Router
	DELETE(path string, handler handler.FuncHandler) Router
	RESOURCE(resource string, handler handler.ResourceHandler) Router
	Static(prefix, root string) Router
}
