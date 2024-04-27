package app

import (
	"context"

	"github.com/go-jedi/portfolio/internal/handler/user"
	"github.com/go-jedi/portfolio/internal/repository"
	userRepository "github.com/go-jedi/portfolio/internal/repository/user"
	"github.com/go-jedi/portfolio/internal/service"
	userService "github.com/go-jedi/portfolio/internal/service/user"
)

func (s *serverProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serverProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serverProvider) UserHandler(ctx context.Context) *user.Handler {
	if s.userHandler == nil {
		s.userHandler = user.NewHandler(
			s.UserService(ctx),
			s.validator,
		)
	}

	return s.userHandler
}
