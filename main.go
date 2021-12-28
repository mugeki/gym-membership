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
		&_membershipProductsRepo.MembershipProducts{},

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
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	membershipProductsRepo := _driverFactory.NewMembershipProductsRepository(db)
	membershipProductsUsecase := _membershipProductsUsecase.NewMembershipProductsUsecase(membershipProductsRepo)
	membershipProductsCtrl := _membershipProductsController.NewMembershipProductsController(membershipProductsUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userCtrl,
		MembershipProductsController:     *membershipProductsCtrl,
	}
	routesInit.RegisterRoute(e)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":"+port)
}