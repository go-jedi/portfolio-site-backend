package review

import (
	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Get(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER REVIEW) Get...",
	)

	result, err := h.reviewService.Get(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное получение отзывов",
		"result":  result,
	})
}
