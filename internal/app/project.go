package app

import (
	"context"

	"github.com/go-jedi/portfolio/internal/handler/project"
	"github.com/go-jedi/portfolio/internal/repository"
	projectRepository "github.com/go-jedi/portfolio/internal/repository/project"
	"github.com/go-jedi/portfolio/internal/service"
	projectService "github.com/go-jedi/portfolio/internal/service/project"
)

func (s *serverProvider) ProjectRepository(ctx context.Context) repository.ProjectRepository {
	if s.projectRepository == nil {
		s.projectRepository = projectRepository.NewRepository(s.DBClient(ctx))
	}

	return s.projectRepository
}

func (s *serverProvider) ProjectService(ctx context.Context) service.ProjectService {
	if s.projectService == nil {
		s.projectService = projectService.NewService(
			s.ProjectRepository(ctx),
			s.ImageRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.projectService
}

func (s *serverProvider) ProjectHandler(ctx context.Context) *project.Handler {
	if s.projectHandler == nil {
		s.projectHandler = project.NewHandler(
			s.ProjectService(ctx),
			s.validator,
		)
	}

	return s.projectHandler
}
