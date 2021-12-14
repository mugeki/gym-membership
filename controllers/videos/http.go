package videos

import (
	"gym-membership/business/videos"
	controller "gym-membership/controllers"
	"gym-membership/controllers/videos/request"
	"gym-membership/controllers/videos/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type VideoController struct {
	videoUsecase videos.Usecase
}

func NewVideoController(Usecase videos.Usecase) *VideoController {
	return &VideoController{
		videoUsecase: Usecase,
	}
}

func (ctrl *VideoController) GetAll(c echo.Context) error {
	data, err := ctrl.videoUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.FromDomainArray(data)
	return controller.NewSuccessResponse(c, res)
}

func (ctrl *VideoController) Insert(c echo.Context) error {
	req := request.Videos{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	adminId := 1 //temporary adminID
	data, err := ctrl.videoUsecase.Insert(req.ToDomain(), uint(adminId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *VideoController) UpdateVideoByID(c echo.Context) error {
	req := request.Videos{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	videoId, _ := strconv.Atoi(c.Param("idVideo"))
	adminId := 1 //temporary adminID
	data, err := ctrl.videoUsecase.UpdateVideoByID(uint(videoId), req.ToDomain(), uint(adminId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}
