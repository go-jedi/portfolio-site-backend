package image

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Create(ctx context.Context, id int, filenames []string) error {
	logger.Info(
		"(REPOSITORY IMAGE) Create...",
		zap.Int("id", id),
		zap.Strings("filenames", filenames),
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(projectIDColumn, filenameColumn)

	for _, file := range filenames {
		builder = builder.Values(id, file)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "image_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
