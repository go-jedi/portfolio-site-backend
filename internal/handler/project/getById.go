package project

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) GetById(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER PROJECT) GetById...",
	)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.validator.Var(id, "required,min=1,number")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := h.projectService.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное получение проекта",
		"result":  result,
	})
}
