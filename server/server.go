package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxdstyle/icarus/server/engines"
	"github.com/zxdstyle/icarus/server/handler"
	"github.com/zxdstyle/icarus/server/options"
	"github.com/zxdstyle/icarus/server/router"
)

type Server struct {
	engine engines.Engine
	option options.Option
}

func New(engine engines.Engine) *Server {
	return &Server{
		engine: engine,
		option: options.Default(),
	}
}

func (s *Server) WithOption(opt options.Option) {
	s.option = opt
}

func (s *Server) Group(prefix string) router.Router {
	return s.engine.Group(prefix)
}

func (s *Server) GET(path string, handler handler.FuncHandler) router.Router {
	return s.engine.GET(path, handler)
}

func (s *Server) POST(path string, handler handler.FuncHandler) router.Router {
	return s.engine.POST(path, handler)
}

func (s *Server) PUT(path string, handler handler.FuncHandler) router.Router {
	return s.engine.PUT(path, handler)
}

func (s *Server) DELETE(path string, handler handler.FuncHandler) router.Router {
	return s.engine.DELETE(path, handler)
}

func (s *Server) RESOURCE(path string, handler handler.ResourceHandler) router.Router {
	return s.engine.RESOURCE(path, handler)
}

func (s *Server) Use(funcHandler fiber.Handler) router.Router {
	return s.engine.Use(funcHandler)
}

func (s *Server) Run() error {
	return s.engine.ListenAndServe(s.option.Addr)
}

func (s *Server) Shutdown() error {
	return s.engine.Shutdown()
}
