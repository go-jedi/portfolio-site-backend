package user

import "github.com/go-jedi/portfolio/internal/service"

type Handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}
