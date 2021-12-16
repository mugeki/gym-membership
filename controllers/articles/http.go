package articles

import (
	"gym-membership/business/articles"
	controller "gym-membership/controllers"
	"gym-membership/controllers/articles/request"

	// "gym-membership/controllers/articles/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUsecase articles.Usecase
}

func NewVideoController(Usecase articles.Usecase) *ArticleController {
	return &ArticleController{
		articleUsecase: Usecase,
	}
}

func (ctrl *ArticleController) GetAll(c echo.Context) error {
	data, err := ctrl.articleUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	// res := response.FromDomainArray(data)
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *ArticleController) Insert(c echo.Context) error {
	req := request.Articles{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	adminId := 1 //temporary adminID
	data, err := ctrl.articleUsecase.Insert(req.ToDomain(), uint(adminId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *ArticleController) UpdateArticleByID(c echo.Context) error {
	req := request.Articles{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	articleId, _ := strconv.Atoi(c.Param("idVideo"))
	adminId := 1 //temporary adminID
	data, err := ctrl.articleUsecase.UpdateArticleByID(uint(articleId), req.ToDomain(), uint(adminId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}
