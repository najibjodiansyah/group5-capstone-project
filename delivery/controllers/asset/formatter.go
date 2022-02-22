package asset

type RequestAssetFormat struct{
	Name			string	`json:"name" form:"name"`
	Description		string	`json:"description" form:"description"`
	Category		string		`json:"category" form:"category"`
	Quantity		int  `json:"quantity" form:"quantity"`
	Picture			string	`json:"picture" form:"picture"`
}