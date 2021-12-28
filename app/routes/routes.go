package routes

import (
	"gym-membership/controllers/membership_products"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
	MembershipProductsController membership_products.MembershipProductsController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	membership_products := e.Group("membership_products")
	membership_products.POST("", ctrlList.MembershipProductsController.Insert)
	membership_products.GET("/:idMembershipProducts ", ctrlList.MembershipProductsController.GetByUserID)
}