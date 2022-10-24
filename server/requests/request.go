package requests

import "github.com/zxdstyle/liey/pkg/server/helper/ua"

type Request interface {
	GetResourceID() uint
	GetRouteParam(field string) string
	GetQuery(key string, def ...string) string
	ScanQueries(pointer any) error
	Validate(pointer any) error
	Bind(pointer any) error
	Redirect(path string, status int)
	Headers() map[string]string
	ScanHeaders(pointer any) error
	IP() string
	UserAgent() *ua.UserAgent
}
