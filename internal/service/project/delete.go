package project

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Delete(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(SERVICE PROJECT) Delete...",
		zap.Int("id", id),
	)

	result, err := s.projectRepository.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
