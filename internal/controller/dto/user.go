package dto

type NameParam struct {
	Name string `uri:"name" validate:"required,gte=5,lte=20,name"`
}
