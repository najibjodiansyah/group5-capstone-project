package asset

import (
	"bytes"
	"capstone-project/entities"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestCreate(t *testing.T) {
	t.Run("Test Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":1,
			"description":"ram 8 gb dan ssd 128gb",
			"categoryid":1,
			"quantity":2,
			"picture":"image.jpeg",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Create())(context)) {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
	t.Run("Test Error Upload Picture", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		name, _ := w.CreateFormField("name")
		name.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		categoryid, _ := w.CreateFormField("categoryid")
		categoryid.Write([]byte("1"))

		quantity, _ := w.CreateFormField("quantity")
		quantity.Write([]byte("2"))

		picture, _ := w.CreateFormFile("picture", "picture.jpg")
		picture.Write([]byte("string apapun"))

		w.Close()
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		req.Header.Set(echo.HeaderContentType, w.FormDataContentType())

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Create())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, res.Code)
		}
	})
	t.Run("Test Error Picture Extention", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		name, _ := w.CreateFormField("name")
		name.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		categoryid, _ := w.CreateFormField("categoryid")
		categoryid.Write([]byte("1"))

		quantity, _ := w.CreateFormField("quantity")
		quantity.Write([]byte("2"))

		picture, _ := w.CreateFormFile("picture", "picture.svg")
		picture.Write([]byte("string apapun"))

		w.Close()
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		req.Header.Set(echo.HeaderContentType, w.FormDataContentType())

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/api/v1/assets")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Create())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Create Asset", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":"laptop",
			"description":"ram 8 gb dan ssd 128gb",
			"categoryid":1,
			"quantity":2,
			"picture":"",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Create())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	// error generate item
	// success create asset, success upload to S3
}

func TestGetById(t *testing.T) {
	t.Run("Test Error Convert Id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/assets")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.GetById())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Repository GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("2")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.GetById())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	// if total == 0
	// if total > 0
	t.Run("Test Success GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.GetById())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Test Error Get Count Asset Used", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockErrorAssetRepository{})
		if assert.NoError(t, (AssetController.GetById())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Error Convert Id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("/assets")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Update())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"description":false,
		})
		req := httptest.NewRequest(http.MethodPut, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Update())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	})
	t.Run("Test Error Repository GetById", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
		})
		req := httptest.NewRequest(http.MethodPut, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("2")

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.Update())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response entities.User
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Repository Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"description":"string",
		})
		req := httptest.NewRequest(http.MethodPut, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.Update()(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
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
	t.Run("Test Success with nul value array", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"description":"",
		})
		req := httptest.NewRequest(http.MethodPut, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.Update()(context)) {
			bodyResponses := res.Body.String()
			fmt.Println("success with null value",bodyResponses)
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update asset", response.Message)
		}
	})
	t.Run("Test Success Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"description":"release terbaru",
		})
		req := httptest.NewRequest(http.MethodPut, "/",  bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets/")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.Update()(context)) {
			bodyResponses := res.Body.String()
			var response Responses
			fmt.Println("successaja",bodyResponses)

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update asset", response.Message)
		}
	})
	
}
func TestGetAll(t *testing.T) {
	t.Run("Test Success Get All Asset", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?page=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.GetAll()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all asset", response.Message)
		}
	})
	t.Run("Test Error Failed to Fetch Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.GetAll()(context)) {
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
	t.Run("Test Error Failed to Convert Category Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?category=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.GetAll()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert category id", response.Message)
		}
	})
	t.Run("Test Error Failed to Convert Page Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?page=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.GetAll()(context)) {
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
	t.Run("Test Error Get Count Asset Used", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?page=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AssetController := New(mockErrorAssetRepository{})
		if assert.NoError(t, (AssetController.GetAll())(context)) {
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Test Error Failed to Convert Page Data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?page=a", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)
		context.SetPath("api/v1/assets")

		AssetController := New(mockAssetRepository{})
		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, AssetController.GetAll()(context)) {
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
	t.Run("Test Error Get Count Asset Used", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?page=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		
		context := e.NewContext(req, res)

		AssetController := New(mockAssetRepository{})
		if assert.NoError(t, (AssetController.GetAll())(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	//IF TOTAL == 0
	//IF STATUS TERSEDIA
	//IF STATUS TIDAK TERSEDIA
}
// =========================== mocking ===========================

type mockAssetRepository struct{}

func (ma mockAssetRepository) Create(asset entities.Asset) (entities.Asset, int, error) {
	if asset.CategoryId != 1 {
		return asset, 0, fmt.Errorf("failed to create asset")
	}
	return asset, 1, nil
}


func (ma mockAssetRepository) GenerateItem(assetName string, assetId int) error {
	return nil
}

func (ma mockAssetRepository) GetById(assetId int) (entities.Asset, error) {
	var asset entities.Asset
	if assetId != 1 {
		return asset, fmt.Errorf("failed get asset by id")
	}
	return asset, nil
}

func (ma mockAssetRepository) GetAll(page int, category int, keyword string)([]entities.Asset,int,error){
	var (
		asset []entities.Asset
		totalAsset int)
		asset = append(asset, entities.Asset{})
	if page != 1 {
		return asset,totalAsset,fmt.Errorf("failed get all asset")
	}
	return asset, totalAsset, nil
}

func (ma mockAssetRepository) GetCountAssetUsed(assetId int) (int, error) {
	var total int
	return total, nil
}

func (ma mockAssetRepository) Update(idasset int, asset entities.Asset)(entities.Asset, error){
	if asset.Description == "string"{
		return asset, fmt.Errorf("failed to update asset")
	}
	return asset, nil
}

type mockErrorAssetRepository struct{}

func (ma mockErrorAssetRepository) Create(asset entities.Asset) (entities.Asset, int, error) {
	if asset.CategoryId != 1 {
		return asset, 0, fmt.Errorf("failed to create asset")
	}
	return asset, 1, nil
}


func (ma mockErrorAssetRepository) GenerateItem(assetName string, assetId int) error {
	return fmt.Errorf("error generate item")
}

func (ma mockErrorAssetRepository) GetById(assetId int) (entities.Asset, error) {
	var asset entities.Asset
	if assetId != 1 {
		return asset, fmt.Errorf("failed get asset by id")
	}
	return asset, nil
}

func (ma mockErrorAssetRepository) GetAll(page int, category int, keyword string)([]entities.Asset,int,error){
	var (
		asset []entities.Asset
		totalAsset int)

		asset = append(asset, entities.Asset{})
	if page != 1 {
		return asset,totalAsset,fmt.Errorf("failed get all asset")
	}
	return asset, totalAsset, nil
}

func (ma mockErrorAssetRepository) GetCountAssetUsed(assetId int) (int, error) {
	var total int = 2
	return total, errors.New("failed get count asset used")
}

func (ma mockErrorAssetRepository) Update(idasset int, asset entities.Asset)(entities.Asset, error){
	if asset.Description == "string"{
		return asset, fmt.Errorf("failed to update asset")
	}
	return asset, nil
}

