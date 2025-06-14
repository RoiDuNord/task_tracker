package storage

import (
	"fmt"
	"io_bound/models"
	"log/slog"
	"sync"
	"time"
)

type Tasks struct {
	mu      sync.RWMutex
	taskMap map[string]models.Task
}

func NewTasks() *Tasks {
	return &Tasks{
		taskMap: make(map[string]models.Task),
	}
}

func (ts *Tasks) Create(t models.Task) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, ok := ts.taskMap[t.ID]; ok {
		err := fmt.Errorf("task with ID %s already exists", t.ID)
		slog.Error("failed to create task", "error", err, "id", t.ID)
		return err
	}

	ts.taskMap[t.ID] = t
	slog.Info("task created", "id", t.ID, "creation_time", t.CreationTime, "status", t.Status)

	return nil
}

func (ts *Tasks) GetByID(id string) (models.Task, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	task, ok := ts.taskMap[id]
	if !ok {
		err := fmt.Errorf("task ID %s not found", id)
		slog.Error("failed to get task", "error", err, "id", id)
		return models.Task{}, err
	}

	if task.Status == models.StatusRunning {
		task.Duration = float32(time.Since(task.CreationTime).Seconds())
	}

	slog.Info("task found", "id", id, "duration", task.Duration)
	return task, nil
}

func (ts *Tasks) Update(t models.Task) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, ok := ts.taskMap[t.ID]; !ok {
		err := fmt.Errorf("task with ID %s not existed", t.ID)
		slog.Error("failed to update task", "error", err, "id", t.ID)
		return err
	}

	ts.taskMap[t.ID] = t
	slog.Info("task updated", "id", t.ID)

	return nil
}

func (ts *Tasks) Delete(id string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, exists := ts.taskMap[id]; !exists {
		err := fmt.Errorf("task with ID %s not found", id)
		slog.Error("failed to delete task", "error", err, "id", id)
		return err
	}

	if ts.taskMap[id].Status == models.StatusRunning && ts.taskMap[id].CancelFunc != nil {
		ts.taskMap[id].CancelFunc()
		slog.Info("task cancelled before deletion", "id", id)
	}

	delete(ts.taskMap, id)
	slog.Info("task deleted", "id", id)

	return nil
}
