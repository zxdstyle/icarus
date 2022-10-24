package responses

import "github.com/gofiber/fiber/v2"

type (
	Response interface {
		StatusCode() int
		Response(ctx *fiber.Ctx) error
	}
)
