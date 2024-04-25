package project

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Get(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER PROJECT) Get...",
	)

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.validator.Var(page, "required,min=1,number")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.validator.Var(limit, "required,min=1,number")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, params, err := h.projectService.Get(c.UserContext(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "успешное получение проектов",
			"result":  []project.Get{},
			"params":  project.Params{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное получение проектов",
		"result":  result,
		"params":  params,
	})
}
