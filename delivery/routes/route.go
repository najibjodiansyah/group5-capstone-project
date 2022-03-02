package routes

import (
	"capstone-project/delivery/controllers/application"
	"capstone-project/delivery/controllers/asset"
	"capstone-project/delivery/controllers/auth"
	"capstone-project/delivery/controllers/category"
	"capstone-project/delivery/controllers/item"
	"capstone-project/delivery/controllers/procurement"
	"capstone-project/delivery/controllers/statistic"
	"capstone-project/delivery/controllers/user"
	"capstone-project/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userConstroller *user.UserController,
	authController *auth.AuthController,
	assetController *asset.AssetController,
	itemController *item.ItemController,
	appController *application.ApplicationController,
	procurementController *procurement.ProcurementController,
	categoryController *category.CategoryController,
	statisticController *statistic.StatisticController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	api.GET("/employees", userConstroller.GetEmployees())
	// api.GET("/users/:id", userConstroller.GetById(), middlewares.JWTMiddleware())
	// api.PUT("/users/:id", userConstroller.Update(), middlewares.JWTMiddleware())
	// api.DELETE("/users/:id", userConstroller.Delete(), middlewares.JWTMiddleware())

	//Auth
	api.POST("/login", authController.Login())

	//Asset
	api.POST("/assets", assetController.Create())
	api.GET("/assets/:id", assetController.GetById())
	api.GET("/assets", assetController.GetAll())

	//Application
	api.POST("/applications", appController.Create(), middlewares.JWTMiddleware())
	api.PUT("/applications/:id", appController.UpdateStatus(), middlewares.JWTMiddleware())
	api.GET("/applications/:id", appController.GetById())
	api.GET("/applications", appController.GetAll())
	api.GET("/users/:id/applications/activity", appController.UsersApplicationActivity())
	api.GET("/users/:id/applications/history", appController.UsersApplicationHistory())
	//Item
	api.GET("/items", itemController.Get())
	api.GET("/items/:id", itemController.GetById())
	api.PUT("/items/:id", itemController.Update())
	api.GET("/items/:id/usage", itemController.GetItemUsageHistory())

	//Procurement
	api.POST("/procurements", procurementController.Create(), middlewares.JWTMiddleware())
	api.GET("/procurements", procurementController.Get())
	api.GET("/procurements/:id", procurementController.GetById())
	api.PUT("/procurements/:id", procurementController.Update(), middlewares.JWTMiddleware())

	//category
	api.GET("/categories", categoryController.Get())

	//statistic
	api.GET("/statistics", statisticController.Get())
}
