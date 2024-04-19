package review

import (
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/repository"
)

const (
	tableName = "reviews"

	idColumn        = "id"
	authorColumn    = "author"
	messageColumn   = "message"
	ratingColumn    = "rating"
	deletedColumn   = "deleted"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ReviewRepository {
	return &repo{db: db}
}
