package category

import (
	categoryRepo "capstone-project/repository/category"
	"net/http"

	response "capstone-project/delivery/commons"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	repository categoryRepo.Category
}

func New(category categoryRepo.Category) *CategoryController {
	return &CategoryController{repository: category}
}

func (cc CategoryController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get all categories from database
		categories, err := cc.repository.Get()

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all categories", categories))
	}
}
