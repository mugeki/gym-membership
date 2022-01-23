package drivers

import (
	usersDomain "gym-membership/business/users"
	usersDB "gym-membership/drivers/databases/users"

	videosDomain "gym-membership/business/videos"
	videosDB "gym-membership/drivers/databases/videos"

	calendarsDomain "gym-membership/business/calendars"
	calendarsAPI "gym-membership/drivers/calendarsApi"

	"google.golang.org/api/calendar/v3"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) usersDomain.Repository {
	return usersDB.NewMySQLRepo(conn)
}

func NewVideoRepository(conn *gorm.DB) videosDomain.Repository {
	return videosDB.NewMySQLRepo(conn)
}

func NewCalendarsRepository(calendarService *calendar.Service) calendarsDomain.Repository {
	return calendarsAPI.NewCalendarsApi(calendarService)
}
