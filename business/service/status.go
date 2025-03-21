package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
)

func NewStatus(source, target repository.Status, log *slog.Logger) *Status {
	return &Status{
		source: source,
		target: target,
		log:    log,
	}
}

type Status struct {
	source repository.Status
	target repository.Status
	log    *slog.Logger
}

func (sts *Status) Task() execute.TaskInfo {
	return execute.TaskInfo{
		Name: "[status] 统计状态数据同步",
		Func: sts.exec,
		Cron: time.Minute,
	}
}

func (sts *Status) exec(ctx context.Context) error {
	return nil
}
