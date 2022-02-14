package user

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"log"
)

type UserRepository struct{
	db *sql.DB
}

func New(db *sql.DB) *UserRepository{
	return &UserRepository{db: db}
}

func (ur *UserRepository) checkEmailExistence(email string) (id int64, err error) {
	stmt, err := ur.db.Prepare("select id from users where email = ?")

	if err != nil {
		return 0, err
	}

	res, err := stmt.Query(email)

	if err != nil {
		return 0, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {
	id, err := ur.checkEmailExistence(user.Email)
	if err != nil {
		return user, errors.New("internal server error")
	}

	if id != 0 {
		return user, errors.New("User Already exist") // User A:ready exist
	}
	stmt, err := ur.db.Prepare("insert into users(name, email, password) values(?,?,?)")
	if err != nil {
		log.Println(err)
		return user, errors.New("internal server error")
	}

	_, errr := stmt.Exec(user.Name, user.Email, user.Password)
	if errr != nil {
		log.Println(errr)
		return user, errors.New("internal server error")
	}

	return user, nil
}