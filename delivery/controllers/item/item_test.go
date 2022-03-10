package item

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

func TestGetItems(t *testing.T) {
	t.Run("success get items", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/?page=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, itemController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all items", response.Message)
		}
	})
	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/?page=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")

		itemController := New(mockErrorItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, itemController.Get()(context)) {
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
	t.Run("failed to convert page", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/?page=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, itemController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert page", response.Message)
		}
	})
	t.Run("failed to convert category id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/?category=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, itemController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert category_id", response.Message)
		}
	})

}

func TestGetById(t *testing.T) {
	t.Run("success get item by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetById()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get item", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")
		context.SetParamNames("id")
		context.SetParamValues("a")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetById()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockErrorItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetById()(context)) {
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

func TestGetItemUsageHistory(t *testing.T) {
	t.Run("success get item usage history", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items/usage")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetItemUsageHistory()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get item", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items/usage")
		context.SetParamNames("id")
		context.SetParamValues("a")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetItemUsageHistory()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/items/usage")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockErrorItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.GetItemUsageHistory()(context)) {
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

func TestUpdate(t *testing.T) {
	t.Run("success update item", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"availableStatus": "tersedia",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.Update()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update item", response.Message)
		}

	})
	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"availableStatus": "tersedia",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/items")
		context.SetParamNames("id")
		context.SetParamValues("a")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.Update()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"availableStatus": 1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.Update()(context)) {
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
			"availableStatus": "tersedia",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockOtherItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.Update()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to update data", response.Message)
		}

	})
	t.Run("failed to fetch item", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"availableStatus": "tersedia",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/items")
		context.SetParamNames("id")
		context.SetParamValues("1")

		itemController := New(mockErrorItemRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, itemController.Update()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to fetch data by id", response.Message)
		}

	})

}

type mockItemRepository struct{}

func (m mockItemRepository) Get(availableStatus string, category int, keyword string, page int) ([]entities.ItemResponseFormat, int, error) {
	return []entities.ItemResponseFormat{}, 1, nil
}

func (m mockItemRepository) GetItemUsageHistory(id int) (entities.ItemUsageHistory, error) {
	return entities.ItemUsageHistory{}, nil
}
func (m mockItemRepository) GetById(id int) (entities.ItemResponseFormat, error) {
	return entities.ItemResponseFormat{}, nil
}
func (m mockItemRepository) GetByIdUpdate(id int) (entities.Item, error) {
	return entities.Item{}, nil
}
func (m mockItemRepository) Update(id int, item entities.Item) error {
	return nil
}

type mockErrorItemRepository struct{}

func (m mockErrorItemRepository) Get(availableStatus string, category int, keyword string, page int) ([]entities.ItemResponseFormat, int, error) {
	return []entities.ItemResponseFormat{}, 1, fmt.Errorf("error")
}

func (m mockErrorItemRepository) GetItemUsageHistory(id int) (entities.ItemUsageHistory, error) {
	return entities.ItemUsageHistory{}, fmt.Errorf("error")
}
func (m mockErrorItemRepository) GetById(id int) (entities.ItemResponseFormat, error) {
	return entities.ItemResponseFormat{}, fmt.Errorf("error")
}
func (m mockErrorItemRepository) GetByIdUpdate(id int) (entities.Item, error) {
	return entities.Item{}, fmt.Errorf("error")
}
func (m mockErrorItemRepository) Update(id int, item entities.Item) error {
	return fmt.Errorf("error")
}

type mockOtherItemRepository struct{}

func (m mockOtherItemRepository) Get(availableStatus string, category int, keyword string, page int) ([]entities.ItemResponseFormat, int, error) {
	return []entities.ItemResponseFormat{}, 1, fmt.Errorf("error")
}

func (m mockOtherItemRepository) GetItemUsageHistory(id int) (entities.ItemUsageHistory, error) {
	return entities.ItemUsageHistory{}, fmt.Errorf("error")
}
func (m mockOtherItemRepository) GetById(id int) (entities.ItemResponseFormat, error) {
	return entities.ItemResponseFormat{}, fmt.Errorf("error")
}
func (m mockOtherItemRepository) GetByIdUpdate(id int) (entities.Item, error) {
	return entities.Item{}, nil
}
func (m mockOtherItemRepository) Update(id int, item entities.Item) error {
	return fmt.Errorf("error")
}
