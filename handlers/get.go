package handlers

import (
	"encoding/json"
	"fmt"
	"io_bound/models"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := s.getTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if task.Status == models.StatusRunning && task.StartTime != nil {
		task.Duration = float32(time.Since(*task.StartTime).Seconds())
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		slog.Error("failed to encode task to JSON", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	slog.Info("task info obtained", "id", id)
}

func (s *Server) getTaskByID(id string) (*models.Task, error) {
	tRaw, ok := s.tasks.Load(id)
	if !ok {
		return nil, fmt.Errorf("task not found")
	}
	task, ok := tRaw.(*models.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type")
	}
	return task, nil
}
