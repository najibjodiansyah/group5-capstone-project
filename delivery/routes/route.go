package routes

import (
	"capstone-project/delivery/controllers/auth"
	"capstone-project/delivery/controllers/user"
	"capstone-project/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo, 
	userConstroller *user.UserController,
	authController *auth.AuthController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	api.GET("/users/:id", userConstroller.GetById(), middlewares.JWTMiddleware())
	api.PUT("/users/:id", userConstroller.Update(), middlewares.JWTMiddleware())
	api.DELETE("/users/:id", userConstroller.Delete(),  middlewares.JWTMiddleware())

	//Auth
	api.POST("/login", authController.Login())
}
