package review

import (
	"strconv"

	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetByID(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER REVIEW) GetById...",
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

	result, err := h.reviewService.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное получение отзыва",
		"result":  result,
	})
}
