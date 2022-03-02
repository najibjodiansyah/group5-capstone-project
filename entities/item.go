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
	Category        string `json:"category" form:"category"`
	Name            string `json:"name" form:"name"`
	Description     string `json:"description" form:"description"`
	Picture         string `json:"picture" form:"picture"`
	AvailableStatus string `json:"availableStatus" form:"availableStatus"`
}

type ItemResponseTotal struct {
	TotalPage int                  `json:"totalPage" form:"totalPage"`
	Items     []ItemResponseFormat `json:"items" form:"items"`
}

type ItemUsageHistory struct {
	ID         int        `json:"id" form:"id"`
	Name       string     `json:"name" form:"name"`
	CategoryId int        `json:"categoryId" form:"categoryId"`
	Category   string     `json:"category" form:"category"`
	Picture    string     `json:"picture" form:"picture"`
	Users      []ItemUser `json:"users" form:"users"`
}

type ItemUser struct {
	AssetUser   string `json:"assetUser" form:"assetUser"`
	LendingDate string `json:"lendingDate" form:"lendingDate"`
	UsageStatus string `json:"usageStatus" form:"usageStatus"`
}
