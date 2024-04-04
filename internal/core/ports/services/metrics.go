package services

import (
	"context"
	"pensatta/internal/core/domain"
)

type MetricsService interface {
	GetMetricsByUserID(ctx context.Context, userID uint64) (domain.Metrics, error)
}
