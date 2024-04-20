package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ImageRoutes(router fiber.Router) {
	router.Delete("/image/:id", r.imageHandler.Delete)
}
