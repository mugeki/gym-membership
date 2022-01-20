package routes

import (
	_middleware "gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/controllers"
	"gym-membership/controllers/admins"
	"gym-membership/controllers/articles"
	"gym-membership/controllers/class"
	"gym-membership/controllers/class_transactions"
	"gym-membership/controllers/classifications"
	"gym-membership/controllers/members"
	"gym-membership/controllers/membership_products"
	"gym-membership/controllers/membership_transactions"
	"gym-membership/controllers/payment_accounts"
	"gym-membership/controllers/trainers"
	"gym-membership/controllers/users"
	"gym-membership/controllers/videos"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware                   middleware.JWTConfig
	UserController                  users.UserController
	ClassController                 class.ClassController
	ClassTransactionController      class_transactions.ClassTransactionController
	TrainerController               trainers.TrainerController
	AdminController                 admins.AdminController
	ArticleController               articles.ArticleController
	ClassificationController        classifications.ClassificationController
	VideoController                 videos.VideoController
	MembershipProductsController    membership_products.MembershipProductsController
	MembershipTransactionController membership_transactions.MembershipTransactionController
	MemberController                members.MemberController
	PaymentAccountController        payment_accounts.PaymentAccountController
}

func (ctrlList *ControllerList) RegisterRoute(e *echo.Echo) {
	auth := e.Group("auth")
	auth.POST("/verify-jwt/:token", ctrlList.UserController.VerifyJWT)

	users := e.Group("users")
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)
	users.PUT("", ctrlList.UserController.Update, middleware.JWTWithConfig(ctrlList.JWTMiddleware))

	videos := e.Group("videos", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	videos.GET("", ctrlList.VideoController.GetAll)
	videos.GET("/:idVideo", ctrlList.VideoController.GetByID)
	videos.POST("", ctrlList.VideoController.Insert, SuperAdminValidation())
	videos.PUT("/:idVideo", ctrlList.VideoController.UpdateByID, SuperAdminValidation())
	videos.DELETE("/:idVideo", ctrlList.VideoController.DeleteByID, SuperAdminValidation())

	membership_products := e.Group("membership-products", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	membership_products.GET("", ctrlList.MembershipProductsController.GetAll)
	membership_products.GET("/:id", ctrlList.MembershipProductsController.GetByID)
	membership_products.POST("", ctrlList.MembershipProductsController.Insert, SuperAdminValidation())
	membership_products.DELETE("/:id", ctrlList.MembershipProductsController.DeleteByID, SuperAdminValidation())
	membership_products.PUT("/:id", ctrlList.MembershipProductsController.UpdateByID, SuperAdminValidation())

	class := e.Group("classes", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	class.GET("", ctrlList.ClassController.GetAll)
	class.GET("/:idClass", ctrlList.ClassController.GetByID)
	class.POST("", ctrlList.ClassController.Insert, SuperAdminValidation())
	class.PUT("/:idClass", ctrlList.ClassController.UpdateClassByID, SuperAdminValidation())
	// class.GET("/my-schedule/:idUser", ctrlList.ClassController.ScheduleByID)

	class_transactions := e.Group("transaction-class", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	class_transactions.GET("", ctrlList.ClassTransactionController.GetAll, AdminValidation(), SuperAdminValidation())
	class_transactions.GET("/user", ctrlList.ClassTransactionController.GetAllByUser)
	class_transactions.GET("/active/:idUser", ctrlList.ClassTransactionController.GetActiveClass)
	class_transactions.GET("/:idClassTransaction", ctrlList.ClassTransactionController.GetByID)
	class_transactions.POST("", ctrlList.ClassTransactionController.Insert)
	class_transactions.PUT("/update-status/:idClassTransaction", ctrlList.ClassTransactionController.UpdateStatus, AdminValidation(), SuperAdminValidation())
	class_transactions.PUT("/update-receipt/:idClassTransaction", ctrlList.ClassTransactionController.UpdateReceipt)

	trainers := e.Group("trainers", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	trainers.GET("", ctrlList.TrainerController.GetAll)

	article := e.Group("articles", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	article.GET("", ctrlList.ArticleController.GetAll)
	article.GET("/:idArticle", ctrlList.ArticleController.GetByID)
	article.POST("", ctrlList.ArticleController.Insert, SuperAdminValidation())
	article.DELETE("/:idArticle", ctrlList.ArticleController.DeleteByID, SuperAdminValidation())
	article.PUT("/:idArticle", ctrlList.ArticleController.UpdateArticleByID, SuperAdminValidation())

	classification := e.Group("classification", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	classification.POST("", ctrlList.ClassificationController.Insert, SuperAdminValidation())
	classification.GET("", ctrlList.ClassificationController.GetAll)

	payment_account := e.Group("payment-account", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	payment_account.POST("", ctrlList.PaymentAccountController.Insert, SuperAdminValidation())
	payment_account.GET("", ctrlList.PaymentAccountController.GetAll)

	admins := e.Group("admins")
	admins.POST("", ctrlList.AdminController.Register)
	admins.POST("/login", ctrlList.AdminController.Login)
	admins.PUT("", ctrlList.AdminController.Update, middleware.JWTWithConfig(ctrlList.JWTMiddleware), AdminValidation())

	membership_transactions := e.Group("transaction-membership", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	membership_transactions.GET("", ctrlList.MembershipTransactionController.GetAll, AdminValidation(), SuperAdminValidation())
	membership_transactions.GET("/user", ctrlList.MembershipTransactionController.GetAllByUser)
	membership_transactions.GET("/:idMembershipTransaction", ctrlList.MembershipTransactionController.GetByID)
	membership_transactions.POST("", ctrlList.MembershipTransactionController.Insert)
	membership_transactions.PUT("/update-status/:idMembershipTransaction", ctrlList.MembershipTransactionController.UpdateStatus, AdminValidation(), SuperAdminValidation())
	membership_transactions.PUT("/update-receipt/:idMembershipTransaction", ctrlList.MembershipTransactionController.UpdateReceipt)

	members := e.Group("members", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	members.GET("/:userId", ctrlList.MemberController.GetByUserID)
}

func AdminValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := _middleware.GetUser(c)
			if claims.IsAdmin {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, business.ErrUnauthorized)
			}
		}
	}
}

func SuperAdminValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := _middleware.GetUser(c)
			if claims.IsSuperAdmin {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, business.ErrUnauthorized)
			}
		}
	}
}

func MemberValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := _middleware.GetUser(c)
			if claims.IsMember {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, business.ErrUnauthorized)
			}
		}
	}
}
