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
	Updatedat		time.Time	`json:"updatedat" form:"updatedat"`
}

type ResponseApplication struct {
	Id				int			`json:"id"`
	Employeeid		int			`json:"employeeid"`
	Employeename	string		`json:"employeename"`
	Managerid		int			`json:"managerid"`
	Managername		string		`json:"managername"`
	Assetid			int			`json:"assetid"`
	Assetname		string		`json:"assetname"`
	Itemid			int			`json:"itemid"`
	ItemName		string		`json:"itemname"`
	Photo			string		`json:"photo"`
	Requestdate		time.Time	`json:"requestdate"`
	Returndate		time.Time	`json:"returndate"`
	Specification	string		`json:"specification"`
	Description		string		`json:"description"`
	Status			string		`json:"status"`
	Updatedat		string		`json:"updateat"`
}