package specs

import "time"

type (
	Spec interface {
		// EverySecond 每隔{seconds}秒执行，默认1秒
		EverySecond(seconds ...int)
		// EveryMinute 每隔{minutes}分钟执行，默认1分钟
		EveryMinute(minutes ...int)
		// Hourly 每{hour}小时，每隔一个小时
		Hourly(hour ...int)
		// HourlyAt 每小时的第{minute}分钟
		HourlyAt(minute int)
		// Daily 每天
		Daily(day ...int)
		// DailyAt 每天的第{hour}小时第{minute}分钟
		DailyAt(hour, minute int)
		// Weekly 每周日 00:00 执行一次任务
		Weekly()
		// WeeklyOn 每周{week} {time} 执行一次任务 WeeklyOn(time.Sunday, 13 , 0 )
		WeeklyOn(week time.Weekday, hour, minute int)
		// Monthly 每月第一天 00:00 执行一次任务
		Monthly()
		// MonthlyOn 每月第{day}天 {time} MonthlyOn WeeklyOn(4, "13:00")
		MonthlyOn(day, hour, minute int)
		// Cron 自定义 Cron 计划执行任务
		Cron(spec string)
	}
)
