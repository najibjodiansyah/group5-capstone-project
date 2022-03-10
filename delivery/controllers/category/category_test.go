package category

import (
	"capstone-project/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	t.Run("success get categories", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/categories")

		categoryController := New(mockCategoryRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, categoryController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all categories", response.Message)
		}
	})
	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/categories")

		categoryController := New(mockErrorCategoryRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, categoryController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to fetch data", response.Message)
		}
	})
}

type mockCategoryRepository struct{}

func (m mockCategoryRepository) Get() ([]entities.Category, error) {
	return []entities.Category{}, nil
}

type mockErrorCategoryRepository struct{}

func (m mockErrorCategoryRepository) Get() ([]entities.Category, error) {
	return []entities.Category{}, fmt.Errorf("error")
}
