package app

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/icarus"
	"os"
	"os/signal"
	"syscall"
)

func initConsole() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			startHttpServer()
		},
	}
}

func startHttpServer() {
	go icarus.Server().Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc

	icarus.Server().Shutdown()

	icarus.Shutdown()
}
