package project

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Params(ctx context.Context) (project.Params, error) {
	logger.Info(
		"(REPOSITORY PROJECT) Params...",
	)

	builder := sq.Select(
		"count(*) as page_count, 5 as limit",
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(
			sq.Eq{deletedColumn: false},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return project.Params{}, err
	}

	q := db.Query{
		Name:     "project_repository.Params",
		QueryRaw: query,
	}

	var params project.Params
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&params.PageCount,
		&params.Limit,
	)
	if err != nil {
		return project.Params{}, err
	}

	return params, nil
}
