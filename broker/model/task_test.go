package model_test

import (
	"testing"

	"github.com/Skylli202/go-queue/broker/model"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	task := model.NewTask()

	require.NotNil(t, task)
	require.NotNil(t, task.ID)
	require.NotZero(t, task.ID)
	require.NotNil(t, task.CreatedAt)
	require.NotZero(t, task.CreatedAt)
}
