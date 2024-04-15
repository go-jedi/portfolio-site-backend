package user

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Get(ctx context.Context, id int64) (string, error) {
	logger.Info(
		"(SERVICE) Get...",
		zap.Int64("id", id),
	)

	result, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return result, nil
}
