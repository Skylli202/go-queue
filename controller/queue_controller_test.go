package controller_test

import (
	"Skylli202/go-queue/controller"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueueControllerGetHandler(t *testing.T) {
	tests := []struct {
		name               string
		expectedStatusCode int
	}{
		{name: "empty queue", expectedStatusCode: http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			qc := controller.NewQueueController()

			assert.HTTPStatusCode(qc.GetHandler, http.MethodGet, "/api/queue", nil, http.StatusOK)
		})
	}
}
