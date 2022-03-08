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

	stmt, err := ar.db.Prepare(`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
	inner join categories as c on a.categoryid = c.id  
	where a.id = ?`)
	if err != nil {
		log.Println(err)
		return asset, errors.New("internal server error")
	}

	res, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
		return asset, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return asset, errors.New("internal server error")
	}

	errScan := res.Scan(&asset.Id, &asset.Name, &asset.Description, &asset.CategoryId, &asset.CategoryName, &asset.Quantity, &asset.Picture, &asset.CreatedAt)
	if errScan != nil {
		log.Println(errScan)
		return asset, errors.New("internal server error")
	}

	return asset, nil
}

func (ar *AssetRepository) GetAll(page int, category int, keyword string)([]entities.Asset,int,error){
	var (assets []entities.Asset
	 	totalAsset int
		err error
	 	result *sql.Rows)
		 
	limit := 10
	offset := (page - 1) * limit

	if category != 0 && page != 0 {
		result, err = ar.db.Query(
			`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat from assets as a 
			inner join categories as c on a.categoryid = c.id 
			where a.categoryid= ?
			order by a.id desc
			limit ? offset?`, category, limit, offset)
	}else if keyword != "" && page != 0 {
		word := "%" + keyword + "%"
		query := fmt.Sprintf(
			`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat 
			from assets as a 
			inner join categories as c on a.categoryid = c.id
			where upper(a.name) like '%v' 
			order by a.id desc
			limit %v offset %v`, word, limit, offset)
		result, err = ar.db.Query(query)
	}else if page != 0 {
		result, err = ar.db.Query(
			`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat 
			from assets as a 
			inner join categories as c on a.categoryid = c.id 
			order by a.id desc
			limit ? offset ?`, limit, offset)
	}else{
		result, err = ar.db.Query(
			`select a.id, a.name, a.description, a.categoryid, c.name, a.quantity, a.picture, a.createdat 
			from assets as a 
			inner join categories as c on a.categoryid = c.id 
			order by a.id desc`)
	}

	if err != nil {
		log.Println(err)
		return assets,totalAsset, errors.New("internal server error")
	}

	defer result.Close()

	for result.Next() {
		var asset entities.Asset
		err := result.Scan(&asset.Id, &asset.Name, &asset.Description, &asset.CategoryId, &asset.CategoryName, &asset.Quantity, &asset.Picture, &asset.CreatedAt)
		if err!= nil {
			log.Println(err)
			return assets,totalAsset, errors.New("internal server error")
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
		log.Println(err)
		return assets,totalAsset, errors.New("internal server error")
	}

	defer totalAssetQuery.Close()
	
	for totalAssetQuery.Next() {
		err := totalAssetQuery.Scan(&totalAsset)
		if err != nil {
			log.Println(err)
			return assets,totalAsset, errors.New("internal server error")
		}
	}
	return assets, totalAsset, nil
}

func (ar *AssetRepository) GetCountAssetUsed(assetid int)(int, error){
	stmt, err := ar.db.Prepare(`select count(id) as total from items where availableStatus = "digunakan" or availableStatus = "pemeliharaan" and assetid = ? group by assetid`)
	if err != nil {
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	res, err := stmt.Query(assetid)
	if err != nil {
		log.Println(err)
		return 0, errors.New("internal server error")
	}

	defer res.Close()

	if isExist := res.Next(); !isExist {
		return 0, nil
	}

	var total int

	errScan := res.Scan(&total)
	if errScan != nil {
		log.Println(errScan)
		return 0, errors.New("internal server error")
	}

	return total , nil
}

func (ar *AssetRepository) Update(idasset int, asset entities.Asset)(entities.Asset, error){

	stmt, err := ar.db.Prepare(`UPDATE assets SET description= ? WHERE id = ?`)
	if err != nil {
		log.Println(err)
		return asset, errors.New("internal server error")
	}

	result, err := stmt.Exec(asset.Description, idasset)
	if err != nil {
		log.Println(err)
		return asset, errors.New("internal server error")
	}

	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		log.Println("rows affected is 0 while update data")
		return asset, errors.New("internal server error")
	}

	return asset, nil
}