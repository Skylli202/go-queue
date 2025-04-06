package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey string

const RequestIdKey ctxKey = "request-id"

func RequestIdMiddleware(slogger *slog.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), RequestIdKey, requestId.String()))
		h.ServeHTTP(w, r)
	})
}

func NewRequestIdMiddleware(slogger *slog.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return RequestIdMiddleware(slogger, h)
	}
}
