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
