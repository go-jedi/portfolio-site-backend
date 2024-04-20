package project

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Create(ctx context.Context, dto project.Create) (int, error) {
	logger.Info(
		"(REPOSITORY PROJECT) Create...",
		zap.String("title", dto.Title),
		zap.String("description", dto.Description),
		zap.String("technology", dto.Technology),
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, descriptionColumn, technologyColumn).
		Values(dto.Title, dto.Description, dto.Technology).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "project_repository.Create",
		QueryRaw: query,
	}

	var id int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
