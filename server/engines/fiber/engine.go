package fiber

import (
	json "github.com/json-iterator/go"
	"github.com/zxdstyle/icarus/server/engines"
	"github.com/zxdstyle/icarus/server/handler"
	"github.com/zxdstyle/icarus/server/helper"
	"github.com/zxdstyle/icarus/server/router"
)

type Fiber struct {
	app *fiber.App
}

func NewFiber() (engines.Engine, error) {
	return &Fiber{
		app: fiber.New(fiber.Config{
			EnablePrintRoutes: true,
			JSONDecoder:       json.Unmarshal,
			JSONEncoder:       json.Marshal,
		}),
	}, nil
}

func (f *Fiber) Group(prefix string) router.Router {
	group := f.app.Group(prefix)
	return &Group{
		group,
	}
}

func (f *Fiber) GET(path string, handler handler.FuncHandler) router.Router {
	f.app.Get(path, wrapHandler(handler))
	return f
}

func (f *Fiber) POST(path string, handler handler.FuncHandler) router.Router {
	f.app.Post(path, wrapHandler(handler))
	return f
}

func (f *Fiber) PUT(path string, handler handler.FuncHandler) router.Router {
	f.app.Put(path, wrapHandler(handler))
	return f
}

func (f *Fiber) DELETE(path string, handler handler.FuncHandler) router.Router {
	f.app.Delete(path, wrapHandler(handler))
	return f
}

func (f *Fiber) RESOURCE(resource string, handler handler.ResourceHandler) router.Router {
	base := helper.GetBaseName(resource)
	return f.GET(helper.GetResourceUriIndex(resource, base), handler.List).
		GET(helper.GetResourceUriShow(resource, base), handler.Show).
		POST(helper.GetResourceUriStore(resource, base), handler.Create).
		PUT(helper.GetResourceUriUpdate(resource, base), handler.Update).
		DELETE(helper.GetResourceUriDestroy(resource, base), handler.Destroy)
}

func (f *Fiber) Use(funcHandler fiber.Handler) router.Router {
	f.app.Use(funcHandler)
	return f
}

func (f *Fiber) ListenAndServe(address string) error {
	return f.app.Listen(address)
}

func (f *Fiber) Shutdown() error {
	return f.app.Shutdown()
}

func wrapHandler(handler handler.FuncHandler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		resp := handler(ctx.Context(), newRequest(ctx))
		if resp == nil {
			return nil
		}
		ctx.Status(resp.StatusCode())
		return resp.Response(ctx)
	}
}
