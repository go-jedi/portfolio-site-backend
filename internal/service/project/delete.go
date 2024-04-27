package project

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
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
