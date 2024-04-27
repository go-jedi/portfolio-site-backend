package review

import (
	"github.com/go-jedi/portfolio/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	reviewService service.ReviewService
	validator     *validator.Validate
}

func NewHandler(
	reviewService service.ReviewService,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		reviewService: reviewService,
		validator:     validator,
	}
}
