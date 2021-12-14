package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	videosDomain "gym-membership/business/videos"
	videosDB "gym-membership/drivers/databases/videos"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository{
	return usersDB.NewMySQLRepo(conn)
}

func NewVideoRepository(conn *gorm.DB) videosDomain.Repository{
	return videosDB.NewMySQLRepo(conn)
}