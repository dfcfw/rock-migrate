package middle

import (
	"log/slog"
	"net"

	"github.com/xgfone/ship/v5"
)

func AccessLog(h ship.Handler) ship.Handler {
	return func(c *ship.Context) error {
		remoteAddr, clientIP := c.RemoteAddr(), c.ClientIP()
		directIP, _, _ := net.SplitHostPort(remoteAddr)
		if directIP == "" {
			directIP = remoteAddr
		}
		host, method := c.Host(), c.Method()
		req := c.Request()
		reqURL := req.URL

		attrs := []any{
			slog.String("client_ip", clientIP),
			slog.String("remote_addr", remoteAddr),
			slog.String("method", method),
			slog.String("host", host),
			slog.String("path", reqURL.Path),
		}

		err := h(c)
		if err != nil {
			attrs = append(attrs, slog.String("error", err.Error()))
			c.Warnf("访问接口出错", attrs...)
		} else {
			c.Infof("访问接口", attrs...)
		}

		return err
	}
}
