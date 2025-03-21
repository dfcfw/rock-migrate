package service

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func NewStatus(source, target repository.Status, log *slog.Logger) *Status {
	return &Status{
		source: source,
		target: target,
		log:    log,
	}
}

type Status struct {
	source  repository.Status
	target  repository.Status
	log     *slog.Logger
	syncing atomic.Bool
}

func (sts *Status) Task() execute.TaskInfo {
	return execute.TaskInfo{
		Name: "[status] 定时状态统计数据同步",
		Func: sts.exec,
		Cron: time.Minute,
	}
}

func (sts *Status) exec(parent context.Context) error {
	if !sts.syncing.CompareAndSwap(false, true) {
		return ErrTaskRunning
	}
	defer sts.syncing.Store(false)

	ctx, cancel := context.WithTimeout(parent, 10*time.Minute)
	defer cancel()

	// 查询目的数据库最新数据
	//var lastAt time.Time
	//opt := options.FindOne().
	//	SetSort(bson.D{{Key: "time", Value: -1}})
	//if last, _ := sts.target.FindOne(ctx, bson.M{}, opt); last != nil {
	//	lastAt = last.Time
	//}
	//
	//// 最早不过 180 天
	//day180 := time.Now().Add(-180 * 24 * time.Hour)
	//if lastAt.Before(day180) {
	//	lastAt = day180
	//}

	lastAt := time.Now().Add(-time.Hour)
	var cnt int
	attrs := []any{slog.Time("after_at", lastAt)}
	filter := bson.M{"time": bson.M{"$gte": lastAt}}
	for ips, err := range sts.source.All(ctx, filter, 100) {
		if err != nil {
			return err
		}
		cnt += len(ips)
		_, _ = sts.target.InsertMany(ctx, ips)
	}
	attrs = append(attrs, slog.Int("count", cnt))
	sts.log.Info("定时状态统计数据同步", attrs...)

	return nil
}
