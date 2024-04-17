package review

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Create(ctx context.Context) error {
	logger.Info(
		"(SERVICE REVIEW) Create...",
	)

	err := s.reviewRepository.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}
