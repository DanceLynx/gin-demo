package repository

import (
	"hello/model"

	"github.com/gin-gonic/gin"
)

//userRepository the struct of userRepository
type userRepository struct {
	baseRepository
}

// NewUserRepository create an instance of *userRepository
func NewUserRepository(ctx *gin.Context) *userRepository {
	userRepository := userRepository{}
	userRepository.New(ctx)
	return &userRepository
}

// Create user data into database
func (ur *userRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

// Find user record from database
func (ur *userRepository) Find(user *model.User) error {
	return ur.db.Find(user).Error
}
