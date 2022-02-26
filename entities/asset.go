package entities

import "time"

type Asset struct {
	Id int				`json:"id" form:"id"`
	Name string			`json:"name" form:"name"`
	Description string	`json:"description" form:"description"`
	CategoryId int		`json:"categoryid" form:"categoryid"`
	CategoryName string	`json:"categoryname" form:"categoryname"`
	Quantity int		`json:"quantity" form:"quantity"`
	Picture string		`json:"picture" form:"picture"`
	CreatedAt time.Time	`json:"created at" form:"created at"`
}