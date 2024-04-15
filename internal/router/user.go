package router

import "github.com/gofiber/fiber/v3"

func (r *Router) UserRoutes(router fiber.Router) {
	router.Get("/", r.userHandler.Get)
}
