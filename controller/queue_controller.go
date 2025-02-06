package controller

import (
	"fmt"
	"net/http"
)

type QueueController struct{}

// var _ Controller = (*QueueController)(nil)

// func (c *QueueController) RegisterRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /queue", c.GetHandler)
// }

func (c *QueueController) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "GET /queue success !!")
}
