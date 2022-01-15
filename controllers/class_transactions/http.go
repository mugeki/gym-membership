package class_transactions

import (
	"gym-membership/business/class_transactions"
	controller "gym-membership/controllers"
	"gym-membership/controllers/class_transactions/request"
	"gym-membership/controllers/class_transactions/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ClassTransactionController struct {
	class_transactionsUsecase class_transactions.Usecase
}

func NewClassTransactionController(Usecase class_transactions.Usecase) *ClassTransactionController {
	return &ClassTransactionController{
		class_transactionsUsecase: Usecase,
	}
}

func (ctrl *ClassTransactionController) Insert(c echo.Context) error {
	req := request.ClassTransaction{}
	res := response.ClassTransaction{}
	domain := class_transactions.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.class_transactionsUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *ClassTransactionController) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	status := c.QueryParam("status")
	idUser, _ := strconv.Atoi(c.QueryParam("idUser"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.class_transactionsUsecase.GetAll(status, uint(idUser), page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	res := []response.ClassTransaction{}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *ClassTransactionController) GetActiveClass(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.Param("idUser"))
	data, err := ctrl.class_transactionsUsecase.GetActiveClass(uint(idUser))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *ClassTransactionController) UpdateStatus(c echo.Context) error {
	idClassTransaction, _ := strconv.Atoi(c.Param("idClassTransaction"))
	status := c.QueryParam("status")
	idAdmin, _ := strconv.Atoi(c.QueryParam("admin"))
	_, err := ctrl.class_transactionsUsecase.UpdateStatus(uint(idClassTransaction), uint(idAdmin), status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, nil)
}
