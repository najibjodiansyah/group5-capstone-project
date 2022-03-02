package procurement

import (
	"capstone-project/delivery/middlewares"
	procurementRepo "capstone-project/repository/procurement"
	"log"
	"net/http"
	"strconv"

	response "capstone-project/delivery/commons"
	"capstone-project/entities"

	"github.com/labstack/echo/v4"
)

type ProcurementController struct {
	repository procurementRepo.Procurement
}

func New(procurement procurementRepo.Procurement) *ProcurementController {
	return &ProcurementController{repository: procurement}
}

func (pc ProcurementController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		if role != "employee" {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		var input ProcurementRequestFormat

		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		procurement := entities.Procurement{
			EmployeeId:    id,
			AssetName:     input.AssetName,
			Spesification: input.Spesification,
			Description:   input.Description,
			Status:        "toadmin",
		}

		proc, errRepo := pc.repository.Create(procurement)
		if errRepo != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "Failed to create procurement"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success create procurement", proc))
	}
}

func (pc ProcurementController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get all products from database
		status := c.QueryParam("status")
		// pageInput := c.QueryParam("page")

		// if pageInput == "" {
		// 	pageInput = "0"
		// }
		// page, err := strconv.Atoi(pageInput)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert page"))
		// }
		procurements, err := pc.repository.Get(status)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		// limit := 10
		// totalPage := int(math.Ceil(float64(totalProcurement) / float64(limit)))
		// responseData := entities.ProcurementResponseTotal{
		// 	TotalPage:    totalPage,
		// 	Procurements: procurements,
		// }

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all procurements", procurements))
	}
}

func (pc ProcurementController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		procurementId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		procurement, err := pc.repository.GetById(procurementId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get procurement", procurement))
	}
}

func (pc ProcurementController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		var input Inputstatus

		procId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("Bad Request", "Failed to Bind Input"))
		}

		if input.Status == "tomanager" {
			if role != "admin" {
				return c.JSON(http.StatusBadRequest, response.BadRequest("Bad Request", "Unauthorized Role"))
			}
			procurement := entities.Procurement{
				Status: input.Status,
			}

			if err := pc.repository.Update(procId, procurement); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status to 'tomanager' by admin"))

		} else if input.Status == "toaccept" {
			// jika if status = diterima maka 1 update kolom status, 2 update kolom manager id (dari token id)

			if role != "manager" {
				return c.JSON(http.StatusBadRequest, response.BadRequest("Bad Request", "Unauthorized Role"))
			}

			procurement := entities.Procurement{
				Status:    input.Status,
				ManagerId: userid,
			}

			if err := pc.repository.Update(procId, procurement); err != nil {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
			}

			return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'accept' by manager"))
		} else if input.Status == "decline" {
			if role == "manager" {
				procurement := entities.Procurement{
					Status:    input.Status,
					ManagerId: userid,
				}

				if err := pc.repository.Update(procId, procurement); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by manager"))
			} else if role == "admin" {
				procurement := entities.Procurement{
					Status: input.Status,
				}

				if err := pc.repository.Update(procId, procurement); err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", err.Error()))
				}

				return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update status 'decline' by admin"))
			}

			return c.JSON(http.StatusBadRequest, response.BadRequest("Bad Request", "Unauthorized Role"))

		}

		return c.JSON(http.StatusUnprocessableEntity, response.BadRequest("Bad Request", "status not recognized"))
	}

}
