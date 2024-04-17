package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ProjectRoutes(router fiber.Router) {
	router.Get("/project", r.projectHandler.Create)
}
