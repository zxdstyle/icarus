package schedulers

import (
	"github.com/robfig/cron/v3"
	"github.com/zxdstyle/icarus/schedulers/jobs"
	"github.com/zxdstyle/icarus/schedulers/specs"
)

type (
	ScheduleFunc func(spec specs.Spec) specs.Spec

	Scheduler interface {
		Job(job jobs.Job) specs.Spec
		Entries() []cron.Entry
		Run()
	}
)
