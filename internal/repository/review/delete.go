package review

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Delete(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Delete...",
		zap.Int("id", id),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(deletedColumn, true).
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
		Name:     "review_repository.Delete",
		QueryRaw: query,
	}

	var deletedId int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&deletedId)
	if err != nil {
		return 0, err
	}

	return deletedId, nil
}
