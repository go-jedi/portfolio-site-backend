package app

import (
	"context"

	"github.com/go-jedi/portfolio/internal/handler/review"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"

	reviewRepository "github.com/go-jedi/portfolio/internal/repository/review"
	reviewService "github.com/go-jedi/portfolio/internal/service/review"
)

func (s *serverProvider) ReviewRepository(ctx context.Context) repository.ReviewRepository {
	if s.reviewRepository == nil {
		s.reviewRepository = reviewRepository.NewRepository(s.DBClient(ctx))
	}

	return s.reviewRepository
}

func (s *serverProvider) ReviewService(ctx context.Context) service.ReviewService {
	if s.reviewService == nil {
		s.reviewService = reviewService.NewService(
			s.ReviewRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.reviewService
}

func (s *serverProvider) ReviewHandler(ctx context.Context) *review.Handler {
	if s.reviewHandler == nil {
		s.reviewHandler = review.NewHandler(
			s.ReviewService(ctx),
			s.validator,
		)
	}

	return s.reviewHandler
}
