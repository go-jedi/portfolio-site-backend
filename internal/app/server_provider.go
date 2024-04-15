package app

import (
	"context"
	"log"

	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/platform_common/pkg/db/pg"
	"github.com/go-jedi/platform_common/pkg/db/transaction"

	"github.com/go-jedi/portfolio/internal/config"
	"github.com/go-jedi/portfolio/internal/handler/user"
	"github.com/go-jedi/portfolio/internal/repository"
	"github.com/go-jedi/portfolio/internal/service"

	userRepository "github.com/go-jedi/portfolio/internal/repository/user"
	userService "github.com/go-jedi/portfolio/internal/service/user"
)

type serverProvider struct {
	loggerConfig config.LoggerConfig
	pgConfig     config.PGConfig
	restConfig   config.RESTConfig

	dbClient  db.Client
	txManager db.TxManager

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
		s.userHandler = user.NewHandler(s.UserService(ctx))
	}

	return s.userHandler
}
