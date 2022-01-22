package routes

import (
	"gym-membership/controllers/auth"
	"gym-membership/controllers/calendars"
	"gym-membership/controllers/users"
	"gym-membership/controllers/videos"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware          middleware.JWTConfig
	UserController         users.UserController
	VideoController        videos.VideoController
	CalendarsApiController calendars.CalendarsController
	AuthController		   auth.AuthController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)
	users.GET("/videos", ctrlList.VideoController.GetAll)

	admins := e.Group("admins")
	admins.POST("/videos", ctrlList.VideoController.Insert)
	admins.PUT("/videos/:idVideo", ctrlList.VideoController.UpdateByID)
	admins.DELETE("/videos/:idVideo", ctrlList.VideoController.DeleteByID)

	calendars := e.Group("calendars")
	calendars.GET("", ctrlList.CalendarsApiController.GetAll)

	oauth := e.Group("")
    oauth.POST("/GoogleLogin",ctrlList.AuthController.HandleGoogleLogin)
    oauth.POST("/GoogleCallback",ctrlList.AuthController.HandleGoogleCallback)
}
