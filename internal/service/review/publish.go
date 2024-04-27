package review

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Publish(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(SERVICE REVIEW) Publish...",
		zap.Int("id", id),
	)

	result, err := s.reviewRepository.Publish(ctx, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
