package cronv3

import (
	"slices"
	"time"

	"github.com/robfig/cron/v3"
)

// NewSpecificTimes 定点任务，在指定的几个时间点执行。
//
// 例如仅在：
//
//	2025-01-01 00:00:00
//	2025-03-01 00:00:00
//	2025-03-15 00:00:00
//
// 执行三次就完事的任务。
func NewSpecificTimes(times []time.Time) cron.Schedule {
	slices.SortFunc(times, func(a, b time.Time) int {
		return int(a.Sub(b))
	})

	return &specificTimes{
		times: times,
	}
}

type specificTimes struct {
	times []time.Time
}

func (st *specificTimes) Next(now time.Time) time.Time {
	for idx, at := range st.times {
		if at.After(now) {
			st.times = st.times[idx:]
			return at
		}
	}

	return time.Time{}
}

// NewPeriodicallyTimes 定期执行。
// 不同于 cron 表达式的是，cron/v3 时间粒度（秒级/分钟）是可配置的，
// 不同的级别对应的 cron 表达式位数不一样。
// 此方式可以自适应不同级别的 cron/v3，弊端就是最小间隔为一分钟。
func NewPeriodicallyTimes(du time.Duration) cron.Schedule {
	if du < time.Minute {
		du = time.Minute
	}

	return &periodicallyTimes{
		du: du,
	}
}

type periodicallyTimes struct {
	du time.Duration
}

func (pt *periodicallyTimes) Next(now time.Time) time.Time {
	return now.Add(pt.du)
}
