package auth

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"log"
)

type AuthRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// return repository berbentuk entity saja
func (r *AuthRepository) Login(email string) (entities.User, error) {
	stmt, err := r.db.Prepare(`select id, name, password, role from users where email = ? and deleted_at IS NULL`)

	if err != nil {
		log.Println(err)
		return entities.User{}, errors.New("internal server error")
	}

	res, err := stmt.Query(email)

	if err != nil {
		log.Println(err)
		return entities.User{}, errors.New("internal server error")
	}

	defer res.Close()

	var user entities.User

	if res.Next() {
		err := res.Scan(&user.ID, &user.Name, &user.Password, &user.Role)

		if err != nil {
			log.Println(err)
			return entities.User{}, errors.New("internal server error")
		}
	}

	return user, nil
}