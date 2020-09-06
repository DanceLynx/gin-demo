package service

import (
	"hello/model"
)

type userService struct {
}

var User = &userService{}

func (this *userService) AddUser(user *model.User) error {
	return db.Create(user).Error
}

func (this *userService) GetUserById(id int) (*model.User, error) {
	model := &model.User{}
	err := db.First(model, id).Error
	return model, err
}
