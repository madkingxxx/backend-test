package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()
			ctx := req.Context()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zap.Field{
				zap.Int("status", res.Status),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
				zap.Duration("latency", time.Since(start)),
			}

			if id != "" {
				fields = append(fields, zap.String("request_id", id))
			}

			n := res.Status
			switch {
			case n >= 500:
				utils.Logger.Error(ctx, "Server error", fields...)
			case n >= 400:
				utils.Logger.Warn(ctx, "Client error", fields...)
			case n >= 300:
				utils.Logger.Info(ctx, "Redirection", fields...)
			default:
				utils.Logger.Info(ctx, "Success", fields...)
			}

			return nil
		}
	}
}
