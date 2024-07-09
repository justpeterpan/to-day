package main

import (
	"time"
)

type TaskList struct {
	Version      string    `json:"version"`
	LastModified time.Time `json:"lastModified"`
	Tasks        []Task    `json:"tasks"`
}

type Task struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	CreatedAt  time.Time  `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

func (t *Task) Status() string {
	if t.DeletedAt != nil {
		return "deleted"
	}
	if t.FinishedAt != nil {
		return "done"
	}

	return "open"
}
