package domain

import "time"

type Statistics struct {
	TasksCreated             int
	TasksCompleted           int
	TasksCompletedRate       *float64
	TasksAverageCompleteTime *time.Duration
}

func NewStatistic(
	tasksCreatede int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompleteTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:             tasksCompleted,
		TasksCompleted:           tasksCompleted,
		TasksCompletedRate:       tasksCompletedRate,
		TasksAverageCompleteTime: tasksAverageCompleteTime,
	}
}
