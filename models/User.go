package models

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Role      int
	Status    int
	CreatedAt time.Time
}
