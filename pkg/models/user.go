package models

import "time"

type User struct {
	Id        int64     `json:"-" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password_hash" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Role      int64     `json:"role_id" db:"role_id"`
}
