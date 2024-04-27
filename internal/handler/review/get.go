package review

import (
	"strconv"

	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Get(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER REVIEW) Get...",
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

	result, params, err := h.reviewService.Get(c.UserContext(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "успешное получение отзывов",
			"result":  []review.Review{},
			"params":  review.Params{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное получение отзывов",
		"result":  result,
		"params":  params,
	})
}
