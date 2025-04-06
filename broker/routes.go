package broker

import (
	"log/slog"
	"net/http"
)

// Main function to define the broker's service API surface.
func addRoutes(
	mux *http.ServeMux,
	slogger *slog.Logger,
) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		slogger.InfoContext(r.Context(), "incoming request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)
		w.WriteHeader(http.StatusOK)
	})
}
