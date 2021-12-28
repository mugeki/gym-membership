package routes

import (
	"gym-membership/controllers/class"
	"gym-membership/controllers/trainers"
	"gym-membership/controllers/transactionClass"
	"gym-membership/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware              middleware.JWTConfig
	UserController             users.UserController
	ClassController            class.ClassController
	TrainerController          trainers.TrainerController
	TransactionClassController transactionClass.TransactionClassController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)

	class := e.Group("class")
	class.POST("", ctrlList.ClassController.Insert)
	class.GET("", ctrlList.ClassController.GetAll)
	class.PUT("/:idClass", ctrlList.ClassController.UpdateClassByID)

	class.GET("/active/:idUser", ctrlList.TransactionClassController.GetActiveClass)

	transactionClass := e.Group("transactionClass")
	transactionClass.GET("", ctrlList.TransactionClassController.GetAll)
	transactionClass.POST("", ctrlList.TransactionClassController.Insert)
	transactionClass.PUT("/updateStatus/:idClass", ctrlList.TransactionClassController.GetAll)
	// transactionClass.PUT("/updateKuota/:idClass", ctrlList.TransactionClassController.Insert)

	trainers := e.Group("trainers")
	trainers.GET("", ctrlList.TrainerController.GetAll)
}
