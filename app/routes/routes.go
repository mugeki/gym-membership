package routes

import (
	"gym-membership/controllers/members"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
	memberController members.UserController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	members := e.Group("members")
	members.POST("/members(create)", ctrlList.MembersController.Create)
	members.GET("/members/:user_id ", ctrlList.MembersController.GetByUserID)
}