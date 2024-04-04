package exercise

import (
	"context"
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/repositories"
)

type Service struct {
	er repositories.ExerciseRepository
}

func NewService(er repositories.ExerciseRepository) *Service {
	return &Service{
		er: er,
	}
}

func (s *Service) CreateExercise(ctx context.Context, exercise domain.Exercise, userID uint64) error {
	return s.er.Create(ctx, exercise, userID)
}

func (s *Service) GetExercisesByUserID(ctx context.Context, userID uint64) ([]domain.Exercise, error) {
	return s.er.GetByUserID(ctx, userID)
}

func (s *Service) GetExercisesByUserIDAndExerciseID(ctx context.Context, userID, exerciseID uint64) ([]domain.Exercise, error) {
	return s.er.GetByUserIDAndExerciseID(ctx, userID, exerciseID)
}
