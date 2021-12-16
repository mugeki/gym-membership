package routes

import (
	"gym-membership/controllers/admins"
	"gym-membership/controllers/articles"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	AdminController   admins.AdminController
	ArticleController articles.ArticleController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	admins := e.Group("admins")
	admins.POST("", ctrlList.UserController.Register)
	admins.POST("/login", ctrlList.UserController.Login)

	articles := e.Group("articles")
	articles.GET("", ctrlList.ArticleController.GetAll)
}
