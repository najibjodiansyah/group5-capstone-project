package entities

import "time"

type Applications struct {
	Id				int			`json:"id" form:"id"`
	Employeeid		int			`json:"employeeid" form:"employeeid"`
	Managerid		int			`json:"managerid" form:"managerid"`
	AssetId			int			`json:"assetid" form:"assetid"`
	Itemid			int			`json:"itemid" form:"itemid"`
	Requestdate		time.Time	`json:"requestdate" form:"requestdate"`
	Returndate		time.Time	`json:"returndate" form:"returndate"`
	Activity		string		`json:"activity" form:"activity"`
	Specification	string		`json:"specification" form:"specification"`
	Description		string		`json:"description" form:"description"`
	Status			string		`json:"status" form:"status"`
	Updatedat		string	`json:"updatedat" form:"updatedat"`
}