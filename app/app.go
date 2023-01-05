package app

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/icarus/consoles"
	"github.com/zxdstyle/icarus/container"
	"github.com/zxdstyle/icarus/schedulers"
	"log"
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
	// 注册自定义命令
	if a.kernel.Consoles != nil {
		a.RegisterConsole(a.kernel.Consoles...)
	}

	// 注册scheduler驱动
	container.ProvideValue[schedulers.Scheduler](providers, schedulers.NewCron())

	a.RegisterConsole(
		consoles.HttpProvider{},
		consoles.NewSchedulerProvider(Make[schedulers.Scheduler]()),
		consoles.NewMigrateProvider(a.kernel.Models),
	)

	if a.kernel.Boot != nil {
		a.kernel.Boot(providers)
	}

	if a.kernel.Schedule != nil {
		cron := Make[schedulers.Scheduler]()
		a.kernel.Schedule(cron)
	}

	return a.rootCmd.Execute()
}

func (a *Application) RegisterConsole(commands ...consoles.Console) {
	for idx, _ := range commands {
		a.rootCmd.AddCommand(a.transferConsole(commands[idx]))
	}
}

func (a *Application) transferConsole(cmd consoles.Console) *cobra.Command {
	return &cobra.Command{
		Use:   cmd.Signature(),
		Short: cmd.Description(),
		Run: func(c *cobra.Command, args []string) {
			if err := cmd.Handle(args...); err != nil {
				log.Fatalln(err)
			}
		},
	}
}
