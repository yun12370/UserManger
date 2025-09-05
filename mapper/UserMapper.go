package mapper

import (
	"database/sql"
	"errors"
	"github.com/yun/UserManger/models"
)

type UserMapper struct {
	DB *sql.DB
}

func NewUserMapper(db *sql.DB) *UserMapper {
	return &UserMapper{DB: db}
}

// 查询
func (um *UserMapper) GetUsers(page, pageSize int) ([]*models.UserVO, error) {
	sql := "select id, username, role, status, created_at from users limit ? offset ?"
	rows, err := um.DB.Query(sql, pageSize, ((page - 1) * pageSize))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.UserVO
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Role, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		userVO := models.ToVO(&user)
		users = append(users, &userVO)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// 添加
func (um *UserMapper) CreateUser(user *models.User) error {
	sql := "insert into users(username,password,role,status) values(?,?,?,?)"
	result, err := um.DB.Exec(sql, user.Username, user.Password, user.Role, user.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

// 修改
func (um *UserMapper) UpdateUser(user *models.User) error {
	sql := "update users set username=?,password=?,role=?,status=? where id=?"
	result, err := um.DB.Exec(sql, user.Username, user.Password, user.Role, user.Status, user.ID)
	if err != nil {
		return errors.New("用户名已存在")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("信息未修改或修改用户不存在")
	}
	return nil
}

// 删除
func (um *UserMapper) DeleteUser(id int) error {
	sql := "delete from users where id=?"
	result, err := um.DB.Exec(sql, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("删除用户不存在")
	}
	return nil
}

func (um *UserMapper) GetUserByName(username string) (*models.User, error) {
	user := &models.User{}
	sql := "select * from users where username=? "
	err := um.DB.QueryRow(sql, username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
