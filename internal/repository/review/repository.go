package review

import (
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ReviewRepository {
	return &repo{db: db}
}
