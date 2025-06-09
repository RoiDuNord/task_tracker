package handlers

import (
	"context"
	"io_bound/models"
	"math/rand"
	"sync"
	"time"
)

type Server struct {
	tasks sync.Map
}

func (s *Server) runTask(ctx context.Context, task *models.Task) {
	startTime := time.Now()
	task.Status = models.StatusRunning
	task.StartTime = &startTime

	duration := time.Duration(3+rand.Intn(3)) * time.Minute

	select {
	case <-time.After(duration):
		markDone(task, startTime)
	case <-ctx.Done():
		markCancelled(task, startTime)
	}
}

func markDone(task *models.Task, startTime time.Time) {
	finished := time.Now()
	task.Status = models.StatusDone
	task.FinishTime = &finished
	task.Duration = float32(finished.Sub(startTime).Seconds())
	task.Result = "Task completed successfully"
}

func markCancelled(task *models.Task, startTime time.Time) {
	finished := time.Now()
	task.Status = models.StatusCancelled
	task.FinishTime = &finished
	task.Duration = float32(finished.Sub(startTime).Seconds())
	task.Error = "Task was cancelled"
}
