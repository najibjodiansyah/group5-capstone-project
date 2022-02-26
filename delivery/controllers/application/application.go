package application

import (
	response "capstone-project/delivery/commons"
	"capstone-project/delivery/middlewares"
	"capstone-project/entities"
	applicationRepo "capstone-project/repository/application"
	"fmt"
	"log"
	"net/http"
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

		// kalo admin yang input dapet data employeeid nya dari mana
		// bikin endpoint get all employee
		
		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity,response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		// setaun := strconv.Itoa(int(tahun))
		// fmt.Println(setaun)

		var app entities.Applications

		app.Employeeid = input.Employeeid
		app.AssetId	= input.Assetid
		if input.Returndate == "" {
			fmt.Println("string koshong")
			app.Returndate = time.Now().Add(time.Hour * (24*365))
			fmt.Println(app.Returndate)
		} else {
			fmt.Println("else")
			returndate, _ := time.Parse("2006-01-02", input.Returndate)
			app.Returndate = returndate
			fmt.Println(app.Returndate)
		}
		app.Specification = input.Specification
		app.Description = input.Description
		if role == "admin" {
			app.Status = "toManager"
		} else if role == "employee" {
			app.Status = "toAdmin"
		}

		fmt.Println(input)
		fmt.Println("/",app)

		appId, _, errRepo := ac.repository.Create(app)
		if errRepo != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
		}

		app.Id = appId

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create applications"))
	}
}