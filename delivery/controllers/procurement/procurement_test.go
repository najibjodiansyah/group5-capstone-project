package procurement

import (
	"bytes"
	"capstone-project/delivery/middlewares"
	"capstone-project/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetProcurements(t *testing.T) {
	t.Run("success get procurements", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, procurementController.Get()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all procurements", response.Message)
		}
	})
	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, procurementController.Get()(context)) {
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

func TestGetById(t *testing.T) {
	t.Run("success get procurement by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, procurementController.GetById()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get procurement", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("a")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, procurementController.GetById()(context)) {
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
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, procurementController.GetById()(context)) {
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

func TestCreate(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(1, "employee")
	)
	t.Run("success create product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"assetName":     "Lenovo ideapad gaming gg banget",
			"spesification": "ram 8gb",
			"description":   "untuk kerja",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create procurement", response.Message)
		}

	})
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"assetName":     1,
			"spesification": "ram 8gb",
			"description":   "untuk kerja",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "Failed to Bind Input", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"assetName":     "Lenovo ideapad gaming gg banget",
			"spesification": "ram 8gb",
			"description":   "untuk kerja",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "Failed to create procurement", response.Message)
		}

	})
	t.Run("invalid id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"assetName":     "Lenovo ideapad gaming gg banget",
			"spesification": "ram 8gb",
			"description":   "untuk kerja",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "unauthorized access", response.Message)
		}

	})
	t.Run("not employee", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"assetName":     "Lenovo ideapad gaming gg banget",
			"spesification": "ram 8gb",
			"description":   "untuk kerja",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "unauthorized access", response.Message)
		}

	})

}

func TestUpdate(t *testing.T) {
	// var (
	// 	globalToken, errCreateToken = middlewares.CreateToken(1, "employee")
	// )
	t.Run("invalid id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "tomanager",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "unauthorized access", response.Message)
		}

	})
	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "tomanager",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("a")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
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
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "Failed to Bind Input", response.Message)
		}

	})
	t.Run("tomanager not admin", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "employee")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "tomanager",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "Unauthorized Role", response.Message)
		}

	})
	t.Run("tomanager success", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "tomanager",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update status to 'tomanager' by admin", response.Message)
		}

	})
	t.Run("tomanager failed", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "tomanager",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
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
	t.Run("toaccept not manager", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "toaccept",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "Unauthorized Role", response.Message)
		}

	})
	t.Run("toaccept success", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "manager")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "toaccept",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update status 'accept' by manager", response.Message)
		}

	})
	t.Run("toaccept failed", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "manager")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "toaccept",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
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
	t.Run("decline unauthorize", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "employee")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "decline",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "Unauthorized Role", response.Message)
		}

	})
	t.Run("decline success manager", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "manager")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "decline",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update status 'decline' by manager", response.Message)
		}

	})
	t.Run("decline failed manager", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "manager")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "decline",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
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
	t.Run("decline success admin", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "decline",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update status 'decline' by admin", response.Message)
		}

	})
	t.Run("decline failed admin", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "decline",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
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
	t.Run("invalid status", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "menolak",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/procurements")
		context.SetParamNames("id")
		context.SetParamValues("1")

		procurementController := New(mockErrorProcurementRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(procurementController.Update())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
			assert.Equal(t, "Bad Request", response.Status)
			assert.Equal(t, "status not recognized", response.Message)
		}

	})

}

type mockProcurementRepository struct{}

func (m mockProcurementRepository) Get(status string) ([]entities.ProcurementResponseFormat, error) {
	return []entities.ProcurementResponseFormat{}, nil
}
func (m mockProcurementRepository) GetById(id int) (entities.ProcurementResponseFormat, error) {
	return entities.ProcurementResponseFormat{}, nil
}
func (m mockProcurementRepository) Create(procurement entities.Procurement) (entities.Procurement, error) {
	return entities.Procurement{}, nil
}
func (m mockProcurementRepository) Update(id int, procurement entities.Procurement) error {
	return nil
}

type mockErrorProcurementRepository struct{}

func (m mockErrorProcurementRepository) Get(status string) ([]entities.ProcurementResponseFormat, error) {
	return []entities.ProcurementResponseFormat{}, fmt.Errorf("error")
}
func (m mockErrorProcurementRepository) GetById(id int) (entities.ProcurementResponseFormat, error) {
	return entities.ProcurementResponseFormat{}, fmt.Errorf("error")
}
func (m mockErrorProcurementRepository) Create(procurement entities.Procurement) (entities.Procurement, error) {
	return entities.Procurement{}, fmt.Errorf("error")
}
func (m mockErrorProcurementRepository) Update(id int, procurement entities.Procurement) error {
	return fmt.Errorf("error")
}
