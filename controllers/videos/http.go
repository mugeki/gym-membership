package videos

import (
	"gym-membership/business/videos"
	controller "gym-membership/controllers"
	"gym-membership/controllers/videos/request"
	"gym-membership/controllers/videos/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
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
	title := c.QueryParam("title")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.videoUsecase.GetAll(title, page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.VideosPagination{}
	res.Limit = limit
	res.Offset = offset
	res.TotalData = totalData
	copier.Copy(&res.Videos, &data)
	if len(res.Videos) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, res)
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

	domain := videos.Domain{}
	copier.Copy(&domain, &req)
	data, err := ctrl.videoUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *VideoController) UpdateByID(c echo.Context) error {
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
	domain := videos.Domain{}
	copier.Copy(&domain, &req)
	data, err := ctrl.videoUsecase.UpdateByID(uint(videoId), &domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, data)
}
