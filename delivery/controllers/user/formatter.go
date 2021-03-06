package user

import "time"
type RegisterUserFormat struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Role 	  string `json:"role" form:"role"`
	Avatar	  string `json:"avatar" form:"avatar"`
}

type ResponseUserFormat struct {
	ID		  int `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Role 	  string `json:"role" form:"role"`
	Avatar	  string `json:"avatar" form:"avatar"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}