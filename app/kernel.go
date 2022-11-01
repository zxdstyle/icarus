package app

import (
	"github.com/samber/do"
	"github.com/zxdstyle/icarus/consoles"
	"github.com/zxdstyle/icarus/scheduler"
)

type Kernel struct {
	Name     string
	Boot     func(injector *do.Injector)
	Consoles []consoles.Console
	Schedule func(scheduler scheduler.Scheduler)
}
