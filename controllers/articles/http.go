package articles

import (
	"gym-membership/business/articles"
	controller "gym-membership/controllers"
	"gym-membership/controllers/articles/request"
	"gym-membership/controllers/articles/response"

	// "gym-membership/controllers/articles/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUsecase articles.Usecase
}

func NewArticleController(Usecase articles.Usecase) *ArticleController {
	return &ArticleController{
		articleUsecase: Usecase,
	}
}

func (ctrl *ArticleController) GetAll(c echo.Context) error {
	title := c.QueryParam("title")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.articleUsecase.GetAll(title, page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.Articles{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *ArticleController) Insert(c echo.Context) error {
	req := request.Articles{}
	res := response.Articles{}
	domain := articles.Domain{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	adminId := 1 //temporary adminID
	req.AdminID = uint(adminId)
	copier.Copy(&domain, &req)
	data, err := ctrl.articleUsecase.Insert(&domain)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *ArticleController) UpdateArticleByID(c echo.Context) error {
	req := request.Articles{}
	res := response.Articles{}
	domain := articles.Domain{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	articleId, _ := strconv.Atoi(c.Param("idArticle"))
	adminId := 1 //temporary adminID
	req.AdminID = uint(adminId)
	copier.Copy(&domain, &req)
	data, err := ctrl.articleUsecase.UpdateArticleByID(uint(articleId), &domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

// func (ctrl *ArticleController) DeleteArticleByID(c echo.Context) error {
// 	// println("cek param path", c.QueryParam("id"))
// 	// req := request.Articles{}
// 	// err := c.Bind(&req)
// 	// if err != nil {
// 	// 	return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	// }

// 	// _, err = govalidator.ValidateStruct(req)
// 	// if err != nil {
// 	// 	return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	// }

// 	articleId, _ := strconv.Atoi(c.Param("idArticle"))
// 	// println(articleId, "article id")
// 	adminId := 1 //temporary adminID
// 	req.AdminID = uint(adminId)
// 	data, err := ctrl.articleUsecase.UpdateArticleByID(uint(articleId), req.ToDomain())
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controller.NewSuccessResponse(c, http.StatusOK, data)
// }
