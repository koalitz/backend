package redis

import (
	"context"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/pkg/conf"
	"github.com/redis/go-redis/v9"
	"reflect"
	"time"
)

const month = time.Hour * 24 * 30

var cfg = conf.GetConfig()

type RClient struct {
	client *redis.Client
}

func NewRClient(client *redis.Client) *RClient {
	return &RClient{client: client}
}

// SetSession and all its parameters
func (r *RClient) SetSession(ctx context.Context, sessionId string, info dao.Session) error {
	return r.client.Watch(ctx, func(tx *redis.Tx) error {

		v := reflect.ValueOf(info)
		typeOfFields := v.Type()

		for i := 0; i < v.NumField(); i++ {
			if err := tx.HSetNX(ctx, sessionId, typeOfFields.Field(i).Name,
				v.Field(i).Interface()).Err(); err != nil {
				return err
			}
		}

		return tx.Expire(ctx, sessionId, month).Err()
	}, sessionId)
}

// GetSession and all its parameters
func (r *RClient) GetSession(ctx context.Context, sessionId string) (*dao.Session, error) {
	info := new(dao.Session)
	err := r.client.HGetAll(ctx, sessionId).Scan(info)
	return info, err
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire.
// Returns true if session was expired
func (r *RClient) ExpandExpireSession(ctx context.Context, sessionId string) (bool, error) {
	v, err := r.client.TTL(ctx, sessionId).Result()
	if err == nil && v <= cfg.Session.Duration/2 {
		return r.client.Expire(ctx, sessionId, month).Result()
	}
	return false, err
}

func (r *RClient) DelKeys(ctx context.Context, keys ...string) {
	r.client.Del(ctx, keys...)
}

// EqualsPopCode returns true if code is saved earlier in email and deletes it
func (r *RClient) EqualsPopCode(ctx context.Context, email string, code string) (ok bool, err error) {
	return ok, r.client.Watch(ctx, func(tx *redis.Tx) error {
		ok, err = tx.SIsMember(ctx, email, code).Result()
		if err != nil || !ok {
			return err
		}
		return tx.Del(ctx, email).Err()
	}, email)
}

// SetCodes or add it to existing key
func (r *RClient) SetCodes(ctx context.Context, key string, value ...any) error {
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		err := tx.SAdd(ctx, key, value...).Err()
		if err != nil {
			return err
		}
		return tx.Expire(ctx, key, time.Hour).Err()
	}, key)
}
