package main

import (
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"

	_userUsecase "gym-membership/business/users"
	_userController "gym-membership/controllers/users"
	_userRepo "gym-membership/drivers/databases/users"

	_classService "gym-membership/business/class"
	_classController "gym-membership/controllers/class"
	_classRepo "gym-membership/drivers/databases/class"

	_trainerService "gym-membership/business/trainers"
	_trainerController "gym-membership/controllers/trainers"
	_trainerRepo "gym-membership/drivers/databases/trainers"

	_transactionClassService "gym-membership/business/transactionClass"
	_transactionClassController "gym-membership/controllers/transactionClass"
	_transactionClassRepo "gym-membership/drivers/databases/transactionClass"

	_adminService "gym-membership/business/admins"
	_adminController "gym-membership/controllers/admins"
	_adminRepo "gym-membership/drivers/databases/admins"

	_articleService "gym-membership/business/articles"
	_articleController "gym-membership/controllers/articles"
	_articleRepo "gym-membership/drivers/databases/articles"

	_classificationService "gym-membership/business/classification"
	_classificationController "gym-membership/controllers/classifications"

	_videoService "gym-membership/business/videos"
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
		&_classRepo.Class{},
		&_trainerRepo.Trainers{},
		&_transactionClassRepo.TransactionClass{},
		&_adminRepo.Admins{},
		&_articleRepo.Articles{},
		&_classificationRepo.Classification{},
		&_videoRepo.Videos{},
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

	classRepo := _driverFactory.NewClassRepository(db)
	classUsecase := _classService.NewClassUsecase(classRepo, &configJWT)
	classCtrl := _classController.NewClassController(classUsecase)

	trainerRepo := _driverFactory.NewTrainerRepository(db)
	trainerUsecase := _trainerService.NewTrainerUsecase(trainerRepo)
	trainerCtrl := _trainerController.NewTrainerController(trainerUsecase)

	transactionClassRepo := _driverFactory.NewTransactionClassRepository(db)
	transactionClassUsecase := _transactionClassService.NewTransactionClassUsecase(transactionClassRepo, classRepo)
	transactionClassCtrl := _transactionClassController.NewTransactionClassController(transactionClassUsecase)

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminUsecase := _adminService.NewAdminUsecase(adminRepo, &configJWT)
	adminCtrl := _adminController.NewAdminController(adminUsecase)

	articleRepo := _driverFactory.NewArticleRepository(db)
	classificationRepo := _driverFactory.NewClassificationRepository(db)
	articleUsecase := _articleService.NewArticleUsecase(articleRepo, classificationRepo)
	articleCtrl := _articleController.NewArticleController(articleUsecase)

	classificationUsecase := _classificationService.NewClassificationUsecase(classificationRepo)
	classificationCtrl := _classificationController.NewClassificationController(classificationUsecase)
  
  	videoRepo := _driverFactory.NewVideoRepository(db)
	videoUsecase := _videoService.NewVideoUsecase(videoRepo)
	videoCtrl := _videoController.NewVideoController(videoUsecase)
  
	routesInit := _routes.ControllerList{
		JWTMiddleware:            configJWT.Init(),
		UserController:           *userCtrl,
		AdminController:          *adminCtrl,
		ArticleController:        *articleCtrl,
		ClassificationController: *classificationCtrl,
    VideoController:	*videoCtrl,
    ClassController:            *classCtrl,
		TrainerController:          *trainerCtrl,
		TransactionClassController: *transactionClassCtrl,
	}
	routesInit.RegisterRoute(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
