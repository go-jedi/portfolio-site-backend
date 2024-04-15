package user

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (h *Handler) Get(c fiber.Ctx) error {
	id := 1

	logger.Info(
		"(HANDLER) Get...",
		zap.Int64("id", int64(id)),
	)

	result, err := h.userService.Get(c.UserContext(), int64(id))
	fmt.Println("result: ", result)
	fmt.Println("err: ", err)

	return c.SendString("Hello, World!")
}
