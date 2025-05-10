package logger

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps a zap.Logger and integrates Sentry for error reporting
type Logger struct {
	*zap.Logger
}

// New initializes Sentry (if DSN provided) and returns a configured zap.Logger
func New(sentryDSN string) (*Logger, error) {
	// Initialize Sentry if DSN is provided
	if sentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              sentryDSN,
			AttachStacktrace: true,
		}); err != nil {
			return nil, fmt.Errorf("sentry init: %w", err)
		}
	}

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("zap build: %w", err)
	}

	return &Logger{logger}, nil
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	if l.Logger != nil {
		return l.Logger.Sync()
	}
	return nil
}

// Error logs an error-level message and reports the error to Sentry
func (l *Logger) Error(msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
	if sentry.CurrentHub().Client() != nil {
		sentry.CaptureException(err)
		sentry.Flush(2 * time.Second)
	}
}

// Info logs an info-level message
func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.Logger.Info(msg, fields...)
}
