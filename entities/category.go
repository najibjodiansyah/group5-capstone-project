package entities

type Category struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
