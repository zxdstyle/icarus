package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zxdstyle/icarus/server/engines"
	"github.com/zxdstyle/icarus/server/handler"
	"github.com/zxdstyle/icarus/server/helper"
	"github.com/zxdstyle/icarus/server/middlewares"
	"github.com/zxdstyle/icarus/server/responses"
	"github.com/zxdstyle/icarus/server/router"
	"net/http"
	"time"
)

type Gin struct {
	app   *gin.Engine
	serve *http.Server
}

func New() (engines.Engine, error) {
	return &Gin{
		app: gin.Default(),
	}, nil
}

func (g *Gin) GET(path string, handler handler.FuncHandler) router.Router {
	g.app.GET(path, wrapHandler(handler))
	return g
}

func (g *Gin) POST(path string, handler handler.FuncHandler) router.Router {
	g.app.POST(path, wrapHandler(handler))
	return g
}

func (g *Gin) PUT(path string, handler handler.FuncHandler) router.Router {
	g.app.PUT(path, wrapHandler(handler))
	return g
}

func (g *Gin) DELETE(path string, handler handler.FuncHandler) router.Router {
	g.app.DELETE(path, wrapHandler(handler))
	return g
}

func (g *Gin) RESOURCE(resource string, handler handler.ResourceHandler) router.Router {
	base := helper.GetBaseName(resource)
	return g.GET(helper.GetResourceUriIndex(resource, base), handler.List).
		GET(helper.GetResourceUriShow(resource, base), handler.Show).
		POST(helper.GetResourceUriStore(resource, base), handler.Create).
		PUT(helper.GetResourceUriUpdate(resource, base), handler.Update).
		DELETE(helper.GetResourceUriDestroy(resource, base), handler.Destroy)
}

func (g *Gin) Use(handler middlewares.FuncMiddleware) router.Router {
	g.app.Use(wrapMiddleware(handler))
	return g
}

func (g *Gin) Group(prefix string) router.Router {
	group := g.app.Group(prefix)
	return &Group{
		group,
	}
}

func (g *Gin) ListenAndServe(address string) error {
	srv := &http.Server{
		Addr:    address,
		Handler: g.app,
	}
	g.serve = srv
	return srv.ListenAndServe()
}

func (g *Gin) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return g.serve.Shutdown(ctx)
}

func wrapMiddleware(middleware middlewares.FuncMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := middleware(context.TODO(), newRequest(c))
		parseResponse(c, resp)
	}
}

func wrapHandler(handler handler.FuncHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.TODO()
		resp := handler(ctx, newRequest(c))
		if resp == nil {
			return
		}

		parseResponse(c, resp)
		return
	}
}

func parseResponse(ctx *gin.Context, resp responses.Response) {
	if resp == nil {
		return
	}

	ctx.Status(resp.StatusCode())
	switch resp.(type) {
	case *responses.ApiResponse:
		ctx.JSON(resp.StatusCode(), resp.Content())
	case *responses.RedirectResp:
		ctx.Redirect(resp.StatusCode(), resp.Content().(string))
	}
}
