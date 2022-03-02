package entities

type Procurement struct {
	ID            int    `json:"id" form:"id"`
	EmployeeId    int    `json:"employeeid" form:"employeeid"`
	AssetName     string `json:"assetName" form:"assetName"`
	RequestDate   string `json:"requestDate" form:"requestDate"`
	Spesification string `json:"spesification" form:"spesification"`
	Description   string `json:"description" form:"description"`
	Status        string `json:"status" form:"status"`
	UpdatedAt     string `json:"updatedAt" form:"updatedAt"`
	ManagerId     int    `json:"managerId" form:"managerId"`
}

type ProcurementResponseFormat struct {
	ID            int     `json:"id" form:"id"`
	EmployeeId    int     `json:"employeeid" form:"employeeid"`
	EmployeeName  string  `json:"employeeName" form:"employeeName"`
	AssetName     string  `json:"assetName" form:"assetName"`
	RequestDate   string  `json:"requestDate" form:"requestDate"`
	Spesification string  `json:"spesification" form:"spesification"`
	Description   string  `json:"description" form:"description"`
	Status        string  `json:"status" form:"status"`
	UpdatedAt     string  `json:"updatedAt" form:"updatedAt"`
	ManagerId     *int    `json:"managerId" form:"managerId"`
	ManagerName   *string `json:"managerName" form:"managerName"`
}

type ProcurementResponseTotal struct {
	TotalPage    int                         `json:"totalPage" form:"totalPage"`
	Procurements []ProcurementResponseFormat `json:"procurements" form:"procurements"`
}
