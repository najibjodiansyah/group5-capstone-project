package user

import (
	"capstone-project/entities"
	userRepo "capstone-project/repository/user"
	"fmt"
	"log"
	"net/http"

	// "os"
	// "strconv"
	// "strings"
	// "time"

	response "capstone-project/delivery/commons"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository userRepo.User
}

func New(repository userRepo.User) *UserController {
	return &UserController{repository: repository}
}

func (uc UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterUserFormat

		if err := c.Bind(&input); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data")) // entity tidak bisa diproses
		}

		user := entities.User{}
		user.Name = input.Name
		user.Email = input.Email
		user.Role = input.Role
		user.Avatar = "https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg"

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

// func (uc UserController) GetById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _, err := middlewares.ExtractToken(c)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}

// 		userId, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
// 		}

// 		if userId != id {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}

// 		user, err := uc.repository.GetById(userId)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
// 		}

// 		var responseUser ResponseUserFormat
// 		responseUser.ID = user.ID
// 		responseUser.Name = user.Name
// 		responseUser.Email = user.Email
// 		responseUser.Avatar = user.Avatar
// 		responseUser.Role = user.Role
// 		responseUser.CreatedAt = user.CreatedAt

// 		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user", responseUser))
// 	}
// }

// func (uc UserController) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, role, err := middlewares.ExtractToken(c)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}
// 		// NOTES : misal gada perubahan pas update / update data yang sama
// 		// NOTES : email gaboleh ganti dengan email yang sudah dipakai user lain
// 		user := entities.User{}
// 		if err_bind := c.Bind(&user); err_bind != nil {
// 			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data"))
// 		}

// 		userid, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
// 		}

// 		if userid != id {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}

// 		fmt.Println(role)

// 		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 		if errEncrypt != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
// 		}
// 		user.Password = string(hashedPassword)

// 		updateUser, err := uc.repository.GetById(userid)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
// 		}

// 		if user.Name != "" {
// 			updateUser.Name = user.Name
// 		}
// 		if user.Email != "" {
// 			updateUser.Email = user.Email
// 		}
// 		if user.Password != "" {
// 			updateUser.Password = user.Password
// 		}
// 		if user.Role != "" {
// 			updateUser.Role = user.Role
// 		}
// 		src, file, err := c.Request().FormFile("avatar")
// 		if err != nil {
// 			fmt.Println(err)
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to upload avatar"))
// 		}
// 		ext := strings.Split(file.Filename, ".")
// 		extension := ext[len(ext)-1]
// 		check_extension := strings.ToLower(extension)
// 		if check_extension != "jpg" && check_extension != "png" && check_extension != "jpeg" {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file extention not allowed"))
// 		}
// 		if file.Size == 0 {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "illegal file size"))
// 		} else if file.Size > 1050000 {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "file size exceeded the limit"))
// 		}

// 		file.Filename = fmt.Sprintf("%d-%d.%s", userid, time.Now().Unix(), extension)

// 		sess := session.Must(util.GetAWSSession())

// 		uploader := s3manager.NewUploader(sess)

// 		_, err = uploader.Upload(&s3manager.UploadInput{
// 			Bucket: aws.String(os.Getenv("AWS_BUCKET")),
// 			Key:    aws.String(file.Filename),
// 			Body:   src,
// 		})

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, response.InternalServerError("failed", "Internal server error"))
// 		}
// 		updateUser.Avatar = fmt.Sprintf("https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/%s", file.Filename)

// 		err_update := uc.repository.Update(userid, updateUser)
// 		if err_update != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
// 		}
// 		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success Update user"))
// 	}
// }

// func (uc UserController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _, err := middlewares.ExtractToken(c)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}

// 		userId, errConv := strconv.Atoi(c.Param("id"))
// 		if errConv != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
// 		}

// 		if userId != id {
// 			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
// 		}

// 		errDelete := uc.repository.Delete(userId)
// 		if errDelete != nil {
// 			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
// 		}

// 		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
// 	}
// }

func (uc UserController) GetEmployees() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.repository.GetEmployees()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get employees", users))
	}
}
