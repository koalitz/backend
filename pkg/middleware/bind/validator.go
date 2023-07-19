package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/koalitz/backend/pkg/log"
)

func NewValidator() *validator.Validate {
	v := validator.New()

	if err := v.RegisterValidation("name", validateName); err != nil {
		log.WithErr(err).Warn("can't validate name fields")
	}

	if err := v.RegisterValidation("link", validateLink); err != nil {
		log.WithErr(err).Warn("can't validate email fields")
	}

	if err := v.RegisterValidation("title", validateTitle); err != nil {
		log.WithErr(err).Warn("can't validate email fields")
	}

	if err := v.RegisterValidation("email", validateEmail); err != nil {
		log.WithErr(err).Warn("can't validate email fields")
	}

	if err := v.RegisterValidation("uuid4", validateUUID4); err != nil {
		log.WithErr(err).Warn("can't validate uuid4 fields")
	}

	if err := v.RegisterValidation("enum", validateEnum); err != nil {
		log.WithErr(err).Warn("can't validate enums fields")
	}

	return v
}

func HandleBody[T any](handler func(*gin.Context, T) error, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		var t T
		if err := c.ShouldBindJSON(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)

	}
}

func HandleParam[T any](handler func(*gin.Context, T) error, v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		var t T
		if err := c.ShouldBindUri(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t)
	}
}
