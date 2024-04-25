package review

import "time"

type Review struct {
	ID        int       `json:"id" validate:"required,min=1"`
	Username  string    `json:"username" validate:"required,min=1,max=255"`
	Message   string    `json:"message" validate:"required"`
	Rating    int       `json:"rating" validate:"required,min=0,max=5"`
	CreatedAt time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt time.Time `json:"updated_at" validate:"required,datetime"`
}

type Create struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Message  string `json:"message" validate:"required"`
	Rating   int    `json:"rating" validate:"required,min=0,max=5"`
}

type Params struct {
	PageCount int `json:"page_count" validate:"required,min=1"`
	Limit     int `json:"limit" validate:"required,min=1"`
}
