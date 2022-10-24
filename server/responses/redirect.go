package responses

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type redirect struct {
	Status int
	Path   string
}

func (r *redirect) StatusCode() int {
	return r.Status
}

func (r *redirect) Response(ctx *fiber.Ctx) error {
	return ctx.Redirect(r.Path, r.Status)
}

func Redirect(path string, status int) Response {
	return &redirect{
		Status: http.StatusFound,
		Path:   path,
	}
}
