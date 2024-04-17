package project

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Create(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER PROJECT) Create...",
	)

	err := h.projectService.Create(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "успешное создание проекта",
	})
}
