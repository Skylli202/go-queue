package store_test

import (
	"Skylli202/go-queue/queue"
	"Skylli202/go-queue/store"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SaveEmptyQueue(t *testing.T) {
	tempDir := os.TempDir()
	defer os.RemoveAll(tempDir)

	s := store.NewFileQueueStore(tempDir)
	q := queue.NewInMemoryQueue(queue.FIFO)

	id, err := s.Save(q)
	require.Nil(t, err, "Save should not return an err when saving an empty queue. Error: %v", err)
	require.NotEmpty(t, id, "InMemoryQueueStore.Save() should not return an empty (invalid) UUID.")

	fs := os.DirFS(tempDir)
	file, err := fs.Open(id.String())
	if err != nil {
		t.Fatalf("InMemoryQueueStore should have created a folder for the queue passed to be saved (persisted), but it did not. Err: %v", err)
	}

	info, err := file.Stat()
	if err != nil {
		t.Fatalf("Calling Stat() on the directory created by InMemoryQueueStore.Save() should not return an err, but it did. Err: %v", err)
	}
	assert.True(t, info.IsDir(), "InMemoryQueueStore.Save() should have created directory")
}

func Test_GetQueue(t *testing.T) {
	tempDir := os.TempDir()
	defer os.RemoveAll(tempDir)
	id := uuid.New()
	err := os.Mkdir(path.Join(tempDir, id.String()), 0777)
	require.NoError(t, err, "An error occured while setting up the test.")

	s := store.NewFileQueueStore(tempDir)

	_, err = s.Get(id)
	require.NoErrorf(t, err, "s.Get(%s) should not have returned an error: %v", id.String(), err)
}
