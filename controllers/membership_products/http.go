package membership_products

import (
	"gym-membership/business/membership_products"
	controller "gym-membership/controllers"
	"gym-membership/controllers/membership_products/request"
	"gym-membership/controllers/membership_products/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type MembershipProductsController struct {
	membershipProductsUsecase membership_products.Usecase
}

func NewMembershipProductsController(Usecase membership_products.Usecase) *MembershipProductsController {
	return &MembershipProductsController{
		membershipProductsUsecase: Usecase,
	}
}

func (ctrl *MembershipProductsController) Insert(c echo.Context) error {
	req := request.MembershipProducts{}
	domain := membership_products.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.membershipProductsUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *MembershipProductsController) GetByID(c echo.Context) error {
	idMembershipProducts, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.membershipProductsUsecase.GetByID(uint(idMembershipProducts))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resp := response.MembershipProducts{}
	copier.Copy(&resp, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, resp)
}
