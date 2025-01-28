package queue

type (
	SimpleMessage string
	InMemoryQueue struct {
		store []Message
	}
)

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		store: make([]Message, 0),
	}
}

func (q *InMemoryQueue) Size() int { return len(q.store) }

func (q *InMemoryQueue) Enqueue(m Message) error {
	q.store = append(q.store, m)
	return nil
}
