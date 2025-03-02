package service

import (
	"context"
	"goth-template/database"
	"goth-template/server/repository"
	"log/slog"
)

type UserService struct {
	userRepository *repository.UserRepository
	logger         *slog.Logger
}

func CreateUserService(userRepository *repository.UserRepository, logger *slog.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (userService *UserService) GetUsers(ctx context.Context) ([]database.User, error) {
	userService.logger.Info("UserService.GetUsers()")
	users, err := userService.userRepository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
