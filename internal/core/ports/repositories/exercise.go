package repositories

import (
	"context"
	"pensatta/internal/core/domain"
)

type ExerciseRepository interface {
	Create(ctx context.Context, exercise domain.Exercise, userID uint64) error
	GetByUserID(ctx context.Context, userID uint64) ([]domain.Exercise, error)
	GetByUserIDAndExerciseID(ctx context.Context, userID, exerciseID uint64) ([]domain.Exercise, error)
}
