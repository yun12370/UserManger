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
