package image

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
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
