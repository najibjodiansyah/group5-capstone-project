package entities

import "time"

type User struct{
	ID 			int
	Name 		string 
	Email 		string
	Password 	string
	Avatar		string
	CreatedAt 	time.Time
	DeletedAt 	time.Time
}