package server

import (
	"io_bound/handlers"
	"log/slog"
)

func New() *handlers.Server {
	slog.Info("server created")
	return &handlers.Server{}
}
