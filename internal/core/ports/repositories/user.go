package repositories

import (
	"context"
	"pensatta/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	GetByID(ctx context.Context, id uint64) (domain.User, error)
	Get(ctx context.Context) ([]domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetByRole(ctx context.Context, role string) ([]domain.User, error)
	Update(ctx context.Context, u domain.User) error
	Delete(ctx context.Context, id uint64) error
}
