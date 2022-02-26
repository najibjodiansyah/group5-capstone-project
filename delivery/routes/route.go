package routes

import (
	"capstone-project/delivery/controllers/auth"
	"capstone-project/delivery/controllers/item"
	"capstone-project/delivery/controllers/user"
	"capstone-project/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userConstroller *user.UserController,
	authController *auth.AuthController,
	itemController *item.ItemController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	api.GET("/users/:id", userConstroller.GetById(), middlewares.JWTMiddleware())
	api.PUT("/users/:id", userConstroller.Update(), middlewares.JWTMiddleware())
	api.DELETE("/users/:id", userConstroller.Delete(), middlewares.JWTMiddleware())

	//Item
	api.GET("/items", itemController.Get())
	api.GET("/items/:id", itemController.GetById())
	api.PUT("/items/:id", itemController.Update())

	//Auth
	api.POST("/login", authController.Login())
}
