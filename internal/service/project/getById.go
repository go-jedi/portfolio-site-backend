package project

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
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
