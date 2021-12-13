package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository{
	return usersDB.NewMySQLRepo(conn)
}