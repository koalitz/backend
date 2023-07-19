package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller"
	"github.com/koalitz/backend/internal/repo/postgres"
	redisRepo "github.com/koalitz/backend/internal/repo/redis"
	"github.com/koalitz/backend/internal/service"
	"github.com/koalitz/backend/pkg/client/email"
	"github.com/koalitz/backend/pkg/client/postgresql"
	redisInit "github.com/koalitz/backend/pkg/client/redis"
	"github.com/koalitz/backend/pkg/conf"
	"github.com/koalitz/backend/pkg/log"
	"github.com/koalitz/backend/pkg/middleware/bind"
	"github.com/koalitz/backend/pkg/middleware/errs"
	"github.com/koalitz/backend/pkg/middleware/query"
	"github.com/koalitz/backend/pkg/middleware/session"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := conf.GetConfig()

	pClient, rClient, mailClient := getClients(cfg)

	h, sess := initHandler(pClient, rClient, mailClient, cfg)

	r := gin.New()
	r.MaxMultipartMemory = 5 << 20 // 1 MB

	h.InitRoutes(createSetter(r, sess))

	run(cfg.Listen.Port, r, pClient, rClient)
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client) {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithErr(err).Fatalf("error occurred while running http server")
		}
	}()
	log.Infof("Server Started On Port %d", port)

	<-quit

	log.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithErr(err).Fatal("Server Shutdown Failed")
	}

	if err := rClient.Close(); err != nil {
		log.WithErr(err).Fatal("Redis Connection Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		log.WithErr(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	log.LastInfo("Server Exited Properly")
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *email.MailClient) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisInit.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Password, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.NewMailClient(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	return pClient, rClient, mailClient
}

func initHandler(pClient *ent.Client, rClient *redis.Client, mailClient *email.MailClient, cfg *conf.Config) (*controller.Handler, *session.Auth) {
	pUser := postgres.NewUserStorage(pClient.User)
	pPost := postgres.NewPostStorage(pClient.Post)
	rConn := redisRepo.NewRClient(rClient)

	user := service.NewUserService(pUser, rConn)
	post := service.NewPostService(pPost)

	auth := service.NewAuthService(pUser, rConn)
	sess := session.NewAuth(auth, cfg)

	return controller.NewHandler(
		user,
		auth,
		mailClient,
		post,
		sess,
		cfg,
	), sess
}

func createSetter(r *gin.Engine, sess *session.Auth) *controller.Setter {
	return controller.NewSetter(
		r,
		bind.NewValidator(),
		errs.NewErrHandler(),
		query.NewQueryHandler(),
		sess,
	)
}
