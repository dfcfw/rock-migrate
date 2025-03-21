package business

import (
	"context"
	"io"
	"log/slog"

	"github.com/dfcfw/rock-migrate/library/dynwriter"
)

func NewLog(dwrt dynwriter.Writer, log *slog.Logger) *Log {
	return &Log{
		dwrt: dwrt,
		log:  log,
	}
}

type Log struct {
	dwrt dynwriter.Writer
	log  *slog.Logger
}

func (lg *Log) Attach(ctx context.Context, w io.Writer) bool {
	succeed := lg.dwrt.Attach(w)
	if succeed {
		lg.log.WarnContext(ctx, "新增日志输出")
	} else {
		lg.log.ErrorContext(ctx, "新增日志输出失败")
	}

	return succeed
}

func (lg *Log) Detach(ctx context.Context, w io.Writer) bool {
	succeed := lg.dwrt.Detach(w)
	if succeed {
		lg.log.WarnContext(ctx, "移除日志输出")
	} else {
		lg.log.ErrorContext(ctx, "移除日志输出失败")
	}

	return succeed
}
