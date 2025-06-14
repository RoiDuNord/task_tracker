package main

import (
	"io_bound/server"
	"io_bound/storage"
	"io_bound/task"
	"log/slog"
)

func main() {
	tasks := storage.NewTasks()
	taskManager := task.NewTaskManager(tasks)
	srv := server.NewServer(taskManager)

	if err := srv.Run(); err != nil {
		return
	}

	slog.Info("Server stopped gracefully")
}
