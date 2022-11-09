package consoles

import (
	"github.com/zxdstyle/icarus"
	"os"
	"os/signal"
	"syscall"
)

type HttpProvider struct {
}

func (p HttpProvider) Signature() string {
	return "serve"
}

func (p HttpProvider) Description() string {
	return "Start http server"
}

func (p HttpProvider) Handle() error {
	go icarus.Server().Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc

	icarus.Server().Shutdown()

	icarus.Shutdown()

	return nil
}
