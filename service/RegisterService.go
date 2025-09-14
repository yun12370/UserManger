package service

import (
	"errors"
	"fmt"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/models"
	"golang.org/x/crypto/bcrypt"
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
	if len(username) == 0 || len(username) > 20 {
		return fmt.Errorf("用户名长度必须在1-20位之间")
	}
	if len(password) < 6 || len(password) > 20 {
		return fmt.Errorf("密码长度必须在6-20位之间")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}
	user := &models.User{
		Username: username,
		Password: string(hashed),
		Role:     2,
		Status:   1,
	}
	sysUser, _ := rs.RegisterMapper.GetUserByName(username)
	if sysUser != nil {
		return errors.New("用户名已存在")
	}
	err = rs.RegisterMapper.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}
