package service

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/review"
)

type ProjectService interface {
	Create(ctx context.Context) error
}

type ReviewService interface {
	Create(ctx context.Context, dto review.Create) (int, error)
	Get(ctx context.Context) ([]review.Review, error)
	GetByID(ctx context.Context, id int) (review.Review, error)
}

type UserService interface {
	Get(ctx context.Context, id int64) (string, error)
}
