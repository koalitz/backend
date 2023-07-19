package errs

import (
	"net/http"
)

var (
	EntNotFoundError   = newEntError(http.StatusBadRequest, "Entity is not found")
	EntConstraintError = newEntError(http.StatusBadRequest, "This value already exists")
)

// entError describes all server-known errors
type entError struct {
	status      int
	description string
}

func newEntError(status int, description string) *entError {
	return &entError{status: status, description: description}
}

// Error implements the Error type
func (e entError) Error() string {
	return e.description
}

func (e entError) GetInfo(err error) *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Description: e.description,
		Err:         err,
	}
}
