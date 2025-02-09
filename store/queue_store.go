package store

import (
	"Skylli202/go-queue/queue"
	"errors"
	"fmt"
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

var ErrQueueAlreadySaved = errors.New("Queue already saved")

func (s *InMemoryQueueStore) Save(q queue.Queue) (uuid.UUID, error) {
	id := uuid.New()

	path := path.Join(s.rootPath, id.String())
	fmt.Printf("path: %s\n", path)
	_, err := os.Stat(path)
	fmt.Printf("err: %v\n", err)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating directory: " + path)
		err := os.MkdirAll(path, 0774)
		if err != nil {
			return uuid.Nil, err
		}
	} else if err != nil {
		fmt.Printf("error: %v", err)
		return uuid.Nil, err
	}

	return id, nil
}
