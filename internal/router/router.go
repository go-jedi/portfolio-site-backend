package router

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/handler/image"
	"github.com/go-jedi/portfolio/internal/handler/project"
	"github.com/go-jedi/portfolio/internal/handler/review"
	"github.com/go-jedi/portfolio/internal/handler/user"
)

type Router struct {
	app            *fiber.App
	projectHandler *project.Handler
	imageHandler   *image.Handler
	reviewHandler  *review.Handler
	userHandler    *user.Handler
}

func NewRouter(
	app *fiber.App,
	projectHandler *project.Handler,
	imageHandler *image.Handler,
	reviewHandler *review.Handler,
	userHandler *user.Handler,
) *Router {
	return &Router{
		app:            app,
		projectHandler: projectHandler,
		imageHandler:   imageHandler,
		reviewHandler:  reviewHandler,
		userHandler:    userHandler,
	}
}

func (r *Router) InitRoutes() error {
	v1 := r.app.Group("/v1")
	{
		r.ProjectRoutes(v1)
		r.ImageRoutes(v1)
		r.ReviewRoutes(v1)
		r.UserRoutes(v1)
	}

	return nil
}
