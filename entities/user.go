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
type UserResponseFormat struct {
	Id        int    
	Name      string 
	Email     string 
	Avatar	  string 
	CreatedAt time.Time 
}