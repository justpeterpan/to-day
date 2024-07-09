package main

import (
	"testing"
	"time"
)

func TestTaskStatus(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		task     Task
		expected string
	}{
		{
			name: "Open task",
			task: Task{
				CreatedAt: now,
			},
			expected: "open",
		},
		{
			name: "Done task",
			task: Task{
				CreatedAt:  now,
				FinishedAt: &now,
			},
			expected: "done",
		},
		{
			name: "Deleted task",
			task: Task{
				CreatedAt: now,
				DeletedAt: &now,
			},
			expected: "deleted",
		},
		{
			name: "Deleted takes precedence over done",
			task: Task{
				CreatedAt:  now,
				FinishedAt: &now,
				DeletedAt:  &now,
			},
			expected: "deleted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.Status(); got != tt.expected {
				t.Errorf("Task.Status() = %v, want %v", got, tt.expected)
			}
		})
	}
}
