package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	adminsDomain "gym-membership/business/admins"
	adminsDB "gym-membership/drivers/databases/admins"

	articlesDomain "gym-membership/business/articles"
	articlesDB "gym-membership/drivers/databases/articles"

	classificationDomain "gym-membership/business/classification"
	classificationDB "gym-membership/drivers/databases/classifications"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository {
	return usersDB.NewMySQLRepo(conn)
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
