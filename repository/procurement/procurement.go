package procurement

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type ProcurementRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ProcurementRepository {
	return &ProcurementRepository{db: db}
}

func (r *ProcurementRepository) Get(status string) ([]entities.ProcurementResponseFormat, error) {
	// var totalProcurement int
	var err error
	var result *sql.Rows
	if status != "" {
		result, err = r.db.Query(`select p.id, p.employeeid, u.name, p.assetName, p.requestDate, p.status, p.description, p.spesification, p.managerid, us.name,p.updatedAt from procurements p
		join users u on p.employeeid=u.id
		left join users us on p.managerid=us.id where p.status=?`, status)
	} else {
		result, err = r.db.Query(`select p.id, p.employeeid, u.name, p.assetName, p.requestDate, p.status, p.description, p.spesification, p.managerid, us.name,p.updatedAt from procurements p
		join users u on p.employeeid=u.id
		left join users us on p.managerid=us.id`)
	}
	if err != nil {
		fmt.Println("1", err)
		log.Fatal(err)
	}
	// limit := 10
	// offset := (page - 1) * limit
	// if status != "" && page != 0 {
	// 	result, err = r.db.Query(`select p.id, p.employeeid, u.name, p.assetName, p.requestDate, p.status, p.updatedAt from procurements p
	// 	join users u on p.employeeid=u.id where p.status=? limit ? offset ?`, status, limit, offset)
	// } else if page != 0 {
	// 	result, err = r.db.Query(`select p.id, p.employeeid, u.name, p.assetName, p.requestDate, p.status, p.updatedAt from procurements p
	// 	join users u on p.employeeid=u.id limit ? offset ?`, limit, offset)
	// }
	// if err != nil {
	// 	fmt.Println("1", err)
	// 	log.Fatal(err)
	// }

	var procurements []entities.ProcurementResponseFormat

	defer result.Close()

	for result.Next() {
		var procurement entities.ProcurementResponseFormat
		err := result.Scan(&procurement.ID, &procurement.EmployeeId, &procurement.EmployeeName, &procurement.AssetName, &procurement.RequestDate, &procurement.Status, &procurement.Description, &procurement.Spesification, &procurement.ManagerId, &procurement.ManagerName, &procurement.UpdatedAt)
		if err != nil {
			fmt.Println("2", err)
			log.Fatal("error di scan get procurement")
		}
		procurements = append(procurements, procurement)
	}

	// var totalProcurementQuery *sql.Rows
	// if status != "" {
	// 	totalProcurementQuery, err = r.db.Query(`select count(id) from procurements where status=? `, status)
	// } else {
	// 	totalProcurementQuery, err = r.db.Query(`select count(id) from procurements`)
	// }
	// if err != nil {
	// 	fmt.Println("Get 2", err)
	// 	return procurements, totalProcurement, err
	// }

	// defer totalProcurementQuery.Close()

	// for totalProcurementQuery.Next() {
	// 	err := totalProcurementQuery.Scan(&totalProcurement)
	// 	if err != nil {
	// 		fmt.Println("3", err)
	// 		return procurements, totalProcurement, err
	// 	}
	// }
	return procurements, nil

}

func (r *ProcurementRepository) GetById(id int) (entities.ProcurementResponseFormat, error) {
	var procurement entities.ProcurementResponseFormat
	stmt, err := r.db.Prepare(`select p.id, p.employeeid, u.name, p.assetName, p.requestDate, p.status, p.description, p.spesification, p.managerid, us.name,p.updatedAt from procurements p
	join users u on p.employeeid=u.id
	left join users us on p.managerid=us.id where p.id=?`)
	if err != nil {
		//log.Fatal(err)
		log.Println(err)
		return procurement, fmt.Errorf("gagal prepare db")
	}

	result, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
		return procurement, fmt.Errorf("gagal query procurement")
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&procurement.ID, &procurement.EmployeeId, &procurement.EmployeeName, &procurement.AssetName, &procurement.RequestDate, &procurement.Status, &procurement.Description, &procurement.Spesification, &procurement.ManagerId, &procurement.ManagerName, &procurement.UpdatedAt)
		if err != nil {
			log.Println(err)
			return procurement, err
		}
		return procurement, nil
	}
	return procurement, fmt.Errorf("procurement not found")
}

func (r *ProcurementRepository) Create(procurement entities.Procurement) (entities.Procurement, error) {
	stmt, err := r.db.Prepare("insert into procurements(employeeid, assetName, spesification, description, status) values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return procurement, errors.New("internal server error")
	}

	_, errr := stmt.Exec(procurement.EmployeeId, procurement.AssetName, procurement.Spesification, procurement.Description, procurement.Status)
	if errr != nil {
		log.Println(errr)
		return procurement, errors.New("internal server error")
	}

	return procurement, nil
}

func (r *ProcurementRepository) Update(id int, procurement entities.Procurement) error {
	var res sql.Result
	var err error
	if procurement.ManagerId != 0 {
		res, err = r.db.Exec(`update procurements set status =?, updatedat = ?, managerid=? where id = ?`, procurement.Status, time.Now(), procurement.ManagerId, id)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		res, err = r.db.Exec(`update procurements set status =?, updatedat = ? where id = ?`, procurement.Status, time.Now(), id)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	notAffected, _ := res.RowsAffected()
	if notAffected == 0 {
		log.Println("rows affected is 0 while update procuremnt")
		return errors.New("internal server error")
	}
	return nil
}
