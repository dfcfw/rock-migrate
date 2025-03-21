package business

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dfcfw/rock-migrate/datalayer/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

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

// AddCron 添加到定时任务
func (tip *ThreatIP) AddCron(parent context.Context) {
	if !tip.syncing.CompareAndSwap(false, true) {
		return
	}
	defer tip.syncing.Store(false)

	ctx, cancel := context.WithTimeout(parent, time.Hour)
	defer cancel()

	if err := tip.sync(ctx); err != nil {
		tip.log.WarnContext(ctx, "同步恶意IP表错误", slog.Any("error", err))
	} else {
		tip.log.InfoContext(ctx, "同步恶意IP结束")
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

	filter := bson.M{"last_at": bson.M{"$gt": lastAt}}
	opt := options.Find().SetLimit(100)
	for ips, err := range tip.source.All(ctx, filter, opt) {
		if err != nil {
			tip.log.ErrorContext(ctx, "同步 threat_ip 出错", slog.Any("error", err))
			return err
		}
		_, _ = tip.target.InsertMany(ctx, ips)
	}

	return nil
}
