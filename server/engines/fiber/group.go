package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxdstyle/liey/pkg/server/handler"
	"github.com/zxdstyle/liey/pkg/server/helper"
	"github.com/zxdstyle/liey/pkg/server/router"
)

type Group struct {
	group fiber.Router
}

func (g Group) Group(prefix string) router.Router {
	return &Group{group: g.group.Group(prefix)}
}

func (g Group) GET(path string, handler handler.FuncHandler) router.Router {
	g.group.Get(path, wrapHandler(handler))
	return g
}

func (g Group) POST(path string, handler handler.FuncHandler) router.Router {
	g.group.Post(path, wrapHandler(handler))
	return g
}

func (g Group) PUT(path string, handler handler.FuncHandler) router.Router {
	g.group.Put(path, wrapHandler(handler))
	return g
}

func (g Group) DELETE(path string, handler handler.FuncHandler) router.Router {
	g.group.Delete(path, wrapHandler(handler))
	return g
}

func (g Group) RESOURCE(resource string, handler handler.ResourceHandler) router.Router {
	base := helper.GetBaseName(resource)
	return g.GET(helper.GetResourceUriIndex(resource, base), handler.List).
		GET(helper.GetResourceUriShow(resource, base), handler.Show).
		POST(helper.GetResourceUriStore(resource, base), handler.Create).
		PUT(helper.GetResourceUriUpdate(resource, base), handler.Update).
		DELETE(helper.GetResourceUriDestroy(resource, base), handler.Destroy)
}

func (g Group) Use(handler2 fiber.Handler) router.Router {
	g.group.Use(handler2)
	return g
}
