package statistic

import (
	statisticRepo "capstone-project/repository/statistic"
	"net/http"

	response "capstone-project/delivery/commons"

	"github.com/labstack/echo/v4"
)

type StatisticController struct {
	repository statisticRepo.Statistic
}

func New(statistic statisticRepo.Statistic) *StatisticController {
	return &StatisticController{repository: statistic}
}

func (sc StatisticController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get all categories from database
		statistic, err := sc.repository.Get()

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all statistic", statistic))
	}
}
