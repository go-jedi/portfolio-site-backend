package app

import (
	"context"

	"github.com/go-jedi/portfolio/internal/repository"

	imageRepository "github.com/go-jedi/portfolio/internal/repository/image"
)

func (s *serverProvider) ImageRepository(ctx context.Context) repository.ImageRepository {
	if s.imageRepository == nil {
		s.imageRepository = imageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.imageRepository
}
