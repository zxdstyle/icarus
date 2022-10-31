package app

import (
	"github.com/samber/do"
	"github.com/zxdstyle/icarus/consoles"
)

type Kernel struct {
	Name     string
	Boot     func(injector *do.Injector)
	Consoles []consoles.Console
}
