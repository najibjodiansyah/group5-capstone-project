package user

import (
	"capstone-project/entities"
	userRepo "capstone-project/repository/user"
	"fmt"
	"net/http"

	response "capstone-project/delivery/commons"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)
type UserController struct {
	repository userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{repository: user}
}


func (uc *UserController)Register() echo.HandlerFunc{
	return func(c echo.Context) error {
	
	var input RegisterUserFormat
	
	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data")) // entity tidak bisa diproses
	}

	user := entities.User{}
	user.Name = input.Name
	user.Email = input.Email

	hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if errEncrypt != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
	}
	user.Password = string(hashedPassword)

	// create user to database
	_, err := uc.repository.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user"))
	}
}