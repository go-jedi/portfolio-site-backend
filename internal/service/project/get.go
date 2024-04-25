package project

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Get(ctx context.Context, page int, limit int) ([]project.Get, project.Params, error) {
	logger.Info(
		"(SERVICE PROJECT) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	params, err := s.projectRepository.Params(ctx)
	if err != nil {
		return nil, project.Params{}, err
	}

	result, err := s.projectRepository.Get(ctx, page, limit)
	if err != nil {
		return nil, project.Params{}, err
	}

	return result, params, nil
}
