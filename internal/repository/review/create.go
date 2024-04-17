package review

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
)

//const (
//	tableName       = "reviews"
//	idColumn        = "id"
//	authorColumn    = "author"
//	messageColumn   = "message"
//	ratingColumn    = "rating"
//	createdAtColumn = "created_at"
//	updatedAtColumn = "updated_at"
//)

func (r *repo) Create(_ context.Context) error {
	logger.Info(
		"(REPOSITORY REVIEW) Create...",
	)

	return nil
}
