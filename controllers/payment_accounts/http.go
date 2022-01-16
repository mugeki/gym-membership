package payment_accounts

import (
	paymentAccount "gym-membership/business/payment_accounts"
	controller "gym-membership/controllers"
	"gym-membership/controllers/payment_accounts/request"
	"gym-membership/controllers/payment_accounts/response"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type PaymentAccountController struct {
	paymentAccountUsecase paymentAccount.Usecase
}

func NewPaymentAccountController(Usecase paymentAccount.Usecase) *PaymentAccountController {
	return &PaymentAccountController{
		paymentAccountUsecase: Usecase,
	}
}

func (ctrl *PaymentAccountController) GetAll(c echo.Context) error {
	data, err := ctrl.paymentAccountUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.PaymentAccount{}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *PaymentAccountController) Insert(c echo.Context) error {
	req := request.PaymentAccount{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domain := paymentAccount.Domain{}
	copier.Copy(&domain, &req)
	data, err := ctrl.paymentAccountUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.PaymentAccount{}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
