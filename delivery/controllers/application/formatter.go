package application

import "capstone-project/entities"

type InputApp struct {
	Employeeid int			`json:"employeeid"`
	Assetid int				`json:"assetid"`
	Returndate string		`json:"returndate"`
	Activity string			`json:"activity"`
	Specification string	`json:"specification"`
	Description string		`json:"description"`
}


type Inputstatus struct {
	Status 		string	`json:"status"`
}

type responseAll struct{
	Totalpage int				`json:"totalpage"`	
	App []entities.Applications	`json:"applications"`
}