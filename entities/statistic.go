package entities

type Statistic struct {
	TotalAsset       int `json:"totalAsset" form:"totalAsset"`
	TotalUsed        int `json:"totalUsed" form:"totalUsed"`
	TotalAvailable   int `json:"totalAvailable" form:"totalAvailable"`
	TotalMaintenance int `json:"totalMaintenance" form:"totalMaintenance"`
}
