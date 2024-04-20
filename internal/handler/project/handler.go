package project

import (
	"github.com/go-playground/validator/v10"

	"github.com/go-jedi/portfolio/internal/service"
)

type Handler struct {
	projectService service.ProjectService
	validator      *validator.Validate
}

func NewHandler(
	projectService service.ProjectService,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		projectService: projectService,
		validator:      validator,
	}
}
