package members

import (
	"gym-membership/business/members"
	controller "gym-membership/controllers"
	"gym-membership/controllers/members/response"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type MemberController struct {
	memberUsecase members.Usecase
}

func NewMemberController(Usecase members.Usecase) *MemberController {
	return &MemberController{
		memberUsecase: Usecase,
	}
}

func (ctrl *MemberController) GetByUserID(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	data, err := ctrl.memberUsecase.GetByUserID(uint(userId))
	res := response.Member{}
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}