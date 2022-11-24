package repository

import (
	"context"

	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) (id uint64, err error)
	FindById(ctx context.Context, id int) (user *ent.User, err error)
}

type userRepository struct {
	client *ent.Client
}

// CreateUser implements UserRepository
func (ur *userRepository) CreateUser(ctx context.Context, name string, email string) (id uint64, err error) {
	b := ur.client.User.Create()
	// Add all fields.
	b.SetName(name)
	b.SetEmail(email)
	e, err := b.Save(ctx)
	if err != nil {
		return 0, err
	}

	q := ur.client.User.Query().Where(user.ID(e.ID))
    r, err := q.Only(ctx)
	if err != nil {
		// This should never happen.
		return 0, err
	}

	return uint64(r.ID), nil
}

// FindById implements UserRepository
func (ur *userRepository) FindById(ctx context.Context, id int) (user *ent.User, err error) {
	e, err := ur.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}
