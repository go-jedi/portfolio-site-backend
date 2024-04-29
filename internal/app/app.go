package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/go-jedi/portfolio/internal/config"
	"github.com/go-jedi/portfolio/internal/router"
	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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
		a.initValidator,
		a.initFileServer,
		a.initCors,
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

func (a *App) initValidator(ctx context.Context) error {
	a.serverProvider.Validator(ctx)

	return nil
}

func (a *App) initFileServer(_ context.Context) error {
	fileServerDir := a.serverProvider.FileServerConfig().FileServerDir()
	fileServerPrefix := a.serverProvider.FileServerConfig().FileServerPrefix()

	a.restServer.Static(
		fmt.Sprintf("/%s", fileServerPrefix),
		fileServerDir,
		fiber.Static{
			ByteRange: true,
		},
	)

	return nil
}

func (a *App) initCors(_ context.Context) error {
	// Инициализация CORS
	corsOrigin := a.serverProvider.CORSConfig().ORIGIN()
	corsMethod := a.serverProvider.CORSConfig().METHOD()
	corsHeader := a.serverProvider.CORSConfig().HEADER()
	corsCredential := a.serverProvider.CORSConfig().CREDENTIAL()
	corsMaxAge := a.serverProvider.CORSConfig().MAXAGE()

	a.restServer.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigin,
		AllowMethods:     corsMethod,
		AllowHeaders:     corsHeader,
		AllowCredentials: corsCredential,
		MaxAge:           corsMaxAge,
	}))

	return nil
}

func (a *App) initRouter(ctx context.Context) error {
	// Инициализация Handlers
	projectHandler := a.serverProvider.ProjectHandler(ctx)
	imageHandler := a.serverProvider.ImageHandler(ctx)
	reviewHandler := a.serverProvider.ReviewHandler(ctx)
	userHandler := a.serverProvider.UserHandler(ctx)

	// Инициализация роутов
	r := router.NewRouter(
		a.restServer,
		projectHandler,
		imageHandler,
		reviewHandler,
		userHandler,
	)
	err := r.InitRoutes()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runRESTServer() error {
	// Инициализация сертификатов
	certFile := a.serverProvider.CERTConfig().CertFile()
	certKeyFile := a.serverProvider.CERTConfig().CertKeyFile()

	logger.Info(fmt.Sprintf("REST server is running on %s", a.serverProvider.RESTConfig().Address()))

	// запуск сервера
	err := a.restServer.Listen(
		fmt.Sprintf(
			":%s",
			strings.Split(a.serverProvider.RESTConfig().Address(), ":")[1],
		),
		fiber.ListenConfig{
			CertFile:    certFile,
			CertKeyFile: certKeyFile,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
