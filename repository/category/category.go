package category

import (
	"capstone-project/entities"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Get() ([]entities.Category, error) {
	result, err := r.db.Query(`select id, name from categories`)
	if err != nil {
		return nil, err
	}
	var categories []entities.Category
	for result.Next() {
		var category entities.Category
		err := result.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
