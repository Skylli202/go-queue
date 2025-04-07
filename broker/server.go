package broker

import (
	"log/slog"
	"net/http"

	"github.com/Skylli202/go-queue/broker/middleware"
	"github.com/Skylli202/go-queue/broker/store"
)

type BrokerServerConfig struct{}

// Broker server constructor: it create the service server.
// All dependencies are passed as argument, as pointer to
// to allow 'nil' value in test that does not requires all
// dependencies.
//
// It is responsible for all the top-level HTTP tasks:
// CORS, auth middleware, logging, etc. It also call the
// routes.go to define the service's API surface.
func NewBrokerServer(
	config *BrokerServerConfig,
	slogger *slog.Logger,
	stores *store.Stores,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, slogger, stores)
	var handler http.Handler = mux
	handler = middleware.Metric(slogger, handler)
	handler = middleware.RequestIdMiddleware(slogger, handler)
	// handler = middleware.SimpleHeaderMiddleware(slogger, handler)
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)

	return handler
}
