package routes

import (
	"gym-membership/controllers/admins"
	"gym-membership/controllers/articles"
	"gym-membership/controllers/class"
	"gym-membership/controllers/class_transactions"
	"gym-membership/controllers/classifications"
	"gym-membership/controllers/members"
	"gym-membership/controllers/membership_products"
	"gym-membership/controllers/membership_transactions"
	"gym-membership/controllers/trainers"
	"gym-membership/controllers/users"
	"gym-membership/controllers/videos"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware              middleware.JWTConfig
	UserController             users.UserController
	ClassController            class.ClassController
	ClassTransactionController class_transactions.ClassTransactionController
	TrainerController          trainers.TrainerController
	AdminController          admins.AdminController
	ArticleController        articles.ArticleController
	ClassificationController classifications.ClassificationController
	VideoController videos.VideoController
	MembershipProductsController membership_products.MembershipProductsController
	MembershipTransactionController	membership_transactions.MembershipTransactionController
	MemberController	members.MemberController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)
  	users.GET("/videos", ctrlList.VideoController.GetAll)

	membership_products := e.Group("membership-products")
	membership_products.POST("", ctrlList.MembershipProductsController.Insert)
	membership_products.GET("", ctrlList.MembershipProductsController.GetAll)
	membership_products.GET("/:id", ctrlList.MembershipProductsController.GetByID)
	membership_products.DELETE("/:id", ctrlList.MembershipProductsController.DeleteByID)
	membership_products.PUT("/:id", ctrlList.MembershipProductsController.UpdateByID)

	class := e.Group("class")
	class.POST("", ctrlList.ClassController.Insert)
	class.GET("", ctrlList.ClassController.GetAll)
	class.PUT("/:idClass", ctrlList.ClassController.UpdateClassByID)
	class.GET("/my-schedule/:idUser", ctrlList.ClassController.ScheduleByID)

	class_transactions := e.Group("transaction-class")
	class_transactions.GET("", ctrlList.ClassTransactionController.GetAll)
	class_transactions.POST("", ctrlList.ClassTransactionController.Insert)
	class_transactions.PUT("/update-status/:idClassTransaction", ctrlList.ClassTransactionController.UpdateStatus)
	class_transactions.GET("/active/:idUser", ctrlList.ClassTransactionController.GetActiveClass)

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

	transactionMembership := e.Group("transaction-membership")
	transactionMembership.GET("", ctrlList.MembershipTransactionController.GetAll)
	transactionMembership.POST("", ctrlList.MembershipTransactionController.Insert)
	transactionMembership.PUT("/update-status/:idClass", ctrlList.MembershipTransactionController.UpdateStatus)

	members := e.Group("members")
	members.GET("/:userId", ctrlList.MemberController.GetByUserID)
}
