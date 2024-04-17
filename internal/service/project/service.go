package project

import (
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"
)

type serv struct {
	projectRepository repository.ProjectRepository
	txManager         db.TxManager
}

func NewService(
	projectRepository repository.ProjectRepository,
	txManager db.TxManager,
) service.ProjectService {
	return &serv{
		projectRepository: projectRepository,
		txManager:         txManager,
	}
}
