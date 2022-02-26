package application

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type ApplicationRepository struct{
	db *sql.DB
}

func NewApplication(db *sql.DB) *ApplicationRepository{
	return &ApplicationRepository{db: db}
}

func (ar *ApplicationRepository)Create(app entities.Applications)(int,entities.Applications,error){

	stmt, err := ar.db.Prepare(`insert into applications(employeeid, managerid, assetid, itemid, returnDate, spesification, description, status)
	 							values(?,?,?,?,?,?,?,?)
								`)			
	if err != nil {
	log.Println(err)
	return 0, app, errors.New("internal server error")
	}

	fmt.Println(app)

	res, err := stmt.Exec(app.Employeeid, app.Managerid, app.AssetId, app.Itemid, app.Returndate, app.Specification, app.Description, app.Status)
	if err != nil {
		log.Println(err)
		return 0, app, errors.New("internal server error")
	}	
	appid, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, app, errors.New("internal server error")
	}		
	return int(appid),entities.Applications{},nil
}
func (ar *ApplicationRepository)UpdateStatus(applicationid int)(entities.Applications,error){
	fmt.Println("Not implemented")
	return entities.Applications{},nil
}
func (ar *ApplicationRepository)Get(status string,category int,date time.Time,orderbydate time.Time,longestdate time.Time,page int) (entities.Applications, error){
	fmt.Println("Not implemented")
	return entities.Applications{},nil
}
func (ar *ApplicationRepository)GetUserStatus(userid int, status string)(entities.Applications,error){
	fmt.Println("Not Implemented")
	return entities.Applications{},nil
}