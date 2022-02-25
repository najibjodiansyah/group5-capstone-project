package entities

import "time"

type Asset struct {
	Id int
	Name string
	Description string
	Category Category
	Quantity int
	Picture string
	CreatedAt time.Time
}

type Category struct {
	Id int
	Name string
}
