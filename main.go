package main

import (
	"capstone-project/config"
	_authController "capstone-project/delivery/controllers/auth"
	_userController "capstone-project/delivery/controllers/user"
	"capstone-project/delivery/routes"
	_authRepo "capstone-project/repository/auth"
	_userRepo "capstone-project/repository/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initialize database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := config.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize model
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)

	// initialize controller
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)

	// create new echo
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS())
	routes.RegisterPath(e, userController, authController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
