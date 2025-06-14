package server

import (
	"fmt"
	"io_bound/config"
	"io_bound/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type AbstractTaskManager interface {
	CreateNewTask(taskName string) (models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	DeleteTask(id string) error
}

type Server struct {
	tasker AbstractTaskManager
}

func NewServer(tm AbstractTaskManager) *Server {
	slog.Info("server created")
	return &Server{
		tasker: tm,
	}
}

func (srv *Server) setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/tasks", srv.Create)
	r.Get("/tasks/{id}", srv.Get)
	r.Delete("/tasks/{id}", srv.Delete)

	slog.Info("router setup completed", "routes", []string{"POST /tasks", "GET /tasks/{id}", "DELETE /tasks/{id}"})
	return r
}

func (srv *Server) Run(cfg config.Config) error {
	router := srv.setupRouter()
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	slog.Info("server starting", "host", cfg.Host, "port", cfg.Port, "address", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		slog.Error("server failed to start", "error", err, "address", addr)
		return fmt.Errorf("failed to start server on %s: %w", addr, err)
	}

	return nil
}
