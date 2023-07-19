package dto

import (
	"mime/multipart"
)

type PostInfo struct {
	Title   string                `form:"title,omitempty" validate:"required,gte=3,lte=70"`
	Place   string                `form:"place,omitempty" validate:"required,lte=100"`
	Summary string                `form:"summary,omitempty" validate:"required,lte=1024"`
	Image   *multipart.FileHeader `form:"file,omitempty"`
}

type Limit struct {
	Limit int `uri:"limit"`
}

type Title struct {
	Title string `uri:"title"`
}

type ID struct {
	ID int `uri:"id"`
}
