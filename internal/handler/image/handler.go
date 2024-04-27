package image

import (
	"github.com/go-jedi/portfolio/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	imageService service.ImageService
	validator    *validator.Validate
}

func NewHandler(
	imageService service.ImageService,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		imageService: imageService,
		validator:    validator,
	}
}
