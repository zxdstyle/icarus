package fiber

import (
	"errors"
	"github.com/gofiber/fiber/v2"
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
	ctx   *fiber.Ctx
	agent *ua.UserAgent
	once  sync.Once
}

func newRequest(ctx *fiber.Ctx) *request {
	return &request{
		ctx: ctx,
	}
}

func (r *request) GetResourceID() (id uint) {
	params := r.ctx.Route().Params
	if len(params) == 0 {
		return
	}

	return cast.ToUint(r.ctx.Params(params[0]))
}

func (r *request) GetRouteParam(field string) string {
	return r.ctx.Params(field)
}

func (r *request) GetQuery(key string, def ...string) string {
	return r.ctx.Query(key, def...)
}

func (r *request) ScanQueries(pointer any) error {
	if err := r.ctx.QueryParser(pointer); err != nil {
		return err
	}
	if err := validator.Validate(pointer); err != nil {
		return errors.New(validator.Translate(err))
	}
	return nil
}

func (r *request) Bind(pointer any) error {
	return r.ctx.BodyParser(pointer)
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
	return r.ctx.ReqHeaderParser(pointer)
}

func (r *request) IP() string {
	return r.ctx.IP()
}

func (r *request) UserAgent() *ua.UserAgent {
	r.once.Do(func() {
		headers := r.ctx.GetReqHeaders()
		r.agent = ua.New(headers[UserAgentKey])
	})

	return r.agent
}

func (r *request) Next() error {
	return r.ctx.Next()
}

func (r *request) SetHeader(key, value string) {
	r.ctx.Set(key, value)
}

func (r *request) GetHeader(key string) string {
	headers := r.ctx.GetReqHeaders()
	val, ok := headers[key]
	if ok {
		return val
	}
	return ""
}

func (r *request) Value(key string) any {
	return r.ctx.Context().Value(key)
}

func (r *request) Context(key string, value any) {
	r.ctx.Context().SetUserValue(key, value)
}

func (r *request) Method() string {
	return r.ctx.Method()
}
