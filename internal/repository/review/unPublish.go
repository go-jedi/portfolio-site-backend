package review

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (r *repo) UnPublish(ctx context.Context, id int) (int, error) {
	logger.Info(
		"(REPOSITORY REVIEW) UnPublish...",
		zap.Int("id", id),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(isPublishColumn, false).
		Where(
			sq.And{
				sq.Eq{idColumn: id},
				sq.Eq{isPublishColumn: true},
				sq.Eq{deletedColumn: false},
			},
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "review_repository.UnPublish",
		QueryRaw: query,
	}

	var unPublishedID int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&unPublishedID)
	if err != nil {
		return 0, err
	}

	return unPublishedID, nil
}
