package review

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"

	"github.com/go-jedi/portfolio/internal/model/review"
)

func (s *serv) Get(ctx context.Context) ([]review.Review, error) {
	logger.Info(
		"(SERVICE REVIEW) Create...",
	)

	result, err := s.reviewRepository.Get(ctx)
	if err != nil {
		return []review.Review{}, err
	}

	return result, nil
}
