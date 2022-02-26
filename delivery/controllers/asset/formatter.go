package asset

import (
	"capstone-project/entities"
	"time"
)

type RequestAssetFormat struct{
	Name			string	`json:"name" form:"name"`
	Description		string	`json:"description" form:"description"`
	Category		int		`json:"category" form:"category"`
	Quantity		int 	`json:"quantity" form:"quantity"`
	Picture			string	`json:"picture" form:"picture"`
}

type ResponeAssetFormat struct{
	Id				int			`json:"id" form:"id"`
	Name			string		`json:"name" form:"name"`
	Description		string		`json:"description" form:"description"`
	Categoryid		int			`json:"categoryid" form:"categoryid"`
	Category		string		`json:"categoryname" form:"categoryname"`
	Quantity		int			`json:"quantity" form:"quantity"`
	Picture			string		`json:"picture" form:"picture"`
	CreatedAt		time.Time	`json:"created at" form:"created at"`
}

type responseAll struct{
	Totalpage int			`json:"totalpage"`	
	Assets []entities.Asset	`json:"assets"`
}