package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Metric(slogger *slog.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			h.ServeHTTP(w, r)

			id, ok := r.Context().Value(RequestIdKey).(string)
			// Defensive coding in case middleware haven't been called in the correct order
			if !ok {
				slogger.Error("request do not have a request id")
				id = "no-uuid-in-context"
			}
			duration := time.Since(start)
			slogger.WithGroup("general").InfoContext(r.Context(), "general information",
				slog.String("remote Addr", r.RemoteAddr),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("duration", duration.String()),
				slog.String(string(RequestIdKey), id),
			)
		},
	)
}
