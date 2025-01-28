package queue

import "fmt"

type (
	SimpleMessage string
	InMemoryQueue struct {
		store     []Message
		queueType QueueType
	}
	QueueType int
)

const (
	FIFO QueueType = iota
	LIFO
	Priority
)

func NewInMemoryQueue(queueType QueueType) *InMemoryQueue {
	return &InMemoryQueue{
		store:     make([]Message, 0),
		queueType: queueType,
	}
}

func (q *InMemoryQueue) Size() int { return len(q.store) }

func (q *InMemoryQueue) Enqueue(m Message) error {
	q.store = append(q.store, m)
	return nil
}

func (q *InMemoryQueue) Dequeue() (*Message, error) {
	if q.Size() == 0 {
		return nil, &EmptyQueueError{}
	}

	switch q.queueType {
	case FIFO:
		pop := q.store[0]
		q.store = q.store[1:]
		return &pop, nil
	default:
		return nil, &NotImplementedQueueType{queueType: q.queueType}
	}
}

// Implementation of the error interface for the various custom error types.
// EmptyQueueError is returned to signal that the queue is empty and that the
// operation cannot be performed.
type EmptyQueueError struct{}

func (e *EmptyQueueError) Error() string {
	return "Nothing to dequeue. Queue is empty."
}

// is returned to signal that the queue got QueueType that is either
// unsupported, deprecated or simply not implemented. In theory this error
// should not be returned often.
type NotImplementedQueueType struct {
	queueType QueueType
}

func (e *NotImplementedQueueType) Error() string {
	return fmt.Sprintf("QueueType %2d is not implemented, supported or is deprecated.", e.queueType)
}
