package asset

import (
	response "capstone-project/delivery/commons"
	"capstone-project/entities"
	assetRepo "capstone-project/repository/asset"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AssetController struct {
	repository assetRepo.Asset
}

func New(repository assetRepo.Asset) *AssetController {
	return &AssetController{repository: repository}
}

func (ac AssetController)Create()echo.HandlerFunc {
	return func(c echo.Context) error {
	
		var input RequestAssetFormat

		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		asset := entities.Asset{}
		asset.Name = input.Name
		asset.Description = input.Description
		asset.Category.Id = input.Category
		asset.Quantity = input.Quantity
		asset.Picture = input.Picture

		_, assetId, err := ac.repository.Create(asset)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
		}

		assetid := assetId

		for i := 1; i <= asset.Quantity; i++ {
			id := strconv.Itoa(i)
			assetName :=  asset.Name +" "+ id
			if err := ac.repository.GenerateItem(assetName, assetid); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create asset and generate item"))
	}
}

func (ac AssetController)GetById()echo.HandlerFunc{
	return func(c echo.Context) error {

		assetId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		asset, err := ac.repository.GetById(assetId)
		if err != nil{
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}
		
		var Responseasset ResponeAssetFormat
		Responseasset.Id = asset.Id
		Responseasset.Name = asset.Name
		Responseasset.Description = asset.Description
		Responseasset.Categoryid = asset.Category.Id
		Responseasset.Category = asset.Category.Name
		Responseasset.Quantity = asset.Quantity
		Responseasset.Picture = asset.Picture
		Responseasset.CreatedAt = asset.CreatedAt
		
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset by id",Responseasset))
	}
}

func (ac AssetController)GetAll()echo.HandlerFunc{
	return func(c echo.Context) error {
		var assets []ResponeAssetFormat
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all asset", assets))
	}
}