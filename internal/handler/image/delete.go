package image

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Delete(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER IMAGE) Delete...",
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

	result, err := h.imageService.Delete(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "успешное удаление картинки",
		"result":  result,
	})
}
