package entities

type Item struct {
	ID              int    `json:"id" form:"id"`
	AssetId         int    `json:"assetId" form:"assetId"`
	EmployeeId      int    `json:"employeeId" form:"employeeId"`
	Name            string `json:"name" form:"name"`
	AvailableStatus string `json:"availableStatus" form:"availableStatus"`
}

type ItemResponseFormat struct {
	ID              int    `json:"id" form:"id"`
	CategoryId      int    `json:"categoryId" form:"categoryId"`
	Category        int    `json:"category" form:"category"`
	Name            string `json:"name" form:"name"`
	Picture         int    `json:"picture" form:"picture"`
	AvailableStatus string `json:"availableStatus" form:"availableStatus"`
}

type ItemResponseTotal struct {
	TotalPage int `json:"totalItem" form:"totalItem"`
	Items     []ItemResponseFormat
}
