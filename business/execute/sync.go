package execute

import (
	"context"
	"log/slog"
	"time"

	"github.com/dfcfw/rock-migrate/library/cronv3"
	"github.com/robfig/cron/v3"
)

type Executor interface {
	Add(parent context.Context, tasks ...Tasker)
}

type TaskFunc func(context.Context) error

type TaskInfo struct {
	Name string        // 任务名
	Func TaskFunc      // 调用函数
	Cron time.Duration // 如果是定时任务，则定时任务间隔
}

// Tasker 同步任务
type Tasker interface {
	// Task 获取任务信息
	Task() TaskInfo
}

func New(crontab *cronv3.Crontab, log *slog.Logger) Executor {
	return &executor{
		crontab: crontab,
		log:     log,
	}
}

type executor struct {
	crontab *cronv3.Crontab
	log     *slog.Logger
}

func (wrk *executor) Add(parent context.Context, tasks ...Tasker) {
	for _, task := range tasks {
		info := task.Task()
		wrk.add(parent, info)
	}
}

func (wrk *executor) add(parent context.Context, task TaskInfo) {
	name, fn := task.Name, task.Func
	exec := wrk.wrapFunc(parent, name, fn)
	if du := task.Cron; du <= 0 {
		go exec()
	} else {
		spec := cronv3.NewPeriodicallyTimes(du)
		wrk.crontab.Schedule(name, spec, cron.FuncJob(exec))
	}
}

func (wrk *executor) wrapFunc(parent context.Context, name string, fn TaskFunc) func() {
	return func() {
		attrs := []any{slog.String("name", name)}
		defer func() {
			if cause := recover(); cause != nil {
				attrs = append(attrs, slog.Any("cause", cause))
				wrk.log.Error("任务执行 panic", attrs...)
			}
		}()

		if err := fn(parent); err != nil {
			attrs = append(attrs, slog.Any("error", err))
			wrk.log.Warn("任务执行出错", attrs...)
		} else {
			wrk.log.Debug("任务执行结束", attrs...)
		}
	}
}
