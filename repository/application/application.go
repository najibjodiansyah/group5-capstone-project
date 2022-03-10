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

	stmt, err := ar.db.Prepare(`insert into applications(employeeid, assetid, returnDate, spesification, description, status)
	 							values(?,?,?,?,?,?)
								`)			
	if err != nil {
	log.Println(err)
	return 0, app, errors.New("internal server error")
	}

	res, err := stmt.Exec(app.Employeeid, app.AssetId, app.Returndate, app.Specification, app.Description, app.Status)
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

func (ar *ApplicationRepository)UpdateStatus(applicationid int, status string, managerid *int,itemid *int)(error){
	if managerid != nil {
		res, err := ar.db.Exec(`update applications set status =?, updatedat = ?, managerid = ? where id = ?`,status, time.Now(), managerid, applicationid)
		if err != nil {
			log.Println(err)
			return errors.New("internal server error")
		}

		notAffected, _ := res.RowsAffected()
		if notAffected == 0 {
			log.Println("rows affected is 0 while update application")
			return  errors.New("internal server error")
		}

		return nil

	}else if itemid != nil {
		res, err := ar.db.Exec(`update applications set status =?, updatedat =?, itemid =? where id = ?`,status, time.Now(), itemid, applicationid)
		if err != nil {
			log.Println(err)
			return errors.New("internal server error")
		}

		notAffected, _ := res.RowsAffected()
		if notAffected == 0 {
			log.Println("rows affected is 0 while update application")
			return  errors.New("internal server error")
		}

		return nil

	}else{
		res, err := ar.db.Exec(`update applications set status =?, updatedat = ? where id = ?`,status, time.Now(), applicationid)
		if err != nil {
			log.Println(err)
			return err
		}

		notAffected, _ := res.RowsAffected()
		if notAffected == 0 {
			log.Println("rows affected is 0 while update application")
			return  errors.New("internal server error")
		}

		return nil
	}	
}

func (ar *ApplicationRepository)GetAsset(applicationid int)(int,error){
	// cek avaliability item, select id from item where assetid =? and availstatus tersedia order by id asc limit 1
	// return itemid nya buat di pake klo ngebalikin statusnya
	stmt, err := ar.db.Prepare(`SELECT assetid FROM applications WHERE id =? `)
	if err != nil{
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	res, err := stmt.Query(applicationid)
	if err != nil {
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return 0, errors.New("internal server error")
	}

	var Assetid int

	errScan := res.Scan(&Assetid)
	if errScan != nil {
		return 0, errors.New("internal server error")
	}

	return Assetid, nil
}

func (ar *ApplicationRepository)AvailabilityItem(assetid int)(int,error){
	// cek avaliability item, select id from item where assetid =? and availstatus tersedia order by id asc limit 1
	// return itemid nya buat di pake klo ngebalikin statusnya
	stmt, err := ar.db.Prepare(`SELECT id FROM items WHERE assetid =? AND availableStatus =? ORDER BY id ASC LIMIT 1`)
	if err != nil{
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	res, err := stmt.Query(assetid,"tersedia")
	if err != nil {
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return 0, errors.New("internal server error")
	}

	var Itemid int

	errScan := res.Scan(&Itemid)
	if errScan != nil {
		return 0, errors.New("internal server error")
	}

	return Itemid, nil
}

func (ar *ApplicationRepository)UpdateItem(itemid *int, availStatus string, employeeid int)error{
	// ubah availabilitystatus di item sesuai availstatus, where id =? dan employeeid =?
	if employeeid == -1 {
	stmt, err := ar.db.Prepare(`UPDATE items SET availableStatus=?, employee = NULL where id =?`)
	if err != nil {
		log.Print(err)
		return errors.New("internal server error")
	}

	res, err := stmt.Exec(availStatus,&itemid)
	if err != nil {
		log.Print(err)
		return errors.New("internal server error")
	}

	notAffected, _ := res.RowsAffected()
	if notAffected == 0 {
		log.Println("rows affected is 0 while update application")
		return  errors.New("internal server error")
	}

	return nil

	} else {
	stmt, err := ar.db.Prepare(`UPDATE items SET availableStatus=?, employee=? where id =?`)
	if err != nil {
		log.Print(err)
		return errors.New("internal server error")
	}

	res, err := stmt.Exec(availStatus,employeeid,&itemid)
	if err != nil {
		log.Print(err)
		return errors.New("internal server error")
	}

	notAffected, _ := res.RowsAffected()
	if notAffected == 0 {
		log.Println("rows affected is 0 while update application")
		return errors.New("internal server error")
	}
	
	return nil
	}
}

func (ar *ApplicationRepository)GetById(id int)(entities.ResponseApplication,error){
	var app entities.ResponseApplication

	stmt, err := ar.db.Prepare(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
	where ap.id = ? `)
	if err != nil {
		log.Println(err)
		return app, errors.New("internal server error")
	}

	res, err := stmt.Query(id)
	if err != nil {
		fmt.Println("2",err)
		return app, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		log.Println(isExist)
		return app, errors.New("internal server error")
	}

	errScan := res.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
	if errScan != nil {
		log.Println(err)
		return app, errors.New("internal server error")
	}

	return app, nil
}

func (ar *ApplicationRepository)GetAll(status string,category int,date string,orderbydate string,longestdate string) ([]entities.ResponseApplicationWithDuration,int, error){
	var (	apps []entities.ResponseApplicationWithDuration
			totalApp int
			err error
			result *sql.Rows
			query string 	)

	query = `select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id`
	
	if category != 0 {
		query += " where ass.categoryid = ?"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query(category)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if status != ""{
		query += " where ap.status = ?"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query(status)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if date != ""{
		query += " where date(ap.requestdate) = ?"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query(date)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if orderbydate == "asc"{
		query += " order by ap.requestdate asc"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query()
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if orderbydate == "desc"{
		query += " order by ap.requestdate desc"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query()
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if longestdate == "asc"{
		query += " order by ap.returndate asc"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query()
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else if longestdate == "desc"{
		query += " order by ap.returndate desc"
		stmt, err := ar.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
		result, err = stmt.Query()
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	} else {
		result, err = ar.db.Query(query)
		if err != nil {
			log.Println(err)
			return nil, 0, errors.New("internal server error")
		}
	}

	for result.Next() {
		var app entities.ResponseApplicationWithDuration

		err := result.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
		if err!= nil {
			log.Println(err)
			return apps, totalApp, errors.New("internal server error")
		}

		diftime := app.Returndate.Sub(time.Now())

		duration := fmt.Sprintf("%v days", int(diftime.Hours()/24))

		app.Duration = duration
		
		apps = append(apps, app)
	}

	var totalAppQuery *sql.Rows

	if category != 0 {
		totalAppQuery, err = ar.db.Query(`select count(a.id) 
		from applications as a 
		inner join assets as c on a.assetid = c.id 
		where c.categoryid= ?`, category)
	}else if status != "" {
		totalAppQuery, err = ar.db.Query(`select count(a.id) 
		from applications as a 
		where a.status= ?`, status)
	}else if date != "" {
		totalAppQuery, err = ar.db.Query(`select count(a.id) 
		from applications as a 
		where date(a.requestdate)= ?`, date)
	}else {
		totalAppQuery, err = ar.db.Query(`select count(id) 
		from applications`)
	}
	if err != nil {
		log.Println(err)
		return apps, totalApp, errors.New("internal server error")
	}

	defer totalAppQuery.Close()
	
	for totalAppQuery.Next() {
		err := totalAppQuery.Scan(&totalApp)
		if err != nil {
			log.Println(err)
			return apps, totalApp, errors.New("internal server error")
		}
	}

	defer result.Close()

	return apps, totalApp, nil
}

func (ar *ApplicationRepository)UsersApplicationHistory(userid int)([]entities.ResponseApplication,error){
	var apps []entities.ResponseApplication

	stmt, err := ar.db.Prepare(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
	where ap.employeeid = ? and ap.status = ?`)
	if err != nil {
		log.Println(err)
		return apps, errors.New("internal server error")
	}

	res, err := stmt.Query(userid, "donereturn")
	if err != nil {
		log.Println(err)
		return apps, errors.New("internal server error")
	}

	for res.Next() {
		var app entities.ResponseApplication
		err := res.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
		if err!= nil {
			fmt.Println("~~~converting NULL to string is unsupported~~~")
			return apps, errors.New("internal server error")
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (ar *ApplicationRepository)UsersApplicationActivity(userid int)([]entities.ResponseApplication,error){
	var apps []entities.ResponseApplication

	stmt, err := ar.db.Prepare(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
	where ap.employeeid = ? and not ap.status = ?`)
	if err != nil {
		log.Println(err)
		return apps, errors.New("internal server error")
	}

	res, err := stmt.Query(userid, "donereturn")
	if err != nil {
		log.Println(err)
		return apps, errors.New("internal server error")
	}

	for res.Next() {
		var app entities.ResponseApplication

		err := res.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
		if err!= nil {
			log.Println(err)
			return apps, errors.New("internal server error")
		}
		apps = append(apps, app)
	}

	return apps, nil
}