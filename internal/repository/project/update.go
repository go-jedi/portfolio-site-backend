package project

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
	"go.uber.org/zap"
)

func (r *repo) Update(ctx context.Context, dto project.Update) (int, error) {
	logger.Info(
		"(REPOSITORY PROJECT) Update...",
		zap.Int("id", dto.ID),
		zap.String("title", dto.Title),
		zap.String("description", dto.Description),
		zap.String("technology", dto.Technology),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(titleColumn, dto.Title).
		Set(descriptionColumn, dto.Description).
		Set(technologyColumn, dto.Technology).
		Where(
			sq.And{
				sq.Eq{idColumn: dto.ID},
				sq.Eq{deletedColumn: false},
			},
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "project_repository.Update",
		QueryRaw: query,
	}

	var updatedID int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&updatedID)
	if err != nil {
		return 0, err
	}

	return updatedID, nil
}
