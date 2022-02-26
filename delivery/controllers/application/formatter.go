package application

type InputApp struct {
	Employeeid int			`json:"employeeid"`
	Assetid int				`json:"assetid"`
	Returndate string	`json:"returndate"`
	Activity string			`json:"activity"`
	Specification string	`json:"specification"`
	Description string		`json:"description"`
}