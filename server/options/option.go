package options

import "github.com/gookit/config/v2"

type Option struct {
	Name string
	Addr string
}

func Default() Option {
	return Option{
		Addr: config.String("app.addr", ":8080"),
	}
}
