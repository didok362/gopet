package users_service

import (
	"context"
	"gopet/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
}

func NewUsersService(
	UsersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: UsersRepository,
	}
}
