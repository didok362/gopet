package users_service

import (
	"context"
	"fmt"
	"gopet/internal/core/domain"
	core_errors "gopet/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
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

	users, err := s.usersRepository.GetUsers(
		ctx,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get users from repostiry: %w", err)
	}

	return users, nil
}
