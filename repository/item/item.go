package item

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type ItemRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Get(availableStatus string, category int, keyword string, page int) ([]entities.ItemResponseFormat, int, error) {
	var totalItem int
	var err error
	var result *sql.Rows
	limit := 10
	offset := (page - 1) * limit
	if availableStatus != "" && page != 0 {
		result, err = r.db.Query(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
		join assets a ON i.assetid = a.id
		join categories c ON a.categoryid = c.id where i.availableStatus=? limit ? offset ?`, availableStatus, limit, offset)
	} else if category != 0 && page != 0 {
		result, err = r.db.Query(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
		join assets a ON i.assetid = a.id
		join categories c ON a.categoryid = c.id where a.categoryid=? limit ? offset ?`, category, limit, offset)
	} else if keyword != "" && page != 0 {
		kw := "%" + keyword + "%"
		query := fmt.Sprintf(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
		join assets a ON i.assetid = a.id
		join categories c ON a.categoryid = c.id where upper(i.name) like '%v' limit %v offset %v`, kw, limit, offset)
		result, err = r.db.Query(query)
	} else if page != 0 {
		result, err = r.db.Query(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
		join assets a ON i.assetid = a.id
		join categories c ON a.categoryid = c.id limit ? offset ?`, limit, offset)
	}
	// stmt, err := r.db.Prepare(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
	// join assets a ON i.assetid = a.id
	// join categories c ON a.categoryid = c.id`)
	if err != nil {
		fmt.Println("1", err)
		log.Fatal(err)
	}

	var items []entities.ItemResponseFormat

	defer result.Close()

	for result.Next() {
		var item entities.ItemResponseFormat
		err := result.Scan(&item.ID, &item.Name, &item.CategoryId, &item.Category, &item.Picture, &item.AvailableStatus)
		if err != nil {
			fmt.Println("2", err)
			log.Fatal("error di scan get item")
		}
		items = append(items, item)
	}

	var totalItemQuery *sql.Rows
	if availableStatus != "" {
		totalItemQuery, err = r.db.Query(`select count(id) from items where availableStatus=? `, availableStatus)
	} else if category != 0 {
		totalItemQuery, err = r.db.Query(`select count(i.id) from items i
		join assets a on i.assetid =  a.id where a.categoryid=?`, category)
	} else if keyword != "" {
		kw := "%" + keyword + "%"
		query := fmt.Sprintf(`select count(id) from items where upper(name) like '%v'`, kw)
		totalItemQuery, err = r.db.Query(query)
	} else {
		totalItemQuery, err = r.db.Query(`select count(id) from items`)
	}
	if err != nil {
		fmt.Println("Get 2", err)
		return items, totalItem, err
	}

	defer totalItemQuery.Close()

	for totalItemQuery.Next() {
		err := totalItemQuery.Scan(&totalItem)
		if err != nil {
			fmt.Println("3", err)
			return items, totalItem, err
		}
	}
	return items, totalItem, nil

}

func (r *ItemRepository) GetById(id int) (entities.ItemResponseFormat, error) {
	var item entities.ItemResponseFormat
	stmt, err := r.db.Prepare(`select i.id, i.name, a.categoryid, c.name, a.picture, i.availablestatus from items i
	join assets a ON i.assetid = a.id
	join categories c ON a.categoryid = c.id where i.id=?`)
	if err != nil {
		//log.Fatal(err)
		return item, fmt.Errorf("gagal prepare db")
	}

	result, err := stmt.Query(id)
	if err != nil {
		return item, fmt.Errorf("gagal query item")
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.CategoryId, &item.Category, &item.Picture, &item.AvailableStatus)
		if err != nil {
			return item, err
		}
		return item, nil
	}

	return item, fmt.Errorf("user not found")
}

func (r *ItemRepository) GetByIdUpdate(id int) (entities.Item, error) {
	var item entities.Item
	stmt, err := r.db.Prepare(`select id, assetid, employee, name, availablestatus from items where id=?`)
	if err != nil {
		//log.Fatal(err)
		return item, fmt.Errorf("gagal prepare db")
	}

	result, err := stmt.Query(id)
	if err != nil {
		return item, fmt.Errorf("gagal query item")
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&item.ID, &item.AssetId, &item.EmployeeId, &item.Name, &item.AvailableStatus)
		if err != nil {
			return item, err
		}
		return item, nil
	}

	return item, fmt.Errorf("user not found")
}

func (r *ItemRepository) Update(id int, item entities.Item) error {
	stmt, err := r.db.Prepare("UPDATE items SET availableStatus= ? WHERE id = ?")
	if err != nil {
		return errors.New("internal server error")
	}
	result, err := stmt.Exec(item.AvailableStatus, id)
	if err != nil {
		return errors.New("internal server error")
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return errors.New("internal server error")
	}
	return nil
}

func (r *ItemRepository) GetItemUsageHistory(id int) (entities.ItemUsageHistory, error) {
	var itemHistory entities.ItemUsageHistory
	// stmt, err := r.db.Prepare(`select i.id, i.name, a.categoryid, c.name, a.picture from items i
	// join assets a ON i.assetid = a.id
	// join categories c ON a.categoryid = c.id where i.id=?`)
	// if err != nil {
	// 	//log.Fatal(err)
	// 	return itemHistory, fmt.Errorf("gagal prepare db")
	// }

	// result, err := stmt.Query(id)
	// if err != nil {
	// 	return itemHistory, fmt.Errorf("gagal query itemHistory")
	// }
	result, err := r.db.Query(`select i.id, i.name, a.categoryid, c.name, a.picture from items i
	join assets a ON i.assetid = a.id
	join categories c ON a.categoryid = c.id where i.id=?`, id)
	if err != nil {
		fmt.Println("1", err)
		log.Fatal(err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		return itemHistory, fmt.Errorf("item not found")
	}
	err = result.Scan(&itemHistory.ID, &itemHistory.Name, &itemHistory.CategoryId, &itemHistory.Category, &itemHistory.Picture)
	if err != nil {
		return itemHistory, err
	}

	// for result.Next() {
	// 	err := result.Scan(&itemHistory.ID, &itemHistory.Name, &itemHistory.CategoryId, &itemHistory.Category, &itemHistory.Picture)
	// 	if err != nil {
	// 		return itemHistory, err
	// 	}
	// }

	statusDigunakan := "digunakan"
	statusDkembalikan := "dikembalikan"
	res, err := r.db.Query(`select u.name, updatedAt, status from applications a
	join users u on a.employeeid=u.id where a.itemid=? and (a.status =? or a.status=?)`, id, statusDigunakan, statusDkembalikan)

	if err != nil {
		fmt.Println("1", err)
		log.Fatal(err)
	}

	// stmt, err = r.db.Prepare(`select u.name, updatedAt, status from applications a
	// join users u on a.employeeid=u.id where a.itemid=? and a.status =? or a.status=?`)
	// if err != nil {
	// 	//log.Fatal(err)
	// 	return itemHistory, fmt.Errorf("gagal prepare db")
	// }
	// result, err = stmt.Query(id, statusDigunakan, statusDkembalikan)
	// if err != nil {
	// 	return itemHistory, fmt.Errorf("gagal query item History")
	// }
	defer res.Close()

	var users []entities.ItemUser

	for res.Next() {
		var user entities.ItemUser
		err := res.Scan(&user.AssetUser, &user.LendingDate, &user.UsageStatus)
		if err != nil {
			fmt.Println("2", err)
			log.Fatal("error di scan get item History")
		}
		users = append(users, user)
	}
	itemHistory.Users = users

	return itemHistory, nil
}
