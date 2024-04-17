package project

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Create(ctx context.Context) error {
	logger.Info(
		"(SERVICE PROJECT) Create...",
	)

	err := s.projectRepository.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}
