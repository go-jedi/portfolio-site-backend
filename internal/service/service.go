package service

import (
	"context"
	"mime/multipart"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/model/review"
)

type ProjectService interface {
	Create(ctx context.Context, dto project.Create, files []*multipart.FileHeader) error
	Get(ctx context.Context, page int, limit int) ([]project.Get, project.Params, error)
	GetByID(ctx context.Context, id int) (project.Get, error)
	Update(ctx context.Context, dto project.Update) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type ImageService interface {
	Delete(ctx context.Context, id int) (int, error)
}

type ReviewService interface {
	Create(ctx context.Context, dto review.Create) (int, error)
	Get(ctx context.Context, page int, limit int) ([]review.Review, review.Params, error)
	GetByID(ctx context.Context, id int) (review.Review, error)
	Publish(ctx context.Context, id int) (int, error)
	UnPublish(ctx context.Context, id int) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type UserService interface {
	Get(ctx context.Context, id int64) (string, error)
}
