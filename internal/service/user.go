package service

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller/dao"
	"time"
)

type UserPostgres interface {
	FindMe(ctx context.Context, sess *dao.Session) (*dao.Me, error)
	FindUserByID(ctx context.Context, id int) (*ent.User, error)
}

type UserService struct {
	postgres UserPostgres
	redis    UserRedis
}

func NewUserService(postgres UserPostgres, redis UserRedis) *UserService {
	return &UserService{postgres: postgres, redis: redis}
}

func (u *UserService) FindUserByID(id int) (*ent.User, error) {
	return u.postgres.FindUserByID(context.Background(), id)
}

// FindMe returns the detail information about user
func (u *UserService) FindMe(sess *dao.Session) (*dao.Me, error) {
	user, err := u.postgres.FindMe(context.Background(), sess)
	return user, err
}

type UserRedis interface {
	ContainsKeys(ctx context.Context, keys ...string) (int64, error)
	SetVariable(ctx context.Context, key string, value any, exp time.Duration) error
}

// ContainsKeys of redis by key
func (u *UserService) ContainsKeys(keys ...string) (int64, error) {
	return u.redis.ContainsKeys(context.Background(), keys...)
}

// SetVariable of redis by key, his value and exploration time
func (u *UserService) SetVariable(key string, value any, exp time.Duration) error {
	return u.redis.SetVariable(context.Background(), key, value, exp)
}
