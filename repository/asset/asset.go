package asset

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type AssetRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (ar *AssetRepository) Create(asset entities.Asset)(entities.Asset,int,error){
	stmt, err := ar.db.Prepare("insert into assets(name, description, categoryid, quantity, picture) values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return asset,0, errors.New("internal server error")
	}

	res, errr := stmt.Exec(asset.Name, asset.Description, asset.CategoryId, asset.Quantity, asset.Picture)
	if errr != nil {
		log.Println(errr)
		return asset,0, errors.New("internal server error")
	}

	id,_ := res.LastInsertId()


	return asset, int(id), nil
}

func (ar *AssetRepository) GenerateItem(assetName string, assetId int) error {
	stmt, err := ar.db.Prepare("insert into items(name, assetId, availableStatus) values(?,?,?)")
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	_, errr := stmt.Exec(assetName,assetId,"tersedia")
	if errr != nil {
		log.Println(errr)
		return errors.New("internal server error")
	}

	return nil
}

func (ar *AssetRepository) GetById(id int)(entities.Asset,error){
	var asset entities.Asset
	stmt, err := ar.db.Prepare("select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a inner join categories as c on a.categoryid = c.id  where a.id = ?")
	if err != nil {
		fmt.Println(err)
		return asset, errors.New("internal server error")
	}
	res, err := stmt.Query(id)
	if err != nil {
		return asset, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return asset, errors.New("internal server error")
	}

	errScan := res.Scan(&asset.Id, &asset.Name, &asset.Description, &asset.CategoryId, &asset.CategoryName, &asset.Quantity, &asset.Picture, &asset.CreatedAt)
	if errScan != nil {
		return asset, errScan
	}

	return asset, nil
}

func (ar *AssetRepository) GetAll(page int, category int)([]entities.Asset,int,error){
	var assets []entities.Asset
	// var condition string = ""

	// if category != 0 {
	// 	condition += " where a.id = ?"
	// }

	// if status != "" {
	// 	switch status {
	// 	case "digunakan":
	// 		condition += " and a.status = digunakan"
	// 	case "tersedia":
	// 		condition += " and a.status = tersedia"
	// 	case "pemeliharaan":
	// 		condition += " and a.status = pemeliharaan"
	// 	}
	// }
	// if category != 0 {
	// 	query := "select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a inner join categories as c on a.categoryid = c.id where a.categoryid= ?"
	// }

	
	// stmt, err := ar.db.Prepare(query)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// res, err := stmt.Query(status,category)
	// if err != nil {
	// 	fmt.Println("2",err)
	// 	return assets, errors.New("internal server error")
	// }
	var totalAsset int
	var err error
	var result *sql.Rows
	limit := 10
	offset := (page - 1) * limit

	if category == 0 && page == 0 {
		result, err = ar.db.Query(`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
			inner join categories as c on a.categoryid = c.id`)
	}else if page != 0 && category == 0{
		result, err = ar.db.Query(`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
			inner join categories as c on a.categoryid = c.id limit ? offset?` ,limit, offset)
	}else if category != 0 && page == 0{
		result, err = ar.db.Query(`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
			inner join categories as c on a.categoryid = c.id 
			where a.categoryid= ?`, category)
	} else if category != 0 && page != 0 {
		result, err = ar.db.Query(
			`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
			inner join categories as c on a.categoryid = c.id 
			where a.categoryid= ?
			limit ? offset?`, category, limit, offset)
	}
	if err != nil {
		fmt.Println("Get 1", err)
		return assets,totalAsset, err
	}

	defer result.Close()

	for result.Next() {
		var asset entities.Asset
		err := result.Scan(&asset.Id, &asset.Name, &asset.Description, &asset.CategoryId, &asset.CategoryName, &asset.Quantity, &asset.Picture, &asset.CreatedAt)
		if err!= nil {
			return assets,totalAsset, err
		}
		assets = append(assets, asset)
	}
	var totalAssetQuery *sql.Rows
	if category == 0 {
		totalAssetQuery, err = ar.db.Query(`select count(id) from assets`)
	}else {
		totalAssetQuery, err = ar.db.Query(`select count(id) from assets where categoryid= ?`, category)
	}
	if err != nil {
		fmt.Println("Get 2", err)
		return assets,totalAsset, err
	}

	defer totalAssetQuery.Close()
	
	for totalAssetQuery.Next() {
		err := totalAssetQuery.Scan(&totalAsset)
		if err != nil {
			return assets,totalAsset, err
		}
	}
	return assets, totalAsset, nil
}

func (ar *AssetRepository)GetCountAssetUsed(assetid int)(int, error){
	stmt, err := ar.db.Prepare(`select count(id) as total from items where availableStatus = "digunakan" or availableStatus = "pemeliharaan" and assetid = ? group by assetid`)
	if err != nil {
		fmt.Println("Get 1", err)
		return 0, err
	}

	res, err := stmt.Query(assetid)
	if err != nil {
		fmt.Println("Get 2", err)
		return 0, err
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return 0, nil
	}

	var total int

	errScan := res.Scan(&total)
	if errScan != nil {
		return 0, errScan
	}

	return total , nil
}