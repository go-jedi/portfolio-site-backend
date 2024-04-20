package project

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Get(ctx context.Context, page int, limit int) ([]project.Get, error) {
	logger.Info(
		"(SERVICE PROJECT) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	result, err := s.projectRepository.Get(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}
