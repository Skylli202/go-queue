package store

import "database/sql"

type Stores struct {
	TaskStore *TaskStore
}

func NewStores(db *sql.DB) *Stores {
	taskStore := NewTaskStore(db)
	return &Stores{
		TaskStore: taskStore,
	}
}
