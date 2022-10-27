package app

import "github.com/zxdstyle/icarus/container"

type Kernel struct {
	Providers map[string]any
}

type Application struct {
	kernel *Kernel
}

func New(kernel *Kernel) *Application {
	return &Application{
		kernel: kernel,
	}
}

func (a *Application) Run() error {
	if a.kernel.Providers != nil && len(a.kernel.Providers) > 0 {
		for name, provider := range a.kernel.Providers {
			container.ProvideNamedValue(providers, name, provider)
		}
	}

	return initConsole().Execute()
}
