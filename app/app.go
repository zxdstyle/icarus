package app

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/icarus/consoles"
)

type Application struct {
	kernel  *Kernel
	rootCmd *cobra.Command
}

func New(kernel *Kernel) *Application {
	return &Application{
		kernel: kernel,
		rootCmd: &cobra.Command{
			Use: kernel.Name,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		},
	}
}

func (a *Application) Run() error {
	a.RegisterConsole(a.kernel.Consoles...)

	a.RegisterConsole(consoles.HttpProvider{})

	a.kernel.Boot(providers)

	return initConsole().Execute()
}

func (a *Application) RegisterConsole(cmds ...consoles.Console) {
	for _, cmd := range cmds {
		a.rootCmd.AddCommand(a.transferConsole(cmd))
	}
}

func (a *Application) transferConsole(cmd consoles.Console) *cobra.Command {
	return &cobra.Command{
		Use: cmd.Signature(),
		Run: func(c *cobra.Command, args []string) {
			if err := cmd.Handle(); err != nil {
			}
		},
	}
}
