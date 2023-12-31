package errs

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/koalitz/backend/ent"
	"net/http"
	"strings"
)

// newValidError creates a new StandardError and returns it
func newValidError(errs validator.ValidationErrors) *AbstractError {
	fields := make(map[string]string)

	for _, err := range errs {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "email":
			fields[field] = fmt.Sprintf("%s is not the correct email", field)
		case "required":
			fields[field] = fmt.Sprintf("%s should not be empty", field)
		case "numeric":
			fields[field] = fmt.Sprintf("%s must be a number", field)
		case "gte":
			fields[field] = fmt.Sprintf("%s must be greater or equal %s", field, err.Param())
		case "lte":
			fields[field] = fmt.Sprintf("%s must be lesser or equal %s", field, err.Param())
		case "len":
			fields[field] = fmt.Sprintf("%s must have a length of %s", field, err.Param())
		case "gt":
			fields[field] = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "lt":
			fields[field] = fmt.Sprintf("%s must be lesser than %s", field, err.Param())
		case "name":
			fields[field] = fmt.Sprintf("%s is not valid name", field)
		case "password":
			fields[field] = fmt.Sprintf("%s must contain only english letters and _ character", field)
		case "enum":
			fields[field] = fmt.Sprintf("%s may contain %s values", field, strings.ReplaceAll(err.Param(), "*", ", "))
		}
	}

	return &AbstractError{
		Status: http.StatusUnprocessableEntity,
		Fields: fields,
		Err:    errs,
	}
}

func newValidErrorEnt(v *ent.ValidationError) *AbstractError {
	return &AbstractError{
		Status: http.StatusUnprocessableEntity,
		Fields: map[string]string{
			v.Name: fmt.Sprintf("%s is incorrect", v.Name),
		},
		Err: v,
	}
}
