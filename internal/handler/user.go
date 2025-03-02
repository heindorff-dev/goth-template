package handler

import (
	"goth-template/internal/service"
	"log/slog"
)

type UserHandler struct {
	userService *service.UserService
	logger      *slog.Logger
}

func NewUserHandler(userService *service.UserService, logger slog.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      &logger,
	}
}
