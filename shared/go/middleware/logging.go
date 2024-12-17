package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// CtxAttrs should be a map values wanted
// from a context
type CtxAttrs map[string]struct{}

// Logs request start/end with optional timing and attributes extracted from the request context
func LoggingMiddleware(logger *slog.Logger, ctxAttrs CtxAttrs) func(http.Handler) http.Handler {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			h.ServeHTTP(w, r)

			end := time.Since(start)

			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo.Level(),
				"processed request",
				slog.String("method", r.Method),
				slog.String("host", r.Host),
				slog.String("url", r.URL.Path),
				slog.String("query", r.URL.RawQuery),
				slog.String("agent", r.UserAgent()),
				slog.String("duration", end.String()),
			)
		})
	}
}
