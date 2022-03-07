package asset

import (
	response "capstone-project/delivery/commons"
	"capstone-project/entities"
	assetRepo "capstone-project/repository/asset"
	"capstone-project/util"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		asset.CategoryId = input.Categoryid
		asset.Quantity = input.Quantity
		src, file, err := c.Request().FormFile("picture")
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to upload picture"))
		}
		ext := strings.Split(file.Filename, ".")
		extension := ext[len(ext)-1]
		check_extension := strings.ToLower(extension)
		if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file extention not allowed"))
		}
		if file.Size == 0 {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "illegal file size"))
		} else if file.Size > 1050000 {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file size exceeded the limit"))
		}

		file.Filename = fmt.Sprintf("%d-%d.%s", asset.Id, time.Now().Unix(), extension)

		sess := session.Must(util.GetAWSSession())

		uploader := s3manager.NewUploader(sess)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET")),
			Key:    aws.String(file.Filename),
			Body:   src,
		})

		// detect failure while uploading file
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, response.InternalServerError("failed", "Internal server error"))
		}
		asset.Picture = fmt.Sprintf("https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/%s", file.Filename)

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
		
		var Responseasset ResponseAssetFormat
		Responseasset.Id = asset.Id
		Responseasset.Name = asset.Name
		Responseasset.Description = asset.Description
		Responseasset.Categoryid = asset.CategoryId
		Responseasset.Category = asset.CategoryName
		Responseasset.Quantity = asset.Quantity
		Responseasset.Picture = asset.Picture
		Responseasset.CreatedAt = asset.CreatedAt

		total, err := ac.repository.GetCountAssetUsed(asset.Id)
			if err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
			}
			
			if total == 0 || total < asset.Quantity {
				Responseasset.Availability = "tersedia"
			}else {
				Responseasset.Availability = "tidak tersedia"
			}
			
		
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset by id",Responseasset))
	}
}

func (ac AssetController)GetAll()echo.HandlerFunc{
	return func(c echo.Context) error {

		var category,pagination,status,keyword string
		category = c.QueryParam("category")
		pagination = c.QueryParam("page")
		status = c.QueryParam("status") 
		keyword = c.QueryParam("keyword")
		if category == "" {
			category = "0"
		}
		if pagination == "" {
			pagination = "0"
		}

		categoryid, err := strconv.Atoi(category)
			if err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert category_id"))
			}

		page, err := strconv.Atoi(pagination)
			if err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert category_id"))
			}

		assets,totalAsset, err := ac.repository.GetAll(page,categoryid,keyword)
		if err!= nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		limit := 10
		totalPage := int(math.Ceil(float64(totalAsset) / float64(limit)))

		var assetResp []ResponseAssetFormat
		for _, asset := range assets {
			var Responseasset ResponseAssetFormat
			total, err := ac.repository.GetCountAssetUsed(asset.Id)
			if err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
			}
			
			if total == 0 || total < asset.Quantity {
				Responseasset.Availability = "tersedia"
			}else {
				Responseasset.Availability = "tidak tersedia"
			}
			
			Responseasset.Id = asset.Id
			Responseasset.Name = asset.Name
			Responseasset.Description = asset.Description
			Responseasset.Categoryid = asset.CategoryId
			Responseasset.Category = asset.CategoryName
			Responseasset.Quantity = asset.Quantity
			Responseasset.Picture = asset.Picture
			Responseasset.CreatedAt = asset.CreatedAt
			if status != "" && status == "tersedia" {
				if Responseasset.Availability == "tidak tersedia" {
					continue
				}
			} else if status == "tidaktersedia" && status != "" {
				if Responseasset.Availability == "tersedia" {
					continue
				}
			}
			assetResp = append(assetResp, Responseasset)
		}

		var responseData responseAll
		responseData.Totalpage = totalPage
		responseData.Assets = assetResp

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all asset", responseData))
	}
}

func (ac AssetController) Update()echo.HandlerFunc{
	return func(c echo.Context) error {
		asset := entities.Asset{}
		if err_bind := c.Bind(&asset); err_bind != nil {
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data"))
		}

		assetid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		current_asset, err := ac.repository.GetById(assetid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}

		if asset.Description != "" {
			current_asset.Description = asset.Description
		}

		// src, file, err := c.Request().FormFile("picture")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to upload avatar"))
		// }
		// if file == nil{
		// 	current_asset.Picture = asset.Picture
		// }else{
		// 	ext := strings.Split(file.Filename, ".")
		// 	extension := ext[len(ext)-1]
		// 	check_extension := strings.ToLower(extension)
		// 	if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
		// 		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file extention not allowed"))
		// 	}
		// 	if file.Size == 0 {
		// 		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "illegal file size"))
		// 	} else if file.Size > 1050000 {
		// 		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file size exceeded the limit"))
		// 	}
	
		// 	file.Filename = fmt.Sprintf("%d-%d.%s", assetid, time.Now().Unix(), extension)
	
		// 	sess := session.Must(util.GetAWSSession())
	
		// 	uploader := s3manager.NewUploader(sess)
	
		// 	_, err = uploader.Upload(&s3manager.UploadInput{
		// 		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		// 		Key:    aws.String(file.Filename),
		// 		Body:   src,
		// 	})
	
		// 	// detect failure while uploading file
		// 	if err != nil {
		// 		return c.JSON(http.StatusInternalServerError, response.InternalServerError("failed", "Internal server error"))
		// 	}
		// 	current_asset.Picture = fmt.Sprintf("https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/%s", file.Filename)
		// }
		
		_, err_update := ac.repository.Update(assetid, current_asset)
		if err_update != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to update data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update asset"))
	}
}