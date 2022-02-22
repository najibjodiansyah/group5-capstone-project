package asset

import (
	"capstone-project/entities"
	"database/sql"
	"errors"
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

	res, errr := stmt.Exec(asset.Name, asset.Description, asset.Category, asset.Quantity, asset.Picture)
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
	stmt, err := ar.db.Prepare("select a.id,a.name,a.description,c.name,a.quantity,a.picture from assets as a inner join categories as c on a.categoryid = c.id  where id = ?")
	if err != nil {
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

	errScan := res.Scan(&asset.Id, &asset.Name, &asset.Name, &asset.Description, &asset.Category, &asset.Picture)
	if errScan != nil {
		return asset, errScan
	}

	return asset, nil
}
func (ar *AssetRepository) GetAll(status string, category string)(entities.Asset,error){
	var asset entities.Asset
	return asset, nil
}