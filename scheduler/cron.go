package scheduler

import "github.com/robfig/cron/v3"

type Cron struct {
	instance *cron.Cron
}

func NewCron() Scheduler {
	return &Cron{
		instance: cron.New(cron.WithSeconds()),
	}
}

func (c *Cron) Job(job Job) Spec {
	c.instance.AddJob()
}
