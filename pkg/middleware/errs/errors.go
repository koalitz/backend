package errs

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/pkg/log"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type MyError interface {
	GetInfo() *AbstractError
}

type AbstractError struct {
	Status      int               `json:"-"`
	Description string            `json:"description,omitempty"`
	Fields      map[string]string `json:"fields,omitempty"`
	Err         error             `json:"-"`
}

type ErrHandler struct {
	log *log.Logger
}

func NewErrHandler() *ErrHandler {
	return &ErrHandler{log: log.NewLogger(log.ErrLevel, &log.JSONFormatter{}, true)}
}

func (e *ErrHandler) HandleError(handler func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err == nil {
			return
		}

		var my *AbstractError

		switch err.(type) {
		case StandardError:
			my = err.(StandardError).GetInfo()

		case validator.ValidationErrors:
			my = newValidError(err.(validator.ValidationErrors))

		case *ent.ValidationError:
			my = newValidErrorEnt(err.(*ent.ValidationError))

		case *ent.NotFoundError:
			my = EntNotFoundError.GetInfo(err.(*ent.NotFoundError))

		case *ent.ConstraintError:
			my = EntConstraintError.GetInfo(err.(*ent.ConstraintError))

		case redis.Error:
			redisErr := err.(redis.Error)
			switch err.(redis.Error) {
			case redis.Nil:
				my = RedisNilError.GetInfo(redisErr)

			case redis.TxFailedErr:
				my = RedisTxError.GetInfo(redisErr)

			default:
				my = RedisError.GetInfo(redisErr)
			}
		default:
			my = &AbstractError{
				Status:      http.StatusInternalServerError,
				Description: "Server exception was occurred",
				Err:         err,
			}
		}

		l := e.log.WithErr(my.Err)
		if my.Fields == nil {
			l.Err(my.Description)
		} else {
			l.Err(my.Fields)
		}

		c.JSON(my.Status, my)
	}
}
