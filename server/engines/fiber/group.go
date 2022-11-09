package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxdstyle/icarus/server/handler"
	"github.com/zxdstyle/icarus/server/helper"
	"github.com/zxdstyle/icarus/server/middlewares"
	"github.com/zxdstyle/icarus/server/router"
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

func (g Group) Use(handler middlewares.FuncMiddleware) router.Router {
	g.group.Use(wrapMiddleware(handler))
	return g
}

func (g Group) Static(prefix, root string) router.Router {
	g.group.Static(prefix, root)
	return g
}
