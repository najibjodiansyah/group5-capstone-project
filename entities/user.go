package entities

import "time"

type User struct {
	ID        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Avatar    string `json:"avatar" form:"avatar"`
	CreatedAt time.Time
	DeletedAt time.Time
}
