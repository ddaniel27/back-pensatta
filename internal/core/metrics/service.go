package metrics

import (
	"context"
	"fmt"
	"math"
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/repositories"
)

type Service struct {
	mr repositories.MetricsRepository
}

func NewService(mr repositories.MetricsRepository) *Service {
	return &Service{
		mr: mr,
	}
}

func (s *Service) GetMetricsByUserID(ctx context.Context, userID uint64) (domain.Metrics, error) {
	scores, dimensions, err := s.mr.GetByUserID(ctx, userID)
	if err != nil {
		return domain.Metrics{}, err
	}

	return calculateMetrics(scores, dimensions), nil
}

func calculateMetrics(scores []uint64, dimensions []uint64) domain.Metrics {
	sv := make(map[string]interface{})
	av := map[string]interface{}{
		"1": uint64(0),
		"2": uint64(0),
		"3": uint64(0),
	}
	dc := map[string]uint64{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
	}

	result := domain.Metrics{
		SpiderValues:       sv,
		AppropiationValues: av,
	}

	for i := 0; i < len(scores); i++ {
		dim, score := dimensions[i], scores[i]
		dimStr := fmt.Sprintf("%d", dim)

		switch {
		case score < 60:
			av["1"] = av["1"].(uint64) + 1
		case score < 80:
			av["2"] = av["2"].(uint64) + 1
		default:
			av["3"] = av["3"].(uint64) + 1
		}

		if _, ok := sv[dimStr]; !ok {
			sv[dimStr] = float64(score)
			dc[dimStr] = 1
			continue
		}

		currentSum := (sv[dimStr].(float64) * float64(dc[dimStr])) + float64(score)
		dc[dimStr] = dc[dimStr] + 1
		sv[dimStr] = currentSum / float64(dc[dimStr])
		sv[dimStr] = math.Round(sv[dimStr].(float64)*100) / 100
	}

	return result
}
