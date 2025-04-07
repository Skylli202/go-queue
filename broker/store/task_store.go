package store

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Task is the Model (in the Repository Pattern context)
type Task struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Payload   []byte    `json:"payload"`
}

// TODO: Implement this function to auto-generate the UUID & CreatedAt
func NewTask() *Task {
	return &Task{}
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %q; CreatedAt: %q; Payload: %s", t.ID, t.CreatedAt, string(t.Payload))
}

// TaskStore is the implementation of the respository pattern for the Model Task.
type TaskStore struct {
	db *sql.DB
}

// TODO: Add a *slogger.Logger in argument & add useful login
func NewTaskStore(db *sql.DB) *TaskStore {
	// TODO: Move this into an init method or something, so the error can be properly
	// handled in the `cmd/broker/web/main.go:run()` instead of an ugly panic here.
	const create string = `
	CREATE TABLE IF NOT EXISTS task (
		id TEXT NOT NULL PRIMARY KEY,
		created_at TEXT NOT NULL,
		payload BLOB
	);`

	if _, err := db.Exec(create); err != nil {
		panic(err)
	}

	return &TaskStore{
		db: db,
	}
}

// HACK: Might have to rename this method later when a Store interface will be
// required as more Model are created within the application.
func (s *TaskStore) InsertTask(t Task) (*Task, error) {
	const dtFmt string = "2006-01-02 15:04:05"
	const create string = "INSERT INTO task (id, created_at, payload) VALUES (?,?,?);"

	// TODO: Add a bit of defensive programming here:
	// - [ ] Check that ID is not null, if it is, initialize it.
	// - [ ] Check that CreatedAt is not null, if it is, initialize it.

	_, err := s.db.Exec(
		create,
		t.ID,
		t.CreatedAt.UTC().Format(dtFmt),
		t.Payload,
	)
	if err != nil {
		return nil, err
	} else {
		slog.Info("Task created", "ID", t.ID, "created_at", t.CreatedAt.Format(dtFmt))
	}

	return &t, nil
}
