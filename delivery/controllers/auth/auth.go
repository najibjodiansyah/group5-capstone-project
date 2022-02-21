package auth

import (
	response "capstone-project/delivery/commons"
	"capstone-project/delivery/middlewares"
	"capstone-project/entities"
	authRepo "capstone-project/repository/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	repository authRepo.Auth
}

func New(repository authRepo.Auth) *AuthController {
	return &AuthController{repository: repository}
}


func (ac AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {

	var input LoginRequestFormat

	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data")) // entity tidak bisa diproses
	}

	loginData, err := ac.repository.Login(input.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
	}

	// detect unauthorized login (email unknown)
	if loginData == (entities.User{}) {
		return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Email is unknown"))
	}

	// detect unauhorized login (password mismatch)
	if err = bcrypt.CompareHashAndPassword([]byte(loginData.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("failed", "Password does not match"))
	}

	token, err := middlewares.CreateToken(loginData.ID,loginData.Role)

	// detect failure in creating token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerError("failed", err.Error()))
	}

	var responseFormat LoginResponseFormat
	responseFormat.Id = loginData.ID
	responseFormat.Name = loginData.Name
	responseFormat.Token = token
	responseFormat.Role = loginData.Role

	return c.JSON(http.StatusOK, response.SuccessOperation("success", "success login", responseFormat))
}
}