package application

import (
	"capstone-project/entities"
	"time"
)

type Application interface {
	Create(entities.Applications)(int, entities.Applications,error)
	UpdateStatus(applicationid int)(entities.Applications,error)
	Get(status string,
		category int,
		date time.Time,
		orderbydate time.Time,
		longestdate time.Time,
		page int)(entities.Applications, error)
	GetUserStatus(userid int, status string)(entities.Applications,error)
}