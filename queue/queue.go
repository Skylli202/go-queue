package queue

type (
	Message interface{}
	Queue   interface {
		Enqueue(Message) error
		Dequeue() (Message, error)
		Peek()
		Size() int
		IsEmpty() bool
	}
)
