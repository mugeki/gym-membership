package routes

import (
	"gym-membership/controllers/admins"
	"gym-membership/controllers/articles"
	"gym-membership/controllers/classifications"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware            middleware.JWTConfig
	UserController           users.UserController
	AdminController          admins.AdminController
	ArticleController        articles.ArticleController
	ClassificationController classifications.ClassificationController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	admin := e.Group("admin")
	admin.POST("", ctrlList.AdminController.Register)
	admin.POST("/login", ctrlList.AdminController.Login)

	article := e.Group("article")
	article.GET("", ctrlList.ArticleController.GetAll)
	article.POST("", ctrlList.ArticleController.Insert)
	// article.DELETE("/:idArticle", ctrlList.ArticleController.DeleteArticleByID)
	article.PUT("/:idArticle", ctrlList.ArticleController.UpdateArticleByID)

	classification := e.Group("classification")
	classification.POST("", ctrlList.ClassificationController.Insert)
	classification.GET("", ctrlList.ClassificationController.GetAll)
}
