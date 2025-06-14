package server

import (
	"encoding/json"

	"log/slog"
	"net/http"
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

	if req.TaskName == "" {
		slog.Error("task name is empty", "request", req)
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	t, err := s.tasker.CreateNewTask(req.TaskName)
	if err != nil {
		slog.Error("failed to create new task", "error", err, "task_name", req.TaskName)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		slog.Error("failed to encode Task to JSON", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
