package asset

import (
	"time"
)

type RequestAssetFormat struct{
	Name			string	`json:"name" form:"name"`
	Description		string	`json:"description" form:"description"`
	Categoryid		int		`json:"categoryid" form:"categoryid"`
	Quantity		int 	`json:"quantity" form:"quantity"`
	Picture			string	`json:"picture" form:"picture"`
}

type ResponseAssetFormat struct{
	Id				int			`json:"id" form:"id"`
	Name			string		`json:"name" form:"name"`
	Description		string		`json:"description" form:"description"`
	Categoryid		int			`json:"categoryid" form:"categoryid"`
	Category		string		`json:"categoryname" form:"categoryname"`
	Quantity		int			`json:"quantity" form:"quantity"`
	Picture			string		`json:"picture" form:"picture"`
	CreatedAt		time.Time	`json:"createdat" form:"createdat"`
	Availability	string		`json:"availability" form:"availability"`
}
type responseAll struct{
	Totalpage int				`json:"totalpage"`	
	Assets []ResponseAssetFormat	`json:"assets"`
}