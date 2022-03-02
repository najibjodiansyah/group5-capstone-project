package procurement

type ProcurementRequestFormat struct {
	Employeeid    int    `json:"employeeid" form:"employeeid"`
	AssetName     string `json:"assetName" form:"assetName"`
	Spesification string `json:"spesification" form:"spesification"`
	Description   string `json:"description" form:"description"`
}

type Inputstatus struct {
	Status string `json:"status"`
}
