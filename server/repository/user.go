package repository

import (
	"context"
	"goth-template/database"
)

type UserRepository struct {
	client *database.DatabaseClient
}

func NewUserRepository(c *database.DatabaseClient) *UserRepository {
	return &UserRepository{client: c}
}

func (r UserRepository) Test(ctx context.Context) ([]database.User, error) {
	users, err := r.client.Query(ctx).ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
