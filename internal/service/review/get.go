package review

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Get(ctx context.Context, page int, limit int) ([]review.Review, error) {
	logger.Info(
		"(SERVICE REVIEW) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	result, err := s.reviewRepository.Get(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}
