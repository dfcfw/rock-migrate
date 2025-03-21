package service

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewFlow(source repository.Flow, target repository.Flow, log *slog.Logger) *Flow {
	return &Flow{
		source: source,
		target: target,
		log:    log,
	}
}

type Flow struct {
	source  repository.Flow
	target  repository.Flow
	log     *slog.Logger
	syncing atomic.Bool
}

func (flw *Flow) Task() execute.TaskInfo {
	return execute.TaskInfo{
		Name: "[flow] 定时同步流量统计数据",
		Func: flw.exec,
		Cron: time.Minute,
	}
}

func (flw *Flow) exec(parent context.Context) error {
	if !flw.syncing.CompareAndSwap(false, true) {
		return ErrTaskRunning
	}
	defer flw.syncing.Store(false)

	ctx, cancel := context.WithTimeout(parent, 10*time.Minute)
	defer cancel()

	// 查询目的数据库最新数据
	var lastAt time.Time
	opt := options.FindOne().
		SetSort(bson.D{{Key: "time", Value: -1}})
	if last, _ := flw.target.FindOne(ctx, bson.M{}, opt); last != nil {
		lastAt = time.UnixMilli(last.Time)
	}

	// 最早不过 180 天
	day180 := time.Now().Add(-180 * 24 * time.Hour)
	if lastAt.Before(day180) {
		lastAt = day180
	}

	var cnt int
	attrs := []any{slog.Time("after_at", lastAt)}
	filter := bson.M{"time": bson.M{"$gt": lastAt.UnixMilli()}}
	for ips, err := range flw.source.All(ctx, filter, 100) {
		if err != nil {
			return err
		}
		cnt += len(ips)
		_, _ = flw.target.InsertMany(ctx, ips)
	}
	attrs = append(attrs, slog.Int("count", cnt))
	flw.log.Info("定时同步流量统计数据", attrs...)

	return nil
}
