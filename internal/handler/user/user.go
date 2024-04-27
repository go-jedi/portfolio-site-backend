package user

import (
	"fmt"

	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func (h *Handler) Get(c fiber.Ctx) error {
	id := 1

	logger.Info(
		"(HANDLER USER) Get...",
		zap.Int64("id", int64(id)),
	)

	result, err := h.userService.Get(c.UserContext(), int64(id))
	fmt.Println("result: ", result)
	fmt.Println("err: ", err)

	return c.SendString("Hello, World!")
}
