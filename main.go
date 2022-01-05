package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"

	_userUsecase "gym-membership/business/users"
	_userController "gym-membership/controllers/users"
	_userRepo "gym-membership/drivers/databases/users"

	_videoService "gym-membership/business/videos"
	_videoController "gym-membership/controllers/videos"
	_videoRepo "gym-membership/drivers/databases/videos"

	_adminRepo "gym-membership/drivers/databases/admins"
	_classificationRepo "gym-membership/drivers/databases/classifications"

	_calendarsApiService "gym-membership/business/calendars"
	_calendarsApiController "gym-membership/controllers/calendars"
	_calendarsApiRepo "gym-membership/drivers/calendarsApi"

	_middleware "gym-membership/app/middleware"
	_routes "gym-membership/app/routes"
	_dbDriver "gym-membership/drivers/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/calendar/v3"
	"gorm.io/gorm"

	"golang.org/x/oauth2"
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

	// ctx := context.Background()
	// credentialsGoogle := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	// client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsGoogle))
	token := _calendarsApiRepo.RequestToken()
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	fmt.Println("client =========== ", client)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// fmt.Println(client, "client print")
	// defer client.Close()
	calendarService, err := calendar.New(client)
	if err != nil {
		fmt.Println("error: ", err)
	}
	// Sets the name for the new bucket.
	// bucketName := "my-new-bucket"

	// // Creates a Bucket instance.
	// bucket := client.Bucket(bucketName)

	// // Creates the new bucket.
	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()
	// if err := bucket.Create(ctx, "primal-archive-319313", nil); err != nil {
	// 	log.Fatalf("Failed to create bucket: %v", err)
	// }

	// fmt.Printf("Bucket %v created.\n", bucketName)
	EXPIRE, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: int64(EXPIRE),
	}
	e := echo.New()
	calendarsApiRepo := _calendarsApiRepo.NewCalendarsApi(calendarService)
	calendarsApiUsecase := _calendarsApiService.NewCalendarUsecase(calendarsApiRepo)
	calendarsApiCtrl := _calendarsApiController.NewCalendarsController(calendarsApiUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	videoRepo := _driverFactory.NewVideoRepository(db)
	videoUsecase := _videoService.NewVideoUsecase(videoRepo)
	videoCtrl := _videoController.NewVideoController(videoUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:          configJWT.Init(),
		UserController:         *userCtrl,
		VideoController:        *videoCtrl,
		CalendarsApiController: *calendarsApiCtrl,
	}
	routesInit.RegisterRoute(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
