package user

import (
	"capstone-project/entities"
	userRepo "capstone-project/repository/user"
	"fmt"
	"net/http"
	"strconv"

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


func (uc UserController)Register() echo.HandlerFunc{
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

func (uc UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		user, err := uc.repository.GetById(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		var responseUser ResponseUserFormat
		responseUser.ID = user.ID
		responseUser.Name = user.Name
		responseUser.Email = user.Email
		responseUser.Avatar = user.Avatar
		responseUser.CreatedAt = user.CreatedAt

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user", responseUser))
	}
}

func (uc UserController)Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// NOTES : misal gada perubahan pas update / update data yang sama
		// NOTES : email gaboleh ganti dengan email yang sudah dipakai user lain
		user := entities.User{}
		if err_bind := c.Bind(&user); err_bind != nil {
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data")) 
		}
		// getting the id
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if errEncrypt != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
		}
		user.Password = string(hashedPassword)

		updateUser, err := uc.repository.GetById(userid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}
		
		if user.Name != "" {
			updateUser.Name = user.Name
		}
		if user.Email != "" {
			updateUser.Email = user.Email
		}
		if user.Password != "" {
			updateUser.Password = user.Password
		}
		if user.Avatar != "" {
			updateUser.Avatar = user.Avatar
		}

		err_update := uc.repository.Update(userid, updateUser)
		if err_update != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success Update user"))
	}
}

func (uc UserController)Delete()echo.HandlerFunc{
	return func(c echo.Context) error {
		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// delete user based on id from database
		errDelete := uc.repository.Delete(userId)
		if errDelete != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}