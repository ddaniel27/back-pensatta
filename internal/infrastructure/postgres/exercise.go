package postgres

import (
	"context"
	"pensatta/internal/core/domain"
	"pensatta/internal/infrastructure/postgres/models"
	"time"

	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{
		db: db,
	}
}

func (er *ExerciseRepository) Create(ctx context.Context, exercise domain.Exercise, userID uint64) error {
	var averages models.CalificationModel
	var totalRecords int64
	record := toRecordModel(exercise, userID)
	if err := er.db.WithContext(ctx).Create(&record).Error; err != nil {
		return err
	}

	if err := er.db.WithContext(ctx).Where("user_id = ?", userID).First(&averages).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}

		averages = models.CalificationModel{
			UserID:       userID,
			AverageScore: uint64(exercise.Score),
			AverageTime:  exercise.Time,
		}

		return er.db.WithContext(ctx).Create(&averages).Error
	}

	if err := er.db.
		WithContext(ctx).
		Model(&models.RecordModel{}).
		Where("user_id = ?", userID).
		Count(&totalRecords).Error; err != nil {
		return err
	}

	averages.NewAverages(exercise.Score, exercise.Time, totalRecords)

	return er.db.WithContext(ctx).Save(&averages).Error
}

func (er *ExerciseRepository) GetByUserID(ctx context.Context, userID uint64) ([]domain.Exercise, error) {
	var exercises []domain.Exercise
	err := er.db.Where("user_id = ?", userID).Find(&exercises).Error
	return exercises, err
}

func (er *ExerciseRepository) GetByUserIDAndExerciseID(
	ctx context.Context,
	userID, exerciseID uint64,
) ([]domain.Exercise, error) {
	var exercises []domain.Exercise
	err := er.db.Where("user_id = ? AND exercise_id = ?", userID, exerciseID).Find(&exercises).Error
	return exercises, err
}

func toRecordModel(exercise domain.Exercise, userID uint64) models.RecordModel {
	return models.RecordModel{
		UserID:     userID,
		ExerciseID: exercise.ID,
		Score:      uint64(exercise.Score),
		Time:       exercise.Time,
		Date:       time.Now(),
	}
}
