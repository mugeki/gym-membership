package main

import (
	"os"
	"strconv"

	_driverFactory "gym-membership/drivers"
	"gym-membership/helper/encrypt"

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

	_classTransactionUsecase "gym-membership/business/class_transactions"
	_classTransactionController "gym-membership/controllers/class_transactions"
	_classTransactionRepo "gym-membership/drivers/databases/class_transactions"

	_membershipTransactionUsecase "gym-membership/business/membership_transactions"
	_membershipTransactionController "gym-membership/controllers/membership_transactions"
	_membershipTransactionRepo "gym-membership/drivers/databases/membership_transactions"

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
	_classificationRepo "gym-membership/drivers/databases/classifications"

	_videoUsecase "gym-membership/business/videos"
	_videoController "gym-membership/controllers/videos"
	_videoRepo "gym-membership/drivers/databases/videos"

	_paymentAccountUsecase "gym-membership/business/payment_accounts"
	_paymentAccountController "gym-membership/controllers/payment_accounts"
	_paymentAccountRepo "gym-membership/drivers/databases/payment_accounts"

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
		&_classTransactionRepo.ClassTransaction{},
		&_adminRepo.Admins{},
		&_articleRepo.Articles{},
		&_classificationRepo.Classification{},
		&_videoRepo.Videos{},
		&_membershipTransactionRepo.MembershipTransactions{},
		&_memberRepo.Members{},
		&_paymentAccountRepo.PaymentAccount{},
	)
	hashed, _ := encrypt.Hash("admin321")
	admin := _adminRepo.Admins{
		Model: gorm.Model{ID:1},
		Username: "admin321",
		Password: hashed,
		Email: "admin321@gmail.com",
		FullName: "Super Admin",
		Gender: "male",
		Telephone: "0881122334455",
		Address: "Jl. H. Sadiah, Joglo",
		UrlImage: "https://images.unsplash.com/photo-1584952811565-c4c4031805a8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80",
		IsSuperAdmin: true,
	}
	classification := []_classificationRepo.Classification{
		{
			ID: 1, 
			Name:"Workout Tips",
		},
		{
			ID: 2, 
			Name:"Health Tips",
		},
	}
	trainer := []_trainerRepo.Trainers{
		{
			ID: 1,
			Fullname: "John Doe", 
			UrlImage: "https://images.unsplash.com/photo-1584952811565-c4c4031805a8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80",
		},
		{
			ID: 2,
			Fullname:"Jane Doe", 
			UrlImage: "https://images.unsplash.com/photo-1550345332-09e3ac987658?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80",
		},
	}
	payment := []_paymentAccountRepo.PaymentAccount{
		{
			Model: gorm.Model{ID:1},
			Name: "BRI",
			NoCard: "75757182839481727",
			OwnerName: "PT Subur Jaya",
			Desc: "This payment need a manual confirmation by uploading image of receipt",
		},
		{
			Model: gorm.Model{ID:2},
			Name: "BCA",
			NoCard: "75757182839481727",
			OwnerName: "PT Subur Jaya",
			Desc: "This payment need a manual confirmation by uploading image of receipt",
		},
		{
			Model: gorm.Model{ID:3},
			Name: "LINKAJA",
			NoCard: "75757182839481727",
			OwnerName: "PT Subur Jaya",
			Desc: "This payment need a manual confirmation by uploading image of receipt",
		},
		{
			Model: gorm.Model{ID:4},
			Name: "GOPAY",
			NoCard: "75757182839481727",
			OwnerName: "PT Subur Jaya",
			Desc: "This payment need a manual confirmation by uploading image of receipt",
		},
	}

	db.Create(&admin)
	db.Create(&classification)
	db.Create(&trainer)
	db.Create(&payment)
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
		AllowOrigins: []string{
			"http://localhost",
			"http://ec2-13-58-52-197.us-east-2.compute.amazonaws.com",
			"http://ec2-18-222-186-208.us-east-2.compute.amazonaws.com",
			"http://gymbro.my.id",
			"http://gymbro-admin.my.id",
		},
		AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType, 
				echo.HeaderAccept,
				echo.HeaderAccessControlAllowCredentials,
				echo.HeaderAccessControlAllowOrigin,
				echo.HeaderAuthorization,
			},
	}))

	memberRepo := _driverFactory.NewMemberRepository(db)
	memberUsecase := _memberUsecase.NewMemberUsecase(memberRepo)
	memberCtrl := _memberController.NewMemberController(memberUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, memberRepo, &configJWT)
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

	paymentAccountRepo := _driverFactory.NewPaymentAccountRepository(db)
	paymentAccountUsecase := _paymentAccountUsecase.NewPaymentAccountUsecase(paymentAccountRepo)
	paymentAccountCtrl := _paymentAccountController.NewPaymentAccountController(paymentAccountUsecase)

	classTransactionRepo := _driverFactory.NewClassTransactionRepository(db)
	classTransactionUsecase := _classTransactionUsecase.NewClassTransactionUsecase(classTransactionRepo, classRepo)
	classTransactionCtrl := _classTransactionController.NewClassTransactionController(classTransactionUsecase)

	membershipTransactionRepo := _driverFactory.NewMembershipTransactionRepository(db)
	membershipTransactionUsecase := _membershipTransactionUsecase.NewMembershipTransactionUsecase(membershipTransactionRepo, membershipProductsRepo, memberRepo)
	membershipTransactionCtrl := _membershipTransactionController.NewMembershipTransactionController(membershipTransactionUsecase)

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
		JWTMiddleware:                   configJWT.Init(),
		UserController:                  *userCtrl,
		MembershipProductsController:    *membershipProductsCtrl,
		AdminController:                 *adminCtrl,
		ArticleController:               *articleCtrl,
		ClassificationController:        *classificationCtrl,
		VideoController:                 *videoCtrl,
		ClassController:                 *classCtrl,
		TrainerController:               *trainerCtrl,
		ClassTransactionController:      *classTransactionCtrl,
		MemberController:                *memberCtrl,
		MembershipTransactionController: *membershipTransactionCtrl,
		PaymentAccountController:        *paymentAccountCtrl,
	}
	routesInit.RegisterRoute(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Start(":" + port)
}
