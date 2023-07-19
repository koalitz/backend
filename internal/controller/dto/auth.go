package dto

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithCode struct {
	Email string `json:"email,omitempty" validate:"required,email"`
	Code  string `json:"code,omitempty" validate:"required,len=5"`
}
