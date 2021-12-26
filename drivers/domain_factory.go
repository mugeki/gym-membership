package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	classDomain "gym-membership/business/class"
	classDB "gym-membership/drivers/databases/class"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository {
	return usersDB.NewMySQLRepo(conn)
}

func NewClassRepository(conn *gorm.DB) classDomain.Repository {
	return classDB.NewMySQLRepo(conn)
}
