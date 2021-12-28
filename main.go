package main

import (
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"

	_userService "gym-membership/business/users"
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
		&_classRepo.Class{},
		&_trainerRepo.Trainers{},
		&_transactionClassRepo.TransactionClass{},
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

	classRepo := _driverFactory.NewClassRepository(db)
	classUsecase := _classService.NewClassUsecase(classRepo, &configJWT)
	classCtrl := _classController.NewClassController(classUsecase)

	trainerRepo := _driverFactory.NewTrainerRepository(db)
	trainerUsecase := _trainerService.NewTrainerUsecase(trainerRepo)
	trainerCtrl := _trainerController.NewTrainerController(trainerUsecase)

	transactionClassRepo := _driverFactory.NewTransactionClassRepository(db)
	transactionClassUsecase := _transactionClassService.NewTransactionClassUsecase(transactionClassRepo)
	transactionClassCtrl := _transactionClassController.NewTransactionClassController(transactionClassUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userCtrl,
		ClassController:      *classCtrl,
		TrainerController:    *trainerCtrl,
		TransactionClassCtrl: *transactionClassCtrl,
	}
	routesInit.RegisterRoute(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
