package user

import "capstone-project/entities"

type User interface{
	Register(entities.User)(entities.User, error)
	GetById(id int)(entities.User, error)
	Update(id int, user entities.User) error
	Delete(id int) error
	GetEmployees()([]entities.Employee,error)
}