package responses

import (
	"net/http"
)

type RedirectResp struct {
	Status int
	Path   string
}

func (r *RedirectResp) StatusCode() int {
	return r.Status
}

func (r *RedirectResp) Content() any {
	return r.Path
}

func Redirect(path string, status int) Response {
	return &RedirectResp{
		Status: http.StatusFound,
		Path:   path,
	}
}
