package project

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) Get(ctx context.Context, page int, limit int) ([]project.Get, error) {
	logger.Info(
		"(REPOSITORY PROJECT) Get...",
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	const queryPaths = `
		COALESCE((
			SELECT json_agg(ag.*)::JSONB
			FROM (
				SELECT id, project_id, filename, created_at, updated_at
				FROM images
				WHERE project_id = projects.id
				AND deleted = FALSE
			) ag
		), '[]') as paths
	`

	builder := sq.Select(
		idColumn,
		titleColumn,
		descriptionColumn,
		technologyColumn,
		createdAtColumn,
		updatedAtColumn,
		queryPaths,
	).PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(
			sq.Eq{deletedColumn: false},
		).
		OrderBy(fmt.Sprintf("%s DESC", idColumn)).
		Offset(uint64(limit * (page - 1))).
		Limit(uint64(limit))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "project_repository.Get",
		QueryRaw: query,
	}

	var projects []project.Get
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := project.Get{}

		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Technology, &p.CreatedAt, &p.UpdatedAt, &p.Paths)
		if err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}

	return projects, nil
}
