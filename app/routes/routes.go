package routes

import (
	"gym-membership/controllers/class"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	ClassController class.ClassController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	class := e.Group("class")
	class.POST("", ctrlList.ClassController.Insert)
	class.PUT("/updateKuota/:idClass", ctrlList.ClassController.UpdateKuota)
}
