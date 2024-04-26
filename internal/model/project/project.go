package project

import (
	"time"

	"github.com/go-jedi/portfolio/internal/model/image"
)

type Project struct {
	ID          int       `json:"id" validate:"required,min=1"`
	Title       string    `json:"title" validate:"required,min=1"`
	Description string    `json:"description" validate:"required,min=1"`
	Technology  string    `json:"technology" validate:"required,min=1"`
	CreatedAt   time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt   time.Time `json:"updated_at" validate:"required,datetime"`
}

type Create struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
	Technology  string `json:"technology" validate:"required,min=1"`
}

type Get struct {
	ID          int         `json:"id" validate:"required,min=1"`
	Title       string      `json:"title" validate:"required,min=1"`
	Description string      `json:"description" validate:"required,min=1"`
	Technology  string      `json:"technology" validate:"required,min=1"`
	CreatedAt   time.Time   `json:"created_at" validate:"required,datetime"`
	UpdatedAt   time.Time   `json:"updated_at" validate:"required,datetime"`
	Paths       []image.Get `json:"paths" validate:"required"`
}

type Params struct {
	PageCount int `json:"page_count" validate:"required,min=1"`
	Limit     int `json:"limit" validate:"required,min=1"`
}

type Update struct {
	ID          int    `json:"id" validate:"required,min=1"`
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
	Technology  string `json:"technology" validate:"required,min=1"`
}
