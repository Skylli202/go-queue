package broker

import (
	"log/slog"
	"net/http"

	"github.com/Skylli202/go-queue/broker/handler"
)

// Main function to define the broker's service API surface.
func addRoutes(
	mux *http.ServeMux,
	slogger *slog.Logger,
) {
	mux.Handle("/health", handler.HealthHandler(slogger))
}
