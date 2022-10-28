package app

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/icarus/console"
	"github.com/zxdstyle/icarus/container"
)

type Kernel struct {
	Providers map[string]any
	Consoles  []console.Console
}

type Application struct {
	kernel  *Kernel
	rootCmd *cobra.Command
}

func New(kernel *Kernel) *Application {
	return &Application{
		kernel:  kernel,
		rootCmd: &cobra.Command{},
	}
}

func (a *Application) Run() error {
	a.RegisterConsole(a.kernel.Consoles...)

	if a.kernel.Providers != nil && len(a.kernel.Providers) > 0 {
		for name, provider := range a.kernel.Providers {
			container.ProvideNamedValue(providers, name, provider)
		}
	}

	return initConsole().Execute()
}

func (a *Application) RegisterConsole(cmds ...console.Console) {
	for _, cmd := range cmds {
		a.rootCmd.AddCommand(&cobra.Command{
			Use: cmd.Signature(),
			Run: func(c *cobra.Command, args []string) {
				if err := cmd.Handle(); err != nil {
				}
			},
		})
	}
}
