package store

import (
	"Skylli202/go-queue/queue"
	"errors"
	"os"
	"path"

	"github.com/google/uuid"
)

// Store is responsible for the persistency.
type InMemoryQueueStore struct {
	rootPath string
}

func NewFileQueueStore(rootPath string) *InMemoryQueueStore {
	return &InMemoryQueueStore{rootPath: rootPath}
}

var ErrAlreadyExist = errors.New("Queue already exist")

func (s *InMemoryQueueStore) Save(q queue.Queue) (uuid.UUID, error) {
	id := uuid.New()

	path := path.Join(s.rootPath, id.String())
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(path, 0774)
			if err != nil {
				return uuid.Nil, err
			}
			return id, nil
		}
		return uuid.Nil, err
	}

	// Hopefully this is never returned? As collision with UUID v4 is very unlikely...
	return uuid.Nil, ErrAlreadyExist
}
