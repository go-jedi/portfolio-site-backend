package review

import "time"

type Review struct {
	ID        int       `json:"id" validate:"required,min=1"`
	Author    string    `json:"author" validate:"required,min=1,max=255"`
	Message   string    `json:"message" validate:"required"`
	Rating    int       `json:"rating" validate:"required,min=0,max=5"`
	CreatedAt time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt time.Time `json:"updated_at" validate:"required,datetime"`
}

type Create struct {
	Author  string `json:"author" validate:"required,min=1,max=255"`
	Message string `json:"message" validate:"required"`
	Rating  int    `json:"rating" validate:"required,min=0,max=5"`
}
