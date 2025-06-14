package models

import (
	"context"
	"time"
)

type TaskStatus string

const (
	StatusCreated   TaskStatus = "created"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Status       TaskStatus `json:"status"`
	CreationTime time.Time  `json:"creation_time"`
	StartTime    time.Time  `json:"start_time,omitzero"`
	FinishTime   time.Time  `json:"finish_time,omitzero"`
	Duration     float32    `json:"duration(sec),omitempty"`
	Result       string     `json:"result,omitempty"`

	CancelFunc context.CancelFunc `json:"-"`
}
