package routes

import (
	"capstone-project/delivery/controllers/asset"
	"capstone-project/delivery/controllers/auth"
	"capstone-project/delivery/controllers/item"
	"capstone-project/delivery/controllers/procurement"
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
	procurementController *procurement.ProcurementController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	// api.GET("/users/:id", userConstroller.GetById(), middlewares.JWTMiddleware())
	// api.PUT("/users/:id", userConstroller.Update(), middlewares.JWTMiddleware())
	// api.DELETE("/users/:id", userConstroller.Delete(), middlewares.JWTMiddleware())

	//Auth
	api.POST("/login", authController.Login())

	//Asset
	api.POST("/assets", assetController.Create())
	api.GET("/assets/:id", assetController.GetById())
	api.GET("/assets", assetController.GetAll())

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
}
