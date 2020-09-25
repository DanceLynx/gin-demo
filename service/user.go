package service

import (
	"hello/core/log"
	"hello/model"
	"hello/repository"

	"github.com/gin-gonic/gin"
)

type userService struct {
}

// UserService create userService instance
var UserService = userService{}

func (us userService) Create(ctx *gin.Context, user *model.User) error {
	userRepository := repository.NewUserRepository(ctx)
	log.Info(ctx, "create user", user)
	return userRepository.Create(user)
}

func (us userService) Find(ctx *gin.Context, user *model.User) error {
	userRepository := repository.NewUserRepository(ctx)
	return userRepository.Find(user)
}
