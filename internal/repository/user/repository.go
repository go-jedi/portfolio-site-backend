package user

import (
	"context"

	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(_ context.Context, id int64) (string, error) {
	logger.Info(
		"(REPOSITORY USER) Get...",
		zap.Int64("id", id),
	)

	return "", nil
}
