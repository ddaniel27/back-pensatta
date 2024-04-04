package postgres

import (
	"context"

	"gorm.io/gorm"
)

type MetricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) *MetricsRepository {
	return &MetricsRepository{
		db: db,
	}
}

type result struct {
	Score     uint64 `gorm:"column:score"`
	Dimension uint64 `gorm:"column:dimension"`
}

func (mr *MetricsRepository) GetByUserID(ctx context.Context, userID uint64) ([]uint64, []uint64, error) {
	var results []result

	if err := mr.db.WithContext(ctx).Raw(`
	SELECT A.score, B.dimension
	FROM pensatta_record A
	INNER JOIN pensatta_exercise B
	ON A.exercise_id = B.id
	WHERE A.user_id = $1`, userID).Scan(&results).Error; err != nil {
		return nil, nil, err
	}

	r1, r2 := toSlices(results)

	return r1, r2, nil
}

func toSlices(r []result) ([]uint64, []uint64) {
	scores := make([]uint64, 0)
	dimensions := make([]uint64, 0)

	for _, v := range r {
		scores = append(scores, v.Score)
		dimensions = append(dimensions, v.Dimension)
	}

	return scores, dimensions
}
