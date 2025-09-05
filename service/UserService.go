package service

import (
	"errors"
	"fmt"
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
	if user.Username == "" || user.Password == "" {
		return errors.New("用户名或密码不能为空")
	}
	sysUser, _ := us.UserMapper.GetUserByName(user.Username)
	if sysUser != nil {
		return errors.New("用户名已存在")
	}
	err := us.UserMapper.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("用户名或密码不能为空")
	}
	if user.Status < 0 || user.Status > 1 {
		return fmt.Errorf("非法用户状态")
	}
	if user.Role < 0 || user.Role > 2 {
		return fmt.Errorf("非法用户角色")
	}
	err := us.UserMapper.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("非法用户ID")
	}
	err := us.UserMapper.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
