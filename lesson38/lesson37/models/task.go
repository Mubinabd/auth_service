package models

import (
	"time"
)

// TaskType represents the possible task types
type TaskType string

const (
	// TaskTodo represents a "todo" task type
	TaskTodo TaskType = "todo"
	// TaskDone represents a "done" task type
	TaskDone TaskType = "done"
)

// Task represents the "task" table
type Task struct {
	ID          string    `json:"id" db:"id" uri:"id"`
	Name        string    `json:"name" db:"name"`
	Type        TaskType  `json:"type" db:"type"`
	UserID      string    `json:"user_id" db:"user_id"`
	Description string    `json:"description" db:"description"`
	DeadLine    time.Time `json:"deadline" db:"deadline"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   int64     `json:"deleted_at" db:"deleted_at"`
}

type AllTasks struct {
	Count uint
	Tasks []Task
}

type TaskId struct {
	ID string `json:"id" db:"id" uri:"id"`
}
