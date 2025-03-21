package cronv3

import (
	"log/slog"

	"github.com/robfig/cron/v3"
)

func NewLog(log *slog.Logger) cron.Logger {
	return &cronLog{log: log}
}

type cronLog struct {
	log *slog.Logger
}

func (c *cronLog) Info(msg string, keysAndValues ...any) {
	c.log.Info(msg, keysAndValues...)
}

func (c *cronLog) Error(err error, msg string, keysAndValues ...any) {
	c.log.Error(msg, append([]any{"error", err}, keysAndValues...)...)
}
