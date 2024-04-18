package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ReviewRoutes(router fiber.Router) {
	router.Post("/review", r.reviewHandler.Create)
}
