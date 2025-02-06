package api

import (
	"Skylli202/go-queue/controller"
	"net/http"
)

type Controller interface {
	RegisterRoutes(*http.ServeMux)
}

func NewServer(
	qc *controller.QueueController,
) http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux

	// qc.RegisterRoutes(mux)
	mux.HandleFunc("GET /api/queue", qc.GetHandler)
	mux.HandleFunc("POST /api/queue", qc.PostHandler)

	return handler
}
