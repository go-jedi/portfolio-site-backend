package repository

import (
	"context"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/model/review"
)

type ProjectRepository interface {
	Create(ctx context.Context, dto project.Create) (int, error)
	Get(ctx context.Context, page int, limit int) ([]project.Get, error)
	GetByID(ctx context.Context, id int) (project.Get, error)
	Update(ctx context.Context, dto project.Update) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type ImageRepository interface {
	Create(ctx context.Context, id int, paths []string) error
	Delete(ctx context.Context, id int) (int, error)
}

type ReviewRepository interface {
	Create(ctx context.Context, dto review.Create) (int, error)
	Get(ctx context.Context, page int, limit int) ([]review.Review, error)
	GetByID(ctx context.Context, id int) (review.Review, error)
	Publish(ctx context.Context, id int) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (string, error)
}
