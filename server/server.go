package server

import (
	"fmt"
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

func (srv *Server) Run() error {
	r := chi.NewRouter()

	r.Post("/tasks", srv.Create)
	r.Get("/tasks/{id}", srv.Get)
	r.Delete("/tasks/{id}", srv.Delete)

	srvPort := "8080"
	srvAddr := "localhost:" + srvPort

	slog.Info("server started", "address", srvAddr)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", srvPort), r); err != nil {
		slog.Error("Server failed", "error", err)
		return err
	}

	return nil
}
