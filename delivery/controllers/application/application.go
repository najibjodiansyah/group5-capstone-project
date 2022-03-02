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
			fmt.Println("string koshong")
			app.Returndate = time.Now().Add(time.Hour * (24*365))
			fmt.Println(app.Returndate)
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

// buat repo availability((cek pengecekan availibility itemid) urutin id asc limit 1 where status item = tersedia), main di response message
// cek availibility jika ada hasilnya maka ubah status = digunakan, 1 update kolom status, 2 update itemid , 3 update availablestatus pada table item menjadi digunakan

// if status = disetujui, 1 update kolom status, 2 update kolom manager id (dari token id)

func (ac ApplicationController) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		var input Inputstatus

		appid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		// belom bisa ngecek status sebelumnya pas mau update status apakah sudah melewati tahap sebelumnya, 
		// nanti bakal bisa di bypass sama admin langsung setuju dari body endpoint,
		// admin bisa mencurangin sistem

		if input.Status == "tomanager" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			if err := ac.repository.UpdateStatus(appid,input.Status,nil,nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status to 'tomanager' by admin"))
		
		}else if input.Status == "toaccept" {
			// jika if status = diterima maka 1 update kolom status, 2 update kolom manager id (dari token id)
			
			if role != "manager"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			var managerid *int = &userid

			if err := ac.repository.UpdateStatus(appid,input.Status,managerid,nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'accept' by manager"))
		}else if input.Status == "decline" {
			if role == "manager"{
				if err := ac.repository.UpdateStatus(appid,input.Status,&userid,nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by manager"))
			}else if role == "admin"{
				if err := ac.repository.UpdateStatus(appid,input.Status,nil,nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by admin"))
			}

			return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			
		}else if input.Status == "inuse" {
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			assetid, err := ac.repository.GetAsset(appid)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			// dapetin id item yang tersedia
			availitemid, err := ac.repository.AvailabilityItem(assetid)
			if err != nil {
				//-> return msg = no item available
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed check availability item"))
			}

			if availitemid == 0 {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "item is not available"))
			}
			
			// update status di application 
			if err := ac.repository.UpdateStatus(appid,input.Status,nil,&availitemid); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			// panggil method get app by id, buat ambil employeeid dan assetid where id = appid
			app, err := ac.repository.GetById(appid)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			availstatus := "digunakan"

			// panggil method untuk ngubah itemstatus
		
			if err := ac.repository.UpdateItem(&availitemid,availstatus,app.Employeeid); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'inuse' and update item table by admin"))

		}else if input.Status == "toreturn" {
			// harusnya bisa admin, bisa juga employee
			// disini bikin status untuk dikembalikan
			if role != "manager"{
				if err := ac.repository.UpdateStatus(appid, input.Status, nil, nil); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'toreturn'"))
			}
			return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Manager cant ask for return"))
			
		}else if input.Status == "donereturn" {
			// admin get application dimana statusnya adalah toreturn
			// 1. admin mengubah status application menjadi donereturn
			// 2. admin juga harus mengupdate table item dimana itemid yang diambil dari get app by id, update availablestatus pada table item menjadi tersedia
			
			if role != "admin"{
				return c.JSON(http.StatusBadRequest,response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			// update status di application jadi done return
			if err := ac.repository.UpdateStatus(appid, input.Status, nil, nil); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			// panggil method get app by id, buat ambil itemid where id = appid
			app, err := ac.repository.GetById(appid)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed get app for item id"))
			}

			// panggil method untuk ngubah itemstatus
			// employee diupdate jadi null
			availstatus := "tersedia"
			if err := ac.repository.UpdateItem(app.Itemid, availstatus, -1); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed","Failed to update item table"))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'donereturn'"))
		}

		return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "status not recognized"))
	}

}

func (ac ApplicationController) GetById()echo.HandlerFunc{
	return func(c echo.Context)error {
		appid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		app, err := ac.repository.GetById(appid)
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