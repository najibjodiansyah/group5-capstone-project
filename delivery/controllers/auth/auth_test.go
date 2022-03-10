package auth

import (
	"bytes"
	"capstone-project/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("Test Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":1,
			"password":"najib",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
	t.Run("Test Error Repository Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":	"eldy@gmail.com",
			"password":	"najib",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Password Missmatch", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":"najib@jodiansyah.com",
			"password":"eldy",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})		
	t.Run("Test Error Email Unknows", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":"ratu@gmail.com",
			"password":"eldy",
		})
		req := httptest.NewRequest(http.MethodPost, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Success login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"email":"najib@jodiansyah.com",
			"password":"najib",
		})

		req := httptest.NewRequest(http.MethodPost, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockAuthRepository struct{}


func (ma mockAuthRepository) Login(email string) (entities.User, error) {
	var user entities.User
	user.ID = 1
	user.Name = "najib"
	user.Email = "najib@jodiansyah.com"
	user.Password = "$2a$10$1G28p8R0TrBdcjFN5TbziOB8ZYE6ICrNhT2WyNhiNIveMDUtiqIOC"
	user.Avatar = "najib.jpeg"
	user.Role = "employee"
	user.CreatedAt = time.Now()

	if email == "eldy@gmail.com"{
		return user, fmt.Errorf("internal server error")
	}

	if email == "ratu@gmail.com"{
		return entities.User{},nil
	}
	
	return user, nil
}