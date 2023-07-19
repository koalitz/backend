package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/internal/controller/dto"
	"github.com/koalitz/backend/pkg/conf"
	"mime/multipart"
)

type UserService interface {
	FindUserByID(id int) (*ent.User, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	DelKeys(keys ...string)
	CompareHashAndPassword(old, new []byte) error

	AuthUserByEmail(email string) (*ent.User, error)
	CreateUserByEmail(auth dto.EmailWithCode) (*ent.User, error)
}

type PostService interface {
	FindPostByID(id int) (*ent.Post, error)
	FindPostByTitle(title string) ([]*ent.Post, error)
	CreatePost(imageName *string, info dto.PostInfo, id int) (*ent.Post, error)
}

type MailSender interface {
	DialAndSend(subj, body string, to ...string) error
}

type Session interface {
	SetNewCookie(id int, c *gin.Context) error
	ValidateSession(sessionId string) (info *dao.Session, ok bool, err error)
	GenerateSecretCode(length int) string
	GenerateFileName(c *gin.Context, file *multipart.FileHeader) (string, error)
}

type Handler struct {
	user UserService
	auth AuthService
	mail MailSender
	post PostService
	sess Session
	cfg  *conf.Config
}

func NewHandler(user UserService, auth AuthService, mail MailSender, post PostService, sess Session, cfg *conf.Config) *Handler {
	return &Handler{user: user, auth: auth, mail: mail, post: post, sess: sess, cfg: cfg}
}
