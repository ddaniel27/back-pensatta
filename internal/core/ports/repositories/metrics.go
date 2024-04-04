package repositories

import "context"

type MetricsRepository interface {
	GetByUserID(ctx context.Context, userID uint64) ([]uint64, []uint64, error)
}
