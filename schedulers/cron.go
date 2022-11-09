package schedulers

import (
	"context"
	"fmt"
	"github.com/gookit/slog"
	"github.com/robfig/cron/v3"
	"github.com/zxdstyle/icarus/schedulers/jobs"
	"github.com/zxdstyle/icarus/schedulers/specs"
	"log"
)

type Cron struct {
	instance *cron.Cron
	entries  map[cron.EntryID]jobs.Job
}

func NewCron() *Cron {
	return &Cron{
		instance: cron.New(cron.WithSeconds()),
		entries:  make(map[cron.EntryID]jobs.Job),
	}
}

func (c *Cron) Job(job jobs.Job) specs.Spec {
	return specs.NewCronSpec(job, c.registerJob)
}

func (c *Cron) Run() {
	if len(c.entries) == 0 {
		slog.Fatalf("no jobs to schedule")
	}

	fmt.Println("Start the schedule worker...")
	c.instance.Run()
}

func (c *Cron) Entries() []cron.Entry {
	return c.instance.Entries()
}

func (c *Cron) registerJob(format string, job jobs.Job) {
	eid, err := c.instance.AddFunc(format, func() {
		ctx := context.TODO()
		job.Run(ctx)
	})
	if err != nil {
		log.Fatalln(err)
	}
	c.entries[eid] = job
}
