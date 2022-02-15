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

func (ur *UserRepository)GetById(id int)(entities.User, error){
	var user entities.User
	stmt, err := ur.db.Prepare("select id, name, email, password, avatar, created_at from users where id = ? and deleted_at is NULL")
	if err != nil {
		return user, errors.New("internal server error") 
	}
	res, err := stmt.Query(id)
	if err != nil{
		return user, errors.New("internal server error") 
	}
	if isExist := res.Next(); !isExist {
		return user, errors.New("internal server error") 
	}
	errScan := res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt)
	if errScan != nil {
		return user, errScan
	}
	return user, nil
}

func (ur *UserRepository)Update(id int, user entities.User) error {
	stmt, err := ur.db.Prepare("UPDATE users SET name= ?, email= ?, password= ?, avatar= ? WHERE id = ? and deleted_at is NULL")
	if err != nil {
		return  errors.New("internal server error") 
	}
	result, err := stmt.Exec(user.Name, user.Email, user.Password, user.Avatar, id)
	if err != nil {
		return  errors.New("internal server error") 
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		log.Println("rows affected is 0 while delete user")
		return  errors.New("internal server error")
	}
	return nil
}

func (ur *UserRepository) Delete (id int) error {
	stmt, err := ur.db.Prepare("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL")
	if err != nil {
		log.Println(err)
		return  errors.New("internal server error") 
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return  errors.New("internal server error") 
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("rows affected is 0 while delete user")
		return err
	}

	return nil
}