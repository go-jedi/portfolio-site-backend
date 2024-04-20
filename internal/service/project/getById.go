package project

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) GetByID(ctx context.Context, id int) (project.Get, error) {
	logger.Info(
		"(SERVICE PROJECT) GetByID...",
		zap.Int("id", id),
	)

	result, err := s.projectRepository.GetByID(ctx, id)
	if err != nil {
		return project.Get{}, err
	}

	return result, nil
}
