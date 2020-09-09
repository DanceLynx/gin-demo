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

func (this *userService) FindByNameAndPass(username,pass string) (model.User,error) {
	var user model.User
	err := db.Where(model.User{Username: username,Password:pass}).First(&user)
	return user,err.Error
}