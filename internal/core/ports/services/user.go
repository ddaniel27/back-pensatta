package services

import (
	"context"
	"pensatta/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, u domain.User) (string, error)
	GetUser(ctx context.Context, id uint64) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByRole(ctx context.Context, role string) ([]domain.User, error)
	UpdateUser(ctx context.Context, u domain.User) error
	DeleteUser(ctx context.Context, id uint64) error
}
