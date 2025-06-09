package server

import (
	"io_bound/handlers"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func Run() {
	srv := new()
	r := chi.NewRouter()

	r.Post("/tasks", srv.Create)
	r.Get("/tasks/{id}", srv.Get)
	r.Delete("/tasks/{id}", srv.Delete)

	srvAddr := ":8080"
	slog.Info("server started", "address", srvAddr)
	if err := http.ListenAndServe(srvAddr, r); err != nil {
		log.Fatalf("Server error: %v", err)
		return
	}
}

func new() *handlers.Server {
	slog.Info("server created")
	return &handlers.Server{}
}
