package procurement

import "capstone-project/entities"

type Procurement interface {
	Get(status string) ([]entities.ProcurementResponseFormat, error)
	GetById(id int) (entities.ProcurementResponseFormat, error)
	Create(procurement entities.Procurement) (entities.Procurement, error)
	Update(id int, procurement entities.Procurement) error
}
