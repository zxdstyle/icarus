package consoles

import (
	"github.com/zxdstyle/icarus/schedulers"
)

type SchedulerProvider struct {
	scheduler schedulers.Scheduler
}

func NewSchedulerProvider(scheduler schedulers.Scheduler) SchedulerProvider {
	return SchedulerProvider{scheduler}
}

func (s SchedulerProvider) Signature() string {
	return "schedule"
}

func (s SchedulerProvider) Description() string {
	return "Start the schedule worker"
}

func (s SchedulerProvider) Handle(args ...string) error {
	s.scheduler.Run()
	return nil
}
