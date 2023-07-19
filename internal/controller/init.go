package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/pkg/middleware/bind"
	"github.com/koalitz/backend/pkg/middleware/session"
	"net/http"
)

type ErrHandler interface {
	HandleError(handler func(*gin.Context) error) gin.HandlerFunc
}

type SessionHandler interface {
	Session(handler func(*gin.Context, *dao.Session) error) func(c *gin.Context) error
	SessionFunc(c *gin.Context) (*dao.Session, error)
}

type QueryHandler interface {
	HandleQueries() gin.HandlerFunc
}

type Setter struct {
	r     *gin.Engine
	valid *validator.Validate
	erh   ErrHandler
	qh    QueryHandler
	sess  SessionHandler
}

func NewSetter(r *gin.Engine, valid *validator.Validate, erh ErrHandler, qh QueryHandler, sess SessionHandler) *Setter {
	return &Setter{r: r, valid: valid, erh: erh, qh: qh, sess: sess}
}

func (h *Handler) InitRoutes(s *Setter) {
	initMiddlewares(s.r, s.qh)

	rg := s.r.Group(h.cfg.Listen.QueryPath)

	auth := rg.Group("/auth")
	{
		auth.POST("/email", s.erh.HandleError(bind.HandleBody(h.signInByEmail, s.valid)))

		sess := auth.Group("/session")
		{
			sess.GET("", s.erh.HandleError(s.sess.Session(h.getMe)))
			sess.DELETE("", s.erh.HandleError(h.signOut))
		}
	}

	post := rg.Group("/post")
	{
		post.POST("", s.erh.HandleError(session.HandleJSONBody(h.createPost, s.sess.SessionFunc, s.valid)))
		post.GET("/:id", s.erh.HandleError(bind.HandleParam(h.getPostById, s.valid)))
		post.GET("/:title", s.erh.HandleError(bind.HandleParam(h.getPostByTitle, s.valid)))
	}

	post.GET("files/:limit", s.erh.HandleError(bind.HandleParam(h.getImages, s.valid)))

	rg.Static("/file", "./"+h.cfg.Files.Path)

	if h.mail != nil {
		email := rg.Group("/email")
		{
			email.POST("/send-code", s.erh.HandleError(bind.HandleBody(h.sendCodeToEmail, s.valid)))
		}
	}
}

func initMiddlewares(r *gin.Engine, qh QueryHandler) {
	config := cors.Config{
		AllowOrigins:     []string{"https://koalitz.github.io", "http://localhost:3000", "http://localhost:80", "http://localhost"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodPatch, http.MethodDelete, http.MethodPost},
		AllowHeaders:     []string{"Content-Code", "Content-Length", "Cache-Control", "User-Agent", "Accept-Language", "Accept", "DomainName", "Accept-Encoding", "Connection", "Set-Cookie", "Cookie", "Date", "Postman-Token", "Host"},
		AllowCredentials: true,
	}

	r.Use(qh.HandleQueries(), cors.New(config), gin.Recovery())
}
