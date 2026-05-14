package stdlib

import (
	"context"
	"log/slog"
	"os"
)

// slog: structured logging package added in Go 1.21
// Replaces the older log package for structured use cases.

// ---- Default logger ----

func slogBasics() {
	// default logger writes text to stderr
	slog.Info("server started", "addr", ":8080")
	slog.Warn("high memory", "used_mb", 950)
	slog.Error("request failed", "err", "timeout", "url", "/api/users")
	slog.Debug("cache miss", "key", "user:123") // not printed by default (Info level)
}

// ---- Log levels ----
// Debug < Info < Warn < Error
// Default level is Info

// ---- JSON handler ----

func jsonLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // enable debug level
	})
	return slog.New(handler)
}

// ---- Text handler ----

func textLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}

// ---- Default logger replacement ----

func setDefaultLogger() {
	logger := jsonLogger()
	slog.SetDefault(logger) // all slog.Info etc. now use this logger
}

// ---- Attributes and groups ----

func attributesDemo(logger *slog.Logger) {
	// key-value pairs
	logger.Info("user logged in", "user_id", 42, "ip", "1.2.3.4")

	// typed attributes (avoids reflection, faster)
	logger.Info("metrics",
		slog.Int("requests", 1000),
		slog.String("status", "ok"),
		slog.Float64("latency_ms", 12.5),
		slog.Bool("cache_hit", true),
	)

	// group: namespace attributes
	logger.Info("request",
		slog.Group("http",
			slog.String("method", "GET"),
			slog.String("path", "/users"),
			slog.Int("status", 200),
		),
	)
}

// ---- Child loggers with pre-set attributes ----

func childLoggerDemo(logger *slog.Logger) {
	// With: returns a child logger with extra attributes always attached
	requestLogger := logger.With("request_id", "abc-123", "user_id", 42)

	requestLogger.Info("processing request")
	requestLogger.Info("fetching from db")
	requestLogger.Info("returning response")
	// all three log lines include request_id and user_id
}

// ---- Context-aware logging ----

func contextLogging(ctx context.Context, logger *slog.Logger) {
	// slog.InfoContext: passes context to handler (for trace IDs etc.)
	logger.InfoContext(ctx, "handling request", "endpoint", "/users")
}

// ---- Custom handler (example: adding trace ID from context) ----

type traceKey struct{}

func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceKey{}, id)
}

type TraceHandler struct {
	slog.Handler
}

func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	if id, ok := ctx.Value(traceKey{}).(string); ok {
		r.AddAttrs(slog.String("trace_id", id))
	}
	return h.Handler.Handle(ctx, r)
}

func newTraceLogger() *slog.Logger {
	base := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(&TraceHandler{base})
}

/*
slog vs older log package:
- slog is structured (key-value pairs) - machine-parseable
- log is unstructured text
- slog has levels; log doesn't
- slog is composable (handlers, child loggers)

slog vs zerolog/zap (third-party):
- slog is slower than zerolog/zap due to reflection in key-value pairs
- Use slog.Int(), slog.String() etc. for zero-allocation paths
- For hot paths with millions of log lines, zerolog or zap may be worth it
*/
