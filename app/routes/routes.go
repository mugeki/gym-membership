package routes

import (
	"gym-membership/controllers/admins"
	"gym-membership/controllers/articles"
	"gym-membership/controllers/class"
	"gym-membership/controllers/classifications"
	"gym-membership/controllers/trainers"
	"gym-membership/controllers/transactionClass"
	"gym-membership/controllers/users"
	"gym-membership/controllers/videos"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware              middleware.JWTConfig
	UserController             users.UserController
	ClassController            class.ClassController
	TrainerController          trainers.TrainerController
	TransactionClassController transactionClass.TransactionClassController
	AdminController            admins.AdminController
	ArticleController          articles.ArticleController
	ClassificationController   classifications.ClassificationController
	VideoController            videos.VideoController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)
	users.GET("/videos", ctrlList.VideoController.GetAll)

	class := e.Group("class")
	class.POST("", ctrlList.ClassController.Insert)
	class.GET("", ctrlList.ClassController.GetAll)
	class.PUT("/:idClass", ctrlList.ClassController.UpdateClassByID)
	class.GET("/myShcedule/:idUser", ctrlList.ClassController.ScheduleByID)

	transactionClass := e.Group("transaction-class")
	transactionClass.GET("", ctrlList.TransactionClassController.GetAll)
	transactionClass.POST("", ctrlList.TransactionClassController.Insert)
	transactionClass.PUT("/update-status/:idTransactionClass", ctrlList.TransactionClassController.UpdateStatus)
	transactionClass.GET("/active/:idUser", ctrlList.TransactionClassController.GetActiveClass)

	trainers := e.Group("trainers")
	trainers.GET("", ctrlList.TrainerController.GetAll)

	article := e.Group("article")
	article.GET("", ctrlList.ArticleController.GetAll)
	article.POST("", ctrlList.ArticleController.Insert)
	article.DELETE("/:idArticle", ctrlList.ArticleController.DeleteByID)
	article.PUT("/:idArticle", ctrlList.ArticleController.UpdateArticleByID)

	classification := e.Group("classification")
	classification.POST("", ctrlList.ClassificationController.Insert)
	classification.GET("", ctrlList.ClassificationController.GetAll)

	admins := e.Group("admins")
	admins.POST("", ctrlList.AdminController.Register)
	admins.POST("/login", ctrlList.AdminController.Login)
	admins.POST("/videos", ctrlList.VideoController.Insert)
	admins.PUT("/videos/:idVideo", ctrlList.VideoController.UpdateByID)
	admins.DELETE("/videos/:idVideo", ctrlList.VideoController.DeleteByID)
}
