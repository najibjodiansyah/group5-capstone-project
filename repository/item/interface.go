package item

import "capstone-project/entities"

type Item interface {
	Get(availableStatus string, category int, keyword string, page int) ([]entities.ItemResponseFormat, int, error)
	GetById(id int) (entities.ItemResponseFormat, error)
	GetByIdUpdate(id int) (entities.Item, error)
	GetItemUsageHistory(id int) (entities.ItemUsageHistory, error)
	Update(id int, item entities.Item) error
}
