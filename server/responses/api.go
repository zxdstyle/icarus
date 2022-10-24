package responses

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	ApiResponse struct {
		Status  int    `json:"status"`
		Meta    *Meta  `json:"meta,omitempty"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	PaginationMeta struct {
		CurrentPage int   `json:"current_page"`
		PerPage     int   `json:"per_page"`
		Total       int64 `json:"total"`
	}

	Meta struct {
		Pagination PaginationMeta `json:"pagination"`
	}
)

func (r *ApiResponse) SetTotal(total int64) {
	if r.Meta == nil {
		r.Meta = &Meta{Pagination: PaginationMeta{
			CurrentPage: 1,
			PerPage:     20,
			Total:       total,
		}}
	}
	r.Meta.Pagination.Total = total
}

func (r *ApiResponse) StatusCode() int {
	return r.Status
}

func (r *ApiResponse) Response(ctx *fiber.Ctx) error {
	return ctx.JSON(r)
}

func Success(data any) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}

func Failed(err string) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusBadRequest,
		Message: err,
		Data:    nil,
	}
}

func Error(err error) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
		Data:    nil,
	}
}

func Unauthorized(err error) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusUnauthorized,
		Message: err.Error(),
		Data:    nil,
	}
}

func Internal(err string) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusInternalServerError,
		Message: err,
		Data:    nil,
	}
}
