package responses

import (
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

func (r *ApiResponse) Content() any {
	return r
}

func Success(data any) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}

func Failed(msg string) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
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

func Internal(msg string) *ApiResponse {
	return &ApiResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Data:    nil,
	}
}

func NotFound(msg ...string) *ApiResponse {
	message := http.StatusText(http.StatusNotFound)
	if len(msg) > 0 {
		message = msg[1]
	}

	return &ApiResponse{
		Status:  http.StatusNotFound,
		Message: message,
		Data:    nil,
	}
}

func Custom(statusCode int, msg string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Status:  statusCode,
		Message: msg,
		Data:    data,
	}
}
