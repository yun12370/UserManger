package service

import (
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/models"
)

type UserService struct {
	UserMapper *mapper.UserMapper
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{UserMapper: userMapper}
}

func (us *UserService) GetUsers(page, pageSize int) ([]*models.UserVO, error) {
	users, err := us.UserMapper.GetUsers(page, pageSize)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (us *UserService) CreateUser(user *models.User) error {
	err := us.UserMapper.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	err := us.UserMapper.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) DeleteUser(id int) error {
	err := us.UserMapper.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
