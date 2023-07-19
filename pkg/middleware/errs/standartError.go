package errs

import "net/http"

// StandardError describes all server-known errors
type StandardError struct {
	status      int
	description string
	err         error
}

var (
	MailCodeError = newStandardError(http.StatusBadRequest, "Code is not correct, used or expired")
	UnAuthorized  = newStandardError(http.StatusUnauthorized, "You are not logged in")
	EmailError    = newStandardError(http.StatusInternalServerError, "Can't send message to your email")
)

// Error implements the Error type
func (e StandardError) Error() string {
	return e.err.Error()
}

// newStandardError creates a new StandardError and returns it
func newStandardError(status int, description string) StandardError {
	return StandardError{
		status:      status,
		description: description,
	}
}

func (e StandardError) AddErr(err error) StandardError {
	e.err = err
	return e
}

func (e StandardError) GetInfo() *AbstractError {
	return &AbstractError{
		Status:      e.status,
		Description: e.description,
		Err:         e.err,
	}
}
