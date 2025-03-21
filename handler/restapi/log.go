package restapi

import (
	"sync/atomic"

	"github.com/dfcfw/rock-migrate/business"
	"github.com/dfcfw/rock-migrate/library/eventsource"
	"github.com/xgfone/ship/v5"
)

func NewLog(biz *business.Log) *Log {
	return &Log{
		biz: biz,
	}
}

type Log struct {
	biz *business.Log
	lim atomic.Int32
}

func (lg *Log) RegisterRoute(r *ship.RouteGroupBuilder) error {
	r.Route("/sse/log/tail").GET(lg.tail)
	return nil
}

func (lg *Log) tail(c *ship.Context) error {
	w, r := c.ResponseWriter(), c.Request()
	sse := eventsource.Accept(w, r)
	if sse == nil {
		c.Warnf("不是 Server-Sent Events 连接")
		return ship.ErrUnsupportedMediaType
	}

	// 限制观察数
	num := lg.lim.Add(1)
	defer lg.lim.Add(-1)
	if num > 5 {
		c.Errorf("观测窗口超限")
		return ship.ErrTooManyRequests
	}

	ctx := r.Context()
	if !lg.biz.Attach(ctx, sse) {
		c.Errorf("日志观测失败")
		return nil
	}
	c.Warnf("进入日志观测")
	<-sse.Done()
	c.Infof("结束日志观测")
	lg.biz.Detach(ctx, sse)

	return nil
}
