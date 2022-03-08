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
			return err
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
			return err
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
		return 0, err
	}
	res, err := stmt.Query(applicationid)
	if err != nil {
		log.Println(err)
		return 0, err
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
		return 0, err
	}
	res, err := stmt.Query(assetid,"tersedia")
	if err != nil {
		log.Println(err)
		return 0, err
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
		return err
	}

	res, err := stmt.Exec(availStatus,&itemid)
	if err != nil {
		log.Print(err)
		return err
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
		return err
	}

	res, err := stmt.Exec(availStatus,employeeid,&itemid)
	if err != nil {
		log.Print(err)
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
		fmt.Println("1",err)
		return app, errors.New("internal server error")
	}
	res, err := stmt.Query(id)
	if err != nil {
		fmt.Println("2",err)
		return app, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		fmt.Println("diexist",isExist)
		return app, errors.New("internal server error")
	}

	errScan := res.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
	if errScan != nil {
		fmt.Println("3",err)
		return app, errScan
	}

	return app, nil
}

func (ar *ApplicationRepository)GetAll(status string,category int,date string,orderbydate string,longestdate string) ([]entities.ResponseApplicationWithDuration,int, error){
	var apps []entities.ResponseApplicationWithDuration
	var totalApp int
	var err error
	var result *sql.Rows
	// limit := 10
	// offset := (page - 1) * limit
	if category != 0 {
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
	where ass.categoryid = ?`, category)
	} else if status != ""{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		where ap.status = ?`, status)
	} else if date != ""{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		where date(ap.requestdate) = ?`, date)
	} else if orderbydate == "asc"{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		order by ap.requestdate asc`)
	} else if orderbydate == "desc"{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		order by ap.requestdate desc`)
	} else if longestdate == "asc"{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		order by ap.returndate asc`)
	} else if longestdate == "desc"{
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id
		order by ap.returndate desc`)
	} else {
	result, err = ar.db.Query(`select ap.id,ap.employeeid, ap.managerid, ap.assetid, ap.itemid, ap.requestdate, ap.returndate, ap.spesification, ap.description, ap.status, ap.updatedat, u.name, a.name, ass.name, i.name, ass.picture, ass.categoryid, c.name
	FROM applications ap
	JOIN assets as ass ON ap.assetid = ass.id
	JOIN users as u ON ap.employeeid = u.id
	LEFT JOIN items as i ON ap.itemid = i.id
	LEFT JOIN users as a ON ap.managerid = a.id
	JOIN categories as c ON ass.categoryid = c.id`)
	}
	if err != nil {
		fmt.Println("Get 1", err)
		return apps, totalApp, err
	}

	for result.Next() {
		var app entities.ResponseApplicationWithDuration
		err := result.Scan(&app.Id, &app.Employeeid, &app.Managerid, &app.Assetid, &app.Itemid, &app.Requestdate, &app.Returndate, &app.Specification, &app.Description, &app.Status, &app.Updatedat, &app.Employeename, &app.Managername, &app.Assetname, &app.ItemName, &app.Photo, &app.Categoryid, &app.Categoryname)
		if err!= nil {
			return apps, totalApp, err
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
		fmt.Println("Get 2", err)
		return apps, totalApp, err
	}

	defer totalAppQuery.Close()
	
	for totalAppQuery.Next() {
		err := totalAppQuery.Scan(&totalApp)
		if err != nil {
			return apps, totalApp, err
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
			return apps, err
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
	where ap.employeeid = ? and not ap.status = ?
	`)
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
			return apps, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}