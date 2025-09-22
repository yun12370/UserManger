package mapper

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yun/UserManger/models"
)

type HomeMapper struct {
	DB *sql.DB
}

func NewHomeMapper(DB *sql.DB) *HomeMapper {
	return &HomeMapper{
		DB: DB,
	}
}
func (hm *HomeMapper) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	sql := "select * from users where id=?"
	err := hm.DB.QueryRow(sql, id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.AvatarURL)
	if err != nil {
		return nil, fmt.Errorf("getUserByID error: %v", err)
	}
	return user, nil
}
func (hm *HomeMapper) UpdateAvatarURL(avatarUrl string, id int) error {
	sql := "update users set avatar_url=? where id=?"
	res, err := hm.DB.Exec(sql, avatarUrl, id)
	if err != nil {
		return fmt.Errorf("updateAvatarURL error: %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateAvatarURL error: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("更新用户头像失败")
	}
	return nil
}
