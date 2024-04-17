package repository

import "context"

type ProjectRepository interface {
	Create(ctx context.Context) error
}

type ReviewRepository interface {
	Create(ctx context.Context) error
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (string, error)
}
