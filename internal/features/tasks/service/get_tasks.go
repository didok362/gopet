package tasks_service

import (
	"context"
	"fmt"
	"gopet/internal/core/domain"
	core_errors "gopet/internal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must not be negative: %w",
			core_errors.ErrInvalidArgumnet,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must not be negative: %w",
			core_errors.ErrInvalidArgumnet,
		)
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("falied to get tasks from repo: %w", err)
	}

	return tasks, nil
}
