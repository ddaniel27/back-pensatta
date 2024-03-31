package institution

import (
	"context"
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/repositories"
)

type Service struct {
	ir repositories.InstitutionRepository
}

func NewService(ir repositories.InstitutionRepository) *Service {
	return &Service{
		ir: ir,
	}
}

func (s *Service) CreateInstitution(ctx context.Context, institution domain.Institution) error {
	if !validateLanguage(institution.Language) {
		return domain.ErrInvalidLanguage
	}

	return s.ir.Create(ctx, institution)
}

func (s *Service) GetInstitutions(ctx context.Context) ([]domain.Institution, error) {
	return s.ir.Get(ctx)
}

func (s *Service) DeleteInstitution(ctx context.Context, id uint64) error {
	return s.ir.Delete(ctx, id)
}

func validateLanguage(language string) bool {
	switch language {
	case "es", "en", "pt":
		return true
	default:
		return false
	}
}
