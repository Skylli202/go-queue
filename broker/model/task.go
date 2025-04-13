package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Task is the Model (in the Repository Pattern context)
type Task struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Payload   []byte    `json:"payload"`
}

// TODO: Implement this function to auto-generate the UUID & CreatedAt
func NewTask() *Task {
	return &Task{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %q; CreatedAt: %q; Payload: %s", t.ID, t.CreatedAt, string(t.Payload))
}
