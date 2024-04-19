package app

import (
	"context"
	"log"

	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/platform_common/pkg/db/pg"
	"github.com/go-jedi/platform_common/pkg/db/transaction"
	"github.com/go-playground/validator/v10"

	"github.com/go-jedi/portfolio/internal/config"
	"github.com/go-jedi/portfolio/internal/handler/project"
	"github.com/go-jedi/portfolio/internal/handler/review"
	"github.com/go-jedi/portfolio/internal/handler/user"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"
)

type serverProvider struct {
	loggerConfig     config.LoggerConfig
	pgConfig         config.PGConfig
	fileServerConfig config.FileServerConfig
	restConfig       config.RESTConfig

	dbClient  db.Client
	txManager db.TxManager

	validator *validator.Validate

	projectRepository repository.ProjectRepository
	projectService    service.ProjectService
	projectHandler    *project.Handler

	reviewRepository repository.ReviewRepository
	reviewService    service.ReviewService
	reviewHandler    *review.Handler

	userRepository repository.UserRepository
	userService    service.UserService
	userHandler    *user.Handler
}

func newServerProvider() *serverProvider {
	return &serverProvider{}
}

func (s *serverProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serverProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serverProvider) FileServerConfig() config.FileServerConfig {
	if s.fileServerConfig == nil {
		cfg, err := config.NewFileServerConfig()
		if err != nil {
			log.Fatalf("failed to get file server config: %s", err.Error())
		}

		s.fileServerConfig = cfg
	}

	return s.fileServerConfig
}

func (s *serverProvider) RESTConfig() config.RESTConfig {
	if s.restConfig == nil {
		cfg, err := config.NewRESTConfig()
		if err != nil {
			log.Fatalf("failed to get rest config: %s", err.Error())
		}

		s.restConfig = cfg
	}

	return s.restConfig
}

func (s *serverProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serverProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serverProvider) Validator(_ context.Context) *validator.Validate {
	if s.validator == nil {
		s.validator = validator.New(validator.WithRequiredStructEnabled())
	}

	return s.validator
}
