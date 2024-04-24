package router

import "github.com/gofiber/fiber/v3"

func (r *Router) ReviewRoutes(router fiber.Router) {
	router.Post("/review", r.reviewHandler.Create)
	router.Get("/review", r.reviewHandler.Get)
	router.Get("/review/:id", r.reviewHandler.GetByID)
	router.Patch("/review/publish/:id", r.reviewHandler.Publish)
	router.Patch("/review/unpublish/:id", r.reviewHandler.UnPublish)
	router.Delete("/review/:id", r.reviewHandler.Delete)
}
