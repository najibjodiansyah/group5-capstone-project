package user

import "capstone-project/entities"

type User interface{
	Register(entities.User)(entities.User,error)
}