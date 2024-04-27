package review

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Create(ctx context.Context, dto review.Create) (int, error) {
	logger.Info(
		"(SERVICE REVIEW) Create...",
		zap.String("username", dto.Username),
		zap.String("message", dto.Message),
		zap.Int("rating", dto.Rating),
	)

	result, err := s.reviewRepository.Create(ctx, dto)
	if err != nil {
		return 0, err
	}

	return result, nil
}
