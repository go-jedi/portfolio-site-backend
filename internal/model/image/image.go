package image

import "time"

type Image struct {
	ID        int       `json:"id" validate:"required,min=1"`
	ProjectID int       `json:"project_id" validate:"required,min=1"`
	PathFile  string    `json:"path_file" validate:"required,filepath,image"`
	CreatedAt time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt time.Time `json:"updated_at" validate:"required,datetime"`
}

type Get struct {
	ID        int    `json:"id" validate:"required,min=1"`
	ProjectID int    `json:"project_id" validate:"required,min=1"`
	PathFile  string `json:"path_file" validate:"required,filepath,image"`
	CreatedAt string `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" validate:"required"`
}
