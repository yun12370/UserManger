package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yun/UserManger/models"
	"time"
)

var JwtKey = []byte("secret")

type Claims struct {
	UserID int `json:"user-id"`
	Role   int `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token生成失败:" + err.Error())
	}
	return signedString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token解析失败:" + err.Error())
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token无效")
	}
}
