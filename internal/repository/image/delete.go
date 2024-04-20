package image

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Delete(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(REPOSITORY IMAGE) Delete...",
		zap.Int("id", id),
	)

	return 0, nil
}
