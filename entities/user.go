package entities

import "time"

type User struct {
	ID        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Avatar    string `json:"avatar" form:"avatar"`
	Role	  string `json:"role" form:"role"`
	CreatedAt time.Time	`json:"created at" form:"created at"`
	DeletedAt time.Time	`json:"deleted at" form:"deleted at"`
}
