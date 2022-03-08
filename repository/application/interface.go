package application

import (
	"capstone-project/entities"
)

type Application interface {
	Create(entities.Applications)(int, entities.Applications,error)
	UpdateStatus(applicationid int, status string, managerid *int, itemid *int)(error)
	AvailabilityItem(assetid int) (int, error)
	UpdateItem(itemid *int, availStatus string, employeeid int) error
	GetAll(status string,
		category int,
		date string,
		orderbydate string,
		longestdate string)([]entities.ResponseApplicationWithDuration,int, error)
	GetById(id int)(entities.ResponseApplication,error)
	GetAsset(applicationid int)(int,error)
	UsersApplicationHistory(userid int)([]entities.ResponseApplication,error)
	UsersApplicationActivity(userid int)([]entities.ResponseApplication,error)
}