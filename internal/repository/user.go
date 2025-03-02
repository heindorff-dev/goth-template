package repository

import (
	"context"
	"goth-template/database"
	"log/slog"
)

type UserRepository struct {
	client *database.DatabaseClient
	logger *slog.Logger
}

func NewUserRepository(client *database.DatabaseClient, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		client: client,
		logger: logger,
	}
}

func (userRepository UserRepository) GetUsers(ctx context.Context) ([]database.User, error) {
	userRepository.logger.Info("UserRepository.GetUsers()")
	users, err := userRepository.client.Query(ctx).ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
