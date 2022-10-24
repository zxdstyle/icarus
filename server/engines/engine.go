package engines

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxdstyle/liey/pkg/server/router"
)

type Engine interface {
	router.Router
	Use(handler fiber.Handler) router.Router
	Group(prefix string) router.Router
	ListenAndServe(address string) error
	Shutdown() error
}
