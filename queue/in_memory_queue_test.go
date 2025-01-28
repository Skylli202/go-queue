package queue_test

import (
	"Skylli202/go-queue/queue"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewInMemoryQueue(t *testing.T) {
	require.NotNil(t, queue.NewInMemoryQueue(), "unexpected: NewInMemoryQueue() returned nil.")
}

func Test_InMemoryQueue(t *testing.T) {
	q := queue.NewInMemoryQueue()
	require.Equal(t, 0, q.Size(), "InMemoryQueue should have a size of 0 upon instanciation.")
	q.Enqueue("Anakin was supposed to bring balance to force, not to join them.")
	require.Equal(t, 1, q.Size(), "InMemoryQueue should have a size of 1 after enqueueing one message.")
	q.Enqueue("Do or do not, there is no try.")
	q.Enqueue("Fear is the path to the dark side. Fear leads to anger. Anger leads to hate. Hate leads to suffering.")
	q.Enqueue("When I left you, I was but the learner. Now I am the master.")
	require.Equal(t, 4, q.Size(), "InMemoryQueue should have a size of 4 after 4 enqueues.")
}
