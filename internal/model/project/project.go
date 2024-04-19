package project

import "time"

type Project struct {
	ID          int       `json:"id" validate:"required,min=1"`
	Title       string    `json:"author" validate:"required,min=1"`
	Description string    `json:"description" validate:"required,min=1"`
	Technology  string    `json:"technology" validate:"required,min=1"`
	CreatedAt   time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt   time.Time `json:"updated_at" validate:"required,datetime"`
}
