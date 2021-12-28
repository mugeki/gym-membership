package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"
	usersDomain "gym-membership/business/members"
	usersDB "gym-membership/drivers/databases/members"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository{
	return usersDB.NewMySQLRepo(conn)
}