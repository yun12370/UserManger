package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        int       `json:"ID"`
	Username  string    `json:"Username"`
	Password  string    `json:"Password"`
	Role      int       `json:"Role"`
	Status    int       `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type UserVO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Role      int       `json:"role"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func ToVO(u *User) UserVO {
	return UserVO{
		ID:        u.ID,
		Username:  u.Username,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
	}
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"CreatedAt"`
	}{
		Alias:     (*Alias)(&u),
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
func (u UserVO) MarshalJSON() ([]byte, error) {
	type Alias UserVO
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"CreatedAt"`
	}{
		Alias:     (*Alias)(&u),
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
