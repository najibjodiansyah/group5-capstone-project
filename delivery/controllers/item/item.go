package item

import (
	itemRepo "capstone-project/repository/item"
	"math"
	"net/http"
	"strconv"

	response "capstone-project/delivery/commons"
	"capstone-project/entities"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	repository itemRepo.Item
}

func New(item itemRepo.Item) *ItemController {
	return &ItemController{repository: item}
}

func (ic ItemController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get all products from database
		availableStatus := c.QueryParam("availableStatus")
		category := c.QueryParam("category")
		keyword := c.QueryParam("keyword")
		pageInput := c.QueryParam("page")

		if category == "" {
			category = "0"
		}
		if pageInput == "" {
			pageInput = "0"
		}
		categoryId, err := strconv.Atoi(category)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert category_id"))
		}
		page, err := strconv.Atoi(pageInput)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert page"))
		}
		items, totalItem, err := ic.repository.Get(availableStatus, categoryId, keyword, page)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		limit := 10
		totalPage := int(math.Ceil(float64(totalItem) / float64(limit)))
		responseData := entities.ItemResponseTotal{
			TotalPage: totalPage,
			Items:     items,
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all items", responseData))
	}
}

func (ic ItemController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		item, err := ic.repository.GetById(itemId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get item", item))
	}
}
func (ic ItemController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		item := entities.Item{}
		if errBind := c.Bind(&item); errBind != nil {
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("failed", "failed to bind data"))
		}
		// getting the id
		itemId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		updateItem, err := ic.repository.GetByIdUpdate(itemId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data by id"))
		}

		if item.AvailableStatus != "" {
			updateItem.AvailableStatus = item.AvailableStatus
		}

		errUpdate := ic.repository.Update(itemId, updateItem)
		if errUpdate != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to update data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update item"))
	}
}

func (ic ItemController) GetItemUsageHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		item, err := ic.repository.GetItemUsageHistory(itemId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get item", item))
	}
}
