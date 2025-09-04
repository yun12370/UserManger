package mapper

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yun/UserManger/models"
)

type UserMapper struct {
	DB *sql.DB
}

func NewUserMapper(db *sql.DB) *UserMapper {
	return &UserMapper{DB: db}
}

// 查询
func (um *UserMapper) GetUsers(page, pageSize int) ([]*models.User, error) {
	sql := "select * from users limit ? offset ?"
	fmt.Println(sql)
	fmt.Printf("Fetching users with page: %d, page size: %d\n", page, pageSize)
	rows, err := um.DB.Query(sql, pageSize, ((page - 1) * pageSize))
	if err != nil {
		println("Error fetching users: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user models.User
		fmt.Println("222")
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Status, &user.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning user: %v\n", err)
			return nil, err
		}
		fmt.Println("111")
		users = append(users, &user)
		fmt.Println("sadg")
	}
	fmt.Printf("Fetched %d usersdsf\n", len(users))
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
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("修改用户不存在")
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
