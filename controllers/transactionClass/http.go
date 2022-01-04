package transactionClass

import (
	"gym-membership/business/transactionClass"
	controller "gym-membership/controllers"
	"gym-membership/controllers/transactionClass/request"
	"gym-membership/controllers/transactionClass/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type TransactionClassController struct {
	transactionClassUsecase transactionClass.Usecase
}

func NewTransactionClassController(Usecase transactionClass.Usecase) *TransactionClassController {
	return &TransactionClassController{
		transactionClassUsecase: Usecase,
	}
}

func (ctrl *TransactionClassController) Insert(c echo.Context) error {
	req := request.TransactionClass{}
	res := response.TransactionClass{}
	domain := transactionClass.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.transactionClassUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *TransactionClassController) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	status := c.QueryParam("status")
	idUser, _ := strconv.Atoi(c.QueryParam("idUser"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.transactionClassUsecase.GetAll(status, uint(idUser), page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	res := []response.TransactionClass{}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *TransactionClassController) GetActiveClass(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.Param("idUser"))
	// println(idUser, "id user")
	data, err := ctrl.transactionClassUsecase.GetActiveClass(uint(idUser))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *TransactionClassController) UpdateStatus(c echo.Context) error {
	classId, _ := strconv.Atoi(c.Param("idClass"))
	status := c.QueryParam("status")
	stringStatus, err := ctrl.transactionClassUsecase.UpdateStatus(uint(classId), status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, stringStatus)
}
