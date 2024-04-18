package review

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

const (
	tableName = "reviews"

	//idColumn        = "id"
	authorColumn  = "author"
	messageColumn = "message"
	ratingColumn  = "rating"
	//createdAtColumn = "created_at"
	//updatedAtColumn = "updated_at"
)

func (r *repo) Create(ctx context.Context, dto review.Create) (int, error) {
	logger.Info(
		"(REPOSITORY REVIEW) Create...",
		zap.String("author", dto.Author),
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(authorColumn, messageColumn, ratingColumn).
		Values(dto.Author, dto.Message, dto.Rating).
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
