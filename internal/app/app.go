package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/config"
	"github.com/go-jedi/portfolio/internal/router"
	"github.com/go-jedi/portfolio/pkg/logger"
)

type App struct {
	serverProvider *serverProvider
	restServer     *fiber.App
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runRESTServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServerProvider,
		a.initLogger,
		a.initRestServer,
		a.initRouter,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServerProvider(_ context.Context) error {
	a.serverProvider = newServerProvider()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(
		logger.GetCore(
			logger.GetAtomicLevel(a.serverProvider.LoggerConfig().Level()),
		),
	)

	logger.Info("Logger is running")

	return nil
}

func (a *App) initRestServer(_ context.Context) error {
	a.restServer = fiber.New()

	return nil
}

func (a *App) initRouter(ctx context.Context) error {
	// инициализация Handlers
	userHandler := a.serverProvider.UserHandler(ctx)

	// инициализация роутов
	r := router.NewRouter(a.restServer, userHandler)
	err := r.InitRoutes()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runRESTServer() error {
	logger.Info(fmt.Sprintf("REST server is running on %s", a.serverProvider.RESTConfig().Address()))

	// запуск сервера
	err := a.restServer.Listen(
		fmt.Sprintf(
			":%s",
			strings.Split(a.serverProvider.RESTConfig().Address(), ":")[1]),
	)
	if err != nil {
		return err
	}

	return nil
}