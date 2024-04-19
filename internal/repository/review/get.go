package review

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Get(ctx context.Context, page int, limit int) ([]review.Review, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	builder := sq.Select(
		idColumn,
		authorColumn,
		messageColumn,
		ratingColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{deletedColumn: false}).
		OrderBy(idColumn).
		Offset(uint64(limit * (page - 1))).
		Limit(uint64(limit))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "review_repository.Get",
		QueryRaw: query,
	}

	var reviews []review.Review
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := review.Review{}

		err := rows.Scan(&r.ID, &r.Author, &r.Message, &r.Rating, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, r)
	}

	rows.Close()

	return reviews, nil
}
