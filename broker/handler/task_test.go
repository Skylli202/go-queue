package handler_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"testing"

	"github.com/Skylli202/go-queue/broker/handler"
	"github.com/Skylli202/go-queue/broker/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testTaskStore struct {
	store map[uuid.UUID]store.Task
}

func newTestTaskStore() *testTaskStore {
	return &testTaskStore{
		store: make(map[uuid.UUID]store.Task, 10),
	}
}

func (s *testTaskStore) Insert(task store.Task) (*store.Task, error) {
	s.store[task.ID] = task
	return &task, nil
}

var _ store.Store[store.Task] = (*testTaskStore)(nil)

func TestTaskHandler(t *testing.T) {
	slogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	taskStore := newTestTaskStore()

	h := handler.CreateTaskHandler(slogger, taskStore)

	require.HTTPStatusCode(t, h.ServeHTTP, http.MethodPost, "", url.Values{}, http.StatusCreated, "handler returned wrong status code")
	require.Len(t, taskStore.store, 1, "handler did not properly Insert the task in the store")
}
