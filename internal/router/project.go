package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ProjectRoutes(router fiber.Router) {
	router.Post("/project", r.projectHandler.Create)
	router.Get("/project", r.projectHandler.Get)
	router.Get("/project/:id", r.projectHandler.GetById)
}
