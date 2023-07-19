package postgres

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/ent/user"
	"github.com/koalitz/backend/internal/controller/dao"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, sess *dao.Session) (*dao.Me, error) {
	customer, err := r.userClient.Query().Where(user.ID(sess.ID)).
		Select(user.FieldEmail, user.FieldRole, user.FieldFirstName,
			user.FieldLastName, user.FieldCreateTime).Only(ctx)

	if err == nil {
		return dao.TransformToMe(customer), nil
	}

	return nil, err
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}
