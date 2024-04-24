package image

import (
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/internal/repository"
)

const (
	tableName = "images"

	idColumn        = "id"
	projectIDColumn = "project_id"
	filenameColumn  = "filename"
	deletedColumn   = "deleted"
	//createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ImageRepository {
	return &repo{db: db}
}
