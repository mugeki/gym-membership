package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"
	membershipProductsDomain "gym-membership/business/membership_products"
	membershipProductsDB "gym-membership/drivers/databases/membership_products"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository{
	return usersDB.NewMySQLRepo(conn)
}
func NewMembershipProductsRepository(conn *gorm.DB) membershipProductsDomain.Repository{
	return membershipProductsDB.NewMySQLRepo(conn)
}