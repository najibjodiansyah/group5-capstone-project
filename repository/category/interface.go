package category

import "capstone-project/entities"

type Category interface {
	Get() ([]entities.Category, error)
}
