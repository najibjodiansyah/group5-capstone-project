package statistic

import "capstone-project/entities"

type Statistic interface {
	Get() (entities.Statistic, error)
}
