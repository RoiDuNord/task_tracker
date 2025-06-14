package main

import (
	"io_bound/config"
	"io_bound/server"
	"io_bound/storage"
	"io_bound/task"
	"log/slog"
)

func main() {
	cfg := config.DefaultConfig()

	tasks := storage.NewTasks()
	taskManager := task.NewTaskManager(tasks)
	srv := server.NewServer(taskManager)

	if err := srv.Run(cfg); err != nil {
		return
	}

	slog.Info("Server stopped gracefully")
}
