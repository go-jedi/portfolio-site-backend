package app

import (
	"context"

	"github.com/go-jedi/portfolio/internal/handler/image"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"

	imageRepository "github.com/go-jedi/portfolio/internal/repository/image"
	imageService "github.com/go-jedi/portfolio/internal/service/image"
)

func (s *serverProvider) ImageRepository(ctx context.Context) repository.ImageRepository {
	if s.imageRepository == nil {
		s.imageRepository = imageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.imageRepository
}

func (s *serverProvider) ImageService(ctx context.Context) service.ImageService {
	if s.imageService == nil {
		s.imageService = imageService.NewService(
			s.ImageRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.imageService
}

func (s *serverProvider) ImageHandler(ctx context.Context) *image.Handler {
	if s.imageHandler == nil {
		s.imageHandler = image.NewHandler(
			s.ImageService(ctx),
			s.validator,
		)
	}

	return s.imageHandler
}
