package service

import (
	"errors"
	"fmt"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/models"
	"os"
	"path/filepath"
	"time"
)

type HomeService struct {
	HomeMapper *mapper.HomeMapper
}

func NewHomeService(homeMapper *mapper.HomeMapper) *HomeService {
	return &HomeService{
		HomeMapper: homeMapper,
	}
}

func (hs *HomeService) GetUserByID(id int) (*models.User, error) {
	user, err := hs.HomeMapper.GetUserByID(id)
	if user == nil || user.AvatarURL == "" {
		return nil, errors.New("用户头像获取失败")
	}
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (hs *HomeService) UpdateAvatarURL(avatarUrl string, id int) error {
	err := hs.HomeMapper.UpdateAvatarURL(avatarUrl, id)
	if err != nil {
		return err
	}
	return nil
}

func (hs *HomeService) SaveAvatarFile(id int, data []byte, filename string) (string, error) {
	dir := "static/avatar"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	ext := filepath.Ext(filename)
	newName := fmt.Sprintf("%d_%d%s", id, time.Now().Unix(), ext)
	dstPath := filepath.Join(dir, newName)

	if err := os.WriteFile(dstPath, data, 0644); err != nil {
		return "", err
	}
	newURL := "/static/avatar/" + newName

	user, err := hs.HomeMapper.GetUserByID(id)
	if err != nil {
		return "", err
	}
	if user != nil && user.AvatarURL != "" {
		oldPath := "." + user.AvatarURL
		_ = os.Remove(oldPath)
	}
	return newURL, nil
}
