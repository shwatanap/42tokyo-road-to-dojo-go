package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func HttpLogging(msg string, r *http.Request, start time.Time, latency int64) {
	logger, _ := zap.NewDevelopment()
	logger.Info(
		msg,
		// zap.String("OS", r.Context().Value(context.OsKey).(string)),
		zap.String("method", r.Method),
		zap.String("host", r.Host),
		zap.String("path", r.URL.Path),
		zap.Time("start", start),
		zap.Int64("latency", latency),
	)
}

func ErrorLogging(msg string, err error, r *http.Request) {
	logger, _ := zap.NewDevelopment()
	logger.Error(
		msg,
		zap.Error(err),
		zap.String("method", r.Method),
		zap.String("host", r.Host),
		zap.String("path", r.URL.Path),
	)
}
