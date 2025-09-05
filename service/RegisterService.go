package service

import (
	"errors"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/models"
)

type RegisterService struct {
	RegisterMapper *mapper.RegisterMapper
}

func NewRegisterService(registerMapper *mapper.RegisterMapper) *RegisterService {
	return &RegisterService{
		RegisterMapper: registerMapper,
	}
}

func (rs *RegisterService) RegisterUser(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("用户名或密码不能为空")
	}
	user := &models.User{
		Username: username,
		Password: password,
		Role:     2,
		Status:   1,
	}
	sysUser, _ := rs.RegisterMapper.GetUserByName(username)
	if sysUser != nil {
		return errors.New("用户名已存在")
	}
	err := rs.RegisterMapper.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}
