package project

import (
	"errors"

	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"
)

type serv struct {
	projectRepository repository.ProjectRepository
	imageRepository   repository.ImageRepository
	txManager         db.TxManager
}

func NewService(
	projectRepository repository.ProjectRepository,
	imageRepository repository.ImageRepository,
	txManager db.TxManager,
) service.ProjectService {
	return &serv{
		projectRepository: projectRepository,
		imageRepository:   imageRepository,
		txManager:         txManager,
	}
}

var ErrFileServerDirEmpty = errors.New("file server dir is empty")
