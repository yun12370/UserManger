package mapper

import (
	"database/sql"
	"fmt"
	"github.com/yun/UserManger/models"
)

type LoginMapper struct {
	DB *sql.DB
}

func NewLoginMapper(db *sql.DB) *LoginMapper {
	return &LoginMapper{
		DB: db,
	}
}
func (lm *LoginMapper) GetUserByName(username string) (*models.User, error) {
	user := &models.User{}
	sql := "select * from users where username=? "
	err := lm.DB.QueryRow(sql, username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.AvatarURL)
	if err != nil {
		return nil, fmt.Errorf("GetUserByName error: %v", err)
	}
	return user, nil
}
