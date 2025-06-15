package models

import (
	"context"
	"time"
)

const (
	StatusCreated   = "created"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	ResultCompleted = "task completed successfully"
)

type Task struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	CreationTime time.Time `json:"creation_time"`
	StartTime    time.Time `json:"start_time,omitzero"`
	FinishTime   time.Time `json:"finish_time,omitzero"`
	Duration     float32   `json:"duration(sec),omitempty"`
	Result       string    `json:"result,omitempty"`

	CancelFunc context.CancelFunc `json:"-"`
}
