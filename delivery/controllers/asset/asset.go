package asset

import (
	response "capstone-project/delivery/commons"
	"capstone-project/entities"
	assetRepo "capstone-project/repository/asset"
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
		asset.Category = input.Category
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