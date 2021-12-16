package main

import (
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"

	_userService "gym-membership/business/users"
	_userController "gym-membership/controllers/users"
	_userRepo "gym-membership/drivers/databases/users"

	_videoService "gym-membership/business/videos"
	_videoController "gym-membership/controllers/videos"
	_videoRepo "gym-membership/drivers/databases/videos"

	_adminRepo "gym-membership/drivers/databases/admins"
	_classificationRepo "gym-membership/drivers/databases/classifications"

	_middleware "gym-membership/app/middleware"
	_routes "gym-membership/app/routes"
	_dbDriver "gym-membership/drivers/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_adminRepo.Admins{},
		&_classificationRepo.Classifications{},
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

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userService.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	videoRepo := _driverFactory.NewVideoRepository(db)
	videoUsecase := _videoService.NewVideoUsecase(videoRepo)
	videoCtrl := _videoController.NewVideoController(videoUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:		configJWT.Init(),
		UserController:		*userCtrl,
		VideoController:	*videoCtrl,	
	}
	routesInit.RegisterRoute(e)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":"+port)
}