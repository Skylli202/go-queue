package broker

import (
	"log/slog"
	"net/http"

	"github.com/Skylli202/go-queue/broker/handler"
	"github.com/Skylli202/go-queue/broker/store"
)

// Main function to define the broker's service API surface.
func addRoutes(
	mux *http.ServeMux,
	slogger *slog.Logger,
	stores *store.Stores,
) {
	mux.Handle("/health", handler.HealthHandler(slogger))

	mux.Handle("POST /task", handler.CreateTaskHandler(slogger, stores.TaskStore))
}
