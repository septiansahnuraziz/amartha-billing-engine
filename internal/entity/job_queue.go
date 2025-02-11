package entity

import "context"

type (
	// Task name
	Task string

	// TaskQueue is used to enqueue the jobs
	TaskQueue interface {
		Enqueue(ctx context.Context, taskName Task, payload any) error
		Stop()
	}
)

func (t Task) String() string {
	return string(t)
}
