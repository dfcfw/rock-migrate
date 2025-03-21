package cronv3

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func New(log *slog.Logger, opts ...cron.Option) *Crontab {
	return &Crontab{
		log:   log,
		crond: cron.New(opts...),
		names: make(map[string]cron.EntryID, 16),
	}
}

type Crontab struct {
	log   *slog.Logger
	crond *cron.Cron
	mutex sync.Mutex
	names map[string]cron.EntryID
}

func (c *Crontab) Start() {
	c.crond.Start()
}

func (c *Crontab) Run() {
	c.crond.Run()
}

func (c *Crontab) Stop() context.Context {
	return c.crond.Stop()
}

func (c *Crontab) Location() *time.Location {
	return c.crond.Location()
}

func (c *Crontab) AddFunc(name, spec string, cmd func()) error {
	return c.AddJob(name, spec, cron.FuncJob(cmd))
}

func (c *Crontab) AddJob(name string, spec string, job cron.Job) error {
	c.log.Info("添加或修改定时任务", slog.String("name", name))
	wrapJob := c.wrapperJob(name, job.Run)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if id, ok := c.names[name]; ok {
		c.crond.Remove(id)
	}
	id, err := c.crond.AddJob(spec, wrapJob)
	if err == nil {
		c.names[name] = id
	}

	return err
}

func (c *Crontab) Schedule(name string, sch cron.Schedule, job func()) {
	c.log.Info("添加或修改定时任务", slog.String("name", name))
	wrapJob := c.wrapperJob(name, job)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if id, ok := c.names[name]; ok {
		c.crond.Remove(id)
	}
	id := c.crond.Schedule(sch, wrapJob)
	c.names[name] = id
}

func (c *Crontab) Remove(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if id, ok := c.names[name]; ok {
		c.crond.Remove(id)
		delete(c.names, name)
		c.log.Warn("删除定时任务", slog.String("name", name))
	}
}

// Cleanup 清理哪些不再执行的定时任务，该功能主要针对 [NewTimingSchedule] 类型的定时任务。
func (c *Crontab) Cleanup() {
	for _, ent := range c.crond.Entries() {
		if !ent.Next.IsZero() {
			continue
		}
		if name := c.resolveName(ent.ID); name != "" {
			c.Remove(name)
		}
	}
}

func (c *Crontab) resolveName(id cron.EntryID) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for name, eid := range c.names {
		if id == eid {
			return name
		}
	}

	return ""
}

func (c *Crontab) wrapperJob(name string, job func()) cron.Job {
	return cron.FuncJob(func() {
		c.log.Info("定时任务开始执行", slog.String("name", name))
		job()
		c.log.Info("定时任务执行结束", slog.String("name", name))
	})
}
