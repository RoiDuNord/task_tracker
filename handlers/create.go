package handlers

import (
	"context"
	"encoding/json"
	"io_bound/models"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TaskRequest struct {
	TaskName string `json:"taskName"`
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode JSON to TaskRequest", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	creationTime := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	task := &models.Task{
		ID:           id,
		Name:         req.TaskName,
		Status:       models.StatusCreated,
		CreationTime: creationTime,
		CancelFunc:   cancel,
	}

	s.tasks.Store(id, task)

	slog.Info("task created", "id", id)

	go s.runTask(ctx, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]string{"id": id}); err != nil {
		slog.Error("failed to encode task to JSON", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
