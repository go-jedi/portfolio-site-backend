package review

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Get(ctx context.Context, page int, limit int) ([]review.Review, review.Params, error) {
	logger.Info(
		"(SERVICE REVIEW) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	params, err := s.reviewRepository.Params(ctx)
	if err != nil {
		return nil, review.Params{}, err
	}

	result, err := s.reviewRepository.Get(ctx, page, limit)
	if err != nil {
		return nil, review.Params{}, err
	}

	return result, params, nil
}
