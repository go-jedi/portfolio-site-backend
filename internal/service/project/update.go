package project

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Update(ctx context.Context, dto project.Update) (int, error) {
	logger.Info(
		"(SERVICE PROJECT) Update...",
		zap.Int("id", dto.ID),
		zap.String("title", dto.Title),
		zap.String("description", dto.Description),
		zap.String("technology", dto.Technology),
	)

	result, err := s.projectRepository.Update(ctx, dto)
	if err != nil {
		return 0, err
	}

	return result, nil
}
