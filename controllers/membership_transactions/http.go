package membership_transactions

import (
	"gym-membership/business/membership_transactions"
	controller "gym-membership/controllers"
	"gym-membership/controllers/membership_transactions/request"
	"gym-membership/controllers/membership_transactions/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type MembershipTransactionController struct {
	membershipTransactionsUsecase membership_transactions.Usecase
}

func NewMembershipTransactionController(Usecase membership_transactions.Usecase) *MembershipTransactionController {
	return &MembershipTransactionController{
		membershipTransactionsUsecase: Usecase,
	}
}

func (ctrl *MembershipTransactionController) Insert(c echo.Context) error {
	req := request.MembershipTransaction{}
	res := response.MembershipTransaction{}
	domain := membership_transactions.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.membershipTransactionsUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *MembershipTransactionController) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	status := c.QueryParam("status")
	idUser, _ := strconv.Atoi(c.QueryParam("idUser"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.membershipTransactionsUsecase.GetAll(status, uint(idUser), page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	res := []response.MembershipTransaction{}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *MembershipTransactionController) UpdateStatus(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("idProduct"))
	status := c.QueryParam("status")
	// adminId := c.QueryParam("admin")
	err := ctrl.membershipTransactionsUsecase.UpdateStatus(uint(productId), status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, nil)
}
