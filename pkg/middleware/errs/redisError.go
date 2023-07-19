package errs

import "net/http"

// Database errors (templates)
var (
	RedisNilError = newRedisError(http.StatusBadRequest, "Can't find value")
	RedisTxError  = newRedisError(http.StatusInternalServerError, "Operation failed")
	RedisError    = newRedisError(http.StatusInternalServerError, "Can't perform query")
)

// redisError describes all server-known errors
type redisError struct {
	status      int
	description string
}

func newRedisError(status int, description string) redisError {
	return redisError{status: status, description: description}
}

// Error implements the Error type
func (r redisError) Error() string {
	return r.description
}

func (r redisError) GetInfo(err error) *AbstractError {

	return &AbstractError{
		Status:      r.status,
		Description: r.description,
		Err:         err,
	}
}
