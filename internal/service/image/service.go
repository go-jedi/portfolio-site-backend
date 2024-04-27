package image

import (
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"
)

type serv struct {
	imageRepository repository.ImageRepository
	txManager       db.TxManager
}

func NewService(
	imageRepository repository.ImageRepository,
	txManager db.TxManager,
) service.ImageService {
	return &serv{
		imageRepository: imageRepository,
		txManager:       txManager,
	}
}
