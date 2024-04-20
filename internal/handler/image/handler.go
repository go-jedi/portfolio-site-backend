package image

import (
	"github.com/go-playground/validator/v10"

	"github.com/go-jedi/portfolio/internal/service"
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
