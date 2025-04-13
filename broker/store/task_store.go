package store

import (
	"database/sql"
	"log/slog"

	"github.com/Skylli202/go-queue/broker/model"
	_ "github.com/mattn/go-sqlite3"
)

type Store[T any] interface {
	Insert(t T) (*T, error)
}

var _ Store[model.Task] = (*TaskStore)(nil)

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
func (s *TaskStore) Insert(t model.Task) (*model.Task, error) {
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
