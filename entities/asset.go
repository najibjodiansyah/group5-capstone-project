package entities

import "time"

type Asset struct {
	Id int
	Name string
	Description string
	Category string
	Quantity int
	Picture string
	CreatedAt time.Time
}
