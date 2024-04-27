package user

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Get(ctx context.Context, id int64) (string, error) {
	logger.Info(
		"(SERVICE USER) Get...",
		zap.Int64("id", id),
	)

	result, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return result, nil
}
