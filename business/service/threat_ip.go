package service

import (
	"context"
	"errors"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrTaskRunning = errors.New("任务已经在运行")

func NewThreatIP(source, target repository.ThreatIP, log *slog.Logger) *ThreatIP {
	return &ThreatIP{
		source: source,
		target: target,
		log:    log,
	}
}

type ThreatIP struct {
	source  repository.ThreatIP
	target  repository.ThreatIP
	log     *slog.Logger
	syncing atomic.Bool
}

func (tip *ThreatIP) Task() execute.TaskInfo {
	return execute.TaskInfo{
		Name: "[threat_ip] 恶意 IP 数据定时同步",
		Func: tip.exec,
		Cron: time.Minute,
	}
}

func (tip *ThreatIP) exec(parent context.Context) error {
	if !tip.syncing.CompareAndSwap(false, true) {
		return ErrTaskRunning
	}
	defer tip.syncing.Store(false)

	ctx, cancel := context.WithTimeout(parent, 10*time.Minute)
	defer cancel()

	// 查询目的数据库最新数据
	var lastAt time.Time
	opt := options.FindOne().
		SetSort(bson.D{{Key: "last_at", Value: -1}})
	if last, _ := tip.target.FindOne(ctx, bson.M{}, opt); last != nil {
		lastAt = last.LastAt
	}

	// 最早不过 180 天
	day180 := time.Now().Add(-180 * 24 * time.Hour)
	if lastAt.Before(day180) {
		lastAt = day180
	}

	var cnt int
	attrs := []any{slog.Time("after_at", lastAt)}
	filter := bson.M{"last_at": bson.M{"$gt": lastAt}}
	for ips, err := range tip.source.All(ctx, filter, 100) {
		if err != nil {
			return err
		}
		cnt += len(ips)
		_, _ = tip.target.InsertMany(ctx, ips)
	}
	attrs = append(attrs, slog.Int("count", cnt))
	tip.log.Info("恶意 IP 数据定时同步", attrs...)

	return nil
}
