package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/pkg/conf"
)

type UserService interface {
	FindUserByID(id int) (*ent.User, error)
	FindMe(sess *dao.Session) (*dao.Me, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	DelKeys(keys ...string)
	CompareHashAndPassword(old, new []byte) error

	AuthUserByEmail(email string) (*ent.User, error)
}

type MailSender interface {
	DialAndSend(subj, body string, to ...string) error
}

type Session interface {
	SetNewCookie(id int, c *gin.Context) error
	ValidateSession(sessionId string) (info *dao.Session, ok bool, err error)
	GenerateSecretCode(length int) string
}

type Handler struct {
	user UserService
	auth AuthService
	mail MailSender
	sess Session
	cfg  *conf.Config
}

func NewHandler(user UserService, auth AuthService, mail MailSender, sess Session, cfg *conf.Config) *Handler {
	return &Handler{user: user, auth: auth, mail: mail, sess: sess, cfg: cfg}
}
