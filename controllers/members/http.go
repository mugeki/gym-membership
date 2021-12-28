package members

import (
	"gym-membership/business/members"
	controller "gym-membership/controllers"
	"gym-membership/controllers/members/request"
	"net/http"
	"strconv"
	
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type MembersController struct {
	membersUsecase members.Usecase
}

func NewMembersController(Usecase members.Usecase) *MembersController {
	return &MembersController{
		membersUsecase: Usecase,
	}
}

func (ctrl *MembersController) Insert(c echo.Context) error {
	req := request.Members{}
	domain := members.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.membersUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *ClassController) GetByUserID(c echo.Context) error {
	membersId, _ := strconv.Atoi(c.Param("idMembers"))
	stringStatus, err := ctrl.membersUsecase.GetByUserID(int(membersId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, stringStatus)
}
