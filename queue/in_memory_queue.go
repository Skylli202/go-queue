package queue

import (
	"errors"
)

type (
	SimpleMessage string
	InMemoryQueue struct {
		QueueType QueueType
		store     []Message
	}
	QueueType string
)

// Compilation time check that:
var (
	// InMemoryQueue implement Queue interface
	_ Queue = (*InMemoryQueue)(nil)
	// Message implement Message interface
	_ Message = (*SimpleMessage)(nil)
)

const (
	FIFO     QueueType = "FIFO"
	LIFO     QueueType = "LIFO"
	Priority QueueType = "Priority"
)

func NewInMemoryQueue(queueType QueueType) *InMemoryQueue {
	return &InMemoryQueue{
		store:     make([]Message, 0),
		QueueType: queueType,
	}
}

func (q *InMemoryQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *InMemoryQueue) Peek() (*Message, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	return &q.store[0], nil
}

func (q *InMemoryQueue) Size() int { return len(q.store) }

func (q *InMemoryQueue) Enqueue(m Message) error {
	q.store = append(q.store, m)
	return nil
}

func (q *InMemoryQueue) Dequeue() (*Message, error) {
	if q.Size() == 0 {
		return nil, ErrEmptyQueue
	}

	switch q.QueueType {
	case FIFO:
		pop := q.store[0]
		q.store = q.store[1:]
		return &pop, nil
	default:
		return nil, ErrNotImplementedQueueType
	}
}

// Implementation of the error interface for the various custom error types.
// EmptyQueueError is returned to signal that the queue is empty and that the
// operation cannot be performed.

var ErrEmptyQueue = errors.New("queue is empty - nothing to dequeue")

// is returned to signal that the queue got QueueType that is either
// unsupported, deprecated or simply not implemented. In theory this error
// should not be returned often.
var ErrNotImplementedQueueType = errors.New("queue type not implemented")
