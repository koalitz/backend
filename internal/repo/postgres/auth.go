package postgres

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/ent/user"
)

// IDExist returns true if username exists. Panics if error occurred
func (r *UserStorage) IDExist(ctx context.Context, id int) (bool, error) {
	return r.userClient.Query().Where(user.ID(id)).Exist(ctx)
}

// EmailExist returns true if user Exists. Panic if error occurred
func (r *UserStorage) EmailExist(ctx context.Context, email string) (bool, error) {
	return r.userClient.Query().Where(user.Email(email)).Exist(ctx)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on sessions)
func (r *UserStorage) AuthUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.userClient.Query().Where(
		user.EmailEQ(email),
	).Only(ctx)
}

func (r *UserStorage) AddSession(ctx context.Context, id int, sessions ...string) error {
	return r.userClient.Update().AppendSessions(sessions).Where(user.ID(id)).Exec(ctx)
}
