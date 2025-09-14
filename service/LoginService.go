package service

import (
	"fmt"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/utils"
	"golang.org/x/crypto/bcrypt"
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
	// 数据库查询错误
	if err != nil {
		return "", err
	}
	// 用户不存在
	if user == nil {
		return "", fmt.Errorf("用户名或密码错误")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("用户名或密码错误")
	}
	if user.Status != 1 {
		return "", fmt.Errorf("用户被禁用")
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("令牌生成失败")
	}
	return token, nil
}
