package mapper

import (
	"database/sql"
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
		Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
