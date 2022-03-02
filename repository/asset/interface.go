package asset

import "capstone-project/entities"

type Asset interface {
	Create(asset entities.Asset)(entities.Asset,int,error)
	GenerateItem(assetName string,assetId int)error
	GetById(id int)(entities.Asset,error)
	GetAll(page int, category int)([]entities.Asset, int, error)
	GetCountAssetUsed(assetid int)(int, error)
}