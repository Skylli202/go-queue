package handler_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"testing"

	"github.com/Skylli202/go-queue/broker/handler"
	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	slogger := slog.New(slog.NewTextHandler(io.Discard, nil))
	h := handler.HealthHandler(slogger)
	require.HTTPStatusCode(t, h.ServeHTTP, http.MethodGet, "/", url.Values{}, http.StatusOK, "handler returned wrong status code")
}
