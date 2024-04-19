package repository

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/review"
)

type ProjectRepository interface {
	Create(ctx context.Context) error
}

type ReviewRepository interface {
	Create(ctx context.Context, dto review.Create) (int, error)
	Get(ctx context.Context) ([]review.Review, error)
	GetByID(ctx context.Context, id int) (review.Review, error)
	Delete(ctx context.Context, id int) (int, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (string, error)
}
