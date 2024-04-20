package project

import (
	"context"

	"go.uber.org/zap"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (r *repo) GetByID(ctx context.Context, id int) (project.Get, error) {
	logger.Info(
		"(REPOSITORY PROJECT) GetByID...",
		zap.Int("id", id),
	)

	queryPaths := `
		COALESCE((
			SELECT json_agg(ag.*)::JSONB
			FROM (
				SELECT id, project_id, path_file, created_at, updated_at
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
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(
			sq.And{
				sq.Eq{idColumn: id},
				sq.Eq{deletedColumn: false},
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return project.Get{}, err
	}

	q := db.Query{
		Name:     "project_repository.GetByID",
		QueryRaw: query,
	}

	var proj project.Get
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&proj.ID,
		&proj.Title,
		&proj.Description,
		&proj.Technology,
		&proj.CreatedAt,
		&proj.UpdatedAt,
		&proj.Paths,
	)
	if err != nil {
		return project.Get{}, err
	}

	return proj, nil
}
