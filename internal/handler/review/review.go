package review

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Create(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER REVIEW) Create...",
	)

	var dto review.Create
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

	result, err := h.reviewService.Create(c.UserContext(), dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное создание отзыва",
		"result":  result,
	})
}
