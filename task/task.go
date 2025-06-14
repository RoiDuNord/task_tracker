package task

import (
	"context"
	"io_bound/models"
	"log/slog"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type TaskStorage interface {
	Create(t models.Task) error
	GetByID(id string) (models.Task, error)
	Delete(id string) error
	Update(t models.Task) error
}

type TaskManager struct {
	storage TaskStorage
}

func NewTaskManager(storage TaskStorage) *TaskManager {
	return &TaskManager{
		storage: storage,
	}
}

func (tm *TaskManager) DeleteTask(id string) error {
	if err := tm.storage.Delete(id); err != nil {
		return err
	}

	return nil
}

func (tm *TaskManager) GetTaskByID(id string) (models.Task, error) {
	task, err := tm.storage.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (tm *TaskManager) CreateNewTask(taskName string) (models.Task, error) {
	id := uuid.NewString()
	creationTime := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	t := models.Task{
		ID:           id,
		Name:         taskName,
		Status:       models.StatusCreated,
		CreationTime: creationTime,
		CancelFunc:   cancel,
	}

	if err := tm.storage.Create(t); err != nil {
		return models.Task{}, err
	}

	go func() {
		startTime := time.Now()
		duration := time.Duration(3+rand.Intn(3)) * time.Minute
		timer := time.After(duration)

		t.Status = models.StatusRunning
		t.StartTime = startTime
		if err := tm.storage.Update(t); err != nil {
			slog.Info("task update error", "id", t.ID)
			return
		}
		slog.Info("task running", "id", t.ID, "start_time", t.StartTime)

		select {
		case <-timer:
			t = markCompleted(t)
			if err := tm.storage.Update(t); err != nil {
				slog.Info("task update error", "id", t.ID, "duration", t.Duration)
				return
			}
			slog.Info("task completed", "id", t.ID, "duration", t.Duration)

		case <-ctx.Done():
			tm.storage.Delete(t.ID)
			slog.Info("task deleted", "id", t.ID, "duration", t.Duration)
		}
	}()

	return t, nil
}

func markCompleted(task models.Task) models.Task {
	finished := time.Now()
	task.Status = models.StatusCompleted
	task.FinishTime = finished
	task.Duration = float32(finished.Sub(task.CreationTime).Seconds())
	task.Result = "Task completed successfully"
	return task
}
