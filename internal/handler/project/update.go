package project

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Update(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER PROJECT) Update...",
	)

	var dto project.Update
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := h.validator.Struct(dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := h.projectService.Update(c.UserContext(), dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное изменение проекта",
		"result":  result,
	})
}
