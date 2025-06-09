package handlers

import (
	"io_bound/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := s.getTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if task.Status == models.StatusRunning && task.CancelFunc != nil {
		task.CancelFunc()
	}
	s.tasks.Delete(id)

	slog.Info("task deleted", "task_id", id)

	w.WriteHeader(http.StatusNoContent)
}
