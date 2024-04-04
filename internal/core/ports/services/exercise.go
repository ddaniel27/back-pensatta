package services

import (
	"context"
	"pensatta/internal/core/domain"
)

type ExerciseService interface {
	CreateExercise(ctx context.Context, exercise domain.Exercise, userID uint64) error
	GetExercisesByUserID(ctx context.Context, userID uint64) ([]domain.Exercise, error)
	GetExercisesByUserIDAndExerciseID(ctx context.Context, userID, exerciseID uint64) ([]domain.Exercise, error)
}
