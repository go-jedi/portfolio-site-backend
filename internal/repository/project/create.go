package project

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
)

//const (
//	tableName = "projects"
//
//	idColumn          = "id"
//	titleColumn       = "title"
//	descriptionColumn = "description"
//	technologyColumn  = "technology"
//	createdAtColumn   = "created_at"
//	updatedAtColumn   = "updated_at"
//)

func (r *repo) Create(_ context.Context) error {
	logger.Info(
		"(REPOSITORY PROJECT) Create...",
	)

	return nil
}
