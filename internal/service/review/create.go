package review

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Create(ctx context.Context, dto review.Create) (int, error) {
	logger.Info(
		"(SERVICE REVIEW) Create...",
	)

	result, err := s.reviewRepository.Create(ctx, dto)
	if err != nil {
		return 0, err
	}

	return result, nil
}
