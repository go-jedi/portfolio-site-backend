package image

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Delete(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(SERVICE IMAGE) Delete...",
		zap.Int("id", id),
	)

	result, err := s.imageRepository.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
