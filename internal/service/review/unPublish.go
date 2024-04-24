package review

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) UnPublish(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(SERVICE REVIEW) UnPublish...",
		zap.Int("id", id),
	)

	result, err := s.reviewRepository.UnPublish(ctx, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
