package user

import (
	"github.com/go-playground/validator/v10"

	"github.com/go-jedi/portfolio/internal/service"
)

type Handler struct {
	userService service.UserService
	validator   *validator.Validate
}

func NewHandler(userService service.UserService, validator *validator.Validate) *Handler {
	return &Handler{
		userService: userService,
		validator:   validator,
	}
}
