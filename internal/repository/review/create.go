package review

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Create(ctx context.Context, dto review.Create) (int, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Create...",
		zap.String("username", dto.Username),
		zap.String("message", dto.Message),
		zap.Int("rating", dto.Rating),
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernameColumn, messageColumn, ratingColumn).
		Values(dto.Username, dto.Message, dto.Rating).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "review_repository.Create",
		QueryRaw: query,
	}

	var id int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
