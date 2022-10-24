package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxdstyle/liey/pkg/server/handler"
)

type Router interface {
	Group(prefix string) Router
	Use(fiber.Handler) Router
	GET(path string, handler handler.FuncHandler) Router
	POST(path string, handler handler.FuncHandler) Router
	PUT(path string, handler handler.FuncHandler) Router
	DELETE(path string, handler handler.FuncHandler) Router
	RESOURCE(resource string, handler handler.ResourceHandler) Router
}
