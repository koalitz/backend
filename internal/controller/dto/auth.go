package dto

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithCode struct {
	Email     string `json:"email,omitempty" validate:"required,email"`
	Code      string `json:"code,omitempty" validate:"required,len=5"`
	FirstName string `json:"firstName,omitempty" validate:"omitempty,gte=2,lte=32"`
	LastName  string `json:"lastName,omitempty" validate:"omitempty,gte=2,lte=32"`
}
