package logger

import (
	"context"
	"log/slog"
)

// Уникальный тип ключа для контекста (Go way)
type ctxLoggerKey struct{}

// ContextWithLogger кладет логгер в контекст
func ContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

// LoggerFromContext достает логгер из контекста.
// Если логгера там нет, возвращает дефолтный slog.Default(), чтобы код не падал.
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(ctxLoggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
