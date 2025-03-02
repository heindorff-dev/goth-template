package handler

import "goth-template/server/repository"

type UserHandler struct {
	userRepository *repository.UserRepository
}

func CreateUserHandler(r *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository: r}
}
