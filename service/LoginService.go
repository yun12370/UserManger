package service

import (
	"fmt"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/untils"
)

type LoginService struct {
	LoginMapper *mapper.LoginMapper
}

func NewLoginService(loginMapper *mapper.LoginMapper) *LoginService {
	return &LoginService{
		LoginMapper: loginMapper,
	}
}

func (ls *LoginService) LoginUser(username, password string) (string, error) {
	user, err := ls.LoginMapper.GetUserByName(username)
	if err != nil || user == nil {
		return "", fmt.Errorf("用户不存在")
	}
	if user.Password != password {
		return "", fmt.Errorf("密码错误")
	}
	if user.Status != 1 {
		return "", fmt.Errorf("用户被禁用")
	}
	token, err := untils.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("令牌生成失败")
	}
	return token, nil
}
