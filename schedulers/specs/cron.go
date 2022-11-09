package specs

import (
	"fmt"
	"github.com/zxdstyle/icarus/schedulers/jobs"
	"time"
)

type (
	registerFn func(format string, spec jobs.Job)

	CronSpec struct {
		job        jobs.Job
		format     string
		registerFn registerFn
	}
)

func NewCronSpec(job jobs.Job, registerFn registerFn) Spec {
	return &CronSpec{job: job, registerFn: registerFn}
}

func (c *CronSpec) EverySecond(seconds ...int) {
	second := 1
	if len(seconds) > 0 {
		second = seconds[0]
	}
	c.registerFn(fmt.Sprintf("*/%d * * * * ?", second), c.job)
}

func (c *CronSpec) EveryMinute(minutes ...int) {
	minute := 1
	if len(minutes) > 0 {
		minute = minutes[0]
	}
	c.registerFn(fmt.Sprintf("0 */%d * * * ?", minute), c.job)
}

func (c *CronSpec) Hourly(hours ...int) {
	hour := 1
	if len(hours) > 0 {
		hour = hours[0]
	}
	c.registerFn(fmt.Sprintf("0 0 */%d * * ?", hour), c.job)
}

func (c *CronSpec) HourlyAt(minute int) {
	c.registerFn(fmt.Sprintf("0 %d * * * ?", minute), c.job)
}

func (c *CronSpec) Daily(day ...int) {
	d := 1
	if len(day) > 0 {
		d = day[0]
	}
	c.registerFn(fmt.Sprintf("0 0 0 */%d * ?", d), c.job)
}

func (c *CronSpec) DailyAt(hour, minute int) {
	c.registerFn(fmt.Sprintf("0 %d %d * * ?", minute, hour), c.job)
}

func (c *CronSpec) Weekly() {
	c.Daily(7)
}

func (c *CronSpec) WeeklyOn(week time.Weekday, hour, minute int) {
	c.registerFn(fmt.Sprintf("0 %d %d ? * %d", minute, hour, week), c.job)
}

func (c *CronSpec) Monthly() {
	c.registerFn("0 0 0 1 * ?", c.job)
}

func (c *CronSpec) MonthlyOn(day, hour, minute int) {
	c.registerFn(fmt.Sprintf("0 %d %d %d * ?", minute, hour, day), c.job)
}

func (c *CronSpec) Cron(spec string) {
	c.registerFn(spec, c.job)
}
