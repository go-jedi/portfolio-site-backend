package service

import "context"

type ProjectService interface {
	Create(ctx context.Context) error
}

type ReviewService interface {
	Create(ctx context.Context) error
}

type UserService interface {
	Get(ctx context.Context, id int64) (string, error)
}
