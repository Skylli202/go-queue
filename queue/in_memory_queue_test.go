package queue_test

import (
	"Skylli202/go-queue/queue"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewInMemoryQueue(t *testing.T) {
	require.NotNil(t, queue.NewInMemoryQueue(queue.FIFO), "unexpected: NewInMemoryQueue() returned nil.")
}

func Test_InMemoryQueue(t *testing.T) {
	q := queue.NewInMemoryQueue(queue.FIFO)
	m0 := queue.SimpleMessage("Anakin was supposed to bring balance to force, not to join them.")
	require.Equal(t, 0, q.Size(), "InMemoryQueue should have a size of 0 upon instanciation.")
	q.Enqueue(m0)
	require.Equal(t, 1, q.Size(), "InMemoryQueue should have a size of 1 after enqueueing one message.")
	q.Enqueue("Do or do not, there is no try.")
	q.Enqueue("Fear is the path to the dark side. Fear leads to anger. Anger leads to hate. Hate leads to suffering.")
	q.Enqueue("When I left you, I was but the learner. Now I am the master.")
	require.Equal(t, 4, q.Size(), "InMemoryQueue should have a size of 4 after 4 enqueues.")

	actual, err := q.Dequeue()
	require.Equal(t, m0, *actual, "In FIFO, the first queued message is expected to the first dequeued.")
	require.Nil(t, err, "Dequeue a FIFO queue of size 4 should not return an error.")
}

func Test_InMemoryQueueEmptyQueueError(t *testing.T) {
	q := queue.NewInMemoryQueue(queue.FIFO)
	m, err := q.Dequeue()
	require.Nil(t, m, "Dequeueing an empty queue should return a nil pointer for the message")
	require.ErrorIs(t, err, &queue.EmptyQueueError{}, "")
}
