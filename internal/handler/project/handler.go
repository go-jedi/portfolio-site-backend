package project

import (
	"github.com/go-jedi/portfolio/internal/service"
	"github.com/go-playground/validator/v10"
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
