package handler_test

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Skylli202/go-queue/broker/handler"
	"github.com/Skylli202/go-queue/broker/model"
	"github.com/Skylli202/go-queue/broker/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testTaskStore struct {
	store map[uuid.UUID]model.Task
}

func newTestTaskStore() *testTaskStore {
	return &testTaskStore{
		store: make(map[uuid.UUID]model.Task, 10),
	}
}

func (s *testTaskStore) Insert(task model.Task) (*model.Task, error) {
	s.store[task.ID] = task
	return &task, nil
}

type testStoreFailer struct{}

func newTestStoreFailer() *testStoreFailer {
	return &testStoreFailer{}
}

func (testStoreFailer) Insert(model.Task) (*model.Task, error) {
	return nil, errStoreInsertFailure
}

var _ store.Store[model.Task] = (*testTaskStore)(nil)

var (
	errReadFailure        = errors.New("reading failure")
	errStoreInsertFailure = errors.New("insert failure")
)

type ReaderFail struct{}

func (ReaderFail) Read(p []byte) (n int, err error) {
	return 0, errReadFailure
}

func TestTaskHandler(t *testing.T) {
	slogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	taskStore := newTestTaskStore()

	h := handler.CreateTaskHandler(slogger, taskStore)

	t.Run("happy path", func(t *testing.T) {
		require.HTTPStatusCode(t, h.ServeHTTP, http.MethodPost, "", url.Values{}, http.StatusCreated, "handler returned wrong status code")
		require.Len(t, taskStore.store, 1, "handler did not properly Insert the task in the store")
	})

	t.Run("Error handling: reading request's body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", &ReaderFail{})
		res := httptest.NewRecorder()

		h.ServeHTTP(res, req)

		// TODO: add tests to check that we log the error as well
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
	t.Run("Error handling: insert task into store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		h := handler.CreateTaskHandler(slogger, newTestStoreFailer())
		h.ServeHTTP(res, req)

		// TODO: add tests to check that we log the error as well
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}
