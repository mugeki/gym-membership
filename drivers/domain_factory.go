package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	classDomain "gym-membership/business/class"
	classDB "gym-membership/drivers/databases/class"

	trainerDomain "gym-membership/business/trainers"
	trainerDB "gym-membership/drivers/databases/trainers"

	transactionClassDomain "gym-membership/business/transactionClass"
	transactionClassDB "gym-membership/drivers/databases/transactionClass"

	adminsDomain "gym-membership/business/admins"
	adminsDB "gym-membership/drivers/databases/admins"

	articlesDomain "gym-membership/business/articles"
	articlesDB "gym-membership/drivers/databases/articles"

	classificationDomain "gym-membership/business/classification"
	classificationDB "gym-membership/drivers/databases/classifications"

	videosDomain "gym-membership/business/videos"
	videosDB "gym-membership/drivers/databases/videos"


	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository {
	return usersDB.NewMySQLRepo(conn)
}

func NewClassRepository(conn *gorm.DB) classDomain.Repository {
	return classDB.NewMySQLRepo(conn)
}

func NewTrainerRepository(conn *gorm.DB) trainerDomain.Repository {
	return trainerDB.NewMySQLRepo(conn)
}

func NewTransactionClassRepository(conn *gorm.DB) transactionClassDomain.Repository {
	return transactionClassDB.NewMySQLRepo(conn)
}

func NewAdminRepository(conn *gorm.DB) adminsDomain.Repository {
	return adminsDB.NewMySQLRepo(conn)
}

func NewArticleRepository(conn *gorm.DB) articlesDomain.Repository {
	return articlesDB.NewMySQLRepo(conn)
}

func NewClassificationRepository(conn *gorm.DB) classificationDomain.Repository {
	return classificationDB.NewMySQLRepo(conn)
}

func NewVideoRepository(conn *gorm.DB) videosDomain.Repository{
	return videosDB.NewMySQLRepo(conn)
}

