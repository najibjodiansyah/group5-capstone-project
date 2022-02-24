package asset

import "time"

type RequestAssetFormat struct{
	Name			string	`json:"name" form:"name"`
	Description		string	`json:"description" form:"description"`
	Category		int		`json:"category" form:"category"`
	Quantity		int  `json:"quantity" form:"quantity"`
	Picture			string	`json:"picture" form:"picture"`
}

type ResponeAssetFormat struct{
	Id				int
	Name			string	
	Description		string
	Categoryid		int	
	Category		string		
	Quantity		int
	Picture			string
	CreatedAt		time.Time
}