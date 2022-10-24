package fiber

import (
	"errors"
	"github.com/spf13/cast"
	"github.com/zxdstyle/icarus/server/helper/ua"
	"github.com/zxdstyle/icarus/server/requests"
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

func newRequest(ctx *fiber.Ctx) requests.Request {
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
	return r.ctx.QueryParser(pointer)
}

func (r *request) Bind(pointer any) error {
	return r.ctx.BodyParser(pointer)
}

func (r *request) Validate(pointer any) error {
	if err := r.Bind(pointer); err != nil {
		return err
	}
	if err := validate.Struct(pointer); err != nil {
		errs := err.(validator.ValidationErrors)
		return errors.New(errs[0].Translate(trans))
	}
	return nil
}

func (r *request) Redirect(path string, status int) {
	_ = r.ctx.Redirect(path, status)
}

func (r *request) Headers() map[string]string {
	return r.ctx.GetReqHeaders()
}

func (r *request) ScanHeaders(pointer any) error {
	return r.ctx.ReqHeaderParser(pointer)
}

func (r *request) IP() string {
	return r.ctx.IP()
}

func (r *request) UserAgent() *ua.UserAgent {
	r.once.Do(func() {
		headers := r.Headers()
		r.agent = ua.New(headers[UserAgentKey])
	})

	return r.agent
}
