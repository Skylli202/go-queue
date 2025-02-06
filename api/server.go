package api

import (
	"net/http"
)

type Controller interface {
	RegisterRoutes(*http.ServeMux)
}

func NewServer(
	qc *QueueController,
) http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux

	// qc.RegisterRoutes(mux)
	mux.HandleFunc("GET /api/queue", qc.GetHandler)

	return handler
}
