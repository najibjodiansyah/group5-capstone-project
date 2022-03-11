package application

import (
	"bytes"
	"capstone-project/delivery/middlewares"
	"capstone-project/entities"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Test Error Extreact Token", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(0, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :"1", 
    		"Returndate"  :"2022-01-01", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.Create())(context)) {
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})
	t.Run("Test Error Bind Data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :"1", 
    		"Returndate"  :"2022-01-01", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.Create())(context)) {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
	t.Run("Test Success Create", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :1, 
    		"Returndate"  :"2022-01-01", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Success Create with nil returndate Data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :1, 
    		"Returndate"  :"", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Error Repository Create", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :1, 
    		"Returndate"  :"", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.Create())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	// ERROR REPOSITORY
}
func TestUpdateStatus(t *testing.T) {
	t.Run("Test Error Extreact Token", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(0, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :"1", 
    		"Returndate"  :"2022-01-01", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})
	t.Run("Test Error Convert Param", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Employeeid"  :1,  
    		"Assetid"     :"1", 
    		"Returndate"  :"2022-01-01", 
    		"Specification" :"Laptop ram 2 GB",
    		"Description":"untuk mengolah data menggunakan ms.excel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("a")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error bind status data", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]int{
			"status":1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
	t.Run("Test Error tomanager role not admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"tomanager",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error tomanager failed create repo", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"tomanager",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Success tomanager role not admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"tomanager",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Success toaccept role not admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "manager")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toaccept",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Error toaccept role not manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toaccept",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error toaccept failed create repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "manager")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toaccept",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error decline role is not manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"decline",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Success decline with role is manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "manager")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"decline",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test failed decline with role is manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "manager")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"decline",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Success decline with role is admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"decline",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test decline failed with role is admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"decline",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed with role is admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed get asset", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed get availability item", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("2")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse success", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test inuse failed avail item is 0", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("3")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed updatestatus", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("4")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed get by id", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test inuse failed update item", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"inuse",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("6")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test toreturn success role is not manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test toreturn failed role is manager", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "manager")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test toreturn failed update repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"toreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("4")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test donereturn failed role is not admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"donereturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("4")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test donereturn failed update status repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"donereturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("4")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test donereturn failed get by id repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"donereturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	// failed update item
	// success done return
	t.Run("Test ask return failed role is not admin", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"askreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test ask return failed update repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"askreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("4")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test ask return success update repository", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "admin")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"askreturn",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("6")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test status not recognized", func(t *testing.T) {
		token, _ := middlewares.CreateToken(1, "employee")
		requestBody, _ := json.Marshal(map[string]interface{}{
			"status":"ask return",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		e := echo.New()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,middlewares.JWTMiddleware()(ApplicationController.UpdateStatus())(context)) {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Error Convert Id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetById())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetById())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test success repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetById())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}

func TestUserApplicationActivity(t *testing.T) {
	t.Run("Test Error Convert Id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationActivity())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationActivity())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test success repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationActivity())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}

func TestUserApplicationHistory(t *testing.T) {
	t.Run("Test Error Convert Id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/applications")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationHistory())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("5")

		ApplicationController := New(errorMockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationHistory())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test success repository get by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/applications/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.UsersApplicationHistory())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test Success Get All Asset", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetAll())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Error Failed to Fetch Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?category=2", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetAll())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Failed to Convert Category Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?category=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		ApplicationController := New(mockApplicationRepository{})
		if assert.NoError(t,(ApplicationController.GetAll())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockApplicationRepository struct{}

func (mar mockApplicationRepository)Create(app entities.Applications)(int, entities.Applications,error){
	return 1, app, nil
}

func (mar mockApplicationRepository)UpdateStatus(applicationid int, status string, managerid *int, itemid *int)(error){
	if applicationid == 4 {
		return errors.New("failed update status")
	}
	return nil
}
func (mar mockApplicationRepository)AvailabilityItem(assetid int) (int, error){
	var id int
	if assetid == 2 {
		return id, errors.New("failed check availability item")
	}
	if assetid == 3{
		return 0, nil
	}
	if assetid == 6 {
		return 6, nil
	}
	return 1, nil
}

func (mar mockApplicationRepository)UpdateItem(itemid *int, availStatus string, employeeid int) error{
	if *itemid == 6 {
		return errors.New("failed update item")
	}
	return nil
}
func (mar mockApplicationRepository)GetAll(status string,
	category int,
	date string,
	orderbydate string,
	longestdate string)([]entities.ResponseApplicationWithDuration,int, error){
		var app []entities.ResponseApplicationWithDuration
		var id int
		if category == 2 {
			return app, id, errors.New("failed get all id")
		}
	return app,id, nil
}
func (mar mockApplicationRepository)GetById(id int)(entities.ResponseApplication,error){
	var app entities.ResponseApplication
	if id == 5 {
		return app, errors.New("failed get by id")
	}
	return app, nil
}
func (mar mockApplicationRepository)GetAsset(applicationid int)(int,error){
	var id int
	if applicationid == 2 {
		return 2, nil
	}
	if applicationid == 3 {
		return 3, nil
	}
	if applicationid == 6 {
		return 6, nil
	}
	return id, nil
}
func (mar mockApplicationRepository)UsersApplicationHistory(userid int)([]entities.ResponseApplication,error){
	var app []entities.ResponseApplication
	return app, nil
}
func (mar mockApplicationRepository)UsersApplicationActivity(userid int)([]entities.ResponseApplication,error){
	var app []entities.ResponseApplication
	return app, nil
}

//---------------------------------------------------------------

type errorMockApplicationRepository struct{}

func (mar errorMockApplicationRepository)Create(app entities.Applications)(int, entities.Applications,error){
	return 1, app, errors.New("failed")
}

func (mar errorMockApplicationRepository)UpdateStatus(applicationid int, status string, managerid *int, itemid *int)(error){
	
	return errors.New("failed")
}
func (mar errorMockApplicationRepository)AvailabilityItem(assetid int) (int, error){
	var id int
	return id, errors.New("failed")
}

func (mar errorMockApplicationRepository)UpdateItem(itemid *int, availStatus string, employeeid int) error{
	return errors.New("failed")
}
func (mar errorMockApplicationRepository)GetAll(status string,
	category int,
	date string,
	orderbydate string,
	longestdate string)([]entities.ResponseApplicationWithDuration,int, error){
		var app []entities.ResponseApplicationWithDuration
		var id int
	return app,id, errors.New("failed")
}
func (mar errorMockApplicationRepository)GetById(id int)(entities.ResponseApplication,error){
	var app entities.ResponseApplication
	return app, errors.New("failed")
}
func (mar errorMockApplicationRepository)GetAsset(applicationid int)(int,error){
	id := 1
	return id, errors.New("failed")
}
func (mar errorMockApplicationRepository)UsersApplicationHistory(userid int)([]entities.ResponseApplication,error){
	var app []entities.ResponseApplication
	return app, errors.New("failed")
}
func (mar errorMockApplicationRepository)UsersApplicationActivity(userid int)([]entities.ResponseApplication,error){
	var app []entities.ResponseApplication
	return app, errors.New("failed")
}

	
	
	