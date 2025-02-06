package controller

import (
	"Skylli202/go-queue/queue"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type QueueController struct {
	queues map[uuid.UUID]queue.Queue
}

func NewQueueController() *QueueController {
	return &QueueController{
		queues: make(map[uuid.UUID]queue.Queue),
	}
}

func (c *QueueController) GetHandler(w http.ResponseWriter, r *http.Request) {
	type value struct {
		QueueType queue.QueueType `json:"queuetype"`
		Size      int             `json:"size"`
	}
	data := make(map[uuid.UUID]value, len(c.queues))

	for id, q := range c.queues {
		if q, ok := q.(*queue.InMemoryQueue); ok {
			data[id] = value{QueueType: q.QueueType, Size: q.Size()}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err := encode(w, r, http.StatusOK, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *QueueController) PostHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		QueueType queue.QueueType `json:"queuetype"`
	}
	t, err := decode[request](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	q := queue.NewInMemoryQueue(t.QueueType)

	c.queues[id] = q

	type response struct{ Uuid uuid.UUID }
	err = encode(w, r, http.StatusCreated, &response{Uuid: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func encode[T any](w http.ResponseWriter, _ *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
