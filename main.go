package main

import (
	"capstone-project/config"
	"capstone-project/delivery/controllers/application"
	_assetController "capstone-project/delivery/controllers/asset"
	_authController "capstone-project/delivery/controllers/auth"
	_userController "capstone-project/delivery/controllers/user"
	"capstone-project/delivery/routes"
	_appRepo "capstone-project/repository/application"
	_assetRepo "capstone-project/repository/asset"
	_authRepo "capstone-project/repository/auth"
	_userRepo "capstone-project/repository/user"
	"fmt"
	"log"
	"os"
	"time"

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
	assetRepo := _assetRepo.New(db)
	appRepo := _appRepo.NewApplication(db)

	// initialize controller
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	assetController := _assetController.New(assetRepo)
	appController := application.New(appRepo)

	// create new echo
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS())
	routes.RegisterPath(e, userController, authController, assetController, appController)


	fmt.Println(time.Parse("2006-01-02", "2025-01-01"))
	fmt.Println(time.Now().Add(time.Hour * (24*365)))

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
