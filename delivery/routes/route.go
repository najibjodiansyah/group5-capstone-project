package routes

import (
	"capstone-project/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userConstroller *user.UserController) {
	api := e.Group("/api/v1")
	//User
	api.POST("/users", userConstroller.Register())
	api.GET("/users/:id", userConstroller.GetById())
	api.PUT("/users/:id", userConstroller.Update())
	api.DELETE("/users/:id", userConstroller.Delete())
}
