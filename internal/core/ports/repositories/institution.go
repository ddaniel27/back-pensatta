package repositories

import (
	"context"
	"pensatta/internal/core/domain"
)

type InstitutionRepository interface {
	Create(ctx context.Context, institution domain.Institution) error
	Get(ctx context.Context) ([]domain.Institution, error)
	Delete(ctx context.Context, id uint64) error
}
