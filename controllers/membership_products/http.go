package membership_products

import (
	"gym-membership/business"
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
	res := response.MembershipProducts{}
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
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *MembershipProductsController) GetAll(c echo.Context) error {
	data, err := ctrl.membershipProductsUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.MembershipProducts{}
	copier.Copy(&res, &data)
	if len(res) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, nil)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *MembershipProductsController) GetByID(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.membershipProductsUsecase.GetByID(uint(productId))
	if err != nil {
		if err == business.ErrProductNotFound {
			return controller.NewSuccessResponse(c, http.StatusNoContent, nil)
		}
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resp := response.MembershipProducts{}
	copier.Copy(&resp, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, resp)
}

func (ctrl *MembershipProductsController) UpdateByID(c echo.Context) error {
	req := request.MembershipProducts{}
	res := response.MembershipProducts{}
	domain := membership_products.Domain{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	productId, _ := strconv.Atoi(c.Param("id"))
	copier.Copy(&domain, &req)
	data, err := ctrl.membershipProductsUsecase.UpdateByID(uint(productId), &domain)
	if err != nil {
		if err == business.ErrProductNotFound {
			return controller.NewErrorResponse(c, http.StatusNotFound, err)
		} else {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *MembershipProductsController) DeleteByID(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.membershipProductsUsecase.DeleteByID(uint(productId))
	if err != nil {
		if err == business.ErrProductNotFound {
			return controller.NewErrorResponse(c, http.StatusNotFound, err)
		} else {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	}
	return controller.NewSuccessResponse(c, http.StatusOK, nil)
}
