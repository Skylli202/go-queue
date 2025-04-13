package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/Skylli202/go-queue/broker/model"
	"github.com/Skylli202/go-queue/broker/store"
	"github.com/google/uuid"
)

func CreateTaskHandler(slogger *slog.Logger, s store.Store[model.Task]) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t := new(model.Task)
			t.ID = uuid.New()
			t.CreatedAt = time.Now().UTC()
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err != nil {
				slogger.Error("error while reading request body", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			t.Payload = []byte(b)

			t, err = s.Insert(*t)
			if err != nil {
				slogger.Error("error while inserting task into database", "task", t, "error", err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(&t); err != nil {
				slogger.Error("error while encoding task to ResponseWriter", "task", t, "error", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		},
	)
}
