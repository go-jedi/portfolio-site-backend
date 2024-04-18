package review

import (
	"github.com/go-jedi/platform_common/pkg/db"

	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"
)

type serv struct {
	reviewRepository repository.ReviewRepository
	txManager        db.TxManager
}

func NewService(
	reviewRepository repository.ReviewRepository,
	txManager db.TxManager,
) service.ReviewService {
	return &serv{
		reviewRepository: reviewRepository,
		txManager:        txManager,
	}
}
