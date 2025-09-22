package mapper

import (
	"database/sql"
	"fmt"
	"github.com/yun/UserManger/models"
)

type RegisterMapper struct {
	DB *sql.DB
}

func NewRegisterMapper(db *sql.DB) *RegisterMapper {
	return &RegisterMapper{DB: db}
}

func (rm *RegisterMapper) RegisterUser(user *models.User) error {
	sql := "insert into users(username,password,role,status) values(?,?,?,?)"
	result, err := rm.DB.Exec(sql, user.Username, user.Password, user.Role, user.Status)
	if err != nil {
		return fmt.Errorf("注册用户失败")
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取用户ID失败")
	}
	user.ID = int(insertId)
	return nil
}

func (rm *RegisterMapper) GetUserByName(username string) (*models.User, error) {
	user := &models.User{}
	sql := "select * from users where username=?"
	err := rm.DB.QueryRow(sql, username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.AvatarURL)
	if err != nil {
		return nil, fmt.Errorf("GetUserByName error: %v", err)
	}
	return user, nil
}
