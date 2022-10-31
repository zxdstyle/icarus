package requests

import "github.com/zxdstyle/icarus/server/helper/ua"

type Request interface {
	GetResourceID() uint
	GetRouteParam(field string) string
	GetQuery(key string, def ...string) string
	ScanQueries(pointer any) error
	Validate(pointer any) error
	Bind(pointer any) error
	ScanHeaders(pointer any) error
	IP() string
	UserAgent() *ua.UserAgent
	SetHeader(key, value string)
	GetHeader(key string) string
	Value(key string) any
	Context(key string, value any)
	Method() string
}
