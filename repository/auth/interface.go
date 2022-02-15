package auth

import "capstone-project/entities"

type Auth interface {
	Login(email string)(entities.User,error)
}