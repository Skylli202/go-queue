package handler

import (
	"log/slog"
	"net/http"
)

func HealthHandler(slogger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)
}
