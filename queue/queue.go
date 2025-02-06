package queue

type (
	Message interface{}
	Queue   interface {
		Enqueue(Message) error
		Dequeue() (*Message, error)
		Peek() (*Message, error)
		Size() int
		IsEmpty() bool
	}
)
