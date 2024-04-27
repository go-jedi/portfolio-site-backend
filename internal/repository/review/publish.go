package review

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (r *repo) Publish(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Publish...",
		zap.Int("id", id),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(isPublishColumn, true).
		Where(
			sq.And{
				sq.Eq{idColumn: id},
				sq.Eq{isPublishColumn: false},
				sq.Eq{deletedColumn: false},
			},
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "review_repository.Publish",
		QueryRaw: query,
	}

	var publishedID int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&publishedID)
	if err != nil {
		return 0, err
	}

	return publishedID, nil
}
