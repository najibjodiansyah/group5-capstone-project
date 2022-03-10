package application

import (
	response "capstone-project/delivery/commons"
	"capstone-project/delivery/middlewares"
	"capstone-project/entities"
	applicationRepo "capstone-project/repository/application"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ApplicationController struct{
	repository applicationRepo.Application
}

func New(repository applicationRepo.Application) *ApplicationController {
	return &ApplicationController{repository: repository}
}

func (ac ApplicationController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}
		
		var input InputApp

		if role == "employee" {
			input.Employeeid = id
		}
		
		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		var app entities.Applications

		app.Employeeid = input.Employeeid
		app.AssetId	= input.Assetid
		if input.Returndate == "" {
			app.Returndate = time.Now().Add(time.Hour * (24*365))
		} else {
			returndate, _ := time.Parse("2006-01-02", input.Returndate)
			app.Returndate = returndate
		}
		app.Specification = input.Specification
		app.Description = input.Description
		if role == "admin" {
			app.Status = "tomanager"
		} else if role == "employee" {
			app.Status = "toadmin"
		}

		appId, _, errRepo := ac.repository.Create(app)
		if errRepo != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
		}

		app.Id = appId

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create applications"))
	}
}

func (ac ApplicationController) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		var input InputStatus

		appId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		if input.Status == "tomanager" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			if err := ac.repository.UpdateStatus(appId,input.Status,nil,nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status to 'tomanager' by admin"))
		
		}else if input.Status == "toaccept" {
			if role != "manager"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			var managerid *int = &userid

			if err := ac.repository.UpdateStatus(appId,input.Status,managerid,nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'accept' by manager"))
		
		}else if input.Status == "decline" {
			if role == "manager"{
				if err := ac.repository.UpdateStatus(appId,input.Status,&userid,nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by manager"))
			}else if role == "admin"{
				if err := ac.repository.UpdateStatus(appId,input.Status,nil,nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by admin"))
			}

			return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			
		}else if input.Status == "inuse" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			assetId, err := ac.repository.GetAsset(appId)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			availItemId, err := ac.repository.AvailabilityItem(assetId)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed check availability item"))
			}

			if availItemId == 0 {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "item is not available"))
			}
			
			if err := ac.repository.UpdateStatus(appId,input.Status,nil,&availItemId); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			app, err := ac.repository.GetById(appId)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			availstatus := "digunakan"
		
			if err := ac.repository.UpdateItem(&availItemId,availstatus,app.Employeeid); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'inuse' and update item table by admin"))

		}else if input.Status == "toreturn" {
			if role != "manager"{
				if err := ac.repository.UpdateStatus(appId, input.Status, nil, nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'toreturn'"))
			}
			return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Manager cant ask for return"))
			
		}else if input.Status == "donereturn" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			if err := ac.repository.UpdateStatus(appId, input.Status, nil, nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			app, err := ac.repository.GetById(appId)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			availstatus := "tersedia"
			
			if err := ac.repository.UpdateItem(app.Itemid, availstatus, -1); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed","Failed to update item table"))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'donereturn'"))
		}else if input.Status == "askreturn" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			if err := ac.repository.UpdateStatus(appId, input.Status, nil, nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'askreturn'"))
		}

		return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "status not recognized"))
	}

}

func (ac ApplicationController) GetById()echo.HandlerFunc{
	return func(c echo.Context)error {
		appId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		app, err := ac.repository.GetById(appId)
		if err != nil{
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset by id",app))
	}
}

// GetAll
func (ac ApplicationController) GetAll() echo.HandlerFunc{
	return func(c echo.Context)error {
		var status,category,date,orderbydate,longestdate string
		status = c.QueryParam("status")
		category = c.QueryParam("category")
		date = c.QueryParam("date")
		orderbydate = c.QueryParam("orderbydate")
		longestdate = c.QueryParam("longestdate")
		// pagination = c.QueryParam("page")

		if category == "" {
			category = "0"
		}
		// if pagination == "" {
		// 	pagination = "0"
		// }

		categoryid, err := strconv.Atoi(category)
			if err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert category_id"))
			}

		// page, err := strconv.Atoi(pagination)
		// 	if err != nil {
		// 		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert category_id"))
		// 	}
		app,totalAsset, err := ac.repository.GetAll(status,categoryid,date,orderbydate,longestdate)
		if err!= nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		limit := 10
		totalPage := int(math.Ceil(float64(totalAsset) / float64(limit)))

		var responseData responseAll
		responseData.Totalpage = totalPage
		responseData.App = app

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all applications", responseData))
	}
}

func (ac ApplicationController) UsersApplicationActivity() echo.HandlerFunc{
	return func(c echo.Context) error {
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		apps, err := ac.repository.UsersApplicationActivity(userid)
		if err != nil{
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset by id",apps))
	}
}
func (ac ApplicationController) UsersApplicationHistory() echo.HandlerFunc{
	return func(c echo.Context) error {
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		apps, err := ac.repository.UsersApplicationHistory(userid)
		if err != nil{
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset by id",apps))
	}
}