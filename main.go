package main

import (
	"capstone-project/config"
	"capstone-project/delivery/controllers/application"
	_assetController "capstone-project/delivery/controllers/asset"
	_authController "capstone-project/delivery/controllers/auth"
	_categoryController "capstone-project/delivery/controllers/category"
	_itemController "capstone-project/delivery/controllers/item"
	_procurementController "capstone-project/delivery/controllers/procurement"
	_statisticController "capstone-project/delivery/controllers/statistic"
	_userController "capstone-project/delivery/controllers/user"
	"capstone-project/delivery/routes"
	_appRepo "capstone-project/repository/application"
	_assetRepo "capstone-project/repository/asset"
	_authRepo "capstone-project/repository/auth"
	_categoryRepo "capstone-project/repository/category"
	_itemRepo "capstone-project/repository/item"
	_procurementRepo "capstone-project/repository/procurement"
	_statisticRepo "capstone-project/repository/statistic"
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
	itemRepo := _itemRepo.New(db)
	assetRepo := _assetRepo.New(db)
	procurementRepo := _procurementRepo.New(db)
	appRepo := _appRepo.NewApplication(db)
	categoryRepo := _categoryRepo.New(db)
	statisticRepo := _statisticRepo.New(db)

	// initialize controller
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	itemController := _itemController.New(itemRepo)
	assetController := _assetController.New(assetRepo)
	procurementController := _procurementController.New(procurementRepo)
	appController := application.New(appRepo)
	categoryController := _categoryController.New(categoryRepo)
	statisticController := _statisticController.New(statisticRepo)

	// create new echo
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS())
	routes.RegisterPath(e, userController, authController, assetController, itemController, appController, procurementController, categoryController, statisticController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
