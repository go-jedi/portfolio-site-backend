package project

import (
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/repository"
)

const (
	tableName = "projects"

	//idColumn          = "id"
	titleColumn       = "title"
	descriptionColumn = "description"
	technologyColumn  = "technology"
	//createdAtColumn   = "created_at"
	//updatedAtColumn   = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ProjectRepository {
	return &repo{db: db}
}
