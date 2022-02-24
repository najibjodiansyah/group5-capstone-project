package routes

import (
	"capstone-project/delivery/controllers/asset"
	"capstone-project/delivery/controllers/auth"
	"capstone-project/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo, 
	userConstroller *user.UserController,
	authController *auth.AuthController,
	assetController *asset.AssetController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	// api.GET("/users/:id", userConstroller.GetById(), middlewares.JWTMiddleware())
	// api.PUT("/users/:id", userConstroller.Update(), middlewares.JWTMiddleware())
	// api.DELETE("/users/:id", userConstroller.Delete(),  middlewares.JWTMiddleware())

	//Auth
	api.POST("/login", authController.Login())

	//Asset
	api.POST("/assets", assetController.Create())
	api.GET("/assets/:id", assetController.GetById())
}
