package options

type Option struct {
	Name string
	Addr string
}

func Default() Option {
	return Option{
		Addr: ":8080",
	}
}
