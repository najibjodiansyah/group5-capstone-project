package user

import (
	"bytes"
	"capstone-project/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("success create product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "admin",
			"email":    "admin@admin.com",
			"password": "admin",
			"role":     "admin",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/users")

		userController := New(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.Register()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create user", response.Message)
		}

	})
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
			"email":    "admin@admin.com",
			"password": "admin",
			"role":     "admin",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/users")

		userController := New(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.Register()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to bind data", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "admin",
			"email":    "admin@admin.com",
			"password": "admin",
			"role":     "admin",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/users")

		userController := New(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.Register()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "error", response.Message)
		}

	})
}
func TestGetEmployees(t *testing.T) {
	t.Run("success get employees", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/employess")

		userController := New(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.GetEmployees()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get employees", response.Message)
		}
	})
	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/employess")

		userController := New(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.GetEmployees()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "error", response.Message)
		}
	})

}

type mockUserRepository struct{}

func (m mockUserRepository) Update(id int, user entities.User) error {
	return nil
}
func (m mockUserRepository) GetById(id int) (entities.User, error) {
	return entities.User{}, nil
}
func (m mockUserRepository) Register(entities.User) (entities.User, error) {
	return entities.User{}, nil
}
func (m mockUserRepository) Delete(id int) error {
	return nil
}
func (m mockUserRepository) GetEmployees() ([]entities.Employee, error) {
	return []entities.Employee{}, nil
}

type mockErrorUserRepository struct{}

func (m mockErrorUserRepository) Update(id int, user entities.User) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetById(id int) (entities.User, error) {
	return entities.User{}, fmt.Errorf("error")
}
func (m mockErrorUserRepository) Register(entities.User) (entities.User, error) {
	return entities.User{}, fmt.Errorf("error")
}
func (m mockErrorUserRepository) Delete(id int) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetEmployees() ([]entities.Employee, error) {
	return []entities.Employee{}, fmt.Errorf("error")
}
