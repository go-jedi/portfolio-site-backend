package project

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (r *repo) Delete(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(REPOSITORY PROJECT) Delete...",
		zap.Int("id", id),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(deletedColumn, true).
		Set(updatedAtColumn, "NOW()").
		Where(
			sq.And{
				sq.Eq{idColumn: id},
				sq.Eq{deletedColumn: false},
			},
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "project_repository.Delete",
		QueryRaw: query,
	}

	var deletedID int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&deletedID)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}
