package router

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/handler/user"
)

type Router struct {
	app         *fiber.App
	userHandler *user.Handler
}

func NewRouter(app *fiber.App, userHandler *user.Handler) *Router {
	return &Router{
		app:         app,
		userHandler: userHandler,
	}
}

func (r *Router) InitRoutes() error {
	v1 := r.app.Group("/v1")
	{
		r.UserRoutes(v1)
	}

	return nil
}
