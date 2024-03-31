package services

import (
	"context"
	"pensatta/internal/core/domain"
)

type InstitutionService interface {
	CreateInstitution(ctx context.Context, institution domain.Institution) error
	GetInstitutions(ctx context.Context) ([]domain.Institution, error)
	DeleteInstitution(ctx context.Context, id uint64) error
}
