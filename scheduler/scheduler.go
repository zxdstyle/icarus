package scheduler

import (
	"context"
)

type (
	ScheduleFunc func(spec Spec) Spec

	Scheduler interface {
		Job(job Job) Spec
	}

	Job interface {
		// Run 返回错误则会重新调度
		Run(ctx context.Context) error
	}
)
