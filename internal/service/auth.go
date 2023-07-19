package service

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/internal/controller/dto"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AuthPostgres interface {
	IDExist(ctx context.Context, id int) (bool, error)
	EmailExist(ctx context.Context, email string) (bool, error)
	AuthUserByEmail(ctx context.Context, email string) (*ent.User, error)
	AddSession(ctx context.Context, id int, sessions ...string) error
	CreateUserByEmail(ctx context.Context, email dto.EmailWithCode) (*ent.User, error)
}

type AuthService struct {
	postgres AuthPostgres
	redis    AuthRedis
}

func NewAuthService(postgres AuthPostgres, redis AuthRedis) *AuthService {
	return &AuthService{postgres: postgres, redis: redis}
}

// IDExist returns true if user Exists. Panics if error occurred
func (a *AuthService) IDExist(id int) (bool, error) {
	return a.postgres.IDExist(context.Background(), id)
}

// EmailExist returns true if user Exists. Panic if error occurred
func (a *AuthService) EmailExist(email string) (bool, error) {
	return a.postgres.EmailExist(context.Background(), email)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on sessions)
func (a *AuthService) AuthUserByEmail(email string) (*ent.User, error) {
	return a.postgres.AuthUserByEmail(context.Background(), email)
}

func (a *AuthService) AddSession(id int, sessions ...string) error {
	return a.postgres.AddSession(context.Background(), id, sessions...)
}

func (a *AuthService) CreateUserByEmail(auth dto.EmailWithCode) (*ent.User, error) {
	return a.postgres.CreateUserByEmail(context.Background(), auth)
}

type AuthRedis interface {
	SetSession(ctx context.Context, sessionId string, info dao.Session) error
	GetSession(ctx context.Context, sessionId string) (*dao.Session, error)
	ExpandExpireSession(ctx context.Context, sessionId string) (bool, error)
	DelKeys(ctx context.Context, keys ...string)
	EqualsPopCode(ctx context.Context, email string, code string) (bool, error)
	SetCodes(ctx context.Context, key string, value ...any) error
}

// GetSession and all its parameters
func (a *AuthService) GetSession(sessionId string) (*dao.Session, error) {
	return a.redis.GetSession(context.Background(), sessionId)
}

// SetSession and all its parameters
func (a *AuthService) SetSession(sessionId string, info dao.Session) error {
	return a.redis.SetSession(context.Background(), sessionId, info)
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire
func (a *AuthService) ExpandExpireSession(sessionId string) (bool, error) {
	return a.redis.ExpandExpireSession(context.Background(), sessionId)
}

// DelKeys fully deletes session id
func (a *AuthService) DelKeys(keys ...string) {
	a.redis.DelKeys(context.Background(), keys...)
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (a *AuthService) EqualsPopCode(email string, code string) (bool, error) {
	return a.redis.EqualsPopCode(context.Background(), email, code)
}

// SetCodes or add it to existing key
func (a *AuthService) SetCodes(key string, value ...any) error {
	return a.redis.SetCodes(context.Background(), key, value...)
}

func (a *AuthService) CompareHashAndPassword(old, new []byte) error {
	return bcrypt.CompareHashAndPassword(old, new)
}

func (a *AuthService) FormatLanguage(header string) string {
	parts := strings.SplitN(header, ";", 1)
	languages := strings.SplitN(parts[0], ",", 1)
	switch languages[0] {
	case "ru":
		return languages[0]
	default:
		return "en"
	}
}
