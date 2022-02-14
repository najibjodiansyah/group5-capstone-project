package main

import (
	"capstone-project/config"
	_userController "capstone-project/delivery/controllers/user"
	"capstone-project/delivery/routes"
	_userRepo "capstone-project/repository/user"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initialize database connection
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := config.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize model
	userRepo := _userRepo.New(db)

	// initialize controller
	userController := _userController.New(userRepo)

	// create new echo
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS())
	routes.RegisterPath(e, userController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
