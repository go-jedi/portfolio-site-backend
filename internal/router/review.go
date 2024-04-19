package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ReviewRoutes(router fiber.Router) {
	router.Post("/review", r.reviewHandler.Create)
	router.Get("/review", r.reviewHandler.Get)
	router.Get("/review/:id", r.reviewHandler.GetById)
	router.Delete("/review/:id", r.reviewHandler.Delete)
}
