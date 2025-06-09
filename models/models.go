package models

import (
	"context"
	"time"
)

type TaskStatus string

const (
	StatusCreated   TaskStatus = "created"
	StatusRunning   TaskStatus = "running"
	StatusDone      TaskStatus = "done"
	StatusCancelled TaskStatus = "cancelled"
)

type Task struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Status       TaskStatus `json:"status"`
	CreationTime time.Time  `json:"creation_time"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	FinishTime   *time.Time `json:"finish_time,omitempty"`
	Duration     float32    `json:"duration(sec),omitempty"`
	Result       string     `json:"result,omitempty"`
	Error        string     `json:"error,omitempty"`

	CancelFunc context.CancelFunc `json:"-"`
}
