package service

import (
	"context"
	"errors"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dfcfw/rock-migrate/business/execute"
	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

	// 最早不过 30 天
	day30 := time.Now().Add(-30 * 24 * time.Hour)
	if lastAt.Before(day30) {
		lastAt = day30
	}

	var cnt int
	attrs := []any{slog.Time("last_at", lastAt)}
	filter := bson.M{"last_at": bson.M{"$gt": lastAt}}
	for ips, err := range tip.source.All(ctx, filter, 100) {
		if err != nil {
			return err
		}
		cnt += len(ips)
		_, _ = tip.target.InsertMany(ctx, ips)
	}
	attrs = append(attrs, slog.Int("count", cnt))
	tip.log.Info("恶意IP同步完毕", attrs...)

	return nil
}

// Sync 添加到定时任务
func (tip *ThreatIP) dosync() {
	if !tip.syncing.CompareAndSwap(false, true) {
		tip.log.Warn("同步任务已正在运行")
		return
	}
	defer tip.syncing.Store(false)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	if err := tip.sync(ctx); err != nil {
		tip.log.WarnContext(ctx, "同步恶意IP表错误", slog.Any("error", err))
	}
}

func (tip *ThreatIP) sync(ctx context.Context) error {
	// 查询目的数据库最新数据
	var lastAt time.Time
	{
		opt := options.FindOne().
			SetSort(bson.D{{Key: "last_at", Value: -1}})
		if last, _ := tip.target.FindOne(ctx, bson.M{}, opt); last != nil {
			lastAt = last.LastAt
		}
	}

	var cnt int
	attrs := []any{slog.Time("last_at", lastAt)}
	filter := bson.M{"last_at": bson.M{"$gt": lastAt}}
	for ips, err := range tip.source.All(ctx, filter, 100) {
		if err != nil {
			tip.log.ErrorContext(ctx, "同步 threat_ip 出错", slog.Any("error", err))
			return err
		}
		cnt += len(ips)
		_, _ = tip.target.InsertMany(ctx, ips)
	}
	attrs = append(attrs, slog.Int("count", cnt))
	tip.log.Info("同步恶意IP结束", attrs...)

	return nil
}
