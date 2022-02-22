package asset

import "capstone-project/entities"

type Asset interface {
	Create(asset entities.Asset)(entities.Asset,int,error)
	GenerateItem(assetName string,assetId int)error
	GetById(id int)(entities.Asset,error)
	GetAll(status string, category string)(entities.Asset,error)
}