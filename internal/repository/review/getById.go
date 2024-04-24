package review

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) GetByID(ctx context.Context, id int) (review.Review, error) {
	logger.Info(
		"(REPOSITORY REVIEW) GetByID...",
		zap.Int("id", id),
	)

	builder := sq.Select(
		idColumn,
		usernameColumn,
		messageColumn,
		ratingColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(
			sq.And{
				sq.Eq{idColumn: id},
				sq.Eq{isPublishColumn: true},
				sq.Eq{deletedColumn: false},
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return review.Review{}, err
	}

	q := db.Query{
		Name:     "review_repository.GetByID",
		QueryRaw: query,
	}

	var rvw review.Review
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&rvw.ID,
		&rvw.Username,
		&rvw.Message,
		&rvw.Rating,
		&rvw.CreatedAt,
		&rvw.UpdatedAt,
	)
	if err != nil {
		return review.Review{}, err
	}

	return rvw, nil
}
