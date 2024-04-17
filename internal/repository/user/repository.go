package user

import (
	"context"

	"github.com/go-jedi/platform_common/pkg/db"
	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/pkg/logger"
)

//const (
//	tableName = "users"
//
//	idColumn        = "id"
//	nameColumn      = "name"
//	createdAtColumn = "created_at"
//	updatedAtColumn = "updated_at"
//)

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
