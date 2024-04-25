package review

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Params(ctx context.Context) (review.Params, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Params...",
	)

	builder := sq.Select(
		"count(*) as page_count, 5 as limit",
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(
			sq.And{
				sq.Eq{isPublishColumn: true},
				sq.Eq{deletedColumn: false},
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return review.Params{}, err
	}

	q := db.Query{
		Name:     "review_repository.Params",
		QueryRaw: query,
	}

	var params review.Params
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&params.PageCount,
		&params.Limit,
	)
	if err != nil {
		return review.Params{}, err
	}

	return params, nil
}
