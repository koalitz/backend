package bind

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var (
	EmailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	UUID4       = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`)
)

func validateEmail(fl validator.FieldLevel) bool {
	return EmailRegexp.MatchString(fl.Field().String())
}

func validateUUID4(fl validator.FieldLevel) bool {
	return UUID4.MatchString(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	enums := strings.Split(fl.Param(), "*")
	field := fl.Field().String()

	for _, v := range enums {
		if v == field {
			return true
		}
	}

	return false
}
