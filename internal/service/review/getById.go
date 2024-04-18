package review

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) GetByID(ctx context.Context, id int) (review.Review, error) {
	logger.Info(
		"(SERVICE REVIEW) GetByID...",
		zap.Int("id", id),
	)

	result, err := s.reviewRepository.GetByID(ctx, id)
	if err != nil {
		return review.Review{}, err
	}

	return result, nil
}
