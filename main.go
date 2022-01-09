package main

import (
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"

	_userUsecase "gym-membership/business/users"
	_userController "gym-membership/controllers/users"
	_userRepo "gym-membership/drivers/databases/users"

	_membershipProductsUsecase "gym-membership/business/membership_products"
	_membershipProductsController "gym-membership/controllers/membership_products"
	_membershipProductsRepo "gym-membership/drivers/databases/membership_products"

	_classUsecase "gym-membership/business/class"
	_classController "gym-membership/controllers/class"
	_classRepo "gym-membership/drivers/databases/class"

	_trainerUsecase "gym-membership/business/trainers"
	_trainerController "gym-membership/controllers/trainers"
	_trainerRepo "gym-membership/drivers/databases/trainers"

	_transactionClassUsecase "gym-membership/business/transactionClass"
	_transactionClassController "gym-membership/controllers/transactionClass"
	_transactionClassRepo "gym-membership/drivers/databases/transactionClass"

	_transactionMembershipUsecase "gym-membership/business/membership_transactions"
	_transactionMembershipController "gym-membership/controllers/membership_transactions"
	_transactionMembershipRepo "gym-membership/drivers/databases/membership_transactions"

	_memberUsecase "gym-membership/business/members"
	_memberController "gym-membership/controllers/members"
	_memberRepo "gym-membership/drivers/databases/members"

	_adminUsecase "gym-membership/business/admins"
	_adminController "gym-membership/controllers/admins"
	_adminRepo "gym-membership/drivers/databases/admins"

	_articleUsecase "gym-membership/business/articles"
	_articleController "gym-membership/controllers/articles"
	_articleRepo "gym-membership/drivers/databases/articles"

	_classificationUsecase "gym-membership/business/classification"
	_classificationController "gym-membership/controllers/classifications"

	_videoUsecase "gym-membership/business/videos"
	_videoController "gym-membership/controllers/videos"
	_videoRepo "gym-membership/drivers/databases/videos"

	_classificationRepo "gym-membership/drivers/databases/classifications"

	_middleware "gym-membership/app/middleware"
	_routes "gym-membership/app/routes"
	_dbDriver "gym-membership/drivers/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_membershipProductsRepo.MembershipProducts{},
		&_classRepo.Class{},
		&_trainerRepo.Trainers{},
		&_transactionClassRepo.TransactionClass{},
		&_adminRepo.Admins{},
		&_articleRepo.Articles{},
		&_classificationRepo.Classification{},
		&_videoRepo.Videos{},
		&_transactionMembershipRepo.MembershipTransactions{},
		&_memberRepo.Members{},
	)
}

func main() {
	_ = godotenv.Load()
	configDB := _dbDriver.ConfigDB{
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Database: os.Getenv("DB_NAME"),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	EXPIRE, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: int64(EXPIRE),
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
  		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	membershipProductsRepo := _driverFactory.NewMembershipProductsRepository(db)
	membershipProductsUsecase := _membershipProductsUsecase.NewMembershipProductsUsecase(membershipProductsRepo)
	membershipProductsCtrl := _membershipProductsController.NewMembershipProductsController(membershipProductsUsecase)

	classRepo := _driverFactory.NewClassRepository(db)
	classUsecase := _classUsecase.NewClassUsecase(classRepo, &configJWT)
	classCtrl := _classController.NewClassController(classUsecase)

	trainerRepo := _driverFactory.NewTrainerRepository(db)
	trainerUsecase := _trainerUsecase.NewTrainerUsecase(trainerRepo)
	trainerCtrl := _trainerController.NewTrainerController(trainerUsecase)

	transactionClassRepo := _driverFactory.NewTransactionClassRepository(db)
	transactionClassUsecase := _transactionClassUsecase.NewTransactionClassUsecase(transactionClassRepo, classRepo)
	transactionClassCtrl := _transactionClassController.NewTransactionClassController(transactionClassUsecase)

	memberRepo := _driverFactory.NewMemberRepository(db)
	memberUsecase := _memberUsecase.NewMemberUsecase(memberRepo)
	memberCtrl := _memberController.NewMemberController(memberUsecase)

	transactionMembershipRepo := _driverFactory.NewTransactionMembershipRepository(db)
	transactionMembershipUsecase := _transactionMembershipUsecase.NewMembershipTransactionUsecase(transactionMembershipRepo, membershipProductsRepo, memberRepo)
	transactionMembershipCtrl := _transactionMembershipController.NewMembershipTransactionController(transactionMembershipUsecase)

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo, &configJWT)
	adminCtrl := _adminController.NewAdminController(adminUsecase)

	articleRepo := _driverFactory.NewArticleRepository(db)
	classificationRepo := _driverFactory.NewClassificationRepository(db)
	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, classificationRepo)
	articleCtrl := _articleController.NewArticleController(articleUsecase)

	classificationUsecase := _classificationUsecase.NewClassificationUsecase(classificationRepo)
	classificationCtrl := _classificationController.NewClassificationController(classificationUsecase)
  
  	videoRepo := _driverFactory.NewVideoRepository(db)
	videoUsecase := _videoUsecase.NewVideoUsecase(videoRepo)
	videoCtrl := _videoController.NewVideoController(videoUsecase)
  
	routesInit := _routes.ControllerList{
		JWTMiddleware:            configJWT.Init(),
		UserController:           *userCtrl,
   		MembershipProductsController:     *membershipProductsCtrl,
		AdminController:          *adminCtrl,
		ArticleController:        *articleCtrl,
		ClassificationController: *classificationCtrl,
    	VideoController:	*videoCtrl,
    	ClassController:            *classCtrl,
		TrainerController:          *trainerCtrl,
		TransactionClassController: *transactionClassCtrl,
		MemberController: *memberCtrl,
		MembershipTransactionController: *transactionMembershipCtrl,
	}
	routesInit.RegisterRoute(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
