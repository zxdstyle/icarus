package app

import (
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/icarus/app/consoles"
	"github.com/zxdstyle/icarus/console"
	"log"
)

type (
	Kernel struct {
		Name      string
		Providers map[string]any
		Consoles  []console.Console
		Boot      func(providers *do.Injector)
	}
	Application struct {
		kernel  *Kernel
		rootCmd *cobra.Command
	}
)

func New(kernel *Kernel) *Application {
	return &Application{
		kernel: kernel,
		rootCmd: &cobra.Command{
			Use: kernel.Name,
			Run: func(cmd *cobra.Command, args []string) {
				if err := cmd.Help(); err != nil {
					log.Fatalln(err)
				}
			},
		},
	}
}

func (a *Application) Run() error {
	a.RegisterConsole(consoles.HttpProvider{})
	a.RegisterConsole(a.kernel.Consoles...)

	a.kernel.Boot(providers)

	return a.rootCmd.Execute()
}

func (a *Application) RegisterConsole(commands ...console.Console) {
	for _, cmd := range commands {
		a.rootCmd.AddCommand(a.transferCmd(cmd))
	}
}

func (a *Application) transferCmd(cmd console.Console) *cobra.Command {
	return &cobra.Command{
		Use: cmd.Signature(),
		Run: func(c *cobra.Command, args []string) {
			if err := cmd.Handle(); err != nil {
			}
		},
	}
}
