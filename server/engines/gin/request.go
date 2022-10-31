package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zxdstyle/icarus/server/helper/ua"
	"github.com/zxdstyle/icarus/validator"
	"sync"
)

const (
	UserAgentKey = "User-Agent"

	DeviceMobile  = "Mobile"
	DeviceDesktop = "Desktop"
	DeviceOther   = "Other"
)

type request struct {
	ctx   *gin.Context
	agent *ua.UserAgent
	once  sync.Once
}

func newRequest(ctx *gin.Context) *request {
	return &request{
		ctx: ctx,
	}
}

func (r *request) GetResourceID() (id uint) {

	return cast.ToUint(r.ctx.Param("id"))
}

func (r *request) GetRouteParam(field string) string {
	return r.ctx.Param(field)
}

func (r *request) GetQuery(key string, def ...string) string {
	val := r.ctx.Query(key)
	if len(val) > 0 {
		return val
	}
	if len(def) > 0 {
		return def[0]
	}
	return val
}

func (r *request) ScanQueries(pointer any) error {
	if err := r.ctx.BindQuery(pointer); err != nil {
		return err
	}
	if err := validator.Validate(pointer); err != nil {
		return errors.New(validator.Translate(err))
	}
	return nil
}

func (r *request) Bind(pointer any) error {
	return r.ctx.BindJSON(pointer)
}

func (r *request) Validate(pointer any) error {
	if err := r.Bind(pointer); err != nil {
		return err
	}
	if err := validator.Validate(pointer); err != nil {
		return errors.New(validator.Translate(err))
	}
	return nil
}

func (r *request) ScanHeaders(pointer any) error {
	return r.ctx.BindHeader(pointer)
}

func (r *request) IP() string {
	return r.ctx.RemoteIP()
}

func (r *request) UserAgent() *ua.UserAgent {
	r.once.Do(func() {
		r.agent = ua.New(r.ctx.GetHeader(UserAgentKey))
	})

	return r.agent
}

func (r *request) Next() error {
	r.ctx.Next()
	return nil
}

func (r *request) SetHeader(key, value string) {
	r.ctx.Header(key, value)
}

func (r *request) GetHeader(key string) string {
	return r.ctx.GetHeader(key)
}

func (r *request) Value(key string) any {
	return r.ctx.Value(key)
}

func (r *request) Context(key string, value any) {
	r.ctx.Set(key, value)
}

func (r *request) Method() string {
	return r.ctx.Request.Method
}
