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

	valid := govalidator.IsNonNegative(float64(req.Price)) && govalidator.IsNonNegative(float64(req.PeriodTime))
	if !valid {
		return controller.NewErrorResponse(c, http.StatusBadRequest, business.ErrNegativeValue)
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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.membershipProductsUsecase.GetAll(page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.MembershipProducts{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
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

	valid := govalidator.IsNonNegative(float64(req.Price)) && govalidator.IsNonNegative(float64(req.PeriodTime))
	if !valid {
		return controller.NewErrorResponse(c, http.StatusBadRequest, business.ErrNegativeValue)
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
